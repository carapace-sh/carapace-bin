package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openrc",
	Short: "OpenRC init system",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().Bool("no-stop", false, "do not stop any services")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().StringP("override", "o", "", "override the next runlevel to change into when leaving single user or boot runlevels")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().StringP("service", "s", "", "run the service specified with the rest of the arguments")
	rootCmd.Flags().BoolP("sys", "S", false, "output the RC system type, if any")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"override": openrc.ActionRunlevels(),
		"service":  openrc.ActionServices(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionRunlevels(),
	)
}
