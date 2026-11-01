package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetIntlCmd = &cobra.Command{
	Use:   "Get-Intl",
	Short: "display international settings and languages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetIntlCmd).Standalone()
	rootCmd.AddCommand(GetIntlCmd)
}
