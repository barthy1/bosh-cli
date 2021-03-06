package cmd

import (
	"github.com/cppforlife/go-patch/patch"

	boshtpl "github.com/cloudfoundry/bosh-cli/director/template"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
)

type DeleteCmd struct {
	ui          boshui.UI
	envProvider func(string, boshtpl.Variables, patch.Ops) DeploymentDeleter
}

func NewDeleteCmd(ui boshui.UI, envProvider func(string, boshtpl.Variables, patch.Ops) DeploymentDeleter) *DeleteCmd {
	return &DeleteCmd{ui: ui, envProvider: envProvider}
}

func (c *DeleteCmd) Run(stage boshui.Stage, opts DeleteEnvOpts) error {
	c.ui.BeginLinef("Deployment manifest: '%s'\n", opts.Args.Manifest.Path)

	depDeleter := c.envProvider(
		opts.Args.Manifest.Path, opts.VarFlags.AsVariables(), opts.OpsFlags.AsOps())

	return depDeleter.DeleteDeployment(stage)
}
