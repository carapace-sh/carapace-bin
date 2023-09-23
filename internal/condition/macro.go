package condition

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-spec/pkg/macro"
)

type Macro struct {
	macro macro.Macro[Condition]
}

func (m Macro) Parse(s string) Condition {
	return func(c carapace.Context) bool {
		b, err := m.macro.Parse(s)
		if err != nil {
			return false
		}
		return (*b)(c)
	}
}

func (m Macro) Signature() string { return m.macro.Signature() }

func MacroI[A any](f func(arg A) Condition) Macro {
	return Macro{
		macro: macro.MacroI[A, Condition](func(arg A) (*Condition, error) {
			a := f(arg)
			return &a, nil
		}),
	}
}

func MacroN(f func() Condition) Macro {
	return Macro{
		macro: macro.MacroN[Condition](func() (*Condition, error) {
			a := f()
			return &a, nil
		}),
	}
}

func MacroV[A any](f func(args ...A) Condition) Macro {
	return Macro{
		macro: macro.MacroV[A, Condition](func(args ...A) (*Condition, error) {
			a := f(args...)
			return &a, nil
		}),
	}
}
