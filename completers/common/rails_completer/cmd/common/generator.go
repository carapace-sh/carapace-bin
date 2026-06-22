package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rails"
	"github.com/spf13/cobra"
)

func AddGeneratorFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing files")
	cmd.Flags().Bool("no-color", false, "Disable color output")
	cmd.Flags().BoolP("pretend", "p", false, "Run but make no changes")
	cmd.Flags().BoolP("quiet", "q", false, "Suppress status output")
	cmd.Flags().BoolP("skip", "s", false, "Skip existing files")
	cmd.Flags().Bool("skip-collision-check", false, "Skip file collision check")
	cmd.Flags().Bool("skip-namespace", false, "Skip namespace in file paths")
}

func AddAppGeneratorFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("api", false, "Generate API-only")
	cmd.Flags().StringP("css", "c", "tailwind", "CSS processor")
	cmd.Flags().StringP("database", "d", "sqlite3", "Preconfigure for selected database")
	cmd.Flags().Bool("dev", false, "Use local Rails")
	cmd.Flags().Bool("devcontainer", false, "Add Dev Container")
	cmd.Flags().Bool("edge", false, "Use edge Rails")
	cmd.Flags().StringP("javascript", "j", "importmap", "JavaScript bundler")
	cmd.Flags().Bool("main", false, "Use main branch")
	cmd.Flags().Bool("master", false, "Alias for --main")
	cmd.Flags().Bool("minimal", false, "Generate minimal")
	cmd.Flags().Bool("no-rc", false, "Skip .railsrc")
	cmd.Flags().StringP("name", "n", "", "Module name")
	cmd.Flags().String("rc", "", "Custom .railsrc path")
	cmd.Flags().StringP("ruby", "r", "", "Path to the Ruby binary")
	cmd.Flags().Bool("skip-action-cable", false, "Skip Action Cable")
	cmd.Flags().BoolP("skip-action-mailer", "M", false, "Skip Action Mailer")
	cmd.Flags().Bool("skip-action-mailbox", false, "Skip Action Mailbox")
	cmd.Flags().Bool("skip-action-text", false, "Skip Action Text")
	cmd.Flags().Bool("skip-active-job", false, "Skip Active Job")
	cmd.Flags().BoolP("skip-active-record", "O", false, "Skip Active Record")
	cmd.Flags().Bool("skip-active-storage", false, "Skip Active Storage")
	cmd.Flags().BoolP("skip-asset-pipeline", "A", false, "Skip asset pipeline")
	cmd.Flags().Bool("skip-bootsnap", false, "Skip Bootsnap")
	cmd.Flags().Bool("skip-brakeman", false, "Skip Brakeman")
	cmd.Flags().BoolP("skip-bundle", "B", false, "Skip bundle install")
	cmd.Flags().Bool("skip-ci", false, "Skip GitHub Actions")
	cmd.Flags().Bool("skip-collision-check", false, "Skip collision check")
	cmd.Flags().Bool("skip-decrypted-diffs", false, "Skip credential diffs")
	cmd.Flags().Bool("skip-dev-gems", false, "Skip development gems")
	cmd.Flags().Bool("skip-docker", false, "Skip Docker")
	cmd.Flags().BoolP("skip-git", "G", false, "Skip git init")
	cmd.Flags().Bool("skip-hotwire", false, "Skip Hotwire")
	cmd.Flags().BoolP("skip-javascript", "J", false, "Skip JavaScript")
	cmd.Flags().Bool("skip-jbuilder", false, "Skip Jbuilder")
	cmd.Flags().Bool("skip-kamal", false, "Skip Kamal")
	cmd.Flags().Bool("skip-keeps", false, "Skip .keep files")
	cmd.Flags().Bool("skip-namespace", false, "Skip namespacing")
	cmd.Flags().Bool("skip-rubocop", false, "Skip RuboCop")
	cmd.Flags().Bool("skip-solid", false, "Skip Solid")
	cmd.Flags().Bool("skip-system-test", false, "Skip system tests")
	cmd.Flags().BoolP("skip-test", "T", false, "Skip test framework")
	cmd.Flags().Bool("skip-thruster", false, "Skip Thruster")
	cmd.Flags().StringP("template", "m", "", "Path or URL to template")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"css":        rails.ActionCSSOptions(),
		"database":   rails.ActionDatabases(),
		"javascript": rails.ActionJavaScriptOptions(),
		"rc":         carapace.ActionFiles(),
		"template":   carapace.ActionFiles(),
	})
}
