package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbTestTaskCmd = &cobra.Command{
	Use:   "test",
	Short: "Database test tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbTestTaskCmd).Standalone()
	dbTestTaskCmd.Flags().BoolP("help", "h", false, "Show help")
}
