package completer

import (
	"fmt"
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
	Choice      string       `json:"choice,omitempty"`
	Execute     func() error `json:"-"` // TODO should be a method based on the fields
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

func Split(nameVariantGroup string) (string, string, string) {
	nameVariant, group, _ := strings.Cut(nameVariantGroup, "@")
	name, variant, _ := strings.Cut(nameVariant, "/")
	return name, variant, group
}
