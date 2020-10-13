package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "server is cmd app build with golang",
		Long:  `server is cmd app build with golang using gin and gorm`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
