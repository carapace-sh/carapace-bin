package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setencryptiontypeCmd = &cobra.Command{
	Use:   "setencryptiontype",
	Short: "set the default encryption type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setencryptiontypeCmd).Standalone()
	rootCmd.AddCommand(setencryptiontypeCmd)
}
