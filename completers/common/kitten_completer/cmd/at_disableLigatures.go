package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableLigaturesCmd = &cobra.Command{
	Use:   "disable-ligatures",
	Short: "Control ligature rendering",
}

func init() {
	disableLigaturesCmd.AddCommand(atCmd)
	carapace.Gen(disableLigaturesCmd).Standalone()

	disableLigaturesCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(disableLigaturesCmd).FlagCompletion(carapace.ActionMap{})
}