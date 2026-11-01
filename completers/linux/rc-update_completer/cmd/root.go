package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-update",
	Short: "manage OpenRC services in runlevels",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "process all runlevels")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().BoolP("stack", "s", false, "stack a runlevel instead of a service")
	rootCmd.Flags().BoolP("update", "u", false, "force an update of the dependency tree")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("add", "del", "delete", "show").Tag("commands"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				switch c.Args[0] {
				case "add", "del", "delete":
					if len(c.Args) == 1 {
						return openrc.ActionServices()
					}
					return openrc.ActionRunlevels()
				case "show":
					return openrc.ActionRunlevels()
				}
			}
			return carapace.ActionValues()
		}),
	)
}
