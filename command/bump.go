package command

import (
	"flag"
	"fmt"

	"strings"

	"github.com/jcarley/vcloudcfg/config"
	"github.com/jcarley/vcloudcfg/datastore"
	"github.com/larryli/vagrantcloud.v1"
	"github.com/mitchellh/cli"
)

type BumpCommand struct {
	Ui        cli.Ui
	Config    *config.Config
	Datastore *datastore.Datastore
}

func (c *BumpCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("bump", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	var boxName string
	var boxVersion string

	cmdFlags.StringVar(&boxName, "box", "", "The box name")
	cmdFlags.StringVar(&boxVersion, "box-version", "", "The box version")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	// === Use the datastore here to get the box
	record, err := c.Datastore.GetBox(boxName)
	if err != nil {
		panic(err)
	}

	api := vagrantcloud.New(c.Config.Token)
	box := api.Box(c.Config.Username, record.BoxName)
	err = box.Get()
	if err != nil {
		panic(err)
	}

	nextVersion := box.CurrentVersion.Number + 1
	version := box.Version(nextVersion)
	version.Version = record.GetVersion(boxVersion).BoxVersion

	err = version.New()
	if err != nil {
		panic(err)
	}

	provider := version.Provider(vagrantcloud.ProviderVirtualbox)
	provider.OriginalUrl = record.GetVersion(boxVersion).OriginalUrl

	err = provider.New()
	if err != nil {
		panic(err)
	}
	// version.Release()

	c.Ui.Output(fmt.Sprintf("New version %s created for box %s", version.Version, box.Name))

	return 0
}

func (c *BumpCommand) Help() string {
	helpText := `
Usage: vcloudcfg bump --box <name> --box-version <version>

Options:

   --box                   The name of the box
   --box-version VERSION   The specific version of the box to promote.

Creates a new vagrant cloud version and provider for an existing box.  Uses the
values from the Boxfile file.
`
	return strings.TrimSpace(helpText)
}

func (c *BumpCommand) Synopsis() string {
	return "Creates a new version and provider for an existing box."
}
