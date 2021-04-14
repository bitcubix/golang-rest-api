package cmd

import (
	"github.com/bitcubix/golang-rest-api/pkg/cli"
)

var rootCmd = cli.NewCommand(
	"golang-rest-api",
	"RESTful API Boilerplate written in GO",
	"",
)

func Execute() {
	cli.CheckErr(rootCmd.Execute())
}

func init() {
	// TODO rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
