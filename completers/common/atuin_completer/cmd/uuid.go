package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uuidCmd).Standalone()

	uuidCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(uuidCmd)
}
