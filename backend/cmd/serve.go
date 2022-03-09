package cmd

import (
	"box/internal/rest"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "serve",
		Run: serveCmd,
	})
}

func serveCmd(c *cobra.Command, args []string) {
	server, err := rest.NewServer(rest.Config{
		Port: 8080,
	}, logrus.StandardLogger())
	if err != nil {
		logrus.WithError(err).Fatal("Failed to create server")
	}

	logrus.WithFields(logrus.Fields{
		"port": 8080,
	}).Info("Starting REST server")

	if err := server.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
}
