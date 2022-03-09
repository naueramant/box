package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "box",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	logrus.SetLevel(logrus.DebugLevel)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Failed to execute root command")
		os.Exit(1)
	}
}
