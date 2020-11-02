package cmd

import (
	"github.com/gabrielix29/go-rest-api/internal/server"
	"github.com/gabrielix29/go-rest-api/pkg/logger"
	"github.com/gabrielix29/go-rest-api/pkg/model"
	"github.com/spf13/cobra"
)

// migrateCmd represents the version command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		serve := server.New()
		serve.InitDatabase()
		err := serve.Database.AutoMigrate(&model.Book{})
		if err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
