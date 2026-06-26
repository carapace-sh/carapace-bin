package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kuberc_setCmd = &cobra.Command{
	Use:   "set --section (defaults|aliases) --command COMMAND",
	Short: "Set values in the kuberc configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kuberc_setCmd).Standalone()

	kuberc_setCmd.Flags().StringSlice("allowlist-entry", nil, "Allowlist entry the form field=value (can be specified multiple times)")
	kuberc_setCmd.Flags().StringSlice("appendarg", nil, "Argument to append to the command (can be specified multiple times, for aliases only)")
	kuberc_setCmd.Flags().String("command", "", "Command to configure (e.g., 'get', 'create', 'set env')")
	kuberc_setCmd.Flags().String("kuberc", "", "Path to the kuberc file to use for preferences. This can be disabled by exporting KUBECTL_KUBERC=false feature gate or turning off the feature KUBERC=off.")
	kuberc_setCmd.Flags().String("name", "", "Alias name (required for --section=aliases)")
	kuberc_setCmd.Flags().StringSlice("option", nil, "Flag option in the form flag=value (can be specified multiple times)")
	kuberc_setCmd.Flags().Bool("overwrite", false, "Allow overwriting existing entries")
	kuberc_setCmd.Flags().String("policy", "", "Plugin policy to use for exec credential plugins, must be one of 'AllowAll', 'DenyAll' or 'Allowlist'")
	kuberc_setCmd.Flags().StringSlice("prependarg", nil, "Argument to prepend to the command (can be specified multiple times, for aliases only)")
	kuberc_setCmd.Flags().String("section", "", "Section to modify: 'defaults', 'aliases', or 'credentialplugin'")
	kuberc_setCmd.MarkFlagRequired("section")
	kubercCmd.AddCommand(kuberc_setCmd)

	carapace.Gen(kuberc_setCmd).FlagCompletion(carapace.ActionMap{
		"allowlist-entry": carapace.ActionValues(), // TODO
		"appendarg":       carapace.ActionValues(), // TODO
		"command":         carapace.ActionCommands(rootCmd).Split(),
		"kuberc":          carapace.ActionFiles(),
		"option":          carapace.ActionValues(), // TODO
		"policy":          carapace.ActionValues("AllowAll", "DenyAll", "Allowlist"),
		"prependarg":      carapace.ActionValues(), // TODO
		"section":         carapace.ActionValues("defaults", "aliases", "credentialplugin"),
	})
}
