package action

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionBackends completes build backends
func ActionBackends() carapace.Action {
	return carapace.ActionValues(
		"ninja",
		"none",
		"vs",
		"vs2010",
		"vs2012",
		"vs2013",
		"vs2015",
		"vs2017",
		"vs2019",
		"vs2022",
		"vs2026",
		"xcode",
	).Tag("backends")
}

// ActionBuildTypes completes build types
func ActionBuildTypes() carapace.Action {
	return carapace.ActionValues(
		"custom",
		"debug",
		"debugoptimized",
		"minsize",
		"plain",
		"release",
	).Tag("build types")
}

// ActionWrapModes completes wrap modes
func ActionWrapModes() carapace.Action {
	return carapace.ActionValues(
		"default",
		"forcefallback",
		"nodownload",
		"nofallback",
		"nopromote",
	).Tag("wrap modes")
}

// ActionLayouts completes build directory layouts
func ActionLayouts() carapace.Action {
	return carapace.ActionValues("flat", "mirror").Tag("layouts")
}

// ActionNamingSchemes completes target naming schemes
func ActionNamingSchemes() carapace.Action {
	return carapace.ActionValues("classic", "platform").Tag("naming schemes")
}

// ActionOptimizations completes optimization levels
func ActionOptimizations() carapace.Action {
	return carapace.ActionValues("0", "1", "2", "3", "g", "plain", "s").Tag("optimization levels")
}

// ActionUnityBuilds completes unity build modes
func ActionUnityBuilds() carapace.Action {
	return carapace.ActionValues("off", "on", "subprojects").Tag("unity modes")
}

// ActionWarningLevels completes warning levels
func ActionWarningLevels() carapace.Action {
	return carapace.ActionValues("0", "1", "2", "3", "everything").Tag("warning levels")
}

// ActionDefaultLibraries completes default library types
func ActionDefaultLibraries() carapace.Action {
	return carapace.ActionValues("both", "shared", "static").Tag("library types")
}

// ActionBothLibraries completes default both_libraries types
func ActionBothLibraries() carapace.Action {
	return carapace.ActionValues("auto", "shared", "static").Tag("library types")
}

// ActionGenvslite completes genvslite backends
func ActionGenvslite() carapace.Action {
	return carapace.ActionValues("vs2022", "vs2026").Tag("backends")
}

// ActionDistFormats completes archive formats
func ActionDistFormats() carapace.Action {
	return carapace.ActionValues("bztar", "gztar", "xztar", "zip").Tag("formats")
}

// ActionLanguages completes project languages
func ActionLanguages() carapace.Action {
	return carapace.ActionValues(
		"c",
		"cpp",
		"cs",
		"cuda",
		"d",
		"fortran",
		"java",
		"objc",
		"objcpp",
		"rust",
		"vala",
	).Tag("languages")
}

// ActionProjectTypes completes project types
func ActionProjectTypes() carapace.Action {
	return carapace.ActionValues("executable", "library").Tag("project types")
}

// ActionSubprojects completes subproject names from the wrap files
func ActionSubprojects() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		matches, _ := filepath.Glob(filepath.Join(c.Value, "subprojects", "*.wrap"))
		if len(matches) == 0 {
			matches, _ = filepath.Glob("subprojects/*.wrap")
		}

		names := make([]string, 0, len(matches))
		for _, match := range matches {
			name := strings.TrimSuffix(filepath.Base(match), ".wrap")
			if name == "packagefiles" || strings.HasSuffix(name, ".txt") {
				continue
			}
			names = append(names, name)
		}
		return carapace.ActionValues(names...)
	}).Tag("subprojects")
}
