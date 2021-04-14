package cmd

import (
	"github.com/bitcubix/golang-rest-api/internal/server"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		serv, err := server.New()
		if err != nil {
			panic(err)
		}

		err = serv.RunHTTP()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
