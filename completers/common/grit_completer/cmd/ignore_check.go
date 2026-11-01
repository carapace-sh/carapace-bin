package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ignore_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Print whether paths match exclude rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ignore_checkCmd).Standalone()

	ignore_checkCmd.Flags().Bool("all", false, "Also list paths that are not ignored")
	ignore_checkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ignoreCmd.AddCommand(ignore_checkCmd)
}
