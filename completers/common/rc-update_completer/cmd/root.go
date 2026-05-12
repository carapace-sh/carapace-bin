package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-update",
	Short: "add and remove OpenRC services to and from a runlevel",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display usage information and exit")
	rootCmd.PersistentFlags().BoolP("nocolor", "C", false, "Disable color output")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Run quietly")
	rootCmd.PersistentFlags().BoolP("user", "U", false, "Operate on user services and runlevels")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Run verbosely")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Display software version")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(showCmd)
}

var addCmd = &cobra.Command{
	Use:   "add service [runlevel...]",
	Short: "add a service to a runlevel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deleteCmd = &cobra.Command{
	Use:     "delete service [runlevel...]",
	Aliases: []string{"del", "remove", "rm"},
	Short:   "delete a service from a runlevel",
	Run:     func(cmd *cobra.Command, args []string) {},
}

var showCmd = &cobra.Command{
	Use:   "show [runlevel...]",
	Short: "show enabled services and their runlevels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	addCmd.Flags().BoolP("stack", "s", false, "Add the runlevel to the current runlevel stack")
	carapace.Gen(addCmd).PositionalCompletion(
		openrc.ActionServices(),
	)
	carapace.Gen(addCmd).PositionalAnyCompletion(
		openrc.ActionRunlevels(),
	)

	carapace.Gen(deleteCmd).Standalone()
	deleteCmd.Flags().BoolP("all", "a", false, "Remove the service from all runlevels")
	deleteCmd.Flags().BoolP("stack", "s", false, "Remove the runlevel from the current runlevel stack")
	carapace.Gen(deleteCmd).PositionalCompletion(
		openrc.ActionServices(),
	)
	carapace.Gen(deleteCmd).PositionalAnyCompletion(
		openrc.ActionRunlevels(),
	)

	carapace.Gen(showCmd).Standalone()
	showCmd.Flags().BoolP("update", "u", false, "Force an update of the dependency tree cache")
	showCmd.Flags().BoolP("verbose", "v", false, "Show all services")
	carapace.Gen(showCmd).PositionalAnyCompletion(
		openrc.ActionRunlevels(),
	)
}
