package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().BoolP("help", "h", false, "Print help")
	commitCmd.Flags().String("layer", "", "Commit only changes in this layer (mount path relative to repo root)")
	commitCmd.Flags().StringSlice("layer-message", nil, "Per-layer commit message. Takes two values: <path> <message>. Can be specified multiple times")
	commitCmd.Flags().String("link", "", "Commit only changes in this linked repository (mount path relative to repo root)")
	commitCmd.Flags().StringSlice("link-message", nil, "Per-link commit message. Takes two values: <path> <message>. Can be specified multiple times")
	commitCmd.Flags().Bool("stats", false, "Print stats")
	rootCmd.AddCommand(commitCmd)
}
