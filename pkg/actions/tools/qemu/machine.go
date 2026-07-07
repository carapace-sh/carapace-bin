package qemu

import "github.com/carapace-sh/carapace"

// ActionAccels completes QEMU accelerators
//
//	kvm (Linux KVM)
//	tcg (Tiny Code Generator (software emulation))
func ActionAccels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"kvm", "Linux KVM",
		"tcg", "Tiny Code Generator (software emulation)",
		"xen", "Xen",
		"hvf", "macOS Hypervisor.framework",
		"whpx", "Windows Hypervisor Platform",
		"nvmm", "NetBSD virtual machine monitor",
		"nitro", "AWS Nitro Enclaves",
		"mshv", "Microsoft Hyper-V",
	).Tag("accelerators").Uid("qemu", "accel")
}

// ActionBootDrives completes QEMU boot drive letters
//
//	a (floppy)
//	c (hard disk)
func ActionBootDrives() carapace.Action {
	return carapace.ActionValuesDescribed(
		"a", "floppy",
		"b", "floppy",
		"c", "hard disk",
		"d", "CD-ROM",
		"n", "network",
	).Tag("boot drives").Uid("qemu", "boot-drive")
}

// ActionDebugItems completes QEMU debug log items
//
//	cpu (show CPU registers before entering a TB)
//	exec (show trace before each executed TB)
func ActionDebugItems() carapace.Action {
	return carapace.ActionValuesDescribed(
		"out_asm", "show generated host assembly code for each compiled TB",
		"in_asm", "show target assembly code for each compiled TB",
		"op", "show micro ops for each compiled TB",
		"op_opt", "show micro ops after optimization",
		"op_ind", "show micro ops before indirect lowering",
		"op_plugin", "show micro ops before plugin injection",
		"int", "show interrupts/exceptions in short format",
		"exec", "show trace before each executed TB",
		"cpu", "show CPU registers before entering a TB",
		"fpu", "include FPU registers in the 'cpu' logging",
		"mmu", "log MMU-related activities",
		"pcall", "x86 only: show protected mode far calls/returns/exceptions",
		"cpu_reset", "show CPU state before CPU resets",
		"unimp", "log unimplemented functionality",
		"guest_errors", "log when the guest OS does something invalid",
		"page", "dump pages at beginning of user mode emulation",
		"nochain", "do not chain compiled TBs",
		"plugin", "output from TCG plugins",
		"strace", "log every user-mode syscall",
		"tid", "open a separate log file per thread",
		"vpu", "include VPU registers in the 'cpu' logging",
		"invalid_mem", "log invalid memory accesses",
	).Tag("debug items").Uid("qemu", "debug-item")
}
