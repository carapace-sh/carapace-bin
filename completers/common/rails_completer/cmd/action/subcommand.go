package action

import "github.com/carapace-sh/carapace"

// ActionCommands completes top-level Rails commands
func ActionCommands() carapace.Action {
	return carapace.ActionValuesDescribed(
		"new", "Create a new Rails application",
		"generate", "Run a code generator",
		"g", "Alias for generate",
		"destroy", "Undo code generation",
		"d", "Alias for destroy",
		"server", "Start the Rails server",
		"s", "Alias for server",
		"console", "Start the Rails console",
		"c", "Alias for console",
		"runner", "Execute Ruby code in Rails context",
		"r", "Alias for runner",
		"dbconsole", "Start a database console",
		"db", "Alias for dbconsole",
		"test", "Run tests",
		"t", "Alias for test",
		"routes", "List all defined routes",
		"notes", "Search code for annotations",
		"query", "Active Record read-only REPL",
		"devcontainer", "Generate Dev Container setup",
		"plugin", "Create a Rails plugin or engine",
		"about", "Show Rails/Ruby version info",
		"boot", "Boot the application and exit",
		"middleware", "Show Rack middleware stack",
		"initializers", "Print all initializers",
		"restart", "Restart the Rails server",
		"secret", "Generate a secret key",
		"stats", "Show code statistics",
		"version", "Show Rails version",
	).Tag("commands").Uid("rails", "command")
}

// ActionNamespaces completes top-level Rails namespaces
func ActionNamespaces() carapace.Action {
	return carapace.ActionValues(
		"credentials", "db", "encrypted", "assets", "log",
		"tmp", "time", "action_mailbox", "action_text",
		"active_storage", "app", "cache_digests", "css", "dev",
		"importmap", "javascript", "railties", "secrets", "stimulus",
		"tailwindcss", "turbo", "yarn", "zeitwerk",
	).Tag("namespaces").Uid("rails", "namespace")
}
