package cmd

import (
	"fmt"

	"github.com/bitcubix/golang-rest-api/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generate bool

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "with config cmd config file can be set or generated",
	Run: func(cmd *cobra.Command, args []string) {
		if !generate {
			ConfigFile = config.LoadFromFile(ConfigFile)
			if ConfigFile == "no config file found, used defaults" {
				fmt.Println("no config file found, use -g for generate a config file with defaults")
				return
			}
			fmt.Println("current config file is: ", ConfigFile)
			return
		}
		if ConfigFile == "" {
			ConfigFile = "./config.yml"
		}
		config.SetDefaults()
		err := viper.WriteConfigAs(ConfigFile)
		if err != nil {
			panic(fmt.Sprintf("error while generating config file: %v", err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(
		&generate,
		"generate",
		"g",
		false,
		"generate a new config file with defaults",
	)
}
