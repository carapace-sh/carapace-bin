package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	GroupID: "configuration",
	Short:   "Configure environment paths and shell settings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("dir", "", "Custom install directory")
	installCmd.Flags().Bool("skip-aliases", false, "Preserve existing agy/antigravity aliases")
	installCmd.Flags().Bool("skip-path", false, "Do not modify the shell profile PATH")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
