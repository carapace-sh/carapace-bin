package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var decode_entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Decodes NR1 Entitys ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(decode_entityCmd).Standalone()
	decode_entityCmd.Flags().StringP("key", "k", "", "the key you require back from an entity")
	decodeCmd.AddCommand(decode_entityCmd)
}
