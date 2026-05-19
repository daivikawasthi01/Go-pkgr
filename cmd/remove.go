package cmd

import (
	"fmt"

	"github.com/daivikawasthi01/Go-pkgr/internal/lockfile"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [package]",
	Short: "Remove a package",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgName := args[0]

		lf, err := lockfile.Load("pkgr.lock")
		if err != nil {
			return err
		}

		if !lf.Remove(pkgName) {
			return fmt.Errorf("package %q is not installed", pkgName)
		}

		if err := lf.Save("pkgr.lock"); err != nil {
			return err
		}

		fmt.Printf("Removed %s\n", pkgName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
