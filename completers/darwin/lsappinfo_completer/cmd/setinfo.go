package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setinfoCmd = &cobra.Command{
	Use:   "setinfo",
	Short: "Set values for application information items",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setinfoCmd).Standalone()
	setinfoCmd.Flags().String("app", "", "Application specifier")
	rootCmd.AddCommand(setinfoCmd)
}
