/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package commands

import (
	"github.com/nalej/system-model/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"fmt"
	"os"
)

var debugLevel bool
var consoleLogging bool

var rootCmd = &cobra.Command{
	Use:   "system-model",
	Short: "System Model server component",
	Long:  `The System Model server component is responsible of the high level entities in the system.`,
	Version: "unknown-version",

	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debugLevel, "debug", false, "Set debug level")
	rootCmd.PersistentFlags().BoolVar(&consoleLogging, "consoleLogging", false, "Pretty print logging")
}

// Execute the current command.
func Execute() {
	rootCmd.SetVersionTemplate(version.GetVersionInfo())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// SetupLogging sets the debug level and console logging if required.
func SetupLogging() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debugLevel {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if consoleLogging {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}