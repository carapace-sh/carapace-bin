package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scoop",
	Short: "command-line installer for Windows",
	Long:  "https://scoop.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("version", "v", false, "show the current Scoop version")

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		aliases := scoop.Aliases()
		for _, existing := range cmd.Commands() {
			delete(aliases, existing.Name())
		}
		for name, description := range aliases {
			aliasCmd := &cobra.Command{
				Use:   name,
				Short: description,
				Run:   func(cmd *cobra.Command, args []string) {},
			}
			carapace.Gen(aliasCmd).Standalone()
			cmd.AddCommand(aliasCmd)
		}
	})
}
