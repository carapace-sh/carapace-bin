package action

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
)

// scoopExport is the JSON structure output by `scoop export`.
type scoopExport struct {
	Apps    []scoopApp    `json:"apps"`
	Buckets []scoopBucket `json:"buckets"`
}

type scoopApp struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
	Source  string `json:"Source"`
}

type scoopBucket struct {
	Name      string `json:"Name"`
	Source    string `json:"Source"`
	Manifests int    `json:"Manifests"`
}

// ActionInstalledApps completes installed app names with their version as description.
func ActionInstalledApps() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop export")(func(output []byte) carapace.Action {
		var export scoopExport
		if err := json.Unmarshal(output, &export); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, app := range export.Apps {
			vals = append(vals, app.Name, app.Version)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionInstalledBuckets completes installed bucket names with their source as description.
func ActionInstalledBuckets() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop export")(func(output []byte) carapace.Action {
		var export scoopExport
		if err := json.Unmarshal(output, &export); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, bucket := range export.Buckets {
			vals = append(vals, bucket.Name, bucket.Source)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionAvailableApps completes app names available in local buckets.
func ActionAvailableApps() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "$d = if ($env:SCOOP) { $env:SCOOP } else { \"$env:USERPROFILE\\scoop\" }; Get-ChildItem -Path \"$d\\buckets\\*\\bucket\\*.json\" | Select-Object -ExpandProperty BaseName | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var apps []string
		if err := json.Unmarshal(output, &apps); err != nil {
			var app string
			if err2 := json.Unmarshal(output, &app); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			apps = []string{app}
		}
		return carapace.ActionValues(apps...).UniqueList(",")
	})
}

// ActionKnownBuckets completes known bucket names from the static buckets.json list.
func ActionKnownBuckets() carapace.Action {
	return carapace.ActionValuesDescribed(
		"main", "https://github.com/ScoopInstaller/Main",
		"extras", "https://github.com/ScoopInstaller/Extras",
		"versions", "https://github.com/ScoopInstaller/Versions",
		"nirsoft", "https://github.com/ScoopInstaller/Nirsoft",
		"sysinternals", "https://github.com/niheaven/scoop-sysinternals",
		"php", "https://github.com/ScoopInstaller/PHP",
		"nerd-fonts", "https://github.com/matthewjberger/scoop-nerd-fonts",
		"nonportable", "https://github.com/ScoopInstaller/Nonportable",
		"java", "https://github.com/ScoopInstaller/Java",
		"games", "https://github.com/Calinou/scoop-games",
	)
}

// ActionConfigKeys completes configuration setting names.
func ActionConfigKeys() carapace.Action {
	return carapace.ActionValuesDescribed(
		"aria2-enabled", "Use aria2c for downloading of artifacts",
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
		"use_isolated_path", "Isolate Scoop from system PATH",
		"use_lessmsi", "Prefer lessmsi utility over native msiexec",
		"use_sqlite_cache", "Use SQLite database for caching",
		"virustotal_api_key", "VirusTotal API key",
	)
}

// scoopShimInfo is the JSON structure for shim info output.
type scoopShimInfo struct {
	Name   string `json:"Name"`
	Source string `json:"Source"`
}

// ActionShims completes installed shim names with their source as description.
func ActionShims() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop shim list | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var shims []scoopShimInfo
		if err := json.Unmarshal(output, &shims); err != nil {
			var shim scoopShimInfo
			if err2 := json.Unmarshal(output, &shim); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValuesDescribed(shim.Name, shim.Source)
		}
		vals := make([]string, 0)
		for _, shim := range shims {
			vals = append(vals, shim.Name, shim.Source)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// scoopCacheInfo is the JSON structure for cache info output.
type scoopCacheInfo struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
}

// ActionCachedApps completes cached app names with their version as description.
func ActionCachedApps() carapace.Action {
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop cache show | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var cache []scoopCacheInfo
		if err := json.Unmarshal(output, &cache); err != nil {
			var item scoopCacheInfo
			if err2 := json.Unmarshal(output, &item); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValuesDescribed(item.Name, item.Version)
		}
		vals := make([]string, 0)
		for _, item := range cache {
			vals = append(vals, item.Name, item.Version)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
