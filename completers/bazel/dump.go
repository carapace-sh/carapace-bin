package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dumps the internal state of the %{product} server process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()

	dumpCmd.Flags().Bool("action_cache", false, "Dump action cache content.")
	dumpCmd.Flags().Bool("memory", false, "Dump the memory use of the given Skyframe node.")
	dumpCmd.Flags().Bool("noaction_cache", false, "Dump action cache content.")
	dumpCmd.Flags().Bool("nopackages", false, "Dump package cache content.")
	dumpCmd.Flags().Bool("norule_classes", false, "Dump rule classes.")
	dumpCmd.Flags().Bool("norules", false, "Dump rules, including counts and memory usage (if memory is tracked).")
	dumpCmd.Flags().Bool("packages", false, "Dump package cache content.")
	dumpCmd.Flags().Bool("rule_classes", false, "Dump rule classes.")
	dumpCmd.Flags().Bool("rules", false, "Dump rules, including counts and memory usage (if memory is tracked).")
	dumpCmd.Flags().Bool("skyframe", false, "Dump the Skyframe graph.")
	dumpCmd.Flags().Bool("skykey_filter", false, "Regex filter of SkyKey names to output. Only used with --skyframe=deps, rdeps, function_graph.")
	dumpCmd.Flags().Bool("skylark_memory", false, "Dumps a pprof-compatible memory profile to the specified path. To learn more please see https://github.com/google/pprof.")
	rootCmd.AddCommand(dumpCmd)
}
