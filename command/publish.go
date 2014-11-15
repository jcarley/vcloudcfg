package command

import (
	"strings"

	"github.com/jcarley/vcloudcfg/config"
	"github.com/jcarley/vcloudcfg/datastore"
	"github.com/mitchellh/cli"
)

type PublishCommand struct {
	Ui        cli.Ui
	Config    *config.Config
	Datastore *datastore.Datastore
}

func (c *PublishCommand) Run(args []string) int {
	return 0
}

func (c *PublishCommand) Help() string {
	helpText := `
Usage: vcloudcfg publish

	Publishes a box.
`
	return strings.TrimSpace(helpText)
}

func (c *PublishCommand) Synopsis() string {
	return "Publishes a box."
}
