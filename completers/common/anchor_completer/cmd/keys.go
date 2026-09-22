package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Program keypair commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keysCmd).Standalone()

	keysCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(keysCmd)
}
