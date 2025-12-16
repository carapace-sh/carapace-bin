package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-doc",
	Short: "show documentation for package or symbol",
	Long:  "https://pkg.go.dev/cmd/doc",
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
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	rootCmd.Flags().BoolS("all", "all", false, "Show all the documentation for the package")
	rootCmd.Flags().BoolS("c", "c", false, "Respect case when matching symbols")
	rootCmd.Flags().BoolS("cmd", "cmd", false, "Treat a command like a regular package")
	rootCmd.Flags().BoolS("http", "http", false, "Serve HTML docs over HTTP")
	rootCmd.Flags().BoolS("short", "short", false, "One-line representation for each symbol")
	rootCmd.Flags().BoolS("src", "src", false, "Show the full source code for the symbol")
	rootCmd.Flags().BoolS("u", "u", false, "Show documentation for unexported as well")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionDirectories(),
			golang.ActionPackages().UnlessF(condition.CompletingPath),
		).ToA(),
		carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return golang.ActionSymbols(golang.SymbolOpts{
					Package:    c.Args[0],
					Unexported: rootCmd.Flag("u").Changed,
				}).NoSpace()

			case 1:
				return golang.ActionMethodOrFields(golang.MethodOrFieldOpts{
					Package:    c.Args[0],
					Symbol:     c.Parts[0],
					Unexported: rootCmd.Flag("u").Changed,
				})

			default:
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("C").Value.String())
	})
}
