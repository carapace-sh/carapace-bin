package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var expandCmd = &cobra.Command{
	Use:   "expand",
	Short: "Expands macros (wrapper around cargo expand)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(expandCmd).Standalone()

	expandCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	expandCmd.Flags().StringP("program-name", "p", "", "Expand only this program")
	expandCmd.Flags().Bool("stdout", false, "Write to stdout")
	rootCmd.AddCommand(expandCmd)
}
