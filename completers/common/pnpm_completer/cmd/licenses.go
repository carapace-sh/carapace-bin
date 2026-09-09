package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licensesCmd = &cobra.Command{
	Use:     "licenses",
	Short:   "Check the licenses of the installed packages",
	Aliases: []string{"licences"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licensesCmd).Standalone()

	licensesCmd.Flags().BoolP("dev", "D", false, "Only dependencies in \"devDependencies\"")
	licensesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	licensesCmd.Flags().Bool("json", false, "Output the information in JSON format")
	licensesCmd.Flags().Bool("long", false, "Show more details (such as a link to the repo)")
	licensesCmd.Flags().Bool("no-optional", false, "Don't check \"optionalDependencies\"")
	licensesCmd.Flags().BoolP("optional", "O", false, "Only dependencies in \"optionalDependencies\"")
	licensesCmd.Flags().BoolP("prod", "P", false, "Only dependencies in \"dependencies\"")
	licensesCmd.Flags().Bool("production", false, "Only dependencies in \"dependencies\"")
	rootCmd.AddCommand(licensesCmd)
}
