package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/meson_completer/cmd/action"
)

func init() {
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"compile", "Build the project",
			"configure", "Change project options",
			"devenv", "Run commands in developer environment",
			"dist", "Generate release archive",
			"env2mfile", "Convert current environment to a cross or native file",
			"fmt", "Format meson source file",
			"format", "Format meson source file",
			"help", "Print help of a subcommand",
			"init", "Create a new project",
			"install", "Install the project",
			"introspect", "Introspect project",
			"reprotest", "Test if project builds reproducibly",
			"rewrite", "Modify the project definition",
			"setup", "Configure the project",
			"subprojects", "Manage subprojects",
			"test", "Run tests",
			"wrap", "Wrap tools",
		),
	)

	carapace.Gen(setupCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)

	carapace.Gen(configureCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(compileCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(distCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(introspectCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(subprojectsCmd).PositionalCompletion(
		action.ActionSubprojects(),
	)

	carapace.Gen(devenvCmd).PositionalAnyCompletion(
		carapace.ActionExecutables(),
	)

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionValues(
			"compile",
			"configure",
			"devenv",
			"dist",
			"env2mfile",
			"fmt",
			"format",
			"init",
			"install",
			"introspect",
			"reprotest",
			"rewrite",
			"setup",
			"subprojects",
			"test",
			"wrap",
		),
	)

	carapace.Gen(wrap_installCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(wrap_promoteCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)

	carapace.Gen(wrap_infoCmd).PositionalCompletion(
		carapace.ActionFiles(".wrap"),
	)

	carapace.Gen(wrap_statusCmd).PositionalCompletion(
		carapace.ActionFiles(".wrap"),
	)

	carapace.Gen(wrap_updateCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".wrap"),
	)
}
