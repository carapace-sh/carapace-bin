package ar

import "github.com/rsteube/carapace"

// ActionTargets completes targets
//
//	elf64-x86-64
//	elf32-i386
func ActionTargets() carapace.Action {
	return carapace.ActionValues(
		"elf64-x86-64",
		"elf32-i386",
		"elf32-iamcu",
		"elf32-x86-64",
		"pei-i386",
		"pe-x86-64",
		"pei-x86-64",
		"elf64-little",
		"elf64-big",
		"elf32-little",
		"elf32-big",
		"pe-bigobj-x86-64",
		"pe-i386",
		"pdb",
		"elf64-bpfle",
		"elf64-bpfbe",
		"srec",
		"symbolsrec",
		"verilog",
		"tekhex",
		"binary",
		"ihex",
		"plugin",
	)
}
