package action

import (
	"fmt"
	"net/url"

	"github.com/carapace-sh/carapace/pkg/uid"
)

// taskUid creates a UidF function for task values with a namespace prefix.
// Produces UIDs like rails://task/<prefix>/<value> (e.g. rails://task/db/prepare).
func taskUid(prefix string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		return &url.URL{
			Scheme: "rails",
			Host:   "task",
			Path:   fmt.Sprintf("/%s/%s", prefix, s),
		}, nil
	}
}

// namespaceUid creates a UidF function for namespace values with a parent prefix.
// Produces UIDs like rails://namespace/<prefix>/<value> (e.g. rails://namespace/db/migrate).
func namespaceUid(prefix string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		return &url.URL{
			Scheme: "rails",
			Host:   "namespace",
			Path:   fmt.Sprintf("/%s/%s", prefix, s),
		}, nil
	}
}
