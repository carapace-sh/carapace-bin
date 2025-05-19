package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeSeederCmd = &cobra.Command{
	Use:   "make:seeder <name>",
	Short: "Create a new seeder class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeSeederCmd).Standalone()

	rootCmd.AddCommand(makeSeederCmd)
}
