package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/procs_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "procs",
	Short: "a replacement for `ps` written in Rust",
	Long:  "https://github.com/dalance/procs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("and", "a", false, "AND  logic for multi-keyword")
	rootCmd.Flags().StringP("color", "c", "", "Color mode")
	rootCmd.Flags().String("gen-completion", "", "Generate shell completion file")
	rootCmd.Flags().String("gen-completion-out", "", "Generate shell completion file and write to stdout")
	rootCmd.Flags().Bool("gen-config", false, "Generate configuration sample file")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringSliceP("insert", "i", nil, "Insert column to slot")
	rootCmd.Flags().String("interval", "", "Interval to calculate throughput")
	rootCmd.Flags().Bool("json", false, "JSON output")
	rootCmd.Flags().BoolP("list", "l", false, "Show list of kind")
	rootCmd.Flags().String("load-config", "", "Load configuration from file")
	rootCmd.Flags().BoolP("nand", "d", false, "NAND logic for multi-keyword")
	rootCmd.Flags().Bool("no-header", false, "Suppress header")
	rootCmd.Flags().BoolP("nor", "r", false, "NOR  logic for multi-keyword")
	rootCmd.Flags().String("only", "", "Specified column only")
	rootCmd.Flags().BoolP("or", "o", false, "OR   logic for multi-keyword")
	rootCmd.Flags().StringP("pager", "p", "", "Pager mode")
	rootCmd.Flags().Bool("procfs", false, "Path to procfs")
	rootCmd.Flags().String("sorta", "", "Sort column by ascending")
	rootCmd.Flags().String("sortd", "", "Sort column by descending")
	rootCmd.Flags().String("theme", "", "Theme mode")
	rootCmd.Flags().Bool("thread", false, "Show thread")
	rootCmd.Flags().BoolP("tree", "t", false, "Tree view")
	rootCmd.Flags().String("use-config", "", "Use built-in configuration")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().BoolP("watch", "w", false, "Watch mode with default interval (1s)")
	rootCmd.Flags().StringP("watch-interval", "W", "", "Watch mode with custom interval")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":              carapace.ActionValues("auto", "always", "disable").StyleF(style.ForKeyword),
		"gen-completion":     carapace.ActionValues("bash", "elvish", "fish", "powershell", "zsh"),
		"gen-completion-out": carapace.ActionValues("bash", "elvish", "fish", "powershell", "zsh"),
		"insert":             action.ActionColumns(),
		"load-config":        carapace.ActionFiles(),
		"only":               action.ActionColumns(),
		"pager":              carapace.ActionValues("auto", "always", "disable").StyleF(style.ForKeyword),
		"sorta":              action.ActionColumns(),
		"sortd":              action.ActionColumns(),
		"theme":              carapace.ActionValues("auto", "dark", "light"),
		"use-config":         carapace.ActionValues("default", "large"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.Batch(
			docker.ActionContainers().Suppress(".*"), // TODO only active containers
			ps.ActionProcessExecutables(),
			ps.ActionProcessIds(),
		).ToA(),
	)
}
