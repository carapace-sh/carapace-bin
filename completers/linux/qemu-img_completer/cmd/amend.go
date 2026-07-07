package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Update format-specific options of the image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amendCmd).Standalone()

	amendCmd.Flags().StringP("cache", "t", "", "cache mode for FILE (default: writeback)")
	amendCmd.Flags().Bool("force", false, "allow certain unsafe operations")
	amendCmd.Flags().StringP("format", "f", "", "specify FILE format explicitly")
	amendCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	amendCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	amendCmd.Flags().StringP("options", "o", "", "FMT-specific format options (required)")
	amendCmd.Flags().BoolP("progress", "p", false, "show operation progress")
	amendCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.AddCommand(amendCmd)

	carapace.Gen(amendCmd).FlagCompletion(carapace.ActionMap{
		"cache":   qemu.ActionCacheModes(),
		"format":  qemu.ActionImageFormats(),
		"options": carapace.ActionValues(),
	})

	carapace.Gen(amendCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
