// Code generated by counterfeiter. DO NOT EDIT.
package goasyncfakes

import (
	"context"
	"sync"

	goasync "github.com/DanLavine/goasync/v2"
)

type FakeExecuteTask struct {
	ExecuteStub        func(context.Context) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 context.Context
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecuteTask) Execute(arg1 context.Context) error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ExecuteStub
	fakeReturns := fake.executeReturns
	fake.recordInvocation("Execute", []interface{}{arg1})
	fake.executeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeExecuteTask) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeExecuteTask) ExecuteCalls(stub func(context.Context) error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeExecuteTask) ExecuteArgsForCall(i int) context.Context {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeExecuteTask) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecuteTask) ExecuteReturnsOnCall(i int, result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecuteTask) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecuteTask) recordInvocation(key string, args []interface{}) {
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

var _ goasync.ExecuteTask = new(FakeExecuteTask)
