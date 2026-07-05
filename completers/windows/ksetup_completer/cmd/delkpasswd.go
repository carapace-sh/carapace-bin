package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delkpasswdCmd = &cobra.Command{
	Use:   "delkpasswd",
	Short: "delete a Kpasswd server address for the realm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delkpasswdCmd).Standalone()
	rootCmd.AddCommand(delkpasswdCmd)
}
