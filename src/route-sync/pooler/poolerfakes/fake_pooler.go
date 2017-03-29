// This file was generated by counterfeiter
package poolerfakes

import (
	"context"
	"route-sync/pooler"
	"route-sync/route"
	"sync"
)

type FakePooler struct {
	RunStub        func(context.Context, route.Source, route.Router)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 route.Source
		arg3 route.Router
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePooler) Run(arg1 context.Context, arg2 route.Source, arg3 route.Router) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 route.Source
		arg3 route.Router
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub(arg1, arg2, arg3)
	}
}

func (fake *FakePooler) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakePooler) RunArgsForCall(i int) (context.Context, route.Source, route.Router) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2, fake.runArgsForCall[i].arg3
}

func (fake *FakePooler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePooler) recordInvocation(key string, args []interface{}) {
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

var _ pooler.Pooler = new(FakePooler)
