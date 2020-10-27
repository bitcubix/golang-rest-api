package cmd

import (
	"github.com/gabrielix29/go-rest-api/internal/server"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the API server",
	Run: func(cmd *cobra.Command, args []string) {
		serve := server.New()
		serve.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
