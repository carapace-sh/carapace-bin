package ar

import "github.com/carapace-sh/carapace"

// ActionTargets completes targets
//
//	elf64-x86-64
//	elf32-i386
func ActionTargets() carapace.Action {
	return carapace.ActionValues(
		"binary",
		"elf32-big",
		"elf32-i386",
		"elf32-iamcu",
		"elf32-little",
		"elf32-x86-64",
		"elf64-big",
		"elf64-bpfbe",
		"elf64-bpfle",
		"elf64-little",
		"elf64-x86-64",
		"ihex",
		"pdb",
		"pe-bigobj-x86-64",
		"pe-i386",
		"pe-x86-64",
		"pei-i386",
		"pei-x86-64",
		"plugin",
		"srec",
		"symbolsrec",
		"tekhex",
		"verilog",
	)
}
