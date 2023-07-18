package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "show documentation for package or symbol",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docCmd).Standalone()
	docCmd.Flags().SetInterspersed(false)

	docCmd.Flags().BoolS("all", "all", false, "Show all the documentation for the package")
	docCmd.Flags().BoolS("c", "c", false, "Respect case when matching symbols")
	docCmd.Flags().BoolS("cmd", "cmd", false, "Treat a command like a regular package")
	docCmd.Flags().BoolS("short", "short", false, "One-line representation for each symbol")
	docCmd.Flags().BoolS("src", "src", false, "Show the full source code for the symbol")
	docCmd.Flags().BoolS("u", "u", false, "Show documentation for unexported as well")

	rootCmd.AddCommand(docCmd)

	carapace.Gen(docCmd).PositionalCompletion(
		golang.ActionPackages(),
		carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return golang.ActionSymbols(golang.SymbolOpts{
					Package:    c.Args[0],
					Unexported: docCmd.Flag("u").Changed,
				}).NoSpace()

			case 1:
				return golang.ActionMethodOrFields(golang.MethodOrFieldOpts{
					Package:    c.Args[0],
					Symbol:     c.Parts[0],
					Unexported: docCmd.Flag("u").Changed,
				})

			default:
				return carapace.ActionValues()
			}
		}),
	)
}
