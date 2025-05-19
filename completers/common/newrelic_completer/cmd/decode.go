package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes NR1 URL Strings ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(decodeCmd).Standalone()
	rootCmd.AddCommand(decodeCmd)
}
