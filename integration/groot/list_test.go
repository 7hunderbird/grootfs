package groot_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	yaml "gopkg.in/yaml.v2"

	"code.cloudfoundry.org/grootfs/commands/config"
	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/grootfs/integration"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("List", func() {
	var image groot.Image

	BeforeEach(func() {
		sourceImagePath, err := ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		Expect(ioutil.WriteFile(path.Join(sourceImagePath, "foo"), []byte("hello-world"), 0644)).To(Succeed())
		baseImageFile := integration.CreateBaseImageTar(sourceImagePath)
		image = integration.CreateImage(GrootFSBin, StorePath, DraxBin, baseImageFile.Name(), "root-image", 0)
	})

	It("lists all images in the store path", func() {
		images, err := Runner.List()
		Expect(err).NotTo(HaveOccurred())
		Expect(images).To(HaveLen(1))
		Expect(images[0].Path).To(Equal(image.Path))
	})

	Describe("--config global flag", func() {
		var (
			configDir      string
			configFilePath string
		)

		BeforeEach(func() {
			var err error
			configDir, err = ioutil.TempDir("", "")
			Expect(err).NotTo(HaveOccurred())
			configFilePath = path.Join(configDir, "config.yaml")

			cfg := config.Config{
				BaseStorePath: StorePath,
			}

			configYaml, err := yaml.Marshal(cfg)
			Expect(err).NotTo(HaveOccurred())

			Expect(ioutil.WriteFile(configFilePath, configYaml, 0755)).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.RemoveAll(configDir)).To(Succeed())
		})

		Describe("store path", func() {
			It("uses the store path from the config file", func() {
				cmd := exec.Command(GrootFSBin, "--config", configFilePath, "list")

				sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(sess).Should(gexec.Exit(0))
				Expect(sess.Out).To(gbytes.Say(image.Path))
			})
		})
	})

	Context("when the store does not exist", func() {
		It("fails and logs the error", func() {
			logBuffer := gbytes.NewBuffer()
			_, err := Runner.WithStore("/invalid-store").WithStderr(logBuffer).
				List()
			Expect(err).To(HaveOccurred())
			Expect(logBuffer).To(gbytes.Say(`"error":"no store found at /invalid-store"`))
		})
	})

	Context("when there are no existing images", func() {
		BeforeEach(func() {
			Runner.Delete(image.Path)
		})

		It("returns an informative message", func() {
			cmd := exec.Command(GrootFSBin, "--store", StorePath, "list")

			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(sess).Should(gexec.Exit(0))
			Expect(sess.Out).To(gbytes.Say("Store empty"))
		})
	})
})
