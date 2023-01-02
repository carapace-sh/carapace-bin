package styles

import "github.com/rsteube/carapace/pkg/style"

type gh struct {
	Draft             string `desc:"draft pulls" tag:"issue styles"`
	JobFailed         string `desc:"jobs not finished successfully" tag:"job styles"`
	JobInProgress     string `desc:"jobs in progress" tag:"job styles"`
	JobSuccess        string `desc:"jobs successfully finished" tag:"job styles"`
	OwnerOrganization string `desc:"organization" tag:"owner styles"`
	OwnerUser         string `desc:"user" tag:"owner styles"`
	RepoArchived      string `desc:"archived repository" tag:"repo styles"`
	RepoFork          string `desc:"forked repository" tag:"repo styles"`
	RepoLocked        string `desc:"locked repository" tag:"repo styles"`
	RepoMirror        string `desc:"mirror repository" tag:"repo styles"`
	RepoPrivate       string `desc:"private repository" tag:"repo styles"`
	RepoPublic        string `desc:"public repository" tag:"repo styles"`
	RepoTemplate      string `desc:"template repository" tag:"repo styles"`
	StateClosed       string `desc:"closed issues/pulls" tag:"issue styles"`
	StateMerged       string `desc:"merged pulls" tag:"issue styles"`
	StateOpen         string `desc:"open issues/pulls" tag:"issue styles"`
}

var Gh = gh{
	Draft:             style.Gray,
	JobFailed:         style.Red,
	JobInProgress:     style.Yellow,
	JobSuccess:        style.Green,
	OwnerOrganization: style.Blue,
	OwnerUser:         style.Default,
	RepoArchived:      style.Magenta,
	RepoFork:          style.Gray,
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
