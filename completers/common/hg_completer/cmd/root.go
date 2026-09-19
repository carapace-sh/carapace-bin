package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hg",
	Short: "Mercurial distributed SCM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

const (
	group_repository_creation = iota
	group_remote_repository_management
	group_change_creation
	group_change_manipulation
	group_change_organization
	group_file_content_management
	group_change_navigation
	group_working_directory_management
	group_change_import_export
	group_repository_maintenance
	group_help
)

var groups = []*cobra.Group{
	{ID: "creation", Title: "Repository creation"},
	{ID: "remote", Title: "Remote repository management"},
	{ID: "change-creation", Title: "Change creation"},
	{ID: "change-manipulation", Title: "Change manipulation"},
	{ID: "change-organization", Title: "Change organization"},
	{ID: "content", Title: "File content management"},
	{ID: "navigation", Title: "Change navigation"},
	{ID: "working-dir", Title: "Working directory management"},
	{ID: "import-export", Title: "Change import/export"},
	{ID: "maintenance", Title: "Repository maintenance"},
	{ID: "help", Title: "Help"},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddGroup(groups...)

	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("color", "", "when to colorize (boolean, always, auto, never, or debug)")
	rootCmd.Flags().String("config", "", "set/override config option (use 'section.name=value')")
	rootCmd.Flags().String("config-file", "", "load config file to set/override config options")
	rootCmd.Flags().String("cwd", "", "change working directory")
	rootCmd.Flags().Bool("debug", false, "enable debugging output")
	rootCmd.Flags().Bool("debugger", false, "start debugger")
	rootCmd.Flags().String("encoding", "", "set the charset encoding")
	rootCmd.Flags().String("encodingmode", "", "set the charset encoding mode")
	rootCmd.Flags().BoolP("help", "h", false, "display help and exit")
	rootCmd.Flags().Bool("hidden", false, "consider hidden changesets")
	rootCmd.Flags().BoolP("noninteractive", "y", false, "do not prompt, automatically pick the first choice for all prompts")
	rootCmd.Flags().String("pager", "", "when to paginate (boolean, always, auto, or never)")
	rootCmd.Flags().Bool("profile", false, "print command execution profile")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	rootCmd.Flags().StringP("repository", "R", "", "repository root directory or name of overlay bundle file")
	rootCmd.Flags().Bool("time", false, "time how long the command takes")
	rootCmd.Flags().Bool("traceback", false, "always print a traceback on exception")
	rootCmd.Flags().BoolP("verbose", "v", false, "enable additional output")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("true", "false", "always", "auto", "never", "debug"),
		"config": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return hg.ActionConfigKeys()
			default:
				return carapace.ActionValues()
			}
		}),
		"config-file": carapace.ActionFiles(),
		"cwd":         carapace.ActionDirectories(),
		"pager":       carapace.ActionValues("true", "false", "always", "auto", "never"),
		"repository":  carapace.ActionFiles(),
	})
}
