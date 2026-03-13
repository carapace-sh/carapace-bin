package action

import (
	"runtime"

	"github.com/carapace-sh/carapace"
)

func ActionDateFormatSpecifiers() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// see https://github.com/zsh-users/zsh/blob/master/Completion/Unix/Type/_date_formats
		exclusion := map[rune]string{
			'E': "cCgGxXyY",
			'O': "BdeHImMSuUVwWy",
			'-': "OEdegHIjklmMSUz",
			'_': "OEdgHIjmMSUz",
			'0': "Oekl",
			'^': "OEaAbBchP",
			'#': "OEaAbBchpPrXZ",
		}

		specs := map[rune]string{
			'a': "abbreviated day name",
			'A': "full day name",
			'b': "abbreviated month name",
			'h': "abbreviated month name",
			'B': "full month name",
			'c': "preferred locale date and time",
			'C': "2-digit century",
			'd': "day of month (01-31)",
			'D': "American format month/day/year (%m/%d/%y)",
			'e': "day of month ( 1-31)",
			'F': "ISO 8601 year-month-date (%Y-%m-%d)",
			'G': "4-digit ISO 8601 week-based year",
			'g': "2-digit ISO 8601 week-based year",
			'H': "hour (00-23)",
			'I': "hour (01-12)",
			'j': "day of year (001-366)",
			'k': "hour ( 0-23)",
			'l': "hour ( 1-12)",
			'm': "month (01-12)",
			'M': "minute (00-59)",
			'n': "newline",
			'p': "locale dependent AM/PM",
			'r': "locale dependent a.m. or p.m. time (%I:%M:%S %p)",
			'R': "24-hour notation time (%H:%M)",
			's': "seconds since the epoch",
			'S': "seconds (00-60)",
			't': "tab",
			'T': "24-hour notation with seconds (%H:%M:%S)",
			'u': "day of week (1-7, 1=Monday)",
			'U': "week number of current year, Sunday based (00-53)",
			'V': "ISO 8601 week number of current year, week 1 has 4 days in current year (01-53)",
			'w': "day of week (0-6, 0=Sunday)",
			'W': "week number of current year, Monday based (00-53)",
			'x': "locale dependent date representation without time",
			'X': "locale dependent time representation without date",
			'y': "2-digit year (00-99)",
			'Y': "full year",
			'z': "UTC offset",
			'Z': "timezone name",
			'%': "literal %",
		}

		switch runtime.GOOS {
		case "linux":
			specs['N'] = "fractional part of seconds since epoch, in nanoseconds"
			fallthrough
		case "freebsd", "dragonfly", "darwin":
			specs['E'] = "alternate representation"
			specs['O'] = "alternative format modifier"
			specs['-'] = "don't pad numeric values"
			specs['0'] = "left pad numeric values with zeroes"
			specs['_'] = "left pad numeric values with spaces"
		}

		switch runtime.GOOS {
		case "linux", "darwin":
			specs['#'] = "swap case of alphabetic characters"
			specs['^'] = "convert lowercase characters to uppercase"
			specs['P'] = "lower case locale dependent am/pm"
		}

		switch runtime.GOOS {
		case "freebsd", "dragonfly", "darwin":
			specs['v'] = "date in short form (%e-%b-%Y)"
		}

		if c.Getenv("SHELL") == "zsh" { // TODO this is unlikely set
			specs['f'] = "day of month (1-31)"
			specs['K'] = "hour (0-23)"
			specs['L'] = "hour (0-12)"
			specs['N'] = "fractional part of seconds since epoch, in nanoseconds (%9.)"
			specs['.'] = "fractional part of seconds since epoch"
			specs['-'] = "don't pad numeric values"
		}

		actionSpecifiers := carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			vals := make([]string, 0)
			for k, v := range specs {
				vals = append(vals, string(k), v)
			}
			return carapace.ActionValuesDescribed(vals...)
		})

		length := len(c.Value)
		for i := 0; i < length; i++ {
			switch c.Value[i] {
			case '%':
				if i == length-1 {
					return actionSpecifiers.Prefix(c.Value)
				}

				if exclusions, ok := exclusion[rune(c.Value[i+1])]; ok {
					for _, x := range exclusions {
						delete(specs, x)
					}
				}
				i += 1
			}
		}
		return actionSpecifiers.Prefix(c.Value + "%")
	}).NoSpace().Tag("date format specifiers")
}
