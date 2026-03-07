package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module <subcommand> [options] [<module-spec>...]",
	Short: "manage modules (deprecated)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var moduleListCmd = &cobra.Command{
	Use:   "list [module-spec]...",
	Short: "list module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var moduleInfoCmd = &cobra.Command{
	Use:   "info [module-spec]...",
	Short: "print details about module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var moduleEnableCmd = &cobra.Command{
	Use:   "enable <module-spec>...",
	Short: "enable module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var moduleDisableCmd = &cobra.Command{
	Use:   "disable <module-spec>...",
	Short: "disable modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var moduleResetCmd = &cobra.Command{
	Use:   "reset <module-spec>...",
	Short: "reset module state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleCmd).Standalone()

	moduleCmd.AddCommand(moduleListCmd)
	moduleCmd.AddCommand(moduleInfoCmd)
	moduleCmd.AddCommand(moduleEnableCmd)
	moduleCmd.AddCommand(moduleDisableCmd)
	moduleCmd.AddCommand(moduleResetCmd)

	for _, subcmd := range []*cobra.Command{moduleListCmd, moduleInfoCmd} {
		subcmd.Flags().Bool("disabled", false, "show only disabled modules")
		subcmd.Flags().Bool("enabled", false, "show only enabled modules")
	}

	for _, subcmd := range []*cobra.Command{moduleEnableCmd, moduleDisableCmd, moduleResetCmd} {
		subcmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
		subcmd.Flags().Bool("skip-unavailable", false, "allow skipping modules")
	}

	rootCmd.AddCommand(moduleCmd)
}
