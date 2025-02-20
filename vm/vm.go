package vm

import (
	fcom "github.com/hyperbench/hyperbench-common/common"

	"github.com/hyperbench/hyperbench/vm/base"
	"github.com/hyperbench/hyperbench/vm/lua"
)

// Type is the ext name of script file
type Type string

const (
	// LUA vm type of lua.
	LUA Type = "lua"
)

// MasterVM is the part interface of VM that will be called in master
type MasterVM interface {
	// BeforeDeploy will call before deploy contract.
	BeforeDeploy() error
	// DeployContract deploy contract.
	DeployContract() error
	// BeforeGet will call before get context.
	BeforeGet() error
	// GetContext generate context for execute tx in vm.
	GetContext() ([]byte, error)
	// Statistic statistic remote execute info.
	Statistic(from, to int64) (*fcom.RemoteStatistic, error)
	// LogStatus records blockheight and time
	LogStatus() (int64, error)
}

// BaseVM is the part interface of VM that will be called in both master and worker
type BaseVM interface {
	// Type return the vm type.
	Type() string
	// Close close vm.
	Close()
}

// WorkerVM is the part interface of VM that will be called in worker
type WorkerVM interface {
	// BeforeSet will call before set context.
	BeforeSet() error
	// SetContext set context for execute tx in vm, the ctx is generated by GetContext.
	SetContext(ctx []byte) error
	// BeforeRun will call once before run.
	BeforeRun() error
	// Run create and send tx to client.
	Run(ctx fcom.TxContext) (*fcom.Result, error)
	// AfterRun will call once after run.
	AfterRun() error
}

// VM is the integrated interface of Virtual Machine running script
type VM interface {
	BaseVM
	MasterVM
	WorkerVM
}

// NewVM creates a VM according to type and config
func NewVM(vmType string, configBase base.ConfigBase) (VM, error) {
	vm := base.NewVMBase(configBase)
	switch Type(vmType) {
	case LUA:
		return lua.NewVM(vm)
	default:
		return vm, nil
	}
}
