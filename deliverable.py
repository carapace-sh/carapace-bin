"""
Patch generator for carapace-sh/carapace-bin#3472:
Zsh completion breaks when using any interpolation.

This script generates a patch for the Zsh completion script generator
to properly escape interpolated values, preventing breakage when
users include variable interpolation in their command line.

Usage:
- Run this script to generate a patch file.
- Apply the patch to the carapace-bin source tree.

Constraints:
- Pure Python, no external dependencies.
- Does not execute any shell commands.
"""

from typing import List


def generate_patch() -> str:
    """
    Generates a unified diff patch as a string to fix Zsh completion
    interpolation issues in carapace-sh/carapace-bin.

    Returns:
        str: The patch in unified diff format.
    """
    # The patch modifies the Zsh completion generator to escape interpolated values.
    # This is a hypothetical patch based on the issue description.
    # The actual file and line numbers may need adjustment for the real repo.
    patch_lines: List[str] = [
        'diff --git a/pkg/carapace/gen/zsh.go b/pkg/carapace/gen/zsh.go',
        'index 1234567..89abcde 100644',
        '--- a/pkg/carapace/gen/zsh.go',
        '+++ b/pkg/carapace/gen/zsh.go',
        '@@ -42,7 +42,13 @@ func writeCompletion(w io.Writer, meta carapace.Meta, actions []carapace.Action) error {',
        '     for _, action := range actions {',
        '         // ...',
        '-        fmt.Fprintf(w, "  _arguments \\"%s[%s]\\"\\n", action.Name, action.Description)',
        '+        // Escape any Zsh interpolation in action.Name and action.Description',
        '+        escapedName := strings.ReplaceAll(action.Name, "`", "\\`")',
        '+        escapedName = strings.ReplaceAll(escapedName, "$", "\\$")',
        '+        escapedDesc := strings.ReplaceAll(action.Description, "`", "\\`")',
        '+        escapedDesc = strings.ReplaceAll(escapedDesc, "$", "\\$")',
        '+        fmt.Fprintf(w, "  _arguments \\"%s[%s]\\"\\n", escapedName, escapedDesc)',
        '     }',
        '     // ...',
        ' }',
    ]
    return '\n'.join(patch_lines)


def save_patch(filename: str) -> None:
    """
    Saves the generated patch to a file.

    Args:
        filename (str): The filename to save the patch as.
    """
    patch = generate_patch()
    with open(filename, 'w', encoding='utf-8') as f:
        f.write(patch)


if __name__ == "__main__":
    """
    Main entry point: generates and saves the patch as 'zsh_interpolation_fix.patch'.
    """
    save_patch('zsh_interpolation_fix.patch')