package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a package.json file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Output usage information")
	initCmd.Flags().String("name", "", "Set the name field in package.json")
	initCmd.Flags().String("version", "", "Set the version field in package.json")
	initCmd.Flags().String("description", "", "Set the description field in package.json")
	initCmd.Flags().String("author", "", "Set the author field in package.json")
	initCmd.Flags().String("license", "", "Set the license field in package.json")
	initCmd.Flags().String("homepage", "", "Set the homepage field in package.json")
	initCmd.Flags().String("repository", "", "Set the repository field in package.json")
	initCmd.Flags().String("keywords", "", "Set the keywords field in package.json")

	rootCmd.AddCommand(initCmd)
}
