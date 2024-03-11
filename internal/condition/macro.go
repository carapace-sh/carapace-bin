package condition

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-spec/pkg/macro"
)

type Macro struct {
	macro.Macro[Condition]
	Description string
}

func (m Macro) WithDescription(s string) Macro {
	m.Description = s
	return m
}

func (m Macro) Parse(s string) Condition {
	return func(c carapace.Context) bool {
		b, err := m.Macro.Parse(s)
		if err != nil {
			return false
		}
		return (*b)(c)
	}
}

func MacroI[A any](f func(arg A) Condition) Macro {
	return Macro{
		Macro: macro.MacroI[A, Condition](func(arg A) (*Condition, error) {
			a := f(arg)
			return &a, nil
		}),
	}
}

func MacroN(f func() Condition) Macro {
	return Macro{
		Macro: macro.MacroN[Condition](func() (*Condition, error) {
			a := f()
			return &a, nil
		}),
	}
}

func MacroV[A any](f func(args ...A) Condition) Macro {
	return Macro{
		Macro: macro.MacroV[A, Condition](func(args ...A) (*Condition, error) {
			a := f(args...)
			return &a, nil
		}),
	}
}
