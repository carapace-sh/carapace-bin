// package gh contains github related actions
package gh

import "github.com/rsteube/carapace/pkg/cache"

type repo interface {
	host() string
	owner() string
	name() string
}

type HostOpts struct {
	Host string
}

func (o HostOpts) host() string {
	if o.Host == "" {
		return "github.com"
	} else {
		return o.Host
	}
}
func (o HostOpts) owner() string { return "" }
func (o HostOpts) name() string  { return "" }

type OwnerOpts struct {
	Host  string
	Owner string
}

func (o OwnerOpts) host() string {
	if o.Host == "" {
		return "github.com"
	} else {
		return o.Host
	}
}
func (o OwnerOpts) owner() string { return o.Owner }
func (o OwnerOpts) name() string  { return "" }

type RepoOpts struct {
	Host  string
	Owner string
	Name  string
}

func (o RepoOpts) host() string {
	if o.Host == "" {
		return "github.com"
	} else {
		return o.Host
	}
}
func (o RepoOpts) owner() string       { return o.Owner }
func (o RepoOpts) name() string        { return o.Name }
func (o RepoOpts) cacheKey() cache.Key { return cache.String(o.Host, o.Owner, o.Name) }
