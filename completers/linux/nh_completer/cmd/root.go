package cmd

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nh",
	Short: "Nix CLI helper",
	Long:  "https://github.com/nix-community/nh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().Bool("verbose", false, "Show verbose output")
	rootCmd.Flags().Bool("no-nom", false, "Disable nom output")
	rootCmd.Flags().Bool("no-color", false, "Disable colored output")

	rootCmd.AddCommand(
		newOsCmd(),
		newHomeCmd(),
		newCleanCmd(),
		newSearchCmd(),
		newCompletionsCmd(),
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"os", "Manage NixOS configurations",
			"home", "Manage Home Manager configurations",
			"clean", "Clean old profiles and store paths",
			"search", "Search packages",
			"completions", "Generate shell completions",
		).StyleF(style.ForKeyword),
	)
}

func addBuildFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("file", "f", "", "Path to flake")
	cmd.Flags().String("hostname", "", "Hostname/configuration to use")
	cmd.Flags().String("specialisation", "", "NixOS specialisation to use")
	cmd.Flags().Bool("ask", false, "Ask for confirmation")
	cmd.Flags().Bool("dry", false, "Show what would be done")
	cmd.Flags().Bool("update", false, "Update flake inputs")
	cmd.Flags().StringSlice("update-input", nil, "Update a specific flake input")
	cmd.Flags().StringSlice("override-input", nil, "Override a flake input")
	cmd.Flags().StringSlice("extra-args", nil, "Extra arguments passed to nix")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"file":           carapace.ActionFiles(),
		"extra-args":     carapace.ActionValues(),
		"hostname":       actionNixConfigurations(),
		"specialisation": carapace.ActionValues(),
	})
}

func newOsCmd() *cobra.Command {
	osCmd := &cobra.Command{Use: "os", Short: "Manage NixOS configurations", Run: func(cmd *cobra.Command, args []string) {}}
	for _, name := range []string{"switch", "boot", "test", "build"} {
		cmd := &cobra.Command{Use: name, Short: name + " NixOS configuration", Run: func(cmd *cobra.Command, args []string) {}}
		addBuildFlags(cmd)
		osCmd.AddCommand(cmd)
	}
	return osCmd
}

func newHomeCmd() *cobra.Command {
	homeCmd := &cobra.Command{Use: "home", Short: "Manage Home Manager configurations", Run: func(cmd *cobra.Command, args []string) {}}
	for _, name := range []string{"switch", "build"} {
		cmd := &cobra.Command{Use: name, Short: name + " Home Manager configuration", Run: func(cmd *cobra.Command, args []string) {}}
		addBuildFlags(cmd)
		homeCmd.AddCommand(cmd)
	}
	return homeCmd
}

func newCleanCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "clean", Short: "Clean old profiles and store paths", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().Bool("all", false, "Remove all generations")
	cmd.Flags().String("keep", "", "Keep matching generations")
	cmd.Flags().String("keep-since", "", "Keep generations newer than duration")
	cmd.Flags().String("profile", "", "Clean a specific profile")
	cmd.Flags().Bool("ask", false, "Ask for confirmation")
	cmd.Flags().Bool("dry", false, "Show what would be done")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{"profile": carapace.ActionFiles()})
	return cmd
}

func newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "search", Short: "Search packages", Run: func(cmd *cobra.Command, args []string) {}}
	cmd.Flags().String("channel", "", "Channel to search")
	cmd.Flags().Bool("json", false, "Print JSON output")
	cmd.Flags().Bool("no-description", false, "Hide package descriptions")
	carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionValues())
	return cmd
}

func newCompletionsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "completions", Short: "Generate shell completions", Run: func(cmd *cobra.Command, args []string) {}}
	carapace.Gen(cmd).PositionalCompletion(carapace.ActionValues("bash", "elvish", "fish", "powershell", "zsh"))
	return cmd
}

func actionNixConfigurations() carapace.Action {
	return carapace.ActionExecCommand("nix", "eval", "--json", ".#nixosConfigurations", "--apply", "builtins.attrNames")(func(output []byte) carapace.Action {
		var names []string
		if err := json.Unmarshal(output, &names); err != nil {
			return carapace.ActionValues()
		}
		return carapace.ActionValues(names...)
	})
}
