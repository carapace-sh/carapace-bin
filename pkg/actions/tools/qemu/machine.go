package qemu

import "github.com/carapace-sh/carapace"

// ActionAccels completes QEMU accelerators
//
//	kvm (Linux KVM)
//	tcg (Tiny Code Generator (software emulation))
func ActionAccels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"hvf", "macOS Hypervisor.framework",
		"kvm", "Linux KVM",
		"mshv", "Microsoft Hyper-V",
		"nitro", "AWS Nitro Enclaves",
		"nvmm", "NetBSD virtual machine monitor",
		"tcg", "Tiny Code Generator (software emulation)",
		"whpx", "Windows Hypervisor Platform",
		"xen", "Xen",
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
		"cpu", "show CPU registers before entering a TB",
		"cpu_reset", "show CPU state before CPU resets",
		"exec", "show trace before each executed TB",
		"fpu", "include FPU registers in the 'cpu' logging",
		"guest_errors", "log when the guest OS does something invalid",
		"in_asm", "show target assembly code for each compiled TB",
		"int", "show interrupts/exceptions in short format",
		"invalid_mem", "log invalid memory accesses",
		"mmu", "log MMU-related activities",
		"nochain", "do not chain compiled TBs",
		"op", "show micro ops for each compiled TB",
		"op_ind", "show micro ops before indirect lowering",
		"op_opt", "show micro ops after optimization",
		"op_plugin", "show micro ops before plugin injection",
		"out_asm", "show generated host assembly code for each compiled TB",
		"page", "dump pages at beginning of user mode emulation",
		"pcall", "x86 only: show protected mode far calls/returns/exceptions",
		"plugin", "output from TCG plugins",
		"strace", "log every user-mode syscall",
		"tid", "open a separate log file per thread",
		"unimp", "log unimplemented functionality",
		"vpu", "include VPU registers in the 'cpu' logging",
	).Tag("debug items").Uid("qemu", "debug-item")
}
