package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nm",
	Short: "Nm lists the symbols defined or used by an object file, archive, or executabl",
	Long:  "https://pkg.go.dev/cmd/nm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "an alias for -sort address (numeric)")
	rootCmd.Flags().BoolS("size", "size", false, "print symbol size in decimal between address and type")
	rootCmd.Flags().StringS("sort", "sort", "", "sort output in the given order (default name)")
	rootCmd.Flags().BoolS("type", "type", false, "print symbol type after name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"sort": golang.ActionSymbolTypes().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
