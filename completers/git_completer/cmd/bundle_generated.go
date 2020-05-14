package cmd

import (
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use: "bundle",
	Short: "Move objects and refs by archive",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	bundleCmd.Flags().BoolP("verbose", "v", false, "be verbose; must be placed before a subcommand")
    rootCmd.AddCommand(bundleCmd)
}
