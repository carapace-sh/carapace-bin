package pandoc

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionInputFormats completes input formats
func ActionInputFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if !strings.ContainsAny(c.Value, "+-") {
			return carapace.ActionExecCommand("pandoc", "--list-input-formats")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}
		fields := extensionFields(c.Value)
		fields = fields[:len(fields)-1] // omit currently completed extension
		return ActionExtensions(fields[0]).Invoke(c).Filter(fields[1:]...).Prefix(strings.Join(fields, "")).ToA()
	}).Tag("input formats")
}

// ActionOutputFormats completes output formats
func ActionOutputFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if !strings.ContainsAny(c.Value, "+-") {
			return carapace.ActionExecCommand("pandoc", "--list-output-formats")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}
		fields := extensionFields(c.Value)
		fields = fields[:len(fields)-1] // omit currently completed extension
		return ActionExtensions(fields[0]).Invoke(c).Filter(fields[1:]...).Prefix(strings.Join(fields, "")).ToA()
	}).Tag("output formats")
}

func extensionFields(s string) []string {
	fields := []string{s}
	for {
		if last := fields[len(fields)-1]; len(last) == 0 {
			break
		} else {
			if index := strings.IndexAny(last[1:], "+-"); index > -1 {
				fields = append(fields[:len(fields)-1], last[:index+1], last[index+1:])
			} else {
				break
			}
		}
	}
	return fields
}

// ActionFormats completes formats
func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			ActionInputFormats(),
			ActionOutputFormats(),
		).Invoke(c).Merge().ToA()
	})
}

// ActionExtensions completes extensions
func ActionExtensions(format string) carapace.Action {
	return carapace.ActionExecCommand("pandoc", "--list-extensions", format)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	}).Tag("extensions")
}

// ActionHighlightStyles completes highlight styles
func ActionHighlightStyles() carapace.Action {
	return carapace.ActionValues(
		"breezeDark",
		"espresso",
		"haddock",
		"kate",
		"monochrome",
		"pygments",
		"tango",
		"zenburn",
	).Tag("highlight styles")
}
