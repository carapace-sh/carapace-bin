package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ripsecrets",
	Short: "Prevent committing secret keys into your source code",
	Long:  "https://github.com/sirwart/ripsecrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSlice("additional-pattern", nil, "Additional regex pattern used to find secrets")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("install-pre-commit", false, "Install ripsecrets as a pre-commit hook")
	rootCmd.Flags().Bool("only-matching", false, "Print only the matched (non-empty) parts of a matching line")
	rootCmd.Flags().Bool("strict-ignore", false, "Respect .secretsignore file even for files passed as arguments")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
