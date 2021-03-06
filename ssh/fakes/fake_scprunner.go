// This file was generated by counterfeiter
package fakes

import (
	"sync"

	boshdir "github.com/cloudfoundry/bosh-cli/director"
	"github.com/cloudfoundry/bosh-cli/ssh"
)

type FakeSCPRunner struct {
	RunStub        func(ssh.ConnectionOpts, boshdir.SSHResult, ssh.SCPArgs) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 ssh.ConnectionOpts
		arg2 boshdir.SSHResult
		arg3 ssh.SCPArgs
	}
	runReturns struct {
		result1 error
	}
}

func (fake *FakeSCPRunner) Run(arg1 ssh.ConnectionOpts, arg2 boshdir.SSHResult, arg3 ssh.SCPArgs) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 ssh.ConnectionOpts
		arg2 boshdir.SSHResult
		arg3 ssh.SCPArgs
	}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeSCPRunner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeSCPRunner) RunArgsForCall(i int) (ssh.ConnectionOpts, boshdir.SSHResult, ssh.SCPArgs) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2, fake.runArgsForCall[i].arg3
}

func (fake *FakeSCPRunner) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

var _ ssh.SCPRunner = new(FakeSCPRunner)
