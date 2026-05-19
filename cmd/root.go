package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "pkgr",
	Short: "A fast, simple package manager",
	Long:  "pkgr is a CLI package manager to install, remove and list packages",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of pkgr",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pkgr v%s\n", version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
