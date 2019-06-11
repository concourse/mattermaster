// Code generated by counterfeiter. DO NOT EDIT.
package volumefakes

import (
	"sync"

	"github.com/concourse/baggageclaim/volume"
)

type FakeDriver struct {
	CreateCopyOnWriteLayerStub        func(string, string) error
	createCopyOnWriteLayerMutex       sync.RWMutex
	createCopyOnWriteLayerArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createCopyOnWriteLayerReturns struct {
		result1 error
	}
	createCopyOnWriteLayerReturnsOnCall map[int]struct {
		result1 error
	}
	CreateVolumeStub        func(string) error
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 string
	}
	createVolumeReturns struct {
		result1 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyVolumeStub        func(string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		arg1 string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	destroyVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDriver) CreateCopyOnWriteLayer(arg1 string, arg2 string) error {
	fake.createCopyOnWriteLayerMutex.Lock()
	ret, specificReturn := fake.createCopyOnWriteLayerReturnsOnCall[len(fake.createCopyOnWriteLayerArgsForCall)]
	fake.createCopyOnWriteLayerArgsForCall = append(fake.createCopyOnWriteLayerArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateCopyOnWriteLayer", []interface{}{arg1, arg2})
	fake.createCopyOnWriteLayerMutex.Unlock()
	if fake.CreateCopyOnWriteLayerStub != nil {
		return fake.CreateCopyOnWriteLayerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createCopyOnWriteLayerReturns
	return fakeReturns.result1
}

func (fake *FakeDriver) CreateCopyOnWriteLayerCallCount() int {
	fake.createCopyOnWriteLayerMutex.RLock()
	defer fake.createCopyOnWriteLayerMutex.RUnlock()
	return len(fake.createCopyOnWriteLayerArgsForCall)
}

func (fake *FakeDriver) CreateCopyOnWriteLayerCalls(stub func(string, string) error) {
	fake.createCopyOnWriteLayerMutex.Lock()
	defer fake.createCopyOnWriteLayerMutex.Unlock()
	fake.CreateCopyOnWriteLayerStub = stub
}

func (fake *FakeDriver) CreateCopyOnWriteLayerArgsForCall(i int) (string, string) {
	fake.createCopyOnWriteLayerMutex.RLock()
	defer fake.createCopyOnWriteLayerMutex.RUnlock()
	argsForCall := fake.createCopyOnWriteLayerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDriver) CreateCopyOnWriteLayerReturns(result1 error) {
	fake.createCopyOnWriteLayerMutex.Lock()
	defer fake.createCopyOnWriteLayerMutex.Unlock()
	fake.CreateCopyOnWriteLayerStub = nil
	fake.createCopyOnWriteLayerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) CreateCopyOnWriteLayerReturnsOnCall(i int, result1 error) {
	fake.createCopyOnWriteLayerMutex.Lock()
	defer fake.createCopyOnWriteLayerMutex.Unlock()
	fake.CreateCopyOnWriteLayerStub = nil
	if fake.createCopyOnWriteLayerReturnsOnCall == nil {
		fake.createCopyOnWriteLayerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createCopyOnWriteLayerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) CreateVolume(arg1 string) error {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CreateVolume", []interface{}{arg1})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createVolumeReturns
	return fakeReturns.result1
}

func (fake *FakeDriver) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeDriver) CreateVolumeCalls(stub func(string) error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = stub
}

func (fake *FakeDriver) CreateVolumeArgsForCall(i int) string {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	argsForCall := fake.createVolumeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDriver) CreateVolumeReturns(result1 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) CreateVolumeReturnsOnCall(i int, result1 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) DestroyVolume(arg1 string) error {
	fake.destroyVolumeMutex.Lock()
	ret, specificReturn := fake.destroyVolumeReturnsOnCall[len(fake.destroyVolumeArgsForCall)]
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DestroyVolume", []interface{}{arg1})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyVolumeReturns
	return fakeReturns.result1
}

func (fake *FakeDriver) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeDriver) DestroyVolumeCalls(stub func(string) error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = stub
}

func (fake *FakeDriver) DestroyVolumeArgsForCall(i int) string {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	argsForCall := fake.destroyVolumeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDriver) DestroyVolumeReturns(result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDriver) DestroyVolumeReturnsOnCall(i int, result1 error) {
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

func (fake *FakeDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCopyOnWriteLayerMutex.RLock()
	defer fake.createCopyOnWriteLayerMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDriver) recordInvocation(key string, args []interface{}) {
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

var _ volume.Driver = new(FakeDriver)
