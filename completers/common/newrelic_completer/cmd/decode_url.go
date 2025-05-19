package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var decode_urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Decodes NR1 URL Strings ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(decode_urlCmd).Standalone()
	decode_urlCmd.Flags().StringP("param", "p", "", "the query parameter you want to decode")
	decode_urlCmd.Flags().StringP("search", "s", "", "the search key you want returned")
	decodeCmd.AddCommand(decode_urlCmd)
}
