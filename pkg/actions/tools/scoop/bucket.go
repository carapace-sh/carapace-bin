package scoop

import (
	"bytes"
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type scoopBucket struct {
	Name      string `json:"Name"`
	Source    string `json:"Source"`
	Manifests int    `json:"Manifests"`
}

// ActionInstalledBuckets completes installed bucket names with their source as description
//
//	extras (https://github.com/ScoopInstaller/Extras)
//	main (https://github.com/ScoopInstaller/Main)
func ActionInstalledBuckets() carapace.Action {
	return carapace.ActionExecCommand("scoop", "export")(func(output []byte) carapace.Action {
		var export scoopExport
		if err := json.Unmarshal(bytes.TrimPrefix(output, []byte{0xEF, 0xBB, 0xBF}), &export); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, bucket := range export.Buckets {
			vals = append(vals, bucket.Name, bucket.Source)
		}
		return carapace.ActionValuesDescribed(vals...).Tag("buckets")
	})
}

// ActionKnownBuckets completes known bucket names with their source as description
//
//	extras (https://github.com/ScoopInstaller/Extras)
//	main (https://github.com/ScoopInstaller/Main)
func ActionKnownBuckets() carapace.Action {
	return carapace.ActionValuesDescribed(
		"extras", "https://github.com/ScoopInstaller/Extras",
		"games", "https://github.com/ScoopInstaller/Games",
		"java", "https://github.com/ScoopInstaller/Java",
		"main", "https://github.com/ScoopInstaller/Main",
		"nerd-fonts", "https://github.com/matthewjberger/scoop-nerd-fonts",
		"nirsoft", "https://github.com/ScoopInstaller/Nirsoft",
		"nonportable", "https://github.com/ScoopInstaller/Nonportable",
		"php", "https://github.com/ScoopInstaller/PHP",
		"sysinternals", "https://github.com/ScoopInstaller/Sysinternals",
		"versions", "https://github.com/ScoopInstaller/Versions",
	).Tag("buckets")
}
