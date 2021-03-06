package cmd

import (
	"github.com/cppforlife/go-patch/patch"

	boshtpl "github.com/cloudfoundry/bosh-cli/director/template"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
)

type CreateEnvCmd struct {
	ui          boshui.UI
	envProvider func(string, boshtpl.Variables, patch.Ops) DeploymentPreparer
}

func NewCreateEnvCmd(ui boshui.UI, envProvider func(string, boshtpl.Variables, patch.Ops) DeploymentPreparer) *CreateEnvCmd {
	return &CreateEnvCmd{ui: ui, envProvider: envProvider}
}

func (c *CreateEnvCmd) Run(stage boshui.Stage, opts CreateEnvOpts) error {
	c.ui.BeginLinef("Deployment manifest: '%s'\n", opts.Args.Manifest.Path)

	depPreparer := c.envProvider(
		opts.Args.Manifest.Path, opts.VarFlags.AsVariables(), opts.OpsFlags.AsOps())

	return depPreparer.PrepareDeployment(stage)
}
