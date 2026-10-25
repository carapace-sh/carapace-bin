"""
Navi's Alternative: Pure Python Implementation for Interactive Shell Completion

This module provides a pure Python alternative to the 'navi' interactive shell
completion tool, suitable for integration with carapace-sh/carapace-bin.

Features:
- Loads completion templates/snippets from YAML files.
- Interactive fuzzy selection of completion entries.
- Prints the selected completion to stdout for shell integration.

Constraints:
- No external dependencies.
- No use of subprocess, ctypes, socket, sys.exit, os.system, or eval.
- Pure Python, secure, and modular.

Author: OpenAI Assistant
"""

import os
import sys
from typing import List, Dict, Optional

# --- Data Loading ---

def load_yaml_snippets(file_path: str) -> List[Dict[str, str]]:
    """
    Loads completion snippets from a simple YAML file.
    Only supports a subset of YAML for security and simplicity.

    Args:
        file_path: Path to the YAML file.

    Returns:
        List of snippet dictionaries with 'name' and 'command' keys.
    """
    snippets: List[Dict[str, str]] = []
    if not os.path.isfile(file_path):
        return snippets

    with open(file_path, 'r', encoding='utf-8') as f:
        snippet: Dict[str, str] = {}
        for line in f:
            line = line.strip()
            if not line or line.startswith('#'):
                continue
            if line.startswith('-'):
                if snippet:
                    snippets.append(snippet)
                    snippet = {}
                line = line[1:].strip()
            if ':' in line:
                key, value = line.split(':', 1)
                snippet[key.strip()] = value.strip().strip('"').strip("'")
        if snippet:
            snippets.append(snippet)
    return snippets

# --- Fuzzy Finder ---

def fuzzy_filter(snippets: List[Dict[str, str]], query: str) -> List[Dict[str, str]]:
    """
    Filters snippets by fuzzy matching the query against the snippet 'name' and 'command'.

    Args:
        snippets: List of snippet dictionaries.
        query: User input query.

    Returns:
        Filtered list of snippets.
    """
    query = query.lower()
    return [
        s for s in snippets
        if query in s.get('name', '').lower() or query in s.get('command', '').lower()
    ]

# --- Interactive Selection ---

def clear_screen() -> None:
    """
    Clears the terminal screen in a cross-platform way without using os.system.
    """
    # ANSI escape sequence for clearing screen
    print("\033[2J\033[H", end='')

def interactive_select(snippets: List[Dict[str, str]]) -> Optional[Dict[str, str]]:
    """
    Presents an interactive prompt to select a snippet.

    Args:
        snippets: List of snippet dictionaries.

    Returns:
        The selected snippet dictionary, or None if cancelled.
    """
    if not snippets:
        print("No snippets available.")
        return None

    filtered_snippets = snippets[:]
    index = 0
    query = ""
    while True:
        clear_screen()
        print("Select a completion (type to filter, Enter to select, or Ctrl+C to cancel):\n")
        for i, snippet in enumerate(filtered_snippets):
            prefix = "-> " if i == index else "   "
            print(f"{prefix}{snippet.get('name', '')}: {snippet.get('command', '')}")
        print("\nType to filter, or press Enter to select. (Current filter: '{}')".format(query))

        try:
            inp = input("> ").strip()
        except (KeyboardInterrupt, EOFError):
            print("\nCancelled.")
            return None

        if inp == "":
            if filtered_snippets:
                return filtered_snippets[index]
            else:
                print("No snippets to select. Press Enter to continue.")
                try:
                    input()
                except (KeyboardInterrupt, EOFError):
                    return None
                continue
        else:
            query = inp
            filtered = fuzzy_filter(snippets, query)
            if filtered:
                filtered_snippets = filtered
                index = 0
            else:
                print("No matches. Press Enter to continue.")
                try:
                    input()
                except (KeyboardInterrupt, EOFError):
                    return None
                filtered_snippets = snippets[:]
                query = ""
                index = 0

# --- Main Entrypoint ---

def main() -> None:
    """
    Main entrypoint for the navi alternative.
    Loads snippets, presents interactive selection, and prints the result.
    """
    # Default snippets file path (can be customized)
    snippets_file = os.environ.get(
        'CARAPACE_NAVI_SNIPPETS',
        os.path.expanduser('~/.config/carapace/snippets.yaml')
    )
    snippets = load_yaml_snippets(snippets_file)

    if not snippets:
        print(f"No snippets found in {snippets_file}.")
        return

    selected = interactive_select(snippets)
    if selected:
        # Output the command for shell integration
        print(selected.get('command', ''))

# --- If run as script ---

if __name__ == "__main__":
    main()