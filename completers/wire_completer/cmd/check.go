package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var wireCmd = &cobra.Command{
	Use:   "wire",
	Short: "print any Wire errors found",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wireCmd).Standalone()

	wireCmd.Flags().String("tags", "", "append build tags to the default wirebuild")
	rootCmd.AddCommand(wireCmd)
}
