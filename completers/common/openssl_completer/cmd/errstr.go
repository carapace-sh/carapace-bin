package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var errstrCmd = &cobra.Command{
	Use:     "errstr",
	Short:   "Error Number to Error String Conversion",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(errstrCmd).Standalone()

	rootCmd.AddCommand(errstrCmd)
}
