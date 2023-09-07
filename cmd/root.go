/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ficsit-exporter",
	Short: "Exporter ficsit information using a multi-target exporter method",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/ficsit-exporter")
	viper.AddConfigPath("$HOME/.config/ficsit-exporter")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		// only throw an error if its not 'config not found'
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	rootCmd.AddCommand(serve)
}
