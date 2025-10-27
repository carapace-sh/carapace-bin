package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:     "ps [OPTIONS]",
	Short:   "List containers",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(psCmd).Standalone()

	psCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	psCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	psCmd.Flags().String("format", "", "Format output using a custom template:")
	psCmd.Flags().StringP("last", "n", "", "Show n last created containers (includes all states)")
	psCmd.Flags().BoolP("latest", "l", false, "Show the latest created container (includes all states)")
	psCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	psCmd.Flags().BoolP("quiet", "q", false, "Only display container IDs")
	psCmd.Flags().BoolP("size", "s", false, "Display total file sizes")
	rootCmd.AddCommand(psCmd)
}
