package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_envCmd = &cobra.Command{
	Use:   "env",
	Short: "Change environment variables seen by future children",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_envCmd).Standalone()

	at_envCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	atCmd.AddCommand(at_envCmd)

}
