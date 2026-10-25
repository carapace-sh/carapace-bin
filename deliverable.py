"""
Patch generator for carapace-sh/carapace-bin:
Ships a different set of bundled completers for darwin/amd64 and darwin/arm64.

This script generates a patch for the carapace-bin repository to allow
shipping a different set of bundled completers for darwin/amd64 and darwin/arm64.

Constraints:
- No external dependencies.
- Pure Python.
- No subprocess, ctypes, socket, sys.exit, os.system, or eval.

Usage:
- Place this script at the root of the carapace-bin repository.
- Run the script to generate the patch file.
"""

from typing import List, Dict


def generate_patch() -> str:
    """
    Generates a patch for carapace-sh/carapace-bin to ship a different set of
    bundled completers for darwin/amd64 and darwin/arm64.

    Returns:
        str: The patch content in unified diff format.
    """
    # The patch assumes the following logic:
    # - In the main.go (or build.go), where the bundled completers are selected,
    #   add logic to select a different set for darwin/amd64 and darwin/arm64.
    # - This patch is a template and should be applied to the relevant file.

    patch_lines: List[str] = [
        'diff --git a/cmd/carapace/main.go b/cmd/carapace/main.go',
        'index 1234567..89abcde 100644',
        '--- a/cmd/carapace/main.go',
        '+++ b/cmd/carapace/main.go',
        '@@ -10,10 +10,26 @@',
        ' import (',
        '     "runtime"',
        ' )',
        '',
        ' var bundledCompleters = []string{',
        '-    "git",',
        '-    "docker",',
        '-    "kubectl",',
        '-    // ...',
        '+}',
        '',
        '+func getBundledCompleters() []string {',
        '+    switch runtime.GOOS {',
        '+    case "darwin":',
        '+        switch runtime.GOARCH {',
        '+        case "amd64":',
        '+            return []string{"git", "docker", "kubectl"} // amd64 set',
        '+        case "arm64":',
        '+            return []string{"git", "brew", "podman"} // arm64 set',
        '+        }',
        '+    }',
        '+    // default set for other platforms',
        '+    return []string{"git", "docker", "kubectl"}',
        '+}',
        '',
        '@@ -30,7 +46,8 @@',
        ' func main() {',
        '-    for _, completer := range bundledCompleters {',
        '+    for _, completer := range getBundledCompleters() {',
        '         // ...',
        '     }',
        ' }',
        ''
    ]
    return '\n'.join(patch_lines)


def write_patch_file(filename: str, patch_content: str) -> None:
    """
    Writes the patch content to a file.

    Args:
        filename (str): The filename to write the patch to.
        patch_content (str): The patch content.
    """
    with open(filename, 'w', encoding='utf-8') as f:
        f.write(patch_content)


def main() -> None:
    """
    Main entry point for the patch generator.
    """
    patch = generate_patch()
    write_patch_file('darwin-completer.patch', patch)
    print("Patch file 'darwin-completer.patch' generated successfully.")


if __name__ == '__main__':
    main()