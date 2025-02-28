package commands // import "code.cloudfoundry.org/grootfs/commands"

import (
	"fmt"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/lager"

	"code.cloudfoundry.org/grootfs/commands/config"
	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/grootfs/metrics"
	storepkg "code.cloudfoundry.org/grootfs/store"
	"code.cloudfoundry.org/grootfs/store/dependency_manager"
	"code.cloudfoundry.org/grootfs/store/filesystems/overlayxfs"
	"code.cloudfoundry.org/grootfs/store/garbage_collector"
	imageClonerpkg "code.cloudfoundry.org/grootfs/store/image_cloner"
	locksmithpkg "code.cloudfoundry.org/grootfs/store/locksmith"
	errorspkg "github.com/pkg/errors"

	"github.com/urfave/cli"
)

var CleanCommand = cli.Command{
	Name:        "clean",
	Usage:       "clean",
	Description: "Cleans up unused layers",

	Flags: []cli.Flag{
		cli.Int64Flag{
			Name:  "threshold-bytes",
			Usage: "Disk usage of the store directory at which cleanup should trigger",
		},
	},

	Action: func(ctx *cli.Context) error {
		logger := ctx.App.Metadata["logger"].(lager.Logger)
		logger = logger.Session("clean")

		configBuilder := ctx.App.Metadata["configBuilder"].(*config.Builder)
		configBuilder.WithCleanThresholdBytes(ctx.Int64("threshold-bytes"),
			ctx.IsSet("threshold-bytes"))

		cfg, err := configBuilder.Build()
		logger.Debug("clean-config", lager.Data{"currentConfig": cfg})
		if err != nil {
			logger.Error("config-builder-failed", err)
			return cli.NewExitError(err.Error(), 1)
		}

		storePath := cfg.StorePath
		if _, err = os.Stat(storePath); os.IsNotExist(err) {
			err = errorspkg.Errorf("no store found at %s", storePath)
			logger.Error("store-path-failed", err, nil)
			return cli.NewExitError(err.Error(), 0)
		}

		fsDriver := overlayxfs.NewDriver(cfg.StorePath, cfg.TardisBin)

		imageCloner := imageClonerpkg.NewImageCloner(fsDriver, storePath)
		metricsEmitter := metrics.NewEmitter(logger, cfg.MetronEndpoint)
		locksDir := filepath.Join(storePath, storepkg.LocksDirName)
		locksmith := locksmithpkg.NewExclusiveFileSystem(locksDir).WithMetrics(metricsEmitter)
		dependencyManager := dependency_manager.NewDependencyManager(
			filepath.Join(storePath, storepkg.MetaDirName, "dependencies"),
		)

		nsFsDriver, err := createImageDriver(logger, cfg, fsDriver)
		if err != nil {
			logger.Error("failed-to-create-image-driver", err)
			return cli.NewExitError(err.Error(), 1)
		}
		gc := garbage_collector.NewGC(nsFsDriver, imageCloner, dependencyManager)
		sm := storepkg.NewStoreMeasurer(storePath, fsDriver, gc)

		cleaner := groot.IamCleaner(locksmith, sm, gc, metricsEmitter)

		defer func() {
			unusedVolumesSize, err := sm.UnusedVolumesSize(logger)
			if err != nil {
				logger.Error("getting-unused-volumes-size", err)
			}
			metricsEmitter.TryEmitUsage(logger, "UnusedLayersSize", unusedVolumesSize, "bytes")
		}()

		noop, err := cleaner.Clean(logger, cfg.Clean.ThresholdBytes)
		if err != nil {
			logger.Error("cleaning-up-unused-resources", err)
			return cli.NewExitError(err.Error(), 1)
		}

		if noop {
			fmt.Println("threshold not reached: skipping clean")
			return nil
		}

		fmt.Println("clean completed")

		return nil
	},
}
