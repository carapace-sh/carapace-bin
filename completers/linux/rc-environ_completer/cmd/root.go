package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-environ",
	Short: "print OpenRC service environment",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("export", "e", false, "prepend \"export\" to printed variables")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().BoolP("no-escape", "n", false, "do not perform shell sensitive escape on variable values")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("null", "0", false, "end each output line with NUL, not newline")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().StringP("runlevel", "r", "", "add services in runlevel to the list of environments to print")
	rootCmd.Flags().StringP("service", "s", "", "add service to the list of environments to print")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"runlevel": openrc.ActionRunlevels(),
		"service":  openrc.ActionServices(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return openrc.ActionServices()
			}
			return carapace.ActionMultiParts("=", func(c2 carapace.Context) carapace.Action {
				switch len(c2.Parts) {
				case 0:
					return carapace.ActionValues("=").NoSpace()
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
