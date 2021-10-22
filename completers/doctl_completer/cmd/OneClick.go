package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var OneClickCmd = &cobra.Command{
	Use:   "1-click",
	Short: "Display commands that pertain to 1-click applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(OneClickCmd).Standalone()
	rootCmd.AddCommand(OneClickCmd)
}
