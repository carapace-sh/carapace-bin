package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pfctl",
	Short: "control the packet filter and network address translation device",
	Long:  "https://keith.github.io/xcode-manpages/pfctl.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Load only queue rules")
	rootCmd.Flags().StringS("D", "D", "", "Define macro=value")
	rootCmd.Flags().BoolS("E", "E", false, "Enable packet filter and increment reference count")
	rootCmd.Flags().StringS("F", "F", "", "Flush modifier")
	rootCmd.Flags().StringS("K", "K", "", "Kill host/network states")
	rootCmd.Flags().BoolS("M", "M", false, "Enable port to name translation")
	rootCmd.Flags().BoolS("N", "N", false, "Load only NAT rules")
	rootCmd.Flags().BoolS("O", "O", false, "Load only options")
	rootCmd.Flags().BoolS("R", "R", false, "Load only filter rules")
	rootCmd.Flags().StringS("T", "T", "", "Table command")
	rootCmd.Flags().StringS("X", "X", "", "Release pf enable token")
	rootCmd.Flags().StringS("a", "a", "", "Anchor")
	rootCmd.Flags().BoolS("d", "d", false, "Disable packet filter")
	rootCmd.Flags().BoolS("e", "e", false, "Enable packet filter")
	rootCmd.Flags().StringS("f", "f", "", "Load rules from file")
	rootCmd.Flags().BoolS("g", "g", false, "Include debugging output")
	rootCmd.Flags().BoolS("h", "h", false, "Help")
	rootCmd.Flags().StringS("i", "i", "", "Interface")
	rootCmd.Flags().StringS("k", "k", "", "Kill host/network states")
	rootCmd.Flags().BoolS("l", "l", false, "Use local protocol database")
	rootCmd.Flags().BoolS("m", "m", false, "Merge options without resetting")
	rootCmd.Flags().BoolS("n", "n", false, "Only parse rules, do not load")
	rootCmd.Flags().StringS("o", "o", "", "Output level")
	rootCmd.Flags().StringS("p", "p", "", "Device")
	rootCmd.Flags().BoolS("q", "q", false, "Only print errors and warnings")
	rootCmd.Flags().BoolS("r", "r", false, "Reverse DNS lookups on states")
	rootCmd.Flags().StringS("s", "s", "", "Show modifier")
	rootCmd.Flags().StringS("t", "t", "", "Table")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose output")
	rootCmd.Flags().StringS("w", "w", "", "Wait")
	rootCmd.Flags().StringS("x", "x", "", "Debug level")
	rootCmd.Flags().BoolS("z", "z", false, "Clear per-rule statistics")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"F": carapace.ActionValues("nat", "queue", "rules", "states", "Sources", "info", "Tables", "osfp", "all"),
		"f": carapace.ActionFiles(),
		"s": carapace.ActionValues("rules", "nat", "queue", "states", "Sources", "info", "Tables", "osfp", "anchors", "all"),
		"x": carapace.ActionValues("none", "urgent", "misc", "loud", "insane"),
	})
}
