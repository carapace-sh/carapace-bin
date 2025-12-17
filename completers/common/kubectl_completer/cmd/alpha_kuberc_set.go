package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alpha_kuberc_setCmd = &cobra.Command{
	Use:   "set --section (defaults|aliases) --command COMMAND",
	Short: "Set values in the kuberc configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_kuberc_setCmd).Standalone()

	alpha_kuberc_setCmd.Flags().StringSlice("appendarg", nil, "Argument to append to the command (can be specified multiple times, for aliases only)")
	alpha_kuberc_setCmd.Flags().String("command", "", "Command to configure (e.g., 'get', 'create', 'set env')")
	alpha_kuberc_setCmd.Flags().String("kuberc", "", "Path to the kuberc file to use for preferences. This can be disabled by exporting KUBECTL_KUBERC=false feature gate or turning off the feature KUBERC=off.")
	alpha_kuberc_setCmd.Flags().String("name", "", "Alias name (required for --section=aliases)")
	alpha_kuberc_setCmd.Flags().StringSlice("option", nil, "Flag option in the form flag=value (can be specified multiple times)")
	alpha_kuberc_setCmd.Flags().Bool("overwrite", false, "Allow overwriting existing entries")
	alpha_kuberc_setCmd.Flags().StringSlice("prependarg", nil, "Argument to prepend to the command (can be specified multiple times, for aliases only)")
	alpha_kuberc_setCmd.Flags().String("section", "", "Section to modify: 'defaults' or 'aliases'")
	alpha_kuberc_setCmd.MarkFlagRequired("command")
	alpha_kuberc_setCmd.MarkFlagRequired("section")
	alpha_kubercCmd.AddCommand(alpha_kuberc_setCmd)

	carapace.Gen(alpha_kuberc_setCmd).FlagCompletion(carapace.ActionMap{
		"appendarg": carapace.ActionValues(), // TODO
		"command":   carapace.ActionCommands(rootCmd).Split(),
		"kuberc":    carapace.ActionFiles(),
		"option":    carapace.ActionValues(), // TODO
		"section":   carapace.ActionValues("defaults", "aliases"),
	})
}
