package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [toolchain]",
	Short: "Install a toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("format", "text", "Output format (text|json)")
	installCmd.Flags().BoolP("help", "h", false, "Show help information")
	installCmd.Flags().Bool("no-verify", false, "Do not verify the toolchain's PGP signature before installation")
	installCmd.Flags().String("post-install-file", "", "File path to a location for a post installation script")
	installCmd.Flags().String("progress-file", "", "File path where progress information will be written in JSONL format")
	installCmd.Flags().BoolP("use", "u", false, "Mark the newly installed toolchain as in-use")
	installCmd.Flags().Bool("verify", true, "Verify the toolchain's PGP signature before installation")

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"format":            actionFormat(),
		"post-install-file": carapace.ActionFiles(),
		"progress-file":     carapace.ActionFiles(),
	})

	carapace.Gen(installCmd).PositionalCompletion(
		actionToolchains(),
	)
}
