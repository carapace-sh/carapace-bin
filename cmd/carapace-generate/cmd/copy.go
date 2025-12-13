package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace-generate/pkg/completer"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return completer.Optimize(args[0])
	},
}

func init() {
	carapace.Gen(copyCmd).Standalone()

	rootCmd.AddCommand(copyCmd)

	carapace.Gen(copyCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)
}
