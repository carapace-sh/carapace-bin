package action

import "github.com/carapace-sh/carapace"

// ActionCommands completes the composer commands
func ActionCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"about", "Shows a short information about Composer",
		"archive", "Creates an archive of this composer package",
		"audit", "Checks for security vulnerability advisories for installed packages",
		"browse", "Opens the package's repository URL or homepage in your browser",
		"bump", "Increases the lower limit of your composer.json requirements to the currently installed versions",
		"check-platform-reqs", "Check that platform requirements are satisfied",
		"clear-cache", "Clears composer's internal package cache",
		"completion", "Dump the shell completion script",
		"config", "Sets config options",
		"create-project", "Creates new project from a package into given directory",
		"depends", "Shows which packages cause the given package to be installed",
		"diagnose", "Diagnoses the system to identify common errors",
		"dump-autoload", "Dumps the autoloader",
		"exec", "Executes a vendored binary/script",
		"fund", "Discover how to help fund the maintenance of your dependencies",
		"global", "Allows running commands in the global composer dir ($COMPOSER_HOME)",
		"help", "Display help for a command",
		"init", "Creates a basic composer.json file in current directory",
		"install", "Installs the project dependencies from the composer.lock file if present, or falls back on the composer.json",
		"licenses", "Shows information about licenses of dependencies",
		"list", "List commands",
		"outdated", "Shows a list of installed packages that have updates available, including their latest version",
		"policy", "Manages custom dependency policies and their sources",
		"prohibits", "Shows which packages prevent the given package from being installed",
		"reinstall", "Uninstalls and reinstalls the given package names",
		"remove", "Removes a package from the require or require-dev",
		"repository", "Manages repositories",
		"require", "Adds required packages to your composer.json and installs them",
		"run-script", "Runs the scripts defined in composer.json",
		"search", "Searches for packages",
		"self-update", "Updates composer.phar to the latest version",
		"show", "Shows information about packages",
		"status", "Shows a list of locally modified packages",
		"suggests", "Shows package suggestions",
		"update", "Updates your dependencies to the latest version according to composer.json, and updates the composer.lock file",
		"validate", "Validates a composer.json and composer.lock",
	).Tag("commands")
}
