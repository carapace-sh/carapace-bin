package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-depend",
	Short: "show OpenRC service dependencies",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("deptree-file", "F", "", "file to load cached deptree from")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("notrace", "T", false, "don't trace service dependencies")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().BoolP("starting", "a", false, "order services as if runlevel is starting")
	rootCmd.Flags().BoolP("stopping", "o", false, "order services as if runlevel is stopping")
	rootCmd.Flags().BoolP("strict", "s", false, "only use what is in the runlevels")
	rootCmd.Flags().StringP("type", "t", "", "type(s) of dependency to list")
	rootCmd.Flags().BoolP("update", "u", false, "force an update of the dependency tree")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"deptree-file": carapace.ActionFiles(),
		"type":         openrc.ActionDependencyTypes(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		openrc.ActionServices(),
	)
}
