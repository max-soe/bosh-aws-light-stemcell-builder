// Code generated by counterfeiter. DO NOT EDIT.
package resourcesfakes

import (
	"light-stemcell-builder/resources"
	"sync"
)

type FakeAmiDriver struct {
	CreateStub        func(resources.AmiDriverConfig) (resources.Ami, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 resources.AmiDriverConfig
	}
	createReturns struct {
		result1 resources.Ami
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 resources.Ami
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAmiDriver) Create(arg1 resources.AmiDriverConfig) (resources.Ami, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 resources.AmiDriverConfig
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAmiDriver) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeAmiDriver) CreateCalls(stub func(resources.AmiDriverConfig) (resources.Ami, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeAmiDriver) CreateArgsForCall(i int) resources.AmiDriverConfig {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAmiDriver) CreateReturns(result1 resources.Ami, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 resources.Ami
		result2 error
	}{result1, result2}
}

func (fake *FakeAmiDriver) CreateReturnsOnCall(i int, result1 resources.Ami, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 resources.Ami
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 resources.Ami
		result2 error
	}{result1, result2}
}

func (fake *FakeAmiDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAmiDriver) recordInvocation(key string, args []interface{}) {
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

var _ resources.AmiDriver = new(FakeAmiDriver)
