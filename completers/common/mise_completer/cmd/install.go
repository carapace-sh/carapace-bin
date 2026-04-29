package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install [tool@version]...",
	Aliases: []string{"i"},
	Short:   "Install tools from config or arguments",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().BoolP("force", "f", false, "Force reinstall")
	installCmd.Flags().Bool("jobs", false, "Number of jobs to run in parallel")
	rootCmd.AddCommand(installCmd)
}
