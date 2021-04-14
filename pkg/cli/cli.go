package cli

import "github.com/spf13/cobra"

type Command struct {
	*cobra.Command
}

// TODO better way
func NewCommand(use string, short string, long string) *Command {
	return &Command{&cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
	}}
}

func CheckErr(msg interface{}) {
	cobra.CheckErr(msg)
}
