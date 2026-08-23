package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compareCmd)
}

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "perform a backup diff",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareCmd).Standalone()

	compareCmd.Flags().String("D", "", "Limit traversal depth")
	compareCmd.Flags().Bool("E", false, "Don't take exclusions into account")
	compareCmd.Flags().String("I", "", "Ignore paths with a path component equal to name")
	compareCmd.Flags().Bool("U", false, "Ignore logical volume identity")
	compareCmd.Flags().Bool("X", false, "Print output in XML property list format")
	compareCmd.Flags().Bool("a", false, "Compare all supported metadata")
	compareCmd.Flags().Bool("at", false, "Compare extended attributes")
	compareCmd.Flags().Bool("c", false, "Compare creation times")
	compareCmd.Flags().Bool("d", false, "Compare file data forks")
	compareCmd.Flags().Bool("e", false, "Compare ACLs")
	compareCmd.Flags().Bool("f", false, "Compare file flags")
	compareCmd.Flags().Bool("g", false, "Compare GIDs")
	compareCmd.Flags().Bool("m", false, "Compare file modes")
	compareCmd.Flags().Bool("n", false, "No metadata comparison")
	compareCmd.Flags().Bool("s", false, "Compare sizes")
	compareCmd.Flags().Bool("t", false, "Compare modification times")
	compareCmd.Flags().Bool("u", false, "Compare UIDs")

	carapace.Gen(compareCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}