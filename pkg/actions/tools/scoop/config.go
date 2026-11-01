package scoop

import "github.com/carapace-sh/carapace"

// ActionConfigKeys completes configuration setting names
//
//	debug (Additional and detailed output)
//	use_sqlite_cache (Use SQLite database for caching)
func ActionConfigKeys() carapace.Action {
	return carapace.ActionValuesDescribed(
		"aria2-enabled", "Use aria2c for downloading of artifacts",
		"aria2-fallback-enabled", "Fall back to the default downloader when aria2 fails",
		"aria2-max-connection-per-server", "Max connections per server for aria2",
		"aria2-min-split-size", "Min split size for aria2",
		"aria2-options", "Additional aria2 options",
		"aria2-retry-wait", "Seconds between retries for aria2",
		"aria2-split", "Number of connections for aria2",
		"aria2-warning-enabled", "Show/hide aria2 warning",
		"autostash_on_conflict", "Auto-stash uncommitted changes during update",
		"cache_path", "Path for downloads cache",
		"cat_style", "Style for bat when displaying manifests",
		"debug", "Additional and detailed output",
		"default_architecture", "Preferred architecture (64bit/32bit/arm64)",
		"force_update", "Force apps updating to bucket's version",
		"gh_token", "GitHub API token for authenticated requests",
		"global_path", "Path to Scoop global apps directory",
		"hold_update_until", "Hold Scoop updates until a date",
		"ignore_running_processes", "Continue if target app process is running",
		"no_junction", "Don't use 'current' version alias",
		"private_hosts", "Array of private hosts needing additional authentication",
		"proxy", "Proxy settings",
		"root_path", "Path to Scoop root directory",
		"scoop_branch", "Use different branch than master",
		"scoop_repo", "Git repository containing scoop source code",
		"shim", "Choose scoop shim build (kiennq|scoopcs|71)",
		"show_manifest", "Display manifest before install",
		"show_update_log", "Show/hide changed commits on update",
		"update_nightly", "Auto-update nightly versions",
		"use_external_7zip", "Use external 7zip for archives extraction",
		"use_git_history", "Search bucket git history for exact versions when installing",
		"use_isolated_path", "Isolate Scoop from system PATH",
		"use_lessmsi", "Prefer lessmsi utility over native msiexec",
		"use_sqlite_cache", "Use SQLite database for caching",
		"virustotal_api_key", "VirusTotal API key",
	).Tag("config keys")
}
