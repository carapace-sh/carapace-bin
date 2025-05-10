package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install New Relic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().BoolP("assumeYes", "y", false, "use \"yes\" for all questions during install")
	installCmd.Flags().String("localRecipes", "", "a path to local recipes to load instead of service other fetching")
	installCmd.Flags().StringSliceP("recipe", "n", nil, "the name of a recipe to install")
	installCmd.Flags().StringSliceP("recipePath", "c", nil, "the path to a recipe file to install")
	installCmd.Flags().StringSlice("tag", nil, "the tags to add during install, can be multiple. Example: --tag tag1:test,tag2:test")
	installCmd.Flags().BoolP("testMode", "t", false, "fakes operations for UX testing")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"recipePath": carapace.ActionDirectories(),
	})
}
