package storkctl

import (
	"flag"
	"io"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"
)

// NewCommand Create a new storkctl command
func NewCommand(cmdFactory Factory, in io.Reader, out io.Writer, errOut io.Writer) *cobra.Command {

	cmds := &cobra.Command{
		Use:   "storkctl",
		Short: "storkctl manages stork resources",
		PersistentPreRun: func(c *cobra.Command, args []string) {
			err := cmdFactory.UpdateConfig()
			if err != nil {
				util.CheckErr(err)
				return
			}
		},
	}

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: errOut}
	cmdFactory.BindFlags(cmds.PersistentFlags())

	cmds.AddCommand(
		newCreateCommand(cmdFactory, ioStreams),
		newDeleteCommand(cmdFactory, ioStreams),
		newGetCommand(cmdFactory, ioStreams),
		newActivateCommand(cmdFactory, ioStreams),
		newDeactivateCommand(cmdFactory, ioStreams),
		newGenerateCommand(cmdFactory, ioStreams),
		newSuspendCommand(cmdFactory, ioStreams),
		newResumeCommand(cmdFactory, ioStreams),
		newVersionCommand(cmdFactory, ioStreams),
		newTriggerCommand(cmdFactory, ioStreams),
	)

	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	err := flag.CommandLine.Parse([]string{})
	if err != nil {
		util.CheckErr(err)
		return nil
	}

	return cmds
}
