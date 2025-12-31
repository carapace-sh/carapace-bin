package ghostty

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionAsyncBackends completes async backends
//
//	epoll (Use the `epoll` API)
//	io_uring (Use the `io_uring` API)
func ActionAsyncBackends() carapace.Action {
	return carapace.ActionStyledValuesDescribed(
		"auto", "Automatically choose the best backend for the platform based on available options",
		"epoll", "Use the `epoll` API",
		"io_uring", "Use the `io_uring` API",
	).StyleF(style.ForKeyword).
		Tag("async backends")
}
