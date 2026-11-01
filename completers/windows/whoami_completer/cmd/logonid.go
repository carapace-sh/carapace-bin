package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logonidCmd = &cobra.Command{
	Use:   "logonid",
	Short: "display logon ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logonidCmd).Standalone()
	rootCmd.AddCommand(logonidCmd)
}
