//go:build windows && arm64

package shim

import _ "embed"

//go:embed shim_arm64.exe
var content []byte

func body(target string) []byte {
	return content
}
