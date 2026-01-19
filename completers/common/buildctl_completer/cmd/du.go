package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duCmd = &cobra.Command{
	Use:   "du",
	Short: "disk usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duCmd).Standalone()

	duCmd.Flags().StringP("filter", "f", "", "Filter records")
	duCmd.Flags().String("format", "", "Format the output using the given Go template")
	duCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.AddCommand(duCmd)
}
