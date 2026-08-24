package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git-meta_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a string metadata value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().StringP("file", "F", "", "Read value from file")
	setCmd.Flags().BoolP("help", "h", false, "Print help")
	setCmd.Flags().Bool("json", false, "Output as JSON")
	setCmd.Flags().String("timestamp", "", "Override timestamp (milliseconds since epoch, for imports)")
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		action.ActionTarget(),
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)

	carapace.Gen(setCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
