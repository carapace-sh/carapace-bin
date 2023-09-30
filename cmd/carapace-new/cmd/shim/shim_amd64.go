//go:build windows && amd64

package shim

import _ "embed"

//go:embed shim_amd64.exe
var content []byte

func body(target string) []byte {
	return content
}
