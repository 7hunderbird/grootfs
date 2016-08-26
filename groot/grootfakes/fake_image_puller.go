// This file was generated by counterfeiter
package grootfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/lager"
)

type FakeImagePuller struct {
	PullStub        func(logger lager.Logger, spec groot.ImageSpec) (groot.BundleSpec, error)
	pullMutex       sync.RWMutex
	pullArgsForCall []struct {
		logger lager.Logger
		spec   groot.ImageSpec
	}
	pullReturns struct {
		result1 groot.BundleSpec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImagePuller) Pull(logger lager.Logger, spec groot.ImageSpec) (groot.BundleSpec, error) {
	fake.pullMutex.Lock()
	fake.pullArgsForCall = append(fake.pullArgsForCall, struct {
		logger lager.Logger
		spec   groot.ImageSpec
	}{logger, spec})
	fake.recordInvocation("Pull", []interface{}{logger, spec})
	fake.pullMutex.Unlock()
	if fake.PullStub != nil {
		return fake.PullStub(logger, spec)
	} else {
		return fake.pullReturns.result1, fake.pullReturns.result2
	}
}

func (fake *FakeImagePuller) PullCallCount() int {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return len(fake.pullArgsForCall)
}

func (fake *FakeImagePuller) PullArgsForCall(i int) (lager.Logger, groot.ImageSpec) {
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return fake.pullArgsForCall[i].logger, fake.pullArgsForCall[i].spec
}

func (fake *FakeImagePuller) PullReturns(result1 groot.BundleSpec, result2 error) {
	fake.PullStub = nil
	fake.pullReturns = struct {
		result1 groot.BundleSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeImagePuller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pullMutex.RLock()
	defer fake.pullMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeImagePuller) recordInvocation(key string, args []interface{}) {
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

var _ groot.ImagePuller = new(FakeImagePuller)