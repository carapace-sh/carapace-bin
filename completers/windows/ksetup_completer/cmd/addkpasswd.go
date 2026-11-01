package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addkpasswdCmd = &cobra.Command{
	Use:   "addkpasswd",
	Short: "add a Kpasswd server address for the realm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addkpasswdCmd).Standalone()
	rootCmd.AddCommand(addkpasswdCmd)
}
