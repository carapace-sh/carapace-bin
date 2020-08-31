package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	psCmd.Flags().BoolP("all", "a", false, "Show all stopped containers (including those created by the run command)")
	psCmd.Flags().String("filter", "", "Filter services by a property")
	psCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	psCmd.Flags().Bool("services", false, "Display services")
	rootCmd.AddCommand(psCmd)

	carapace.Gen(rootCmd).PositionalAnyCompletion(ActionServices())
}
