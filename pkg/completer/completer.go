package completer

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Completer struct {
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Group       string       `json:"group,omitempty"`
	Package     string       `json:"package,omitempty"`
	Spec        string       `json:"spec,omitempty"`
	Variant     string       `json:"variant,omitempty"`
	Overlay     string       `json:"overlay,omitempty"`
	Url         string       `json:"url,omitempty"`
	Execute     func() error `json:"-"`
}

func (c Completer) FormatImport() string {
	return fmt.Sprintf("\t%s %q", c.Variable(), c.Package)
}

func (c Completer) FormatCompleter() string {
	return fmt.Sprintf(`{
	Name: %q,
	Description: %q,
	Group: %q,
	Package: %q,
	Variant: %q,
	Url: %q,
	Execute: %s,
}`,
		c.Name,
		c.Description,
		c.Group,
		c.Package,
		c.Variant,
		c.Url,
		c.Variable()+".Execute")
}

func (c Completer) Variable() string {
	if c.Variant == "" {
		return varName(strings.Join([]string{c.Group, c.Name}, "__"))
	}
	return varName(strings.Join([]string{c.Group, c.Name, c.Variant}, "__"))
}

func varName(name string) string {
	if name == "go" {
		return "_go"
	}
	if unicode.IsDigit([]rune(name)[0]) {
		name = "_" + name
	}
	return strings.NewReplacer(
		"-", "_",
		".", "_",
	).Replace(name)
}

type Completers []Completer

func (c Completers) Get(variant string) (*Completer, bool) { // TODO Get or call it Lookup?
	sort.Sort(c) // TODO this modifies c
	for _, completer := range c {
		if variant == "" || completer.Variant == variant {
			return &completer, true
		}
	}
	return nil, false
}

func (c Completers) Len() int { return len(c) }
func (c Completers) Less(i, j int) bool {
	if b := strings.Compare(c[i].Name, c[j].Name) < 0; b {
		return b
	}

	groupPriority := map[string]int{
		// user specs
		"user": -5,
		// system specs
		"system": -4,
		// TODO shells? more specific stuff
		// os
		"darwin":  -1,
		"linux":   -1,
		"windows": -1,
		// TODO support pseudo os 'termux'?
	}
	if b := groupPriority[c[i].Group] < groupPriority[c[j].Group]; b {
		return b
	}

	return strings.Compare(c[i].Variant, c[j].Variant) < 0
}
func (c Completers) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
