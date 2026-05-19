package cmd

import (
	"fmt"

	"github.com/daivikawasthi01/Go-pkgr/internal/lockfile"
	"github.com/daivikawasthi01/Go-pkgr/internal/registry"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Install a package",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgName := args[0]

		fmt.Printf("Resolving %s...\n", pkgName)
		pkg, err := registry.Resolve(pkgName)
		if err != nil {
			return err
		}

		lf, err := lockfile.Load("pkgr.lock")
		if err != nil {
			return err
		}

		lf.Add(lockfile.Package{
			Name:    pkg.Name,
			Version: pkg.Version,
			Source:  pkg.Source,
		})

		if err := lf.Save("pkgr.lock"); err != nil {
			return err
		}

		fmt.Printf("Installed %s@%s\n", pkg.Name, pkg.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
