package cmd

import (
	"fmt"

	"github.com/daivikawasthi01/Go-pkgr/internal/lockfile"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		lf, err := lockfile.Load("pkgr.lock")
		if err != nil {
			return err
		}

		if len(lf.Packages) == 0 {
			fmt.Println("No packages installed")
			return nil
		}

		fmt.Printf("Installed packages (%d):\n", len(lf.Packages))
		for name, pkg := range lf.Packages {
			fmt.Printf("  %s@%s\n", name, pkg.Version)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
