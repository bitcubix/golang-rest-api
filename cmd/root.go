package cmd

import "github.com/spf13/cobra"

var ConfigFile string

var rootCmd = &cobra.Command{
	Use:   "golang-rest-api",
	Short: "RESTful API Boilerplate written in GO",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&ConfigFile,
		"file",
		"f",
		"",
		"set the path for config file",
	)
}
