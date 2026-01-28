package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Display diverse information built into the OpenSSL libraries",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolS("configdir", "configdir", false, "Default configuration file directory")
	infoCmd.Flags().BoolS("cpusettings", "cpusettings", false, "CPU settings info")
	infoCmd.Flags().BoolS("dirnamesep", "dirnamesep", false, "Directory-filename separator")
	infoCmd.Flags().BoolS("dsoext", "dsoext", false, "Configured extension for modules")
	infoCmd.Flags().BoolS("listsep", "listsep", false, "List separator character")
	infoCmd.Flags().BoolS("modulesdir", "modulesdir", false, "Default module directory (other than engine modules)")
	infoCmd.Flags().BoolS("seeds", "seeds", false, "Seed sources")
	infoCmd.Flags().BoolS("windowscontext", "windowscontext", false, "Windows install context")
	rootCmd.AddCommand(infoCmd)
}
