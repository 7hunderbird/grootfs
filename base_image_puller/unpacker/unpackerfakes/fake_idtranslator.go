// Code generated by counterfeiter. DO NOT EDIT.
package unpackerfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/base_image_puller/unpacker"
)

type FakeIDTranslator struct {
	TranslateUIDStub        func(id int) int
	translateUIDMutex       sync.RWMutex
	translateUIDArgsForCall []struct {
		id int
	}
	translateUIDReturns struct {
		result1 int
	}
	translateUIDReturnsOnCall map[int]struct {
		result1 int
	}
	TranslateGIDStub        func(id int) int
	translateGIDMutex       sync.RWMutex
	translateGIDArgsForCall []struct {
		id int
	}
	translateGIDReturns struct {
		result1 int
	}
	translateGIDReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIDTranslator) TranslateUID(id int) int {
	fake.translateUIDMutex.Lock()
	ret, specificReturn := fake.translateUIDReturnsOnCall[len(fake.translateUIDArgsForCall)]
	fake.translateUIDArgsForCall = append(fake.translateUIDArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("TranslateUID", []interface{}{id})
	fake.translateUIDMutex.Unlock()
	if fake.TranslateUIDStub != nil {
		return fake.TranslateUIDStub(id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.translateUIDReturns.result1
}

func (fake *FakeIDTranslator) TranslateUIDCallCount() int {
	fake.translateUIDMutex.RLock()
	defer fake.translateUIDMutex.RUnlock()
	return len(fake.translateUIDArgsForCall)
}

func (fake *FakeIDTranslator) TranslateUIDArgsForCall(i int) int {
	fake.translateUIDMutex.RLock()
	defer fake.translateUIDMutex.RUnlock()
	return fake.translateUIDArgsForCall[i].id
}

func (fake *FakeIDTranslator) TranslateUIDReturns(result1 int) {
	fake.TranslateUIDStub = nil
	fake.translateUIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeIDTranslator) TranslateUIDReturnsOnCall(i int, result1 int) {
	fake.TranslateUIDStub = nil
	if fake.translateUIDReturnsOnCall == nil {
		fake.translateUIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.translateUIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeIDTranslator) TranslateGID(id int) int {
	fake.translateGIDMutex.Lock()
	ret, specificReturn := fake.translateGIDReturnsOnCall[len(fake.translateGIDArgsForCall)]
	fake.translateGIDArgsForCall = append(fake.translateGIDArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("TranslateGID", []interface{}{id})
	fake.translateGIDMutex.Unlock()
	if fake.TranslateGIDStub != nil {
		return fake.TranslateGIDStub(id)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.translateGIDReturns.result1
}

func (fake *FakeIDTranslator) TranslateGIDCallCount() int {
	fake.translateGIDMutex.RLock()
	defer fake.translateGIDMutex.RUnlock()
	return len(fake.translateGIDArgsForCall)
}

func (fake *FakeIDTranslator) TranslateGIDArgsForCall(i int) int {
	fake.translateGIDMutex.RLock()
	defer fake.translateGIDMutex.RUnlock()
	return fake.translateGIDArgsForCall[i].id
}

func (fake *FakeIDTranslator) TranslateGIDReturns(result1 int) {
	fake.TranslateGIDStub = nil
	fake.translateGIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeIDTranslator) TranslateGIDReturnsOnCall(i int, result1 int) {
	fake.TranslateGIDStub = nil
	if fake.translateGIDReturnsOnCall == nil {
		fake.translateGIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.translateGIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeIDTranslator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.translateUIDMutex.RLock()
	defer fake.translateUIDMutex.RUnlock()
	fake.translateGIDMutex.RLock()
	defer fake.translateGIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIDTranslator) recordInvocation(key string, args []interface{}) {
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

var _ unpacker.IDTranslator = new(FakeIDTranslator)