package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setRmCmd = &cobra.Command{
	Use:   "set:rm",
	Short: "Remove a member from a set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setRmCmd).Standalone()

	setRmCmd.Flags().BoolP("help", "h", false, "Print help")
	setRmCmd.Flags().Bool("json", false, "Output as JSON")
	setRmCmd.Flags().String("timestamp", "", "Override timestamp (milliseconds since epoch, for imports)")
	rootCmd.AddCommand(setRmCmd)
}
