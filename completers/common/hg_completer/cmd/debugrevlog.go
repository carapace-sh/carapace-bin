package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugrevlogCmd = &cobra.Command{
	Use:    "debugrevlog",
	Short:  "-c|-m|FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrevlogCmd).Standalone()

	debugrevlogCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugrevlogCmd.Flags().String("dir", "", "open directory manifest")
	debugrevlogCmd.Flags().BoolP("dump", "d", false, "dump index data")
	debugrevlogCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	rootCmd.AddCommand(debugrevlogCmd)
}
