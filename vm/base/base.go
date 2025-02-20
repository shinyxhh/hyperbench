package base

import (
	fcom "github.com/hyperbench/hyperbench-common/common"

	"github.com/op/go-logging"
)

// VMBase the base vm for support base config.
type VMBase struct {
	ConfigBase
	Logger *logging.Logger
}

// Type return the vm type.
func (v *VMBase) Type() string {
	return "base"
}

// Close close vm.
func (v *VMBase) Close() {
}

// BeforeDeploy will call before deploy contract.
func (v *VMBase) BeforeDeploy() error {
	return nil
}

// DeployContract deploy contract.
func (v *VMBase) DeployContract() error {
	return nil
}

// BeforeGet will call before get context.
func (v *VMBase) BeforeGet() error {
	return nil
}

// GetContext generate context for execute tx in vm.
func (v *VMBase) GetContext() ([]byte, error) {
	return []byte(""), nil
}

// Statistic statistic remote execute info.
func (v *VMBase) Statistic(from, to int64) (*fcom.RemoteStatistic, error) {
	return &fcom.RemoteStatistic{}, nil
}

// LogStatus records blockheight and time
func (v *VMBase) LogStatus() (int64, error) {
	return 0, nil
}

// BeforeSet will call before set context.
func (v *VMBase) BeforeSet() error {
	return nil
}

// SetContext set context for execute tx in vm, the ctx is generated by GetContext.
func (v *VMBase) SetContext(ctx []byte) error {
	return nil
}

// BeforeRun will call once before run.
func (v *VMBase) BeforeRun() error {
	return nil
}

// Run create and send tx to client.
func (v *VMBase) Run(ctx fcom.TxContext) (*fcom.Result, error) {
	return &fcom.Result{}, nil
}

// AfterRun will call once after run.
func (v *VMBase) AfterRun() error {
	return nil
}

// ConfigBase define base config in vm.
type ConfigBase struct {
	// Path is the path of script file
	Path string
	// Ctx is the context of vm
	Ctx fcom.VMContext
}

// NewVMBase use given config create VMBase.
func NewVMBase(config ConfigBase) *VMBase {
	return &VMBase{
		ConfigBase: config,
		Logger:     fcom.GetLogger("vm"),
	}
}
