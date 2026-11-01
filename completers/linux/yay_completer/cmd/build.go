package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"B"},
	Short:   "",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().BoolP("install", "i", false, "Install built packages")
	buildCmd.Flags().Bool("needed", false, "Do not reinstall the targets already up to date")
	buildCmd.Flags().CountP("nodeps", "d", "Skip dependency version checks (-dd to skip all checks)")

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
