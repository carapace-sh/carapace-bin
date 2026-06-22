package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbFixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Database fixtures tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbFixturesCmd).Standalone()
	dbFixturesCmd.Flags().String("fixtures", "", "Specific fixtures to load (comma-separated)")
	dbFixturesCmd.Flags().BoolP("help", "h", false, "Show help")
}
