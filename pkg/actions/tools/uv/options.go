package uv

import "github.com/carapace-sh/carapace"

// ActionKeyringProviders completes keyring providers
//
//	disabled (Do not use keyring for credential lookup)
//	subprocess (Use the `keyring` command for credential lookup)
func ActionKeyringProviders() carapace.Action {
	return carapace.ActionValuesDescribed(
		"disabled", "Do not use keyring for credential lookup",
		"subprocess", "Use the `keyring` command for credential lookup",
	)
}

// ActionForkStrategies completes fork strategies
//
//	fewest (Optimize for selecting the fewest number of versions for each package.)
//	requires-python (Optimize for selecting latest supported version of each package)
func ActionForkStrategies() carapace.Action {
	return carapace.ActionValuesDescribed(
		"fewest", "Optimize for selecting the fewest number of versions for each package. Older versions may be preferred if they are compatible with a wider range of supported Python versions or platforms",
		"requires-python", "Optimize for selecting latest supported version of each package, for each supported Python version",
	)
}

// ActionIndexStrategies completes index strategies
//
//	first-index (Only use results from the first index that returns a match for a given package name)
//	unsafe-first-match (Search for every package name across all indexes)
//	unsafe-best-match (Search for every package name across all indexes, preferring the "best" version found)
func ActionIndexStrategies() carapace.Action {
	return carapace.ActionValuesDescribed(
		"first-index", "Only use results from the first index that returns a match for a given package name",
		"unsafe-first-match", "Search for every package name across all indexes, exhausting the versions from the first index before moving on to the next",
		"unsafe-best-match", "Search for every package name across all indexes, preferring the \"best\" version found. If a package version is in multiple indexes, only look at the entry for the first index",
	)
}

// ActionLinkModes completes link modes
//
//	clone (Clone (i.e., copy-on-write) packages from the source into the destination)
//	copy (Copy packages from the source into the destination)
func ActionLinkModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"clone", "Clone (i.e., copy-on-write) packages from the source into the destination",
		"copy", "Copy packages from the source into the destination",
		"hardlink", "Hard link packages from the source into the destination",
		"symlink", "Symbolically link packages from the source into the destination",
	)
}

// ActionPreReleases completes pre-release strategies
//
//	disallow (Disallow all pre-release versions)
//	allow (Allow all pre-release versions)
//	if-necessary (Prefer stable versions, falling back to pre-release versions when necessary)
func ActionPreReleases() carapace.Action {
	return carapace.ActionValuesDescribed(
		"disallow", "Disallow all pre-release versions",
		"allow", "Allow all pre-release versions",
		"if-necessary", "Prefer stable versions, falling back to pre-release versions when necessary",
		"explicit", "Prefer stable versions for first-party packages with explicit pre-release specifiers, falling back to pre-release versions when necessary. Disallow pre-release versions for all other packages",
		"if-necessary-or-explicit", "Deprecated alias for `if-necessary`",
	)
}

// ActionResolutions completes resolution strategies
//
//	highest (Resolve the highest compatible version of each package)
//	lowest (Resolve the lowest compatible version of each package)
//	lowest-direct (Resolve the lowest compatible version of any direct dependencies)
func ActionResolutions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"highest", "Resolve the highest compatible version of each package",
		"lowest", "Resolve the lowest compatible version of each package",
		"lowest-direct", "Resolve the lowest compatible version of any direct dependencies, and the highest compatible version of any transitive dependencies",
	)
}
