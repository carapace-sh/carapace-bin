package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setAddCmd = &cobra.Command{
	Use:   "set:add",
	Short: "Add a member to a set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setAddCmd).Standalone()

	setAddCmd.Flags().BoolP("help", "h", false, "Print help")
	setAddCmd.Flags().Bool("json", false, "Output as JSON")
	setAddCmd.Flags().String("timestamp", "", "Override timestamp (milliseconds since epoch, for imports)")
	rootCmd.AddCommand(setAddCmd)
}
