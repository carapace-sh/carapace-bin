package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "melt",
	Short: "melt generates a seed phrase from an SSH key",
	Long:  "https://github.com/charmbracelet/melt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("help", "h", false, "help for melt")
	rootCmd.PersistentFlags().StringP("language", "l", "", "Language (default \"en\")")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"language": os.ActionLanguages(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			ssh.ActionPrivateKeys(),
		).ToA(),
	)
}
