package styles

import "github.com/carapace-sh/carapace/pkg/style"

type gh struct {
	Draft             string `description:"draft pulls" tag:"issue styles"`
	JobFailed         string `description:"jobs not finished successfully" tag:"job styles"`
	JobInProgress     string `description:"jobs in progress" tag:"job styles"`
	JobSuccess        string `description:"jobs successfully finished" tag:"job styles"`
	OwnerOrganization string `description:"organization" tag:"owner styles"`
	OwnerUser         string `description:"user" tag:"owner styles"`
	RepoArchived      string `description:"archived repository" tag:"repo styles"`
	RepoFork          string `description:"forked repository" tag:"repo styles"`
	RepoLocked        string `description:"locked repository" tag:"repo styles"`
	RepoMirror        string `description:"mirror repository" tag:"repo styles"`
	RepoPrivate       string `description:"private repository" tag:"repo styles"`
	RepoPublic        string `description:"public repository" tag:"repo styles"`
	RepoTemplate      string `description:"template repository" tag:"repo styles"`
	StateClosed       string `description:"closed issues/pulls" tag:"issue styles"`
	StateMerged       string `description:"merged pulls" tag:"issue styles"`
	StateOpen         string `description:"open issues/pulls" tag:"issue styles"`
}

var Gh = gh{
	Draft:             style.Of(style.Dim, style.White),
	JobFailed:         style.Red,
	JobInProgress:     style.Yellow,
	JobSuccess:        style.Green,
	OwnerOrganization: style.Blue,
	OwnerUser:         style.Default,
	RepoArchived:      style.Magenta,
	RepoFork:          style.Of(style.Dim, style.White),
	RepoLocked:        style.Bold,
	RepoMirror:        style.Cyan,
	RepoPrivate:       style.Red,
	RepoPublic:        style.Default,
	RepoTemplate:      style.Yellow,
	StateClosed:       style.Red,
	StateMerged:       style.Magenta,
	StateOpen:         style.Green,
}

func init() {
	style.Register("gh", &Gh)
}

func (_gh *gh) ForState(s string, _ style.Context) string {
	switch s {
	case "closed":
		return _gh.StateClosed
	case "merged":
		return _gh.StateMerged
	case "open":
		return _gh.StateOpen
	default:
		return style.Default
	}
}
