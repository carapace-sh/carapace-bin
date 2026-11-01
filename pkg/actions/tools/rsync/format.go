package rsync

import (
	"github.com/carapace-sh/carapace"
)

// ActionFormats completes formats
//
//	a (the remote IP address)
//	b (the number of bytes actually transferred)
func ActionFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"B", "the permission bits of the file",
		"C", "the full-file checksum if it is known for the  file",
		"G", "the gid of the file",
		"L", "the string \" -> SYMLINK\", \" => HARDLINK\", or \"\"",
		"M", "the last-modified time of the file",
		"P", "the module path",
		"U", "the uid of the file",
		"a", "the remote IP address",
		"b", "the number of bytes actually transferred",
		"c", "the total size of the block checksums received for the basis file",
		"f", "the filename",
		"h", "the remote host name",
		"i", "an itemized list of what is being updated",
		"l", "the length of the file in bytes",
		"m", "the module name",
		"n", "the filename",
		"o", "the operation",
		"p", "the process ID of this rsync session",
		"t", "the current date time",
		"u", "the authenticated username or an empty string",
	).Tag("formats")
}
