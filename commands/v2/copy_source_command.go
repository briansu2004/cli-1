package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/commands/flags"
)

type CopySourceCommand struct {
	RequiredArgs        flags.CopySourceArgs `positional-args:"yes"`
	NoRestart           bool                 `long:"no-restart" description:"Override restart of the application in target environment after copy-source completes"`
	Organization        string               `short:"o" description:"Org that contains the target application"`
	Space               string               `short:"s" description:"Space that contains the target application"`
	usage               interface{}          `usage:"CF_NAME copy-source SOURCE-APP TARGET-APP [-s TARGET-SPACE [-o TARGET-ORG]] [--no-restart]\n"`
	relatedCommands     interface{}          `related_commands:"apps, push, restart, target"`
	envCFStagingTimeout interface{}          `environmentName:"CF_STAGING_TIMEOUT" environmentDescription:"Max wait time for buildpack staging, in minutes" environmentDefault:"15"`
	envCFStartupTimeout interface{}          `environmentName:"CF_STARTUP_TIMEOUT" environmentDescription:"Max wait time for app instance startup, in minutes" environmentDefault:"5"`
}

func (_ CopySourceCommand) Setup(config commands.Config, ui commands.UI) error {
	return nil
}

func (_ CopySourceCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
