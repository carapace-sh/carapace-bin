"""
Carapace Launchctl Completer

This module provides a pure Python implementation of a command-line
completion engine for the `launchctl` command, as described in:
https://github.com/carapace-sh/carapace-bin/issues/3360

It parses the arguments and provides possible completions for subcommands,
flags, and arguments, suitable for integration with carapace-sh.

Constraints:
- No external dependencies.
- No subprocess, ctypes, socket.
- Pure Python, secure, and modular.

Author: OpenAI Assistant
"""

from typing import List, Dict, Optional

LAUNCHCTL_COMMANDS: Dict[str, str] = {
    "bootstrap": "Bootstrap services into a domain",
    "bootout": "Remove services from a domain",
    "enable": "Enable specified services",
    "disable": "Disable specified services",
    "kickstart": "Force start/stop of services",
    "kill": "Send signal to services",
    "list": "List services in a domain",
    "print": "Print detailed service info",
    "print-cache": "Print service cache",
    "print-disabled": "Print disabled services",
    "blame": "Show service start times",
    "procinfo": "Show process info",
    "procstats": "Show process statistics",
    "attach": "Attach to service's stdio",
    "config": "Show system configuration",
    "dumpstate": "Dump launchd state",
    "managerpid": "Show launchd's PID",
    "remove": "Remove service from domain",
    "bsexec": "Execute in a bootstrap namespace",
    "asuser": "Execute as a user",
    "submit": "Submit a job definition",
    "error": "Explain error code",
    "help": "Show help",
    "version": "Show version",
}

LAUNCHCTL_FLAGS: Dict[str, str] = {
    "-w": "Write changes to disk",
    "-F": "Force operation",
    "-D": "Specify domain",
    "-g": "Show in graph format",
    "-x": "Show in XML",
    "-p": "Specify property",
    "-t": "Timeout",
    "-u": "User",
    "-s": "Session",
    "-h": "Show help",
    "--help": "Show help",
    "--version": "Show version",
}

def complete_launchctl(args: List[str]) -> List[str]:
    """
    Given a list of arguments (excluding the program name), return possible completions.

    Args:
        args: List of arguments provided so far.

    Returns:
        List of possible completions (strings).
    """
    completions: List[str] = []

    if not args:
        # No arguments yet: suggest commands and global flags
        completions.extend(sorted(LAUNCHCTL_COMMANDS.keys()))
        completions.extend(sorted(LAUNCHCTL_FLAGS.keys()))
        return completions

    current = args[-1]
    prev = args[-2] if len(args) > 1 else ""

    # If current arg starts with '-', suggest flags
    if current.startswith('-'):
        completions.extend(
            [flag for flag in LAUNCHCTL_FLAGS if flag.startswith(current)]
        )
        return completions

    # If first arg is incomplete, suggest matching commands
    if len(args) == 1:
        completions.extend(
            [cmd for cmd in LAUNCHCTL_COMMANDS if cmd.startswith(current)]
        )
        completions.extend(
            [flag for flag in LAUNCHCTL_FLAGS if flag.startswith(current)]
        )
        return completions

    # If first arg is a known command, suggest sub-arguments or flags
    command = args[0]
    if command in LAUNCHCTL_COMMANDS:
        # Suggest flags for the command
        if current.startswith('-'):
            completions.extend(
                [flag for flag in LAUNCHCTL_FLAGS if flag.startswith(current)]
            )
            return completions

        # For some commands, suggest domains or plist files
        if command in ("bootstrap", "bootout", "enable", "disable", "kickstart", "kill", "remove"):
            # Suggest domain names if appropriate
            domains = [
                "system", "user/501", "gui/501", "loginwindow", "background"
            ]
            if len(args) == 2:
                completions.extend([d for d in domains if d.startswith(current)])
                return completions
            # Otherwise, suggest plist files (placeholder, as we can't list files)
            if len(args) == 3:
                completions.append("com.example.service.plist")
                return completions

        # For 'list', suggest domains
        if command == "list" and len(args) == 2:
            domains = [
                "system", "user/501", "gui/501", "loginwindow", "background"
            ]
            completions.extend([d for d in domains if d.startswith(current)])
            return completions

        # For 'error', suggest error codes (as strings)
        if command == "error" and len(args) == 2:
            error_codes = ["2", "3", "5", "78", "112"]
            completions.extend([e for e in error_codes if e.startswith(current)])
            return completions

        # For 'help' and 'version', no further completions
        if command in ("help", "version"):
            return []

        # For other commands, suggest flags
        completions.extend(
            [flag for flag in LAUNCHCTL_FLAGS if flag.startswith(current)]
        )
        return completions

    # If first arg is not a known command, suggest matching commands
    completions.extend(
        [cmd for cmd in LAUNCHCTL_COMMANDS if cmd.startswith(current)]
    )
    completions.extend(
        [flag for flag in LAUNCHCTL_FLAGS if flag.startswith(current)]
    )
    return completions

# Example test cases (for demonstration; in production, use proper unit tests)
if __name__ == "__main__":
    test_cases = [
        [],
        ["b"],
        ["list"],
        ["list", ""],
        ["bootstrap", ""],
        ["bootstrap", "sys"],
        ["kill", ""],
        ["-"],
        ["--v"],
        ["error", ""],
        ["help"],
    ]
    for args in test_cases:
        print(f"Args: {args!r} -> Completions: {complete_launchctl(args)}")