package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grit",
	Short: "A simple Grit-powered CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "general", Title: "Getting started"},
		&cobra.Group{ID: "change", Title: "Making changes"},
		&cobra.Group{ID: "history", Title: "History"},
		&cobra.Group{ID: "branch", Title: "Branches and tags"},
		&cobra.Group{ID: "remote", Title: "Remotes"},
		&cobra.Group{ID: "maintenance", Title: "Maintenance"},
		&cobra.Group{ID: "plumbing", Title: "Plumbing"},
	)

	rootCmd.Flags().BoolS("V", "V", false, "Print version")
	rootCmd.PersistentFlags().String("filter", "", "jq-like expression applied to JSON output (requires `--json`)")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("json", false, "Emit machine-readable JSON instead of human-readable text")
	rootCmd.PersistentFlags().Bool("markdown", false, "Emit agent-friendly Markdown instead of human-readable text")
	rootCmd.PersistentFlags().String("threads", "", "Worker threads for parallel library work (`0` = auto, `1` = serial)")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")
	rootCmd.Flag("V").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"filter": jq.ActionFilters(),
	})
}
