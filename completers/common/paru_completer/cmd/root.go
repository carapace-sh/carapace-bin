package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/paru"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "paru",
	Short: "Feature packed AUR helper",
	Long:  "https://github.com/Morganamilo/paru",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("clean", "c", false, "Remove unneeded dependencies")
	rootCmd.Flags().Bool("gendb", false, "Generates development package DB used for updating")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			a := paru.ActionPackageSearch()
			rootCmd.Flags().Visit(func(f *pflag.Flag) {
				if f.Changed {
					a = carapace.ActionValues()
				}
			})
			return a
		}),
	)

	embed.SubcommandsAsFlags(rootCmd,
		databaseCmd,
		deptestCmd,
		filesCmd,
		getpkgbuildCmd,
		queryCmd,
		removeCmd,
		showCmd,
		syncCmd,
		upgradeCmd,
	)
}
