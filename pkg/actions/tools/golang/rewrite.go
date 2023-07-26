package golang

import "github.com/rsteube/carapace"

// ActionRewrites completes rewrites
//
//	egl (Fixes initializers of EGLDisplay)
//	eglconf (Fixes initializers of EGLConfig)
func ActionRewrites() carapace.Action {
	return carapace.ActionValuesDescribed(
		"buildtag", "Remove +build comments from modules using Go 1.18 or later",
		"cftype", "Fixes initializers and casts of C.*Ref and JNI types",
		"context", "Change imports of golang.org/x/net/context to context",
		"egl", "Fixes initializers of EGLDisplay",
		"eglconf", "Fixes initializers of EGLConfig",
		"gotypes", "Change imports of golang.org/x/tools/go/{exact,types} to go/{constant,types}",
		"jni", "Fixes initializers of JNI's jobject and subtypes",
		"netipv6zone", "Adapt element key to IPAddr, UDPAddr or TCPAddr composite literals",
		"printerconfig", "Add element keys to Config composite literals",
	)
}
