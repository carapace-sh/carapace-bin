"""
iwctl Completer for carapace-sh/carapace-bin

This module provides a pure Python implementation of a command-line completer
for the `iwctl` tool, suitable for integration with carapace-sh/carapace-bin.

Constraints:
- No external dependencies.
- No use of subprocess, ctypes, socket, sys.exit, os.system, or eval.
- Modular, clean, and production-ready code.

Author: OpenAI
"""

from typing import List, Dict, Optional

# Static command structure for iwctl
IWCTL_COMMANDS: Dict[str, Dict[str, List[str]]] = {
    "station": {
        "": ["list", "device", "connect", "scan", "get-networks", "show", "set-property", "disconnect", "get-status"],
        "connect": ["<SSID>"],
        "device": ["<DEVICE>"],
        "scan": ["<DEVICE>"],
        "get-networks": ["<DEVICE>"],
        "show": ["<DEVICE>"],
        "set-property": ["<DEVICE>", "<PROPERTY>", "<VALUE>"],
        "disconnect": ["<DEVICE>"],
        "get-status": ["<DEVICE>"],
    },
    "device": {
        "": ["list", "show", "set-property"],
        "show": ["<DEVICE>"],
        "set-property": ["<DEVICE>", "<PROPERTY>", "<VALUE>"],
    },
    "known-networks": {
        "": ["list", "show", "forget"],
        "show": ["<SSID>"],
        "forget": ["<SSID>"],
    },
    "agent": {
        "": ["on", "off"],
    },
    "help": {},
    "exit": {},
    "quit": {},
    "version": {},
}

IWCTL_TOP_LEVEL_COMMANDS: List[str] = list(IWCTL_COMMANDS.keys())

def iwctl_completer(args: List[str]) -> List[str]:
    """
    Given a list of arguments (the current command line split into tokens),
    return a list of possible completions for the next token.

    Args:
        args (List[str]): The current command line split into tokens.

    Returns:
        List[str]: List of possible completions for the next token.
    """
    if not args or (len(args) == 1 and args[0] == ""):
        # No input yet, suggest top-level commands
        return IWCTL_TOP_LEVEL_COMMANDS

    # Determine the context
    cmd = args[0]
    subargs = args[1:]

    if cmd not in IWCTL_COMMANDS:
        # Partial command, suggest matching top-level commands
        return [c for c in IWCTL_TOP_LEVEL_COMMANDS if c.startswith(cmd)]

    # If no subargs, suggest subcommands
    if not subargs or (len(subargs) == 1 and subargs[0] == ""):
        return IWCTL_COMMANDS[cmd].get("", [])

    # Try to match subcommands
    subcmd = subargs[0]
    subsubargs = subargs[1:]

    # If subcmd is incomplete, suggest matching subcommands
    possible_subs = IWCTL_COMMANDS[cmd].get("", [])
    if subcmd not in possible_subs:
        return [s for s in possible_subs if s.startswith(subcmd)]

    # If subcmd is complete, suggest arguments for that subcommand
    subcmd_args = IWCTL_COMMANDS[cmd].get(subcmd, [])
    # Remove already provided arguments
    remaining_args = subcmd_args[len(subsubargs):]
    return remaining_args

# Example test cases for the completer
def _test_iwctl_completer() -> None:
    assert sorted(iwctl_completer([])) == sorted(IWCTL_TOP_LEVEL_COMMANDS)
    assert sorted(iwctl_completer([""])) == sorted(IWCTL_TOP_LEVEL_COMMANDS)
    assert sorted(iwctl_completer(["sta"])) == ["station"]
    assert sorted(iwctl_completer(["station"])) == sorted(IWCTL_COMMANDS["station"][""])
    assert sorted(iwctl_completer(["station", "c"])) == ["connect"]
    assert sorted(iwctl_completer(["station", "connect"])) == ["<SSID>"]
    assert sorted(iwctl_completer(["station", "connect", ""])) == ["<SSID>"]
    assert sorted(iwctl_completer(["device"])) == sorted(IWCTL_COMMANDS["device"][""])
    assert sorted(iwctl_completer(["device", "set-property", "wlan0"])) == ["<PROPERTY>", "<VALUE>"]
    assert sorted(iwctl_completer(["known-networks"])) == sorted(IWCTL_COMMANDS["known-networks"][""])
    assert sorted(iwctl_completer(["agent"])) == sorted(IWCTL_COMMANDS["agent"][""])
    assert sorted(iwctl_completer(["help"])) == []
    assert sorted(iwctl_completer(["exit"])) == []

if __name__ == "__main__":
    _test_iwctl_completer()