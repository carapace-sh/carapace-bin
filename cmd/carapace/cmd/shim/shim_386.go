//go:build windows && 386

package shim

import _ "embed"

//go:embed shim_386.exe
var content []byte

func body(target string) []byte {
	return content
}
