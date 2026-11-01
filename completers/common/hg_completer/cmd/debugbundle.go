package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugbundleCmd = &cobra.Command{
	Use:    "debugbundle",
	Short:  "FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugbundleCmd).Standalone()

	debugbundleCmd.Flags().BoolP("all", "a", false, "show all details")
	debugbundleCmd.Flags().String("part-type", "", "show only the named part type")
	debugbundleCmd.Flags().Bool("spec", false, "print the bundlespec of the bundle")
	rootCmd.AddCommand(debugbundleCmd)
}
