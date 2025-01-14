// Code generated by counterfeiter. DO NOT EDIT.
package layer_fetcherfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/grootfs/fetcher/layer_fetcher"
	"github.com/containers/image/docker/reference"
	"github.com/containers/image/types"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type FakeManifest struct {
	ConfigBlobStub        func(context.Context) ([]byte, error)
	configBlobMutex       sync.RWMutex
	configBlobArgsForCall []struct {
		arg1 context.Context
	}
	configBlobReturns struct {
		result1 []byte
		result2 error
	}
	configBlobReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	ConfigInfoStub        func() types.BlobInfo
	configInfoMutex       sync.RWMutex
	configInfoArgsForCall []struct {
	}
	configInfoReturns struct {
		result1 types.BlobInfo
	}
	configInfoReturnsOnCall map[int]struct {
		result1 types.BlobInfo
	}
	EmbeddedDockerReferenceConflictsStub        func(reference.Named) bool
	embeddedDockerReferenceConflictsMutex       sync.RWMutex
	embeddedDockerReferenceConflictsArgsForCall []struct {
		arg1 reference.Named
	}
	embeddedDockerReferenceConflictsReturns struct {
		result1 bool
	}
	embeddedDockerReferenceConflictsReturnsOnCall map[int]struct {
		result1 bool
	}
	InspectStub        func(context.Context) (*types.ImageInspectInfo, error)
	inspectMutex       sync.RWMutex
	inspectArgsForCall []struct {
		arg1 context.Context
	}
	inspectReturns struct {
		result1 *types.ImageInspectInfo
		result2 error
	}
	inspectReturnsOnCall map[int]struct {
		result1 *types.ImageInspectInfo
		result2 error
	}
	LayerInfosStub        func() []types.BlobInfo
	layerInfosMutex       sync.RWMutex
	layerInfosArgsForCall []struct {
	}
	layerInfosReturns struct {
		result1 []types.BlobInfo
	}
	layerInfosReturnsOnCall map[int]struct {
		result1 []types.BlobInfo
	}
	LayerInfosForCopyStub        func(context.Context) ([]types.BlobInfo, error)
	layerInfosForCopyMutex       sync.RWMutex
	layerInfosForCopyArgsForCall []struct {
		arg1 context.Context
	}
	layerInfosForCopyReturns struct {
		result1 []types.BlobInfo
		result2 error
	}
	layerInfosForCopyReturnsOnCall map[int]struct {
		result1 []types.BlobInfo
		result2 error
	}
	ManifestStub        func(context.Context) ([]byte, string, error)
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct {
		arg1 context.Context
	}
	manifestReturns struct {
		result1 []byte
		result2 string
		result3 error
	}
	manifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 string
		result3 error
	}
	OCIConfigStub        func(context.Context) (*v1.Image, error)
	oCIConfigMutex       sync.RWMutex
	oCIConfigArgsForCall []struct {
		arg1 context.Context
	}
	oCIConfigReturns struct {
		result1 *v1.Image
		result2 error
	}
	oCIConfigReturnsOnCall map[int]struct {
		result1 *v1.Image
		result2 error
	}
	ReferenceStub        func() types.ImageReference
	referenceMutex       sync.RWMutex
	referenceArgsForCall []struct {
	}
	referenceReturns struct {
		result1 types.ImageReference
	}
	referenceReturnsOnCall map[int]struct {
		result1 types.ImageReference
	}
	SignaturesStub        func(context.Context) ([][]byte, error)
	signaturesMutex       sync.RWMutex
	signaturesArgsForCall []struct {
		arg1 context.Context
	}
	signaturesReturns struct {
		result1 [][]byte
		result2 error
	}
	signaturesReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	SizeStub        func() (int64, error)
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct {
	}
	sizeReturns struct {
		result1 int64
		result2 error
	}
	sizeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	UpdatedImageStub        func(context.Context, types.ManifestUpdateOptions) (types.Image, error)
	updatedImageMutex       sync.RWMutex
	updatedImageArgsForCall []struct {
		arg1 context.Context
		arg2 types.ManifestUpdateOptions
	}
	updatedImageReturns struct {
		result1 types.Image
		result2 error
	}
	updatedImageReturnsOnCall map[int]struct {
		result1 types.Image
		result2 error
	}
	UpdatedImageNeedsLayerDiffIDsStub        func(types.ManifestUpdateOptions) bool
	updatedImageNeedsLayerDiffIDsMutex       sync.RWMutex
	updatedImageNeedsLayerDiffIDsArgsForCall []struct {
		arg1 types.ManifestUpdateOptions
	}
	updatedImageNeedsLayerDiffIDsReturns struct {
		result1 bool
	}
	updatedImageNeedsLayerDiffIDsReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifest) ConfigBlob(arg1 context.Context) ([]byte, error) {
	fake.configBlobMutex.Lock()
	ret, specificReturn := fake.configBlobReturnsOnCall[len(fake.configBlobArgsForCall)]
	fake.configBlobArgsForCall = append(fake.configBlobArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("ConfigBlob", []interface{}{arg1})
	fake.configBlobMutex.Unlock()
	if fake.ConfigBlobStub != nil {
		return fake.ConfigBlobStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.configBlobReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) ConfigBlobCallCount() int {
	fake.configBlobMutex.RLock()
	defer fake.configBlobMutex.RUnlock()
	return len(fake.configBlobArgsForCall)
}

func (fake *FakeManifest) ConfigBlobCalls(stub func(context.Context) ([]byte, error)) {
	fake.configBlobMutex.Lock()
	defer fake.configBlobMutex.Unlock()
	fake.ConfigBlobStub = stub
}

func (fake *FakeManifest) ConfigBlobArgsForCall(i int) context.Context {
	fake.configBlobMutex.RLock()
	defer fake.configBlobMutex.RUnlock()
	argsForCall := fake.configBlobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) ConfigBlobReturns(result1 []byte, result2 error) {
	fake.configBlobMutex.Lock()
	defer fake.configBlobMutex.Unlock()
	fake.ConfigBlobStub = nil
	fake.configBlobReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) ConfigBlobReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.configBlobMutex.Lock()
	defer fake.configBlobMutex.Unlock()
	fake.ConfigBlobStub = nil
	if fake.configBlobReturnsOnCall == nil {
		fake.configBlobReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.configBlobReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) ConfigInfo() types.BlobInfo {
	fake.configInfoMutex.Lock()
	ret, specificReturn := fake.configInfoReturnsOnCall[len(fake.configInfoArgsForCall)]
	fake.configInfoArgsForCall = append(fake.configInfoArgsForCall, struct {
	}{})
	fake.recordInvocation("ConfigInfo", []interface{}{})
	fake.configInfoMutex.Unlock()
	if fake.ConfigInfoStub != nil {
		return fake.ConfigInfoStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.configInfoReturns
	return fakeReturns.result1
}

func (fake *FakeManifest) ConfigInfoCallCount() int {
	fake.configInfoMutex.RLock()
	defer fake.configInfoMutex.RUnlock()
	return len(fake.configInfoArgsForCall)
}

func (fake *FakeManifest) ConfigInfoCalls(stub func() types.BlobInfo) {
	fake.configInfoMutex.Lock()
	defer fake.configInfoMutex.Unlock()
	fake.ConfigInfoStub = stub
}

func (fake *FakeManifest) ConfigInfoReturns(result1 types.BlobInfo) {
	fake.configInfoMutex.Lock()
	defer fake.configInfoMutex.Unlock()
	fake.ConfigInfoStub = nil
	fake.configInfoReturns = struct {
		result1 types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) ConfigInfoReturnsOnCall(i int, result1 types.BlobInfo) {
	fake.configInfoMutex.Lock()
	defer fake.configInfoMutex.Unlock()
	fake.ConfigInfoStub = nil
	if fake.configInfoReturnsOnCall == nil {
		fake.configInfoReturnsOnCall = make(map[int]struct {
			result1 types.BlobInfo
		})
	}
	fake.configInfoReturnsOnCall[i] = struct {
		result1 types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflicts(arg1 reference.Named) bool {
	fake.embeddedDockerReferenceConflictsMutex.Lock()
	ret, specificReturn := fake.embeddedDockerReferenceConflictsReturnsOnCall[len(fake.embeddedDockerReferenceConflictsArgsForCall)]
	fake.embeddedDockerReferenceConflictsArgsForCall = append(fake.embeddedDockerReferenceConflictsArgsForCall, struct {
		arg1 reference.Named
	}{arg1})
	fake.recordInvocation("EmbeddedDockerReferenceConflicts", []interface{}{arg1})
	fake.embeddedDockerReferenceConflictsMutex.Unlock()
	if fake.EmbeddedDockerReferenceConflictsStub != nil {
		return fake.EmbeddedDockerReferenceConflictsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.embeddedDockerReferenceConflictsReturns
	return fakeReturns.result1
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsCallCount() int {
	fake.embeddedDockerReferenceConflictsMutex.RLock()
	defer fake.embeddedDockerReferenceConflictsMutex.RUnlock()
	return len(fake.embeddedDockerReferenceConflictsArgsForCall)
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsCalls(stub func(reference.Named) bool) {
	fake.embeddedDockerReferenceConflictsMutex.Lock()
	defer fake.embeddedDockerReferenceConflictsMutex.Unlock()
	fake.EmbeddedDockerReferenceConflictsStub = stub
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsArgsForCall(i int) reference.Named {
	fake.embeddedDockerReferenceConflictsMutex.RLock()
	defer fake.embeddedDockerReferenceConflictsMutex.RUnlock()
	argsForCall := fake.embeddedDockerReferenceConflictsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsReturns(result1 bool) {
	fake.embeddedDockerReferenceConflictsMutex.Lock()
	defer fake.embeddedDockerReferenceConflictsMutex.Unlock()
	fake.EmbeddedDockerReferenceConflictsStub = nil
	fake.embeddedDockerReferenceConflictsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) EmbeddedDockerReferenceConflictsReturnsOnCall(i int, result1 bool) {
	fake.embeddedDockerReferenceConflictsMutex.Lock()
	defer fake.embeddedDockerReferenceConflictsMutex.Unlock()
	fake.EmbeddedDockerReferenceConflictsStub = nil
	if fake.embeddedDockerReferenceConflictsReturnsOnCall == nil {
		fake.embeddedDockerReferenceConflictsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.embeddedDockerReferenceConflictsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) Inspect(arg1 context.Context) (*types.ImageInspectInfo, error) {
	fake.inspectMutex.Lock()
	ret, specificReturn := fake.inspectReturnsOnCall[len(fake.inspectArgsForCall)]
	fake.inspectArgsForCall = append(fake.inspectArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Inspect", []interface{}{arg1})
	fake.inspectMutex.Unlock()
	if fake.InspectStub != nil {
		return fake.InspectStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.inspectReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) InspectCallCount() int {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	return len(fake.inspectArgsForCall)
}

func (fake *FakeManifest) InspectCalls(stub func(context.Context) (*types.ImageInspectInfo, error)) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = stub
}

func (fake *FakeManifest) InspectArgsForCall(i int) context.Context {
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	argsForCall := fake.inspectArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) InspectReturns(result1 *types.ImageInspectInfo, result2 error) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = nil
	fake.inspectReturns = struct {
		result1 *types.ImageInspectInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) InspectReturnsOnCall(i int, result1 *types.ImageInspectInfo, result2 error) {
	fake.inspectMutex.Lock()
	defer fake.inspectMutex.Unlock()
	fake.InspectStub = nil
	if fake.inspectReturnsOnCall == nil {
		fake.inspectReturnsOnCall = make(map[int]struct {
			result1 *types.ImageInspectInfo
			result2 error
		})
	}
	fake.inspectReturnsOnCall[i] = struct {
		result1 *types.ImageInspectInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) LayerInfos() []types.BlobInfo {
	fake.layerInfosMutex.Lock()
	ret, specificReturn := fake.layerInfosReturnsOnCall[len(fake.layerInfosArgsForCall)]
	fake.layerInfosArgsForCall = append(fake.layerInfosArgsForCall, struct {
	}{})
	fake.recordInvocation("LayerInfos", []interface{}{})
	fake.layerInfosMutex.Unlock()
	if fake.LayerInfosStub != nil {
		return fake.LayerInfosStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.layerInfosReturns
	return fakeReturns.result1
}

func (fake *FakeManifest) LayerInfosCallCount() int {
	fake.layerInfosMutex.RLock()
	defer fake.layerInfosMutex.RUnlock()
	return len(fake.layerInfosArgsForCall)
}

func (fake *FakeManifest) LayerInfosCalls(stub func() []types.BlobInfo) {
	fake.layerInfosMutex.Lock()
	defer fake.layerInfosMutex.Unlock()
	fake.LayerInfosStub = stub
}

func (fake *FakeManifest) LayerInfosReturns(result1 []types.BlobInfo) {
	fake.layerInfosMutex.Lock()
	defer fake.layerInfosMutex.Unlock()
	fake.LayerInfosStub = nil
	fake.layerInfosReturns = struct {
		result1 []types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) LayerInfosReturnsOnCall(i int, result1 []types.BlobInfo) {
	fake.layerInfosMutex.Lock()
	defer fake.layerInfosMutex.Unlock()
	fake.LayerInfosStub = nil
	if fake.layerInfosReturnsOnCall == nil {
		fake.layerInfosReturnsOnCall = make(map[int]struct {
			result1 []types.BlobInfo
		})
	}
	fake.layerInfosReturnsOnCall[i] = struct {
		result1 []types.BlobInfo
	}{result1}
}

func (fake *FakeManifest) LayerInfosForCopy(arg1 context.Context) ([]types.BlobInfo, error) {
	fake.layerInfosForCopyMutex.Lock()
	ret, specificReturn := fake.layerInfosForCopyReturnsOnCall[len(fake.layerInfosForCopyArgsForCall)]
	fake.layerInfosForCopyArgsForCall = append(fake.layerInfosForCopyArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("LayerInfosForCopy", []interface{}{arg1})
	fake.layerInfosForCopyMutex.Unlock()
	if fake.LayerInfosForCopyStub != nil {
		return fake.LayerInfosForCopyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.layerInfosForCopyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) LayerInfosForCopyCallCount() int {
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	return len(fake.layerInfosForCopyArgsForCall)
}

func (fake *FakeManifest) LayerInfosForCopyCalls(stub func(context.Context) ([]types.BlobInfo, error)) {
	fake.layerInfosForCopyMutex.Lock()
	defer fake.layerInfosForCopyMutex.Unlock()
	fake.LayerInfosForCopyStub = stub
}

func (fake *FakeManifest) LayerInfosForCopyArgsForCall(i int) context.Context {
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	argsForCall := fake.layerInfosForCopyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) LayerInfosForCopyReturns(result1 []types.BlobInfo, result2 error) {
	fake.layerInfosForCopyMutex.Lock()
	defer fake.layerInfosForCopyMutex.Unlock()
	fake.LayerInfosForCopyStub = nil
	fake.layerInfosForCopyReturns = struct {
		result1 []types.BlobInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) LayerInfosForCopyReturnsOnCall(i int, result1 []types.BlobInfo, result2 error) {
	fake.layerInfosForCopyMutex.Lock()
	defer fake.layerInfosForCopyMutex.Unlock()
	fake.LayerInfosForCopyStub = nil
	if fake.layerInfosForCopyReturnsOnCall == nil {
		fake.layerInfosForCopyReturnsOnCall = make(map[int]struct {
			result1 []types.BlobInfo
			result2 error
		})
	}
	fake.layerInfosForCopyReturnsOnCall[i] = struct {
		result1 []types.BlobInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) Manifest(arg1 context.Context) ([]byte, string, error) {
	fake.manifestMutex.Lock()
	ret, specificReturn := fake.manifestReturnsOnCall[len(fake.manifestArgsForCall)]
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Manifest", []interface{}{arg1})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.manifestReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeManifest) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeManifest) ManifestCalls(stub func(context.Context) ([]byte, string, error)) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = stub
}

func (fake *FakeManifest) ManifestArgsForCall(i int) context.Context {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	argsForCall := fake.manifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) ManifestReturns(result1 []byte, result2 string, result3 error) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 []byte
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeManifest) ManifestReturnsOnCall(i int, result1 []byte, result2 string, result3 error) {
	fake.manifestMutex.Lock()
	defer fake.manifestMutex.Unlock()
	fake.ManifestStub = nil
	if fake.manifestReturnsOnCall == nil {
		fake.manifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 string
			result3 error
		})
	}
	fake.manifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeManifest) OCIConfig(arg1 context.Context) (*v1.Image, error) {
	fake.oCIConfigMutex.Lock()
	ret, specificReturn := fake.oCIConfigReturnsOnCall[len(fake.oCIConfigArgsForCall)]
	fake.oCIConfigArgsForCall = append(fake.oCIConfigArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("OCIConfig", []interface{}{arg1})
	fake.oCIConfigMutex.Unlock()
	if fake.OCIConfigStub != nil {
		return fake.OCIConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.oCIConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) OCIConfigCallCount() int {
	fake.oCIConfigMutex.RLock()
	defer fake.oCIConfigMutex.RUnlock()
	return len(fake.oCIConfigArgsForCall)
}

func (fake *FakeManifest) OCIConfigCalls(stub func(context.Context) (*v1.Image, error)) {
	fake.oCIConfigMutex.Lock()
	defer fake.oCIConfigMutex.Unlock()
	fake.OCIConfigStub = stub
}

func (fake *FakeManifest) OCIConfigArgsForCall(i int) context.Context {
	fake.oCIConfigMutex.RLock()
	defer fake.oCIConfigMutex.RUnlock()
	argsForCall := fake.oCIConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) OCIConfigReturns(result1 *v1.Image, result2 error) {
	fake.oCIConfigMutex.Lock()
	defer fake.oCIConfigMutex.Unlock()
	fake.OCIConfigStub = nil
	fake.oCIConfigReturns = struct {
		result1 *v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) OCIConfigReturnsOnCall(i int, result1 *v1.Image, result2 error) {
	fake.oCIConfigMutex.Lock()
	defer fake.oCIConfigMutex.Unlock()
	fake.OCIConfigStub = nil
	if fake.oCIConfigReturnsOnCall == nil {
		fake.oCIConfigReturnsOnCall = make(map[int]struct {
			result1 *v1.Image
			result2 error
		})
	}
	fake.oCIConfigReturnsOnCall[i] = struct {
		result1 *v1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) Reference() types.ImageReference {
	fake.referenceMutex.Lock()
	ret, specificReturn := fake.referenceReturnsOnCall[len(fake.referenceArgsForCall)]
	fake.referenceArgsForCall = append(fake.referenceArgsForCall, struct {
	}{})
	fake.recordInvocation("Reference", []interface{}{})
	fake.referenceMutex.Unlock()
	if fake.ReferenceStub != nil {
		return fake.ReferenceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.referenceReturns
	return fakeReturns.result1
}

func (fake *FakeManifest) ReferenceCallCount() int {
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	return len(fake.referenceArgsForCall)
}

func (fake *FakeManifest) ReferenceCalls(stub func() types.ImageReference) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = stub
}

func (fake *FakeManifest) ReferenceReturns(result1 types.ImageReference) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = nil
	fake.referenceReturns = struct {
		result1 types.ImageReference
	}{result1}
}

func (fake *FakeManifest) ReferenceReturnsOnCall(i int, result1 types.ImageReference) {
	fake.referenceMutex.Lock()
	defer fake.referenceMutex.Unlock()
	fake.ReferenceStub = nil
	if fake.referenceReturnsOnCall == nil {
		fake.referenceReturnsOnCall = make(map[int]struct {
			result1 types.ImageReference
		})
	}
	fake.referenceReturnsOnCall[i] = struct {
		result1 types.ImageReference
	}{result1}
}

func (fake *FakeManifest) Signatures(arg1 context.Context) ([][]byte, error) {
	fake.signaturesMutex.Lock()
	ret, specificReturn := fake.signaturesReturnsOnCall[len(fake.signaturesArgsForCall)]
	fake.signaturesArgsForCall = append(fake.signaturesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Signatures", []interface{}{arg1})
	fake.signaturesMutex.Unlock()
	if fake.SignaturesStub != nil {
		return fake.SignaturesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.signaturesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) SignaturesCallCount() int {
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	return len(fake.signaturesArgsForCall)
}

func (fake *FakeManifest) SignaturesCalls(stub func(context.Context) ([][]byte, error)) {
	fake.signaturesMutex.Lock()
	defer fake.signaturesMutex.Unlock()
	fake.SignaturesStub = stub
}

func (fake *FakeManifest) SignaturesArgsForCall(i int) context.Context {
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	argsForCall := fake.signaturesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) SignaturesReturns(result1 [][]byte, result2 error) {
	fake.signaturesMutex.Lock()
	defer fake.signaturesMutex.Unlock()
	fake.SignaturesStub = nil
	fake.signaturesReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) SignaturesReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.signaturesMutex.Lock()
	defer fake.signaturesMutex.Unlock()
	fake.SignaturesStub = nil
	if fake.signaturesReturnsOnCall == nil {
		fake.signaturesReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.signaturesReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) Size() (int64, error) {
	fake.sizeMutex.Lock()
	ret, specificReturn := fake.sizeReturnsOnCall[len(fake.sizeArgsForCall)]
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct {
	}{})
	fake.recordInvocation("Size", []interface{}{})
	fake.sizeMutex.Unlock()
	if fake.SizeStub != nil {
		return fake.SizeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.sizeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakeManifest) SizeCalls(stub func() (int64, error)) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = stub
}

func (fake *FakeManifest) SizeReturns(result1 int64, result2 error) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) SizeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.sizeMutex.Lock()
	defer fake.sizeMutex.Unlock()
	fake.SizeStub = nil
	if fake.sizeReturnsOnCall == nil {
		fake.sizeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.sizeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) UpdatedImage(arg1 context.Context, arg2 types.ManifestUpdateOptions) (types.Image, error) {
	fake.updatedImageMutex.Lock()
	ret, specificReturn := fake.updatedImageReturnsOnCall[len(fake.updatedImageArgsForCall)]
	fake.updatedImageArgsForCall = append(fake.updatedImageArgsForCall, struct {
		arg1 context.Context
		arg2 types.ManifestUpdateOptions
	}{arg1, arg2})
	fake.recordInvocation("UpdatedImage", []interface{}{arg1, arg2})
	fake.updatedImageMutex.Unlock()
	if fake.UpdatedImageStub != nil {
		return fake.UpdatedImageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updatedImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifest) UpdatedImageCallCount() int {
	fake.updatedImageMutex.RLock()
	defer fake.updatedImageMutex.RUnlock()
	return len(fake.updatedImageArgsForCall)
}

func (fake *FakeManifest) UpdatedImageCalls(stub func(context.Context, types.ManifestUpdateOptions) (types.Image, error)) {
	fake.updatedImageMutex.Lock()
	defer fake.updatedImageMutex.Unlock()
	fake.UpdatedImageStub = stub
}

func (fake *FakeManifest) UpdatedImageArgsForCall(i int) (context.Context, types.ManifestUpdateOptions) {
	fake.updatedImageMutex.RLock()
	defer fake.updatedImageMutex.RUnlock()
	argsForCall := fake.updatedImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeManifest) UpdatedImageReturns(result1 types.Image, result2 error) {
	fake.updatedImageMutex.Lock()
	defer fake.updatedImageMutex.Unlock()
	fake.UpdatedImageStub = nil
	fake.updatedImageReturns = struct {
		result1 types.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) UpdatedImageReturnsOnCall(i int, result1 types.Image, result2 error) {
	fake.updatedImageMutex.Lock()
	defer fake.updatedImageMutex.Unlock()
	fake.UpdatedImageStub = nil
	if fake.updatedImageReturnsOnCall == nil {
		fake.updatedImageReturnsOnCall = make(map[int]struct {
			result1 types.Image
			result2 error
		})
	}
	fake.updatedImageReturnsOnCall[i] = struct {
		result1 types.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDs(arg1 types.ManifestUpdateOptions) bool {
	fake.updatedImageNeedsLayerDiffIDsMutex.Lock()
	ret, specificReturn := fake.updatedImageNeedsLayerDiffIDsReturnsOnCall[len(fake.updatedImageNeedsLayerDiffIDsArgsForCall)]
	fake.updatedImageNeedsLayerDiffIDsArgsForCall = append(fake.updatedImageNeedsLayerDiffIDsArgsForCall, struct {
		arg1 types.ManifestUpdateOptions
	}{arg1})
	fake.recordInvocation("UpdatedImageNeedsLayerDiffIDs", []interface{}{arg1})
	fake.updatedImageNeedsLayerDiffIDsMutex.Unlock()
	if fake.UpdatedImageNeedsLayerDiffIDsStub != nil {
		return fake.UpdatedImageNeedsLayerDiffIDsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updatedImageNeedsLayerDiffIDsReturns
	return fakeReturns.result1
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsCallCount() int {
	fake.updatedImageNeedsLayerDiffIDsMutex.RLock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.RUnlock()
	return len(fake.updatedImageNeedsLayerDiffIDsArgsForCall)
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsCalls(stub func(types.ManifestUpdateOptions) bool) {
	fake.updatedImageNeedsLayerDiffIDsMutex.Lock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.Unlock()
	fake.UpdatedImageNeedsLayerDiffIDsStub = stub
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsArgsForCall(i int) types.ManifestUpdateOptions {
	fake.updatedImageNeedsLayerDiffIDsMutex.RLock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.RUnlock()
	argsForCall := fake.updatedImageNeedsLayerDiffIDsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsReturns(result1 bool) {
	fake.updatedImageNeedsLayerDiffIDsMutex.Lock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.Unlock()
	fake.UpdatedImageNeedsLayerDiffIDsStub = nil
	fake.updatedImageNeedsLayerDiffIDsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) UpdatedImageNeedsLayerDiffIDsReturnsOnCall(i int, result1 bool) {
	fake.updatedImageNeedsLayerDiffIDsMutex.Lock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.Unlock()
	fake.UpdatedImageNeedsLayerDiffIDsStub = nil
	if fake.updatedImageNeedsLayerDiffIDsReturnsOnCall == nil {
		fake.updatedImageNeedsLayerDiffIDsReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.updatedImageNeedsLayerDiffIDsReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManifest) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configBlobMutex.RLock()
	defer fake.configBlobMutex.RUnlock()
	fake.configInfoMutex.RLock()
	defer fake.configInfoMutex.RUnlock()
	fake.embeddedDockerReferenceConflictsMutex.RLock()
	defer fake.embeddedDockerReferenceConflictsMutex.RUnlock()
	fake.inspectMutex.RLock()
	defer fake.inspectMutex.RUnlock()
	fake.layerInfosMutex.RLock()
	defer fake.layerInfosMutex.RUnlock()
	fake.layerInfosForCopyMutex.RLock()
	defer fake.layerInfosForCopyMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.oCIConfigMutex.RLock()
	defer fake.oCIConfigMutex.RUnlock()
	fake.referenceMutex.RLock()
	defer fake.referenceMutex.RUnlock()
	fake.signaturesMutex.RLock()
	defer fake.signaturesMutex.RUnlock()
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	fake.updatedImageMutex.RLock()
	defer fake.updatedImageMutex.RUnlock()
	fake.updatedImageNeedsLayerDiffIDsMutex.RLock()
	defer fake.updatedImageNeedsLayerDiffIDsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifest) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ layer_fetcher.Manifest = new(FakeManifest)
