package http

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

type openapi struct {
	Paths map[string]map[string]struct {
		Summary string
	}
}

// ActionOpenApiPaths completes api paths with placeholders for given openapi spec
func ActionOpenApiPaths(spec []byte, method string, placeholderPattern string, match func(c carapace.Context, matchedData map[string]string, segment string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var api openapi
		if err := json.Unmarshal(spec, &api); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		pathsDescribed := make([]string, 0)
		for path, methods := range api.Paths {
			for _method, details := range methods {
				if strings.EqualFold(_method, method) {
					pathsDescribed = append(pathsDescribed, path, details.Summary)
				}
			}
		}
		return ActionApiPathsDescribed(pathsDescribed, placeholderPattern, match)
	})
}

// ActionApiPaths completes api paths with placeholders
func ActionApiPaths(paths []string, placeholderPattern string, match func(c carapace.Context, matchedData map[string]string, segment string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		pathsDescribed := make([]string, 0, len(paths)*2)
		for _, value := range paths {
			pathsDescribed = append(pathsDescribed, value, "")
		}
		return ActionApiPathsDescribed(pathsDescribed, placeholderPattern, match)
	})
}

// ActionApiPathsDescribed completes api paths with placeholders
func ActionApiPathsDescribed(pathsDescribed []string, placeholderPattern string, match func(c carapace.Context, matchedData map[string]string, segment string) carapace.Action) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		if len(pathsDescribed)%2 != 0 {
			return carapace.ActionMessage("pathsDescribed should have even amount of elements")
		}

		placeholder := regexp.MustCompile(placeholderPattern)
		matchedData := make(map[string]string)
		matchedSegments := make(map[string]string)
		staticMatches := make(map[int]bool)

	path:
		for index, path := range pathsDescribed {
			if index%2 != 0 {
				continue
			}
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

			// store segment as path matched so far and this is currently being completed
			if len(segments) == (len(c.Parts) + 1) {
				matchedSegments[segments[len(c.Parts)]] = pathsDescribed[index+1]
			} else {
				if segment := segments[len(c.Parts)]; placeholder.MatchString(segment) {
					matchedSegments[segment] = ""
				} else {
					matchedSegments[segment+"/"] = ""
				}
			}
		}

		actions := make([]carapace.Action, 0, len(matchedSegments))
		for key, value := range matchedSegments {
			if placeholder.MatchString(key) {
				actions = append(actions, match(c, matchedData, key))
			} else {
				actions = append(actions, carapace.ActionValuesDescribed(key, value))
			}
		}
		return carapace.Batch(actions...).Invoke(c).Merge().ToA().NoSpace()
	})
}
