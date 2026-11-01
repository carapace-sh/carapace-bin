package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// settings are the keys `composer config` is able to modify.
var settings = map[string]string{
	"apcu-autoloader":                   "use APCu to cache found/not-found classes",
	"archive-dir":                       "default destination for archives",
	"archive-format":                    "default archive format",
	"audit.abandoned":                   "behavior on abandoned packages",
	"audit.block-abandoned":             "block abandoned packages",
	"audit.block-insecure":              "block packages with security advisories",
	"audit.ignore":                      "advisories to ignore",
	"audit.ignore-abandoned":            "abandoned packages to ignore",
	"audit.ignore-severity":             "severity levels to ignore",
	"audit.ignore-unreachable":          "ignore unreachable repositories",
	"bin-compat":                        "binary proxy type",
	"bin-dir":                           "directory for binaries",
	"bump-after-update":                 "bump requirements after update",
	"cache-dir":                         "cache directory",
	"cache-files-dir":                   "directory for cached package files",
	"cache-files-maxsize":               "maximum size of the files cache",
	"cache-files-ttl":                   "time to live of cached package files",
	"cache-repo-dir":                    "directory for cached repository metadata",
	"cache-ttl":                         "default time to live of cache entries",
	"cache-vcs-dir":                     "directory for cached VCS repositories",
	"cafile":                            "path to a certificate authority file",
	"capath":                            "path to a certificate authority directory",
	"classmap-authoritative":            "autoload classes from the classmap only",
	"data-dir":                          "data directory",
	"description":                       "package description",
	"disable-tls":                       "disable TLS",
	"discard-changes":                   "how to handle local changes in vendor packages",
	"github-domains":                    "GitHub domains",
	"github-expose-hostname":            "expose the hostname in the GitHub user agent",
	"github-protocols":                  "protocols used for GitHub clones",
	"gitlab-domains":                    "GitLab domains",
	"gitlab-protocol":                   "protocol used for GitLab clones",
	"homepage":                          "package homepage",
	"htaccess-protect":                  "create .htaccess files in the cache directories",
	"keywords":                          "package keywords",
	"license":                           "package license",
	"lock":                              "create a lock file",
	"minimum-stability":                 "minimum stability of dependencies",
	"name":                              "package name",
	"notify-on-install":                 "notify repositories about installs",
	"optimize-autoloader":               "optimize the autoloader",
	"platform-check":                    "platform check of the generated autoloader",
	"policy.abandoned.audit":            "audit behavior for abandoned packages",
	"policy.abandoned.block":            "block abandoned packages",
	"policy.advisories.audit":           "audit behavior for advisories",
	"policy.advisories.block":           "block packages with advisories",
	"policy.advisories.ignore-id":       "advisory ids to ignore",
	"policy.advisories.ignore-severity": "severity levels to ignore",
	"policy.ignore-unreachable":         "ignore unreachable policy sources",
	"policy.malware.audit":              "audit behavior for malware",
	"policy.malware.block":              "block malicious packages",
	"policy.malware.block-scope":        "operations to block malicious packages for",
	"policy.malware.ignore-source":      "sources to ignore",
	"prefer-stable":                     "prefer stable versions of dependencies",
	"preferred-install":                 "default install method",
	"prepend-autoloader":                "prepend the composer autoloader",
	"process-timeout":                   "timeout of processes in seconds",
	"secure-http":                       "require HTTPS for remote requests",
	"sort-packages":                     "sort packages when adding them",
	"source-fallback":                   "fall back to source when dist is not available",
	"store-auths":                       "store authentication",
	"type":                              "package type",
	"update-with-minimal-changes":       "only perform necessary changes on update",
	"use-github-api":                    "use the GitHub API",
	"use-include-path":                  "use the include path in the autoloader",
	"use-parent-dir":                    "use the composer.json of a parent directory",
	"vendor-dir":                        "vendor directory",
	"version":                           "package version",
}

// settingGroups are the keys that take a name as suffix (e.g. `extra.branch-alias`).
var settingGroups = map[string]string{
	"allow-plugins.":   "allowed plugins",
	"bearer.":          "bearer tokens",
	"bitbucket-oauth.": "Bitbucket OAuth credentials",
	"custom-headers.":  "custom headers",
	"extra.":           "extra properties",
	"forgejo-token.":   "Forgejo tokens",
	"github-oauth.":    "GitHub OAuth tokens",
	"gitlab-oauth.":    "GitLab OAuth tokens",
	"gitlab-token.":    "GitLab tokens",
	"http-basic.":      "HTTP basic authentication",
	"platform.":        "platform overrides",
	"repositories.":    "repositories",
	"scripts.":         "scripts",
	"suggest.":         "suggested packages",
}

// flatten collects the dot separated paths of all values within given object.
func flatten(prefix string, object map[string]any, keys map[string]string) {
	for key, value := range object {
		if nested, ok := value.(map[string]any); ok && len(nested) > 0 {
			flatten(prefix+key+".", nested, keys)
		} else {
			keys[prefix+key] = "current setting"
		}
	}
}

// ActionConfigKeys completes the setting keys of `composer config`
func ActionConfigKeys(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		for _, name := range []string{"auth", "editor", "list"} {
			if flag := cmd.Flag(name); flag != nil && flag.Value.String() == "true" {
				return carapace.ActionValues()
			}
		}

		keys := make(map[string]string)
		for key, description := range settings {
			keys[key] = description
		}
		for group, description := range settingGroups {
			keys[group] = description
		}

		if path, err := configFile(c, cmd); err == nil {
			if p, err := readProject(path); err == nil {
				flatten("", p.config(), keys)
				flatten("extra.", p.object("extra"), keys)
				flatten("suggest.", p.object("suggest"), keys)
				for name := range p.object("scripts") {
					keys["scripts."+name] = "current setting"
				}
				for _, repository := range repositories(p) {
					keys["repositories."+repository.name] = repository.description
				}
			}
		}

		vals := make([]string, 0, len(keys)*2)
		for key, description := range keys {
			vals = append(vals, key, description)
		}
		return carapace.ActionValuesDescribed(vals...).NoSpace('.')
	}).Tag("config keys")
}

// ActionConfigValues completes the value(s) of the setting given as first argument
func ActionConfigValues() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Args) == 0 {
			return carapace.ActionValues()
		}
		key, values := c.Args[0], c.Args[1:]

		switch {
		case strings.HasPrefix(key, "repositories."), strings.HasPrefix(key, "repository."), strings.HasPrefix(key, "repos."), strings.HasPrefix(key, "repo."):
			if len(values) == 0 {
				return ActionRepositoryTypes()
			}
			return carapace.ActionValues()
		case strings.HasPrefix(key, "allow-plugins."):
			return carapace.ActionValues("false", "true")
		case strings.HasPrefix(key, "platform."):
			return carapace.ActionValues("false")
		}

		switch key {
		case "apcu-autoloader", "audit.block-abandoned", "audit.block-insecure", "audit.ignore-unreachable",
			"classmap-authoritative", "disable-tls", "github-expose-hostname", "htaccess-protect", "lock",
			"notify-on-install", "optimize-autoloader", "policy.abandoned.block", "policy.advisories.block",
			"policy.malware.block", "prefer-stable", "prepend-autoloader", "secure-http", "sort-packages",
			"source-fallback", "update-with-minimal-changes", "use-github-api", "use-include-path":
			return carapace.ActionValues("false", "true")
		case "audit.abandoned", "policy.abandoned.audit", "policy.advisories.audit", "policy.malware.audit":
			return carapace.ActionValues("fail", "ignore", "report")
		case "audit.ignore-severity", "policy.advisories.ignore-severity":
			return carapace.ActionValues("critical", "high", "low", "medium").Filter(values...)
		case "bin-compat":
			return carapace.ActionValues("auto", "full", "proxy", "symlink")
		case "bump-after-update":
			return carapace.ActionValues("dev", "false", "no-dev", "true")
		case "discard-changes":
			return carapace.ActionValues("false", "stash", "true")
		case "github-protocols":
			return carapace.ActionValues("git", "https", "ssh").Filter(values...)
		case "gitlab-protocol":
			return carapace.ActionValues("git", "http", "https")
		case "minimum-stability":
			return carapace.ActionValues("RC", "alpha", "beta", "dev", "stable")
		case "platform-check":
			return carapace.ActionValues("false", "php-only", "true")
		case "policy.malware.block-scope":
			return carapace.ActionValues("all", "install", "update")
		case "preferred-install":
			return carapace.ActionValues("auto", "dist", "source")
		case "store-auths", "use-parent-dir":
			return carapace.ActionValues("false", "prompt", "true")
		case "type":
			return carapace.ActionValues("composer-plugin", "library", "metapackage", "project")
		case "archive-dir", "bin-dir", "cache-dir", "cache-files-dir", "cache-repo-dir", "cache-vcs-dir", "capath", "data-dir", "vendor-dir":
			return carapace.ActionDirectories()
		case "cafile":
			return carapace.ActionFiles()
		}
		return carapace.ActionValues()
	})
}
