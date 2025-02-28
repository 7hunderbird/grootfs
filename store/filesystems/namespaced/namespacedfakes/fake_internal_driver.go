// Code generated by counterfeiter. DO NOT EDIT.
package namespacedfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/base_image_puller"
	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/grootfs/store/image_cloner"
	"code.cloudfoundry.org/lager"
)

type FakeInternalDriver struct {
	CreateImageStub        func(lager.Logger, image_cloner.ImageDriverSpec) (groot.MountInfo, error)
	createImageMutex       sync.RWMutex
	createImageArgsForCall []struct {
		arg1 lager.Logger
		arg2 image_cloner.ImageDriverSpec
	}
	createImageReturns struct {
		result1 groot.MountInfo
		result2 error
	}
	createImageReturnsOnCall map[int]struct {
		result1 groot.MountInfo
		result2 error
	}
	CreateVolumeStub        func(lager.Logger, string, string) (string, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	createVolumeReturns struct {
		result1 string
		result2 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DestroyImageStub        func(lager.Logger, string) error
	destroyImageMutex       sync.RWMutex
	destroyImageArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	destroyImageReturns struct {
		result1 error
	}
	destroyImageReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyVolumeStub        func(lager.Logger, string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	destroyVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	FetchStatsStub        func(lager.Logger, string) (groot.VolumeStats, error)
	fetchStatsMutex       sync.RWMutex
	fetchStatsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	fetchStatsReturns struct {
		result1 groot.VolumeStats
		result2 error
	}
	fetchStatsReturnsOnCall map[int]struct {
		result1 groot.VolumeStats
		result2 error
	}
	HandleOpaqueWhiteoutsStub        func(lager.Logger, string, []string) error
	handleOpaqueWhiteoutsMutex       sync.RWMutex
	handleOpaqueWhiteoutsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 []string
	}
	handleOpaqueWhiteoutsReturns struct {
		result1 error
	}
	handleOpaqueWhiteoutsReturnsOnCall map[int]struct {
		result1 error
	}
	MarkVolumeArtifactsStub        func(lager.Logger, string) error
	markVolumeArtifactsMutex       sync.RWMutex
	markVolumeArtifactsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	markVolumeArtifactsReturns struct {
		result1 error
	}
	markVolumeArtifactsReturnsOnCall map[int]struct {
		result1 error
	}
	MarshalStub        func(lager.Logger) ([]byte, error)
	marshalMutex       sync.RWMutex
	marshalArgsForCall []struct {
		arg1 lager.Logger
	}
	marshalReturns struct {
		result1 []byte
		result2 error
	}
	marshalReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	MoveVolumeStub        func(lager.Logger, string, string) error
	moveVolumeMutex       sync.RWMutex
	moveVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	moveVolumeReturns struct {
		result1 error
	}
	moveVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	VolumePathStub        func(lager.Logger, string) (string, error)
	volumePathMutex       sync.RWMutex
	volumePathArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	volumePathReturns struct {
		result1 string
		result2 error
	}
	volumePathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	VolumesStub        func(lager.Logger) ([]string, error)
	volumesMutex       sync.RWMutex
	volumesArgsForCall []struct {
		arg1 lager.Logger
	}
	volumesReturns struct {
		result1 []string
		result2 error
	}
	volumesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	WriteVolumeMetaStub        func(lager.Logger, string, base_image_puller.VolumeMeta) error
	writeVolumeMetaMutex       sync.RWMutex
	writeVolumeMetaArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 base_image_puller.VolumeMeta
	}
	writeVolumeMetaReturns struct {
		result1 error
	}
	writeVolumeMetaReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInternalDriver) CreateImage(arg1 lager.Logger, arg2 image_cloner.ImageDriverSpec) (groot.MountInfo, error) {
	fake.createImageMutex.Lock()
	ret, specificReturn := fake.createImageReturnsOnCall[len(fake.createImageArgsForCall)]
	fake.createImageArgsForCall = append(fake.createImageArgsForCall, struct {
		arg1 lager.Logger
		arg2 image_cloner.ImageDriverSpec
	}{arg1, arg2})
	fake.recordInvocation("CreateImage", []interface{}{arg1, arg2})
	fake.createImageMutex.Unlock()
	if fake.CreateImageStub != nil {
		return fake.CreateImageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInternalDriver) CreateImageCallCount() int {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	return len(fake.createImageArgsForCall)
}

func (fake *FakeInternalDriver) CreateImageCalls(stub func(lager.Logger, image_cloner.ImageDriverSpec) (groot.MountInfo, error)) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = stub
}

func (fake *FakeInternalDriver) CreateImageArgsForCall(i int) (lager.Logger, image_cloner.ImageDriverSpec) {
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	argsForCall := fake.createImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInternalDriver) CreateImageReturns(result1 groot.MountInfo, result2 error) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = nil
	fake.createImageReturns = struct {
		result1 groot.MountInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) CreateImageReturnsOnCall(i int, result1 groot.MountInfo, result2 error) {
	fake.createImageMutex.Lock()
	defer fake.createImageMutex.Unlock()
	fake.CreateImageStub = nil
	if fake.createImageReturnsOnCall == nil {
		fake.createImageReturnsOnCall = make(map[int]struct {
			result1 groot.MountInfo
			result2 error
		})
	}
	fake.createImageReturnsOnCall[i] = struct {
		result1 groot.MountInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) CreateVolume(arg1 lager.Logger, arg2 string, arg3 string) (string, error) {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateVolume", []interface{}{arg1, arg2, arg3})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInternalDriver) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeInternalDriver) CreateVolumeCalls(stub func(lager.Logger, string, string) (string, error)) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = stub
}

func (fake *FakeInternalDriver) CreateVolumeArgsForCall(i int) (lager.Logger, string, string) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	argsForCall := fake.createVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInternalDriver) CreateVolumeReturns(result1 string, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) CreateVolumeReturnsOnCall(i int, result1 string, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) DestroyImage(arg1 lager.Logger, arg2 string) error {
	fake.destroyImageMutex.Lock()
	ret, specificReturn := fake.destroyImageReturnsOnCall[len(fake.destroyImageArgsForCall)]
	fake.destroyImageArgsForCall = append(fake.destroyImageArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DestroyImage", []interface{}{arg1, arg2})
	fake.destroyImageMutex.Unlock()
	if fake.DestroyImageStub != nil {
		return fake.DestroyImageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyImageReturns
	return fakeReturns.result1
}

func (fake *FakeInternalDriver) DestroyImageCallCount() int {
	fake.destroyImageMutex.RLock()
	defer fake.destroyImageMutex.RUnlock()
	return len(fake.destroyImageArgsForCall)
}

func (fake *FakeInternalDriver) DestroyImageCalls(stub func(lager.Logger, string) error) {
	fake.destroyImageMutex.Lock()
	defer fake.destroyImageMutex.Unlock()
	fake.DestroyImageStub = stub
}

func (fake *FakeInternalDriver) DestroyImageArgsForCall(i int) (lager.Logger, string) {
	fake.destroyImageMutex.RLock()
	defer fake.destroyImageMutex.RUnlock()
	argsForCall := fake.destroyImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInternalDriver) DestroyImageReturns(result1 error) {
	fake.destroyImageMutex.Lock()
	defer fake.destroyImageMutex.Unlock()
	fake.DestroyImageStub = nil
	fake.destroyImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) DestroyImageReturnsOnCall(i int, result1 error) {
	fake.destroyImageMutex.Lock()
	defer fake.destroyImageMutex.Unlock()
	fake.DestroyImageStub = nil
	if fake.destroyImageReturnsOnCall == nil {
		fake.destroyImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) DestroyVolume(arg1 lager.Logger, arg2 string) error {
	fake.destroyVolumeMutex.Lock()
	ret, specificReturn := fake.destroyVolumeReturnsOnCall[len(fake.destroyVolumeArgsForCall)]
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DestroyVolume", []interface{}{arg1, arg2})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyVolumeReturns
	return fakeReturns.result1
}

func (fake *FakeInternalDriver) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeInternalDriver) DestroyVolumeCalls(stub func(lager.Logger, string) error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = stub
}

func (fake *FakeInternalDriver) DestroyVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	argsForCall := fake.destroyVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInternalDriver) DestroyVolumeReturns(result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) DestroyVolumeReturnsOnCall(i int, result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	if fake.destroyVolumeReturnsOnCall == nil {
		fake.destroyVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) FetchStats(arg1 lager.Logger, arg2 string) (groot.VolumeStats, error) {
	fake.fetchStatsMutex.Lock()
	ret, specificReturn := fake.fetchStatsReturnsOnCall[len(fake.fetchStatsArgsForCall)]
	fake.fetchStatsArgsForCall = append(fake.fetchStatsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("FetchStats", []interface{}{arg1, arg2})
	fake.fetchStatsMutex.Unlock()
	if fake.FetchStatsStub != nil {
		return fake.FetchStatsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchStatsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInternalDriver) FetchStatsCallCount() int {
	fake.fetchStatsMutex.RLock()
	defer fake.fetchStatsMutex.RUnlock()
	return len(fake.fetchStatsArgsForCall)
}

func (fake *FakeInternalDriver) FetchStatsCalls(stub func(lager.Logger, string) (groot.VolumeStats, error)) {
	fake.fetchStatsMutex.Lock()
	defer fake.fetchStatsMutex.Unlock()
	fake.FetchStatsStub = stub
}

func (fake *FakeInternalDriver) FetchStatsArgsForCall(i int) (lager.Logger, string) {
	fake.fetchStatsMutex.RLock()
	defer fake.fetchStatsMutex.RUnlock()
	argsForCall := fake.fetchStatsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInternalDriver) FetchStatsReturns(result1 groot.VolumeStats, result2 error) {
	fake.fetchStatsMutex.Lock()
	defer fake.fetchStatsMutex.Unlock()
	fake.FetchStatsStub = nil
	fake.fetchStatsReturns = struct {
		result1 groot.VolumeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) FetchStatsReturnsOnCall(i int, result1 groot.VolumeStats, result2 error) {
	fake.fetchStatsMutex.Lock()
	defer fake.fetchStatsMutex.Unlock()
	fake.FetchStatsStub = nil
	if fake.fetchStatsReturnsOnCall == nil {
		fake.fetchStatsReturnsOnCall = make(map[int]struct {
			result1 groot.VolumeStats
			result2 error
		})
	}
	fake.fetchStatsReturnsOnCall[i] = struct {
		result1 groot.VolumeStats
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) HandleOpaqueWhiteouts(arg1 lager.Logger, arg2 string, arg3 []string) error {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.handleOpaqueWhiteoutsMutex.Lock()
	ret, specificReturn := fake.handleOpaqueWhiteoutsReturnsOnCall[len(fake.handleOpaqueWhiteoutsArgsForCall)]
	fake.handleOpaqueWhiteoutsArgsForCall = append(fake.handleOpaqueWhiteoutsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("HandleOpaqueWhiteouts", []interface{}{arg1, arg2, arg3Copy})
	fake.handleOpaqueWhiteoutsMutex.Unlock()
	if fake.HandleOpaqueWhiteoutsStub != nil {
		return fake.HandleOpaqueWhiteoutsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleOpaqueWhiteoutsReturns
	return fakeReturns.result1
}

func (fake *FakeInternalDriver) HandleOpaqueWhiteoutsCallCount() int {
	fake.handleOpaqueWhiteoutsMutex.RLock()
	defer fake.handleOpaqueWhiteoutsMutex.RUnlock()
	return len(fake.handleOpaqueWhiteoutsArgsForCall)
}

func (fake *FakeInternalDriver) HandleOpaqueWhiteoutsCalls(stub func(lager.Logger, string, []string) error) {
	fake.handleOpaqueWhiteoutsMutex.Lock()
	defer fake.handleOpaqueWhiteoutsMutex.Unlock()
	fake.HandleOpaqueWhiteoutsStub = stub
}

func (fake *FakeInternalDriver) HandleOpaqueWhiteoutsArgsForCall(i int) (lager.Logger, string, []string) {
	fake.handleOpaqueWhiteoutsMutex.RLock()
	defer fake.handleOpaqueWhiteoutsMutex.RUnlock()
	argsForCall := fake.handleOpaqueWhiteoutsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInternalDriver) HandleOpaqueWhiteoutsReturns(result1 error) {
	fake.handleOpaqueWhiteoutsMutex.Lock()
	defer fake.handleOpaqueWhiteoutsMutex.Unlock()
	fake.HandleOpaqueWhiteoutsStub = nil
	fake.handleOpaqueWhiteoutsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) HandleOpaqueWhiteoutsReturnsOnCall(i int, result1 error) {
	fake.handleOpaqueWhiteoutsMutex.Lock()
	defer fake.handleOpaqueWhiteoutsMutex.Unlock()
	fake.HandleOpaqueWhiteoutsStub = nil
	if fake.handleOpaqueWhiteoutsReturnsOnCall == nil {
		fake.handleOpaqueWhiteoutsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.handleOpaqueWhiteoutsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) MarkVolumeArtifacts(arg1 lager.Logger, arg2 string) error {
	fake.markVolumeArtifactsMutex.Lock()
	ret, specificReturn := fake.markVolumeArtifactsReturnsOnCall[len(fake.markVolumeArtifactsArgsForCall)]
	fake.markVolumeArtifactsArgsForCall = append(fake.markVolumeArtifactsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("MarkVolumeArtifacts", []interface{}{arg1, arg2})
	fake.markVolumeArtifactsMutex.Unlock()
	if fake.MarkVolumeArtifactsStub != nil {
		return fake.MarkVolumeArtifactsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.markVolumeArtifactsReturns
	return fakeReturns.result1
}

func (fake *FakeInternalDriver) MarkVolumeArtifactsCallCount() int {
	fake.markVolumeArtifactsMutex.RLock()
	defer fake.markVolumeArtifactsMutex.RUnlock()
	return len(fake.markVolumeArtifactsArgsForCall)
}

func (fake *FakeInternalDriver) MarkVolumeArtifactsCalls(stub func(lager.Logger, string) error) {
	fake.markVolumeArtifactsMutex.Lock()
	defer fake.markVolumeArtifactsMutex.Unlock()
	fake.MarkVolumeArtifactsStub = stub
}

func (fake *FakeInternalDriver) MarkVolumeArtifactsArgsForCall(i int) (lager.Logger, string) {
	fake.markVolumeArtifactsMutex.RLock()
	defer fake.markVolumeArtifactsMutex.RUnlock()
	argsForCall := fake.markVolumeArtifactsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInternalDriver) MarkVolumeArtifactsReturns(result1 error) {
	fake.markVolumeArtifactsMutex.Lock()
	defer fake.markVolumeArtifactsMutex.Unlock()
	fake.MarkVolumeArtifactsStub = nil
	fake.markVolumeArtifactsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) MarkVolumeArtifactsReturnsOnCall(i int, result1 error) {
	fake.markVolumeArtifactsMutex.Lock()
	defer fake.markVolumeArtifactsMutex.Unlock()
	fake.MarkVolumeArtifactsStub = nil
	if fake.markVolumeArtifactsReturnsOnCall == nil {
		fake.markVolumeArtifactsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markVolumeArtifactsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) Marshal(arg1 lager.Logger) ([]byte, error) {
	fake.marshalMutex.Lock()
	ret, specificReturn := fake.marshalReturnsOnCall[len(fake.marshalArgsForCall)]
	fake.marshalArgsForCall = append(fake.marshalArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Marshal", []interface{}{arg1})
	fake.marshalMutex.Unlock()
	if fake.MarshalStub != nil {
		return fake.MarshalStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.marshalReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInternalDriver) MarshalCallCount() int {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return len(fake.marshalArgsForCall)
}

func (fake *FakeInternalDriver) MarshalCalls(stub func(lager.Logger) ([]byte, error)) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.MarshalStub = stub
}

func (fake *FakeInternalDriver) MarshalArgsForCall(i int) lager.Logger {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	argsForCall := fake.marshalArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInternalDriver) MarshalReturns(result1 []byte, result2 error) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.MarshalStub = nil
	fake.marshalReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) MarshalReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.MarshalStub = nil
	if fake.marshalReturnsOnCall == nil {
		fake.marshalReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) MoveVolume(arg1 lager.Logger, arg2 string, arg3 string) error {
	fake.moveVolumeMutex.Lock()
	ret, specificReturn := fake.moveVolumeReturnsOnCall[len(fake.moveVolumeArgsForCall)]
	fake.moveVolumeArgsForCall = append(fake.moveVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("MoveVolume", []interface{}{arg1, arg2, arg3})
	fake.moveVolumeMutex.Unlock()
	if fake.MoveVolumeStub != nil {
		return fake.MoveVolumeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.moveVolumeReturns
	return fakeReturns.result1
}

func (fake *FakeInternalDriver) MoveVolumeCallCount() int {
	fake.moveVolumeMutex.RLock()
	defer fake.moveVolumeMutex.RUnlock()
	return len(fake.moveVolumeArgsForCall)
}

func (fake *FakeInternalDriver) MoveVolumeCalls(stub func(lager.Logger, string, string) error) {
	fake.moveVolumeMutex.Lock()
	defer fake.moveVolumeMutex.Unlock()
	fake.MoveVolumeStub = stub
}

func (fake *FakeInternalDriver) MoveVolumeArgsForCall(i int) (lager.Logger, string, string) {
	fake.moveVolumeMutex.RLock()
	defer fake.moveVolumeMutex.RUnlock()
	argsForCall := fake.moveVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInternalDriver) MoveVolumeReturns(result1 error) {
	fake.moveVolumeMutex.Lock()
	defer fake.moveVolumeMutex.Unlock()
	fake.MoveVolumeStub = nil
	fake.moveVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) MoveVolumeReturnsOnCall(i int, result1 error) {
	fake.moveVolumeMutex.Lock()
	defer fake.moveVolumeMutex.Unlock()
	fake.MoveVolumeStub = nil
	if fake.moveVolumeReturnsOnCall == nil {
		fake.moveVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.moveVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) VolumePath(arg1 lager.Logger, arg2 string) (string, error) {
	fake.volumePathMutex.Lock()
	ret, specificReturn := fake.volumePathReturnsOnCall[len(fake.volumePathArgsForCall)]
	fake.volumePathArgsForCall = append(fake.volumePathArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("VolumePath", []interface{}{arg1, arg2})
	fake.volumePathMutex.Unlock()
	if fake.VolumePathStub != nil {
		return fake.VolumePathStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.volumePathReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInternalDriver) VolumePathCallCount() int {
	fake.volumePathMutex.RLock()
	defer fake.volumePathMutex.RUnlock()
	return len(fake.volumePathArgsForCall)
}

func (fake *FakeInternalDriver) VolumePathCalls(stub func(lager.Logger, string) (string, error)) {
	fake.volumePathMutex.Lock()
	defer fake.volumePathMutex.Unlock()
	fake.VolumePathStub = stub
}

func (fake *FakeInternalDriver) VolumePathArgsForCall(i int) (lager.Logger, string) {
	fake.volumePathMutex.RLock()
	defer fake.volumePathMutex.RUnlock()
	argsForCall := fake.volumePathArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInternalDriver) VolumePathReturns(result1 string, result2 error) {
	fake.volumePathMutex.Lock()
	defer fake.volumePathMutex.Unlock()
	fake.VolumePathStub = nil
	fake.volumePathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) VolumePathReturnsOnCall(i int, result1 string, result2 error) {
	fake.volumePathMutex.Lock()
	defer fake.volumePathMutex.Unlock()
	fake.VolumePathStub = nil
	if fake.volumePathReturnsOnCall == nil {
		fake.volumePathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.volumePathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) Volumes(arg1 lager.Logger) ([]string, error) {
	fake.volumesMutex.Lock()
	ret, specificReturn := fake.volumesReturnsOnCall[len(fake.volumesArgsForCall)]
	fake.volumesArgsForCall = append(fake.volumesArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Volumes", []interface{}{arg1})
	fake.volumesMutex.Unlock()
	if fake.VolumesStub != nil {
		return fake.VolumesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.volumesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInternalDriver) VolumesCallCount() int {
	fake.volumesMutex.RLock()
	defer fake.volumesMutex.RUnlock()
	return len(fake.volumesArgsForCall)
}

func (fake *FakeInternalDriver) VolumesCalls(stub func(lager.Logger) ([]string, error)) {
	fake.volumesMutex.Lock()
	defer fake.volumesMutex.Unlock()
	fake.VolumesStub = stub
}

func (fake *FakeInternalDriver) VolumesArgsForCall(i int) lager.Logger {
	fake.volumesMutex.RLock()
	defer fake.volumesMutex.RUnlock()
	argsForCall := fake.volumesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInternalDriver) VolumesReturns(result1 []string, result2 error) {
	fake.volumesMutex.Lock()
	defer fake.volumesMutex.Unlock()
	fake.VolumesStub = nil
	fake.volumesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) VolumesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.volumesMutex.Lock()
	defer fake.volumesMutex.Unlock()
	fake.VolumesStub = nil
	if fake.volumesReturnsOnCall == nil {
		fake.volumesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.volumesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeInternalDriver) WriteVolumeMeta(arg1 lager.Logger, arg2 string, arg3 base_image_puller.VolumeMeta) error {
	fake.writeVolumeMetaMutex.Lock()
	ret, specificReturn := fake.writeVolumeMetaReturnsOnCall[len(fake.writeVolumeMetaArgsForCall)]
	fake.writeVolumeMetaArgsForCall = append(fake.writeVolumeMetaArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 base_image_puller.VolumeMeta
	}{arg1, arg2, arg3})
	fake.recordInvocation("WriteVolumeMeta", []interface{}{arg1, arg2, arg3})
	fake.writeVolumeMetaMutex.Unlock()
	if fake.WriteVolumeMetaStub != nil {
		return fake.WriteVolumeMetaStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.writeVolumeMetaReturns
	return fakeReturns.result1
}

func (fake *FakeInternalDriver) WriteVolumeMetaCallCount() int {
	fake.writeVolumeMetaMutex.RLock()
	defer fake.writeVolumeMetaMutex.RUnlock()
	return len(fake.writeVolumeMetaArgsForCall)
}

func (fake *FakeInternalDriver) WriteVolumeMetaCalls(stub func(lager.Logger, string, base_image_puller.VolumeMeta) error) {
	fake.writeVolumeMetaMutex.Lock()
	defer fake.writeVolumeMetaMutex.Unlock()
	fake.WriteVolumeMetaStub = stub
}

func (fake *FakeInternalDriver) WriteVolumeMetaArgsForCall(i int) (lager.Logger, string, base_image_puller.VolumeMeta) {
	fake.writeVolumeMetaMutex.RLock()
	defer fake.writeVolumeMetaMutex.RUnlock()
	argsForCall := fake.writeVolumeMetaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInternalDriver) WriteVolumeMetaReturns(result1 error) {
	fake.writeVolumeMetaMutex.Lock()
	defer fake.writeVolumeMetaMutex.Unlock()
	fake.WriteVolumeMetaStub = nil
	fake.writeVolumeMetaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) WriteVolumeMetaReturnsOnCall(i int, result1 error) {
	fake.writeVolumeMetaMutex.Lock()
	defer fake.writeVolumeMetaMutex.Unlock()
	fake.WriteVolumeMetaStub = nil
	if fake.writeVolumeMetaReturnsOnCall == nil {
		fake.writeVolumeMetaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeVolumeMetaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInternalDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createImageMutex.RLock()
	defer fake.createImageMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.destroyImageMutex.RLock()
	defer fake.destroyImageMutex.RUnlock()
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	fake.fetchStatsMutex.RLock()
	defer fake.fetchStatsMutex.RUnlock()
	fake.handleOpaqueWhiteoutsMutex.RLock()
	defer fake.handleOpaqueWhiteoutsMutex.RUnlock()
	fake.markVolumeArtifactsMutex.RLock()
	defer fake.markVolumeArtifactsMutex.RUnlock()
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	fake.moveVolumeMutex.RLock()
	defer fake.moveVolumeMutex.RUnlock()
	fake.volumePathMutex.RLock()
	defer fake.volumePathMutex.RUnlock()
	fake.volumesMutex.RLock()
	defer fake.volumesMutex.RUnlock()
	fake.writeVolumeMetaMutex.RLock()
	defer fake.writeVolumeMetaMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInternalDriver) recordInvocation(key string, args []interface{}) {
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
