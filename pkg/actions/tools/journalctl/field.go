package journalctl

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionUserJournalFields completes user journal fields
//
//	MESSAGE (The human-readable message string)
//	MESSAGE_ID (A 128-bit message identifier ID)
func ActionUserJournalFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"MESSAGE", "The human-readable message string",
		"MESSAGE_ID", "A 128-bit message identifier ID",
		"PRIORITY", "A priority value between 0 and 7",
		"CODE_FILE", "The code file generating this message",
		"CODE_LINE", "The code line generating this message",
		"CODE_FUNC", "The code function generating this message",
		"ERRNO", "The low-level Unix error number causing this entry",
		"INVOCATION_ID", "A randomized, unique 128-bit ID identifying each runtime cycle of the unit",
		"USER_INVOCATION_ID", "A randomized, unique 128-bit ID identifying each runtime cycle of the unit",
		"SYSLOG_FACILITY", "Syslog compatibility field containing the facility",
		"SYSLOG_IDENTIFIER", "Syslog compatibility field containing the identifier",
		"SYSLOG_PID", "Syslog compatibility field containing the pid",
		"SYSLOG_TIMESTAMP", "Syslog compatibility field containing the timestamp",
		"SYSLOG_RAW", "The original contents of the syslog line",
		"TID", "The numeric thread ID (TID)",
		"UNIT", "The name of a unit",
		"USER_UNIT", "The name of a user unit",
	).Tag("user journal fields")
}

// ActionTrustedJournalFields completes trusted journal fields
//
//	_PID (The process ID of the process)
//	_UID (The user ID of the process)
func ActionTrustedJournalFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"_PID", "The process ID of the process",
		"_UID", "The user ID of the process",
		"_GID", "The group ID of the process",
		"_COMM", "The name of the process",
		"_EXE", "The executable path of the process",
		"_CMDLINE", "The command line of the process",
		"_CAP_EFFECTIVE", "The effective capabilities of the process",
		"_AUDIT_SESSION", "The session of the process",
		"_AUDIT_LOGINUID", "The login UID of the process",
		"_SYSTEMD_CGROUP", "The control group path in the systemd hierarchy",
		"_SYSTEMD_SLICE", "The systemd slice unit name",
		"_SYSTEMD_UNIT", "The systemd unit name",
		"_SYSTEMD_USER_UNIT", "The unit name in the systemd user manager",
		"_SYSTEMD_USER_SLICE", "The systemd slice user unit name",
		"_SYSTEMD_SESSION", "The systemd session ID",
		"_SYSTEMD_OWNER_UID", "The owner UID of the systemd user unit or systemd session",
		"_SELINUX_CONTEXT", "The SELinux security context (label)",
		"_SOURCE_REALTIME_TIMESTAMP", "The earliest trusted timestamp of the message",
		"_SOURCE_BOOTTIME_TIMESTAMP", "The earliest trusted timestamp of the message in CLOCK_BOOTTIME clock",
		"_BOOT_ID", "The kernel boot ID for the boot the message was generated in",
		"_MACHINE_ID", "The machine ID of the originating host",
		"_SYSTEMD_INVOCATION_ID", "The invocation ID for the runtime cycle of the unit",
		"_HOSTNAME", "The name of the originating host",
		"_TRANSPORT", "How the entry was received by the journal service",
		"_STREAM_ID", "Specifies a randomized 128-bit ID assigned to the stream connection",
		"_LINE_BREAK", "Indicates that the log message was not terminated with a normal newline character",
		"_NAMESPACE", "Namespace identifier",
	).Tag("trusted journal fields")
}

// ActionKernelJournalFields completes kernel journal fields
//
//	_KERNEL_DEVICE (The kernel device name)
//	_KERNEL_SUBSYSTEM (The kernel subsystem name.)
func ActionKernelJournalFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"_KERNEL_DEVICE", "The kernel device name",
		"_KERNEL_SUBSYSTEM", "The kernel subsystem name.",
		"_UDEV_SYSNAME", "The kernel device name",
		"_UDEV_DEVNODE", "The device node path of this device in /dev/",
		"_UDEV_DEVLINK", "Additional symlink names pointing to the device node in /dev/",
	).Tag("kernel journal fields")
}

// ActionObjectJournalFields completes object journal fields
//
//	OBJECT_PID (The process ID of the progress)
//	OBJECT_UID (The user ID of the progress)
func ActionObjectJournalFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"COREDUMP_UNIT", "Coredump unit",
		"COREDUMP_USER_UNIT", "Coredump user unit",
		"OBJECT_PID", "The process ID of the progress",
		"OBJECT_UID", "The user ID of the progress",
		"OBJECT_GID", "The group ID of the progress",
		"OBJECT_COMM", "The name of the progress",
		"OBJECT_EXE", "The executable path of the progress",
		"OBJECT_CMDLINE", "The command line of the progress",
		"OBJECT_AUDIT_SESSION", "The session of the progress",
		"OBJECT_AUDIT_LOGINUID", "The login UID of the progress",
		"OBJECT_SYSTEMD_CGROUP", "The control group path in the systemd hierarchy",
		"OBJECT_SYSTEMD_SESSION", "The systemd session ID",
		"OBJECT_SYSTEMD_OWNER_UID", "The owner UID of the systemd user unit or systemd session",
		"OBJECT_SYSTEMD_UNIT", "The systemd unit name",
		"OBJECT_SYSTEMD_USER_UNIT", "The unit name in the systemd user manager",
		"OBJECT_SYSTEMD_INVOCATION_ID", "The invocation ID for the runtime cycle of the unit",
	).Tag("object journal fields")

}

// ActionAddressJournalFields completes address journal fields
//
//	__SEQNUM (The sequence number)
//	__SEQNUM_ID (The sequence number ID)
func ActionAddressJournalFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"__CURSOR", "The cursor for the entry",
		"__REALTIME_TIMESTAMP", "The wallclock time when the entry was received",
		"__MONOTONIC_TIMESTAMP", "The monotonic time when the entry was received",
		"__SEQNUM", "The sequence number",
		"__SEQNUM_ID", "The sequence number ID",
	).Tag("address journal fields")
}

// ActionJournalFields completes journal fields
//
//	MESSAGE (The human-readable message string)
//	OBJECT_PID (The process ID of the progress)
func ActionJournalFields() carapace.Action {
	return carapace.Batch(
		ActionUserJournalFields().Style(style.Blue),
		ActionTrustedJournalFields().Style(style.Green),
		ActionKernelJournalFields().Style(style.Magenta),
		ActionObjectJournalFields().Style(style.Yellow),
		ActionAddressJournalFields().Style(style.Cyan),
	).ToA()
}

// ActionJournalFieldValues completes journal field values
func ActionJournalFieldValues(field string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO add more fields and verify
		switch field {
		case "UNIT", "_SYSTEMD_UNIT", "COREDUMP_UNIT", "OBJECT_SYSTEMD_UNIT":
			return systemctl.ActionUnits(systemctl.UnitOpts{User: false, Active: true, Inactive: true})
		case "USER_UNIT", "_SYSTEMD_USER_UNIT", "COREDUMP_USER_UNIT", "OBJECT_SYSTEMD_USER_UNIT":
			return systemctl.ActionUnits(systemctl.UnitOpts{User: true, Active: true, Inactive: true})
		case "_PID", "SYSLOG_PID", "OBJECT_PID":
			return ps.ActionProcessIds()
		case "_TRANSPORT":
			return ActionTransports()
		default:
			return carapace.ActionValues()
		}
	})
}
