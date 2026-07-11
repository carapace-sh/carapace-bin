package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var skill_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List installed skills (preview)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_listCmd).Standalone()

	skill_listCmd.Flags().String("agent", "", "Filter by target agent: {github-copilot|claude-code|cursor|codex|gemini-cli|antigravity|adal|amp|augment|bob|cline|codebuddy|command-code|continue|cortex|crush|deepagents|droid|firebender|goose|iflow-cli|junie|kilo|kimi-cli|kiro-cli|kode|mcpjam|mistral-vibe|mux|neovate|openclaw|opencode|openhands|pi|pochi|qoder|qwen-code|replit|roo|trae|trae-cn|universal|warp|windsurf|zencoder}")
	skill_listCmd.Flags().String("dir", "", "Scan a custom directory for installed skills")
	skill_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	skill_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	skill_listCmd.Flags().String("scope", "", "Filter by installation scope: {project|user}")
	skill_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	skillCmd.AddCommand(skill_listCmd)

	carapace.Gen(skill_listCmd).FlagCompletion(carapace.ActionMap{
		"agent": gh.ActionAgents(),
		"dir":   carapace.ActionDirectories(),
		"jq":    jq.ActionFilters(),
		"json":  gh.ActionSkillListFields().UniqueList(","),
		"scope": carapace.ActionValues("project", "user"),
	})
}
