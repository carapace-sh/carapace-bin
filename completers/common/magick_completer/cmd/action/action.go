package action

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

func ActionCompleteOption(option string) carapace.Action {
	return carapace.ActionExecCommand("magick", "-list", option)(func(output []byte) carapace.Action {
		return carapace.ActionValues(strings.Fields(string(output))...)
	})
}

func ActionColors() carapace.Action {
	return carapace.ActionExecCommand("magick", "-list", "Color")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`^(\S+)\s+(srgb\([^)]+\))\s+(.*)$`)
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")

		if len(lines) < 4 {
			return carapace.ActionValues()
		}

		var completions []string
		for _, line := range lines[3:] {
			if matches := re.FindStringSubmatch(strings.TrimSpace(line)); len(matches) == 4 {
				value, colorValue, compliance := matches[1], matches[2], strings.TrimSpace(matches[3])
				description := colorValue
				if compliance != "" {
					description += " (" + compliance + ")"
				}
				completions = append(completions, value, description)
			}
		}
		return carapace.ActionValuesDescribed(completions...)
	})
}

func ActionFonts() carapace.Action {
	return carapace.ActionExecCommand("magick", "-list", "Font")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^ *Font: (.+)$`)
		matches := re.FindAllSubmatch(output, -1)

		fonts := make([]string, 0, len(matches))
		for _, match := range matches {
			if len(match) > 1 {
				fonts = append(fonts, string(match[1]))
			}
		}

		return carapace.ActionValues(fonts...)
	})
}

func ActionFontFamilies() carapace.Action {
	return carapace.ActionExecCommand("magick", "-list", "Font")(func(output []byte) carapace.Action {
		re := regexp.MustCompile(`(?m)^ *family: (.+)$`)
		matches := re.FindAllSubmatch(output, -1)

		familySet := make(map[string]struct{}, len(matches))
		for _, match := range matches {
			if len(match) > 1 {
				familySet[string(match[1])] = struct{}{}
			}
		}

		families := make([]string, 0, len(familySet))
		for family := range familySet {
			families = append(families, family)
		}

		return carapace.ActionValues(families...)
	})
}
