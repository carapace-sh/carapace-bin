package action

import "github.com/carapace-sh/carapace"

// ActionCommands completes top-level Rails commands
func ActionCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"about", "Show Rails/Ruby version info",
		"boot", "Boot the application and exit",
		"c", "Alias for console",
		"console", "Start the Rails console",
		"d", "Alias for destroy",
		"db", "Alias for dbconsole",
		"dbconsole", "Start a database console",
		"destroy", "Undo code generation",
		"devcontainer", "Generate Dev Container setup",
		"g", "Alias for generate",
		"generate", "Run a code generator",
		"initializers", "Print all initializers",
		"middleware", "Show Rack middleware stack",
		"new", "Create a new Rails application",
		"notes", "Search code for annotations",
		"plugin", "Create a Rails plugin or engine",
		"query", "Active Record read-only REPL",
		"r", "Alias for runner",
		"restart", "Restart the Rails server",
		"routes", "List all defined routes",
		"runner", "Execute Ruby code in Rails context",
		"s", "Alias for server",
		"secret", "Generate a secret key",
		"server", "Start the Rails server",
		"stats", "Show code statistics",
		"t", "Alias for test",
		"test", "Run tests",
		"version", "Show Rails version",
	).Tag("commands").Uid("rails", "command")
}

// ActionNamespaces completes top-level Rails namespaces
func ActionNamespaces() carapace.Action {
	return carapace.ActionValues(
		"action_mailbox", "action_text", "active_storage", "app", "assets",
		"cache_digests", "credentials", "css", "db",
		"dev", "encrypted", "importmap", "javascript", "log",
		"railties", "secrets", "stimulus", "tailwindcss", "time",
		"tmp", "turbo", "yarn", "zeitwerk",
	).Tag("namespaces").Uid("rails", "namespace")
}
