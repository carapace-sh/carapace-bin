package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy an existing stack and its resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	destroyCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	destroyCmd.PersistentFlags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	destroyCmd.PersistentFlags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	destroyCmd.PersistentFlags().String("exec-agent", "", "")
	destroyCmd.PersistentFlags().String("exec-kind", "", "")
	destroyCmd.PersistentFlags().StringP("message", "m", "", "Optional message to associate with the destroy operation")
	destroyCmd.PersistentFlags().IntP("parallel", "p", 2147483647, "Allow P resource operations to run in parallel at once (1 for no parallelism). Defaults to unbounded.")
	destroyCmd.PersistentFlags().BoolP("refresh", "r", false, "Refresh the state of the stack's resources before this update")
	destroyCmd.PersistentFlags().Bool("show-config", false, "Show configuration keys and variables")
	destroyCmd.PersistentFlags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	destroyCmd.PersistentFlags().Bool("show-sames", false, "Show resources that don't need to be updated because they haven't changed, alongside those that do")
	destroyCmd.PersistentFlags().BoolP("skip-preview", "f", false, "Do not perform a preview before performing the destroy")
	destroyCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	destroyCmd.PersistentFlags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	destroyCmd.PersistentFlags().String("suppress-permalink", "", "Suppress display of the state permalink")
	destroyCmd.PersistentFlags().StringArrayP("target", "t", []string{}, "Specify a single resource URN to destroy. All resources necessary to destroy this target will also be destroyed. Multiple resources can be specified using: --target urn1 --target urn2")
	destroyCmd.PersistentFlags().Bool("target-dependents", false, "Allows destroying of dependent targets discovered but not specified in --target list")
	destroyCmd.PersistentFlags().BoolP("yes", "y", false, "Automatically approve and perform the destroy after previewing it")
	destroyCmd.Flag("suppress-permalink").NoOptDefVal = "false"
	rootCmd.AddCommand(destroyCmd)

	carapace.Gen(destroyCmd).FlagCompletion(carapace.ActionMap{
		"stack":       action.ActionStacks(destroyCmd, action.StackOpts{}),
		"config-file": carapace.ActionFiles(),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(destroyCmd)
		}),
		//"event-log": carapace.ActionFiles(),
		"exec-kind": action.ActionExecKinds(),
	})
}
