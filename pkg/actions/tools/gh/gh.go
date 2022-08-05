// package gh contains github related actions
package gh

import "github.com/rsteube/carapace/pkg/cache"

type HostOpts struct {
	Host string
}

func (o HostOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host}
}

type OwnerOpts struct {
	Host  string
	Owner string
}

func (o OwnerOpts) repo() RepoOpts {
	return RepoOpts{Host: o.Host, Owner: o.Owner}

}

type RepoOpts struct {
	Host  string
	Owner string
	Name  string
}

func (o RepoOpts) cacheKey() cache.Key { return cache.String(o.Host, o.Owner, o.Name) }
