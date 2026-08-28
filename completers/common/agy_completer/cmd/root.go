package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "agy",
	Short: "Google Antigravity AI Agent command-line interface",
	Long:  "https://antigravity.google/docs/cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(&cobra.Group{ID: "engine", Title: "Engine"})
	rootCmd.AddGroup(&cobra.Group{ID: "integration", Title: "Integration"})
	rootCmd.AddGroup(&cobra.Group{ID: "configuration", Title: "Configuration"})

	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().StringArray("add-dir", []string{}, "Add a directory to the workspace")
	rootCmd.Flags().String("agent", "", "Agent for the current CLI session")
	rootCmd.Flags().BoolP("continue", "c", false, "Continue the most recent conversation")
	rootCmd.Flags().String("conversation", "", "Resume a previous conversation by ID")
	rootCmd.Flags().Bool("dangerously-skip-permissions", false, "Auto-approve all tool permission requests without prompting")
	rootCmd.Flags().Bool("disable-slash-commands", false, "Disable slash command and skill expansion in print mode")
	rootCmd.Flags().String("effort", "", "Reasoning effort for the current CLI session (low|medium|high)")
	rootCmd.Flags().String("input-format", "text", "Input format for print mode (text, stream-json). stream-json reads one NDJSON message per line from stdin and runs a turn for each; it requires --output-format stream-json")
	rootCmd.Flags().String("json-schema", "", "Optional JSON schema string or path to a schema file to enforce structured output (for stream-json, only applicable to the final result)")
	rootCmd.Flags().String("log-file", "", "Override CLI log file path")
	rootCmd.Flags().String("mode", "", "Set the agent execution mode for this session (accept-edits, plan)")
	rootCmd.Flags().String("model", "", "Model for the current CLI session")
	rootCmd.Flags().Bool("new-project", false, "Create a new project for this session")
	rootCmd.Flags().String("output-format", "text", "Output format for print mode (text, json, stream-json)")
	rootCmd.Flags().StringP("print", "p", "", "Run a single prompt non-interactively")
	rootCmd.Flags().String("print-timeout", "5m0s", "Timeout for print mode wait")
	rootCmd.Flags().String("project", "", "Project ID or project name for the current CLI session")
	rootCmd.Flags().String("prompt", "", "Alias for --print")
	rootCmd.Flags().StringP("prompt-interactive", "i", "", "Run an initial prompt interactively and continue the session")
	rootCmd.Flags().Bool("sandbox", false, "Run in a sandbox with terminal restrictions enabled")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-dir": carapace.ActionDirectories(),
		"agent": carapace.ActionExecCommand("agy", "agents")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines {
				if line != "" {
					parts := strings.Fields(line)
					if len(parts) > 0 {
						vals = append(vals, parts[0])
					}
				}
			}
			return carapace.ActionValues(vals...)
		}),
		"effort":       carapace.ActionValues("low", "medium", "high"),
		"input-format": carapace.ActionValues("text", "stream-json"),
		"json-schema":  carapace.ActionFiles(".json"),
		"log-file":     carapace.ActionFiles(),
		"mode":         carapace.ActionValues("accept-edits", "plan"),
		"model": carapace.ActionExecCommand("agy", "models")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines {
				if line == "" {
					continue
				}
				parts := strings.SplitN(line, "\t", 2)
				if len(parts) == 2 {
					vals = append(vals, parts[0], parts[1])
				} else {
					vals = append(vals, parts[0], "")
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}),
		"output-format": carapace.ActionValues("text", "json", "stream-json"),
		"print-timeout": carapace.ActionValues("30s", "1m", "10m", "1h").Usage("duration (e.g., 5m, 1h30m)"),
	})
}
