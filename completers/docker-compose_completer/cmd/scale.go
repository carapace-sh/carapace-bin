package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:   "scale [SERVICE=REPLICAS...]",
	Short: "Scale services ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scaleCmd).Standalone()

	scaleCmd.Flags().Bool("no-deps", false, "Don't start linked services.")
	rootCmd.AddCommand(scaleCmd)
}
