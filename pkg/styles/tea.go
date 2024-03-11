package styles

import "github.com/carapace-sh/carapace/pkg/style"

type tea struct {
	StateClosed string `description:"closed issues/pulls" tag:"issue styles"`
	StateMerged string `description:"merged pulls" tag:"issue styles"`
	StateOpen   string `description:"open issues/pulls" tag:"issue styles"`
}

var Tea = tea{
	StateClosed: style.Red,
	StateMerged: style.Magenta,
	StateOpen:   style.Green,
}

func init() {
	style.Register("tea", &Tea)
}

func (_tea *tea) ForState(s string, _ style.Context) string {
	switch s {
	case "closed":
		return _tea.StateClosed
	case "merged":
		return _tea.StateMerged
	case "open":
		return _tea.StateOpen
	default:
		return style.Default
	}
}
