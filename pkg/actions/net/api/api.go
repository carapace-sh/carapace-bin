package api

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionApiPaths completes api paths with placeholders
func ActionApiPaths(paths []string, placeholderPattern string, match func(c carapace.Context, matchedData map[string]string, segment string) carapace.Action) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		placeholder := regexp.MustCompile(placeholderPattern)
		matchedData := make(map[string]string)
		matchedSegments := make(map[string]bool)
		staticMatches := make(map[int]bool)

	path:
		for _, path := range paths {
			segments := strings.Split(path, "/")
		segment:
			for index, segment := range segments {
				if index > len(c.Parts)-1 {
					break segment
				} else {
					if segment != c.Parts[index] {
						if !placeholder.MatchString(segment) {
							continue path // skip this path as it doesn't match and is not a placeholder
						} else {
							matchedData[segment] = c.Parts[index] // store entered data for placeholder (overwrite if duplicate)
						}
					} else {
						staticMatches[index] = true // static segment matches so placeholders should be ignored for this index
					}
				}
			}

			if len(segments) < len(c.Parts)+1 {
				continue path // skip path as it is shorter than what was entered (must be after staticMatches being set)
			}

			for key := range staticMatches {
				if segments[key] != c.Parts[key] {
					continue path // skip this path as it has a placeholder where a static segment was matched
				}
			}
			matchedSegments[segments[len(c.Parts)]] = true // store segment as path matched so far and this is currently being completed
		}

		actions := make([]carapace.InvokedAction, 0, len(matchedSegments))
		for key := range matchedSegments {
			actions = append(actions, match(c, matchedData, key).Invoke(c))
		}
		switch len(actions) {
		case 0:
			return carapace.ActionValues()
		case 1:
			return actions[0].ToA()
		default:
			return actions[0].Merge(actions[1:]...).ToA()
		}
	})
}
