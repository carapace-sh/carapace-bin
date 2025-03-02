package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mooCmd = &cobra.Command{
	Use:    "moo",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mooCmd).Standalone()

	mooCmd.Flags().Bool("color", false, "use color")
	rootCmd.AddCommand(mooCmd)
}
