// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/network-policy-client/src/models"
	"github.com/cloudfoundry-incubator/network-policy-client/src/policy_client"
)

type Chunker struct {
	ChunkStub        func(allPolicies []models.Policy) [][]models.Policy
	chunkMutex       sync.RWMutex
	chunkArgsForCall []struct {
		allPolicies []models.Policy
	}
	chunkReturns struct {
		result1 [][]models.Policy
	}
	chunkReturnsOnCall map[int]struct {
		result1 [][]models.Policy
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Chunker) Chunk(allPolicies []models.Policy) [][]models.Policy {
	var allPoliciesCopy []models.Policy
	if allPolicies != nil {
		allPoliciesCopy = make([]models.Policy, len(allPolicies))
		copy(allPoliciesCopy, allPolicies)
	}
	fake.chunkMutex.Lock()
	ret, specificReturn := fake.chunkReturnsOnCall[len(fake.chunkArgsForCall)]
	fake.chunkArgsForCall = append(fake.chunkArgsForCall, struct {
		allPolicies []models.Policy
	}{allPoliciesCopy})
	fake.recordInvocation("Chunk", []interface{}{allPoliciesCopy})
	fake.chunkMutex.Unlock()
	if fake.ChunkStub != nil {
		return fake.ChunkStub(allPolicies)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.chunkReturns.result1
}

func (fake *Chunker) ChunkCallCount() int {
	fake.chunkMutex.RLock()
	defer fake.chunkMutex.RUnlock()
	return len(fake.chunkArgsForCall)
}

func (fake *Chunker) ChunkArgsForCall(i int) []models.Policy {
	fake.chunkMutex.RLock()
	defer fake.chunkMutex.RUnlock()
	return fake.chunkArgsForCall[i].allPolicies
}

func (fake *Chunker) ChunkReturns(result1 [][]models.Policy) {
	fake.ChunkStub = nil
	fake.chunkReturns = struct {
		result1 [][]models.Policy
	}{result1}
}

func (fake *Chunker) ChunkReturnsOnCall(i int, result1 [][]models.Policy) {
	fake.ChunkStub = nil
	if fake.chunkReturnsOnCall == nil {
		fake.chunkReturnsOnCall = make(map[int]struct {
			result1 [][]models.Policy
		})
	}
	fake.chunkReturnsOnCall[i] = struct {
		result1 [][]models.Policy
	}{result1}
}

func (fake *Chunker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chunkMutex.RLock()
	defer fake.chunkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Chunker) recordInvocation(key string, args []interface{}) {
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

var _ policy_client.Chunker = new(Chunker)
