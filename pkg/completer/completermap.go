package completer

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"

	"github.com/carapace-sh/carapace-bridge/pkg/choice"
)

type CompleterMap map[string]Completers

func (c CompleterMap) Lookup(nameVariantGroup string) (*Completer, bool) {
	name, variant, group := Split(nameVariantGroup)
	if variants, ok := c[name]; ok {
		return variants.Lookup(variant, group)
	}
	return nil, false
}

func (c CompleterMap) Filter(filter choice.Choice) CompleterMap {
	if filter.Name == "" && filter.Variant == "" && filter.Group == "" {
		return c
	}

	toFilter := c
	if filter.Name != "" {
		toFilter = make(CompleterMap)
		if v, ok := c[filter.Name]; ok {
			toFilter[filter.Name] = v
		}
	}

	m := make(CompleterMap)
	for name, variants := range toFilter {
		if filter.Name != "" && name != filter.Name {
			continue
		}
		for _, variant := range variants {
			switch {
			case filter.Variant != "" && filter.Variant != variant.Variant:
				continue
			case filter.Group != "" && filter.Group != variant.Group:
				continue
			}
			m[name] = append(m[name], variant)
		}
	}
	return m
}

func (c CompleterMap) Format(pkg, tag string) string {
	if tag != "" {
		tag = fmt.Sprintf("//go:build %s\n", tag)
	}
	return fmt.Sprintf(`%spackage %s

import (
	"github.com/carapace-sh/carapace-bin/pkg/completer"
%s)

%s
`, tag, pkg, c.FormatImports(), c.FormatCompleters())
}

func (c CompleterMap) Merge(other CompleterMap) {
	// TODO use copy?
	for name, otherVariants := range other {
		if variants, ok := c[name]; ok {
			c[name] = append(variants, otherVariants...)
		} else {
			c[name] = otherVariants
		}
	}
}

func (c CompleterMap) FormatImports() string {
	s := make([]string, 0)
	for _, group := range c {
		for _, completer := range group {
			s = append(s, completer.FormatImport())
		}
	}
	sort.Strings(s)
	return strings.Join(s, "\n")
}

func (c CompleterMap) FormatCompleters() string {
	c.SortVariants()

	s := make([]string, 0)
	for _, name := range slices.Sorted(maps.Keys(c)) {
		sVariants := make([]string, 0)
		for _, variant := range c[name] {
			sVariants = append(sVariants, fmt.Sprint(variant.FormatCompleter()))
		}
		s = append(s, fmt.Sprintf("%q: {\n%s,\n}", name, strings.Join(sVariants, ",\n")))
	}
	return fmt.Sprintf(`var completers = completer.CompleterMap{
%s}`, strings.Join(s, ",\n"))
}

func (c CompleterMap) SortVariants() {
	for _, variants := range c {
		sort.Sort(variants)
	}
}
