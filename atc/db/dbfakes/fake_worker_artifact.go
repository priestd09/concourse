// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"

	db "github.com/concourse/concourse/atc/db"
)

type FakeWorkerArtifact struct {
	ChecksumStub        func() string
	checksumMutex       sync.RWMutex
	checksumArgsForCall []struct {
	}
	checksumReturns struct {
		result1 string
	}
	checksumReturnsOnCall map[int]struct {
		result1 string
	}
	CreatedAtStub        func() int
	createdAtMutex       sync.RWMutex
	createdAtArgsForCall []struct {
	}
	createdAtReturns struct {
		result1 int
	}
	createdAtReturnsOnCall map[int]struct {
		result1 int
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
	}
	pathReturns struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerArtifact) Checksum() string {
	fake.checksumMutex.Lock()
	ret, specificReturn := fake.checksumReturnsOnCall[len(fake.checksumArgsForCall)]
	fake.checksumArgsForCall = append(fake.checksumArgsForCall, struct {
	}{})
	fake.recordInvocation("Checksum", []interface{}{})
	fake.checksumMutex.Unlock()
	if fake.ChecksumStub != nil {
		return fake.ChecksumStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checksumReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) ChecksumCallCount() int {
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	return len(fake.checksumArgsForCall)
}

func (fake *FakeWorkerArtifact) ChecksumCalls(stub func() string) {
	fake.checksumMutex.Lock()
	defer fake.checksumMutex.Unlock()
	fake.ChecksumStub = stub
}

func (fake *FakeWorkerArtifact) ChecksumReturns(result1 string) {
	fake.checksumMutex.Lock()
	defer fake.checksumMutex.Unlock()
	fake.ChecksumStub = nil
	fake.checksumReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorkerArtifact) ChecksumReturnsOnCall(i int, result1 string) {
	fake.checksumMutex.Lock()
	defer fake.checksumMutex.Unlock()
	fake.ChecksumStub = nil
	if fake.checksumReturnsOnCall == nil {
		fake.checksumReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.checksumReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorkerArtifact) CreatedAt() int {
	fake.createdAtMutex.Lock()
	ret, specificReturn := fake.createdAtReturnsOnCall[len(fake.createdAtArgsForCall)]
	fake.createdAtArgsForCall = append(fake.createdAtArgsForCall, struct {
	}{})
	fake.recordInvocation("CreatedAt", []interface{}{})
	fake.createdAtMutex.Unlock()
	if fake.CreatedAtStub != nil {
		return fake.CreatedAtStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createdAtReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) CreatedAtCallCount() int {
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	return len(fake.createdAtArgsForCall)
}

func (fake *FakeWorkerArtifact) CreatedAtCalls(stub func() int) {
	fake.createdAtMutex.Lock()
	defer fake.createdAtMutex.Unlock()
	fake.CreatedAtStub = stub
}

func (fake *FakeWorkerArtifact) CreatedAtReturns(result1 int) {
	fake.createdAtMutex.Lock()
	defer fake.createdAtMutex.Unlock()
	fake.CreatedAtStub = nil
	fake.createdAtReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) CreatedAtReturnsOnCall(i int, result1 int) {
	fake.createdAtMutex.Lock()
	defer fake.createdAtMutex.Unlock()
	fake.CreatedAtStub = nil
	if fake.createdAtReturnsOnCall == nil {
		fake.createdAtReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.createdAtReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeWorkerArtifact) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeWorkerArtifact) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorkerArtifact) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
	}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pathReturns
	return fakeReturns.result1
}

func (fake *FakeWorkerArtifact) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeWorkerArtifact) PathCalls(stub func() string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = stub
}

func (fake *FakeWorkerArtifact) PathReturns(result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorkerArtifact) PathReturnsOnCall(i int, result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorkerArtifact) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checksumMutex.RLock()
	defer fake.checksumMutex.RUnlock()
	fake.createdAtMutex.RLock()
	defer fake.createdAtMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWorkerArtifact) recordInvocation(key string, args []interface{}) {
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

var _ db.WorkerArtifact = new(FakeWorkerArtifact)
