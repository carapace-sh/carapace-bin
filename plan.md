# Shell builtin state file

## Issue

[#3940](https://github.com/carapace-sh/carapace-bin/issues/3940): bash completion silently breaks when `CARAPACE_SHELL_FUNCTIONS` exceeds Linux `MAX_ARG_STRLEN` (128 KiB).

When `CARAPACE_BUILTINS` enables shell builtin capture, snippets export shell-internal data as environment variables (`CARAPACE_SHELL_ALIASES`, `CARAPACE_SHELL_BUILTINS`, `CARAPACE_SHELL_FUNCTIONS`, `CARAPACE_SHELL_JOBS`, `CARAPACE_SHELL_VARIABLES`). In shells with many functions (e.g. yc completion defines ~5100 functions, ~195 KB of names), the exported environment makes `execve` fail with `E2BIG`. The bash/zsh/fish bridge invokes carapace through `xargs`; GNU `xargs` returns `126` on exec failure. The current retry logic only handles `rc==1` for quoting fallback, so `data` stays empty and completion silently returns no candidates.

Observed on carapace-bin `v1.8.2`. On master, `f809186f2` gates shell-internal capture behind `CARAPACE_BUILTINS`, which avoids the default case. The failure still exists whenever builtin capture is enabled and the captured data is large, or when any other environment growth causes E2BIG.

## Current behavior

- Snippet generators (`cmd/carapace/cmd/snippet/*.go`) emit shell code that captures shell internals only when `CARAPACE_BUILTINS` contains the shell/group.
- Captured data is exported as environment variables before invoking `carapace <command> <shell> ...`.
- Go consumers in `pkg/actions/shell/shell.go` read `strings.Fields(os.Getenv("CARAPACE_SHELL_*"))`.
- `pkg/actions/os/env.go` hides these internal variables from environment-name completion.
- `CARAPACE_BUILTINS` is documented as experimental (`docs/src/setup/environment.md`, `docs/src/groups.md`, `docs/src/release_notes/v1.8.md`).

## Constraints

- Shell-internal state can only be enumerated by the shell process itself. Carapace runs as a child process and is blocked by the shell until completion returns; it cannot call back into the shell.
- The data must not be passed through the process environment when it can exceed `MAX_ARG_STRLEN`. A fix that retries with the same oversized environment still fails.
- The fix should preserve normal completion latency and not introduce a persistent daemon lifecycle in the first cut.
- Builtins capture is opt-in and experimental; behavior for disabled builtins must remain unchanged.

## Security considerations

- **Names only, never values.** Current capture exports names/identifiers only: aliases, builtin names, function names, job specs/ids, variable names. The state file must contain the same kind of data. It must not contain variable values, function bodies, alias expansions, environment values, or command output beyond the captured identifiers.
- **No variable-content escape hatch by default.** If variable contents are ever needed in the future, they must be guarded by an explicit insecure opt-in value in `CARAPACE_BUILTINS` (for example a dedicated `-insecure`/`variables:insecure` guard). This plan does not implement such a mode.
- **Correct permissions.** State directory must be user-private (`0700`). State files must be user-private (`0600`). On Unix, reject or avoid following symlinks when creating/reading state files. Do not place readable completion state in a shared directory.
- **Predictable path / temp-file safety.** Do not write to a predictable filename in a world-writable directory. If `/tmp` is used as fallback, create a user-private subdirectory with `0700` and never trust pre-existing ownership/mode. Prefer `XDG_RUNTIME_DIR` on Unix when available.
- **Cleanup and stale data.** Stale state files can leak shell metadata (function/variable names) across sessions if left in shared temp locations. Prefer a deterministic per-user/per-shell-session location that is overwritten on use and cleaned on shell exit where possible; document that stale files contain names only.
- **Atomic writes.** Write to a temp file and rename into place, or write with restrictive permissions from creation, to avoid partial reads and permission races.
- **Input validation.** Go readers should treat state as untrusted local input: size limits per key, line-based parsing, ignore empty lines, no code execution, no path traversal from key names.

## Options considered

### 1. Graceful degradation on E2BIG / rc!=1

Retry completion with shell-internal capture disabled when the carapace invocation fails with `126` or another exec failure.

- Pros: small change; prevents silent empty completion; works as safety net.
- Cons: degrades feature when enabled; does not solve the root cause; shell still exports oversized env; repeated failures waste attempts.

### 2. Export size guard / threshold

Only export a captured variable when its serialized size is below a threshold.

- Pros: prevents E2BIG before exec.
- Cons: arbitrary threshold; silent data loss; needs headroom for other env/args; duplicates logic across shell snippets; still env-based.

### 3. Shell state file / directory (recommended)

Write captured shell internals to a user-private state location before invoking carapace. Export only a small pointer env var, e.g. `CARAPACE_SHELL_STATE=/.../carapace/builtins/<id>`. Go reads required keys lazily from that location and falls back to legacy `CARAPACE_SHELL_*` env vars.

- Pros: removes env-size limit; Go can read only needed keys; no daemon; compatible with current one-shot xargs model; easy to secure with file permissions.
- Cons: file I/O on completion; cleanup/lifecycle; old carapace binaries cannot read new snippet state; snippet/test updates across shells.

### 4. Per-key files in a state directory

Use a directory per shell session/invocation containing one file per key (`aliases`, `builtins`, `functions`, `jobs`, `variables`).

- Pros: simplest shell write (`> file`); lazy Go reads by key; no section parser; easy size limits per key.
- Cons: multiple files; directory lifecycle/cleanup.

### 5. Single sectioned/JSON state file

Write one file containing all keys.

- Pros: one file; easier atomic replace.
- Cons: shell-side JSON generation is awkward/unsafe across shells; section format needs parser; Go may read/parse more than needed unless indexed.

### 6. Socket / helper / daemon

Shell pushes state to a local socket or daemon; carapace queries on demand.

- Pros: avoids env size; can centralize state.
- Cons: shell is blocked during completion so it cannot serve callbacks; helper still needs shell-captured data; daemon lifecycle, cleanup, Windows portability, and security surface are significantly more complex.

## Recommendation

Implement **option 3/4: per-key state files in a user-private state directory**, plus **option 1 as a safety net**.

Rationale:

- Per-key files give lazy reads with minimal parsing and no shell JSON dependency.
- A stable per-shell-session directory avoids temp-file leaks on interrupted completions.
- Legacy env fallback keeps old snippets working with new binaries.
- The E2BIG degradation path protects against unexpected env growth and mixed-version cases.

## Proposed design

### Environment variables

- Keep `CARAPACE_SHELL`.
- Add `CARAPACE_SHELL_STATE`: absolute path to a state directory containing captured keys.
- Keep legacy `CARAPACE_SHELL_ALIASES|BUILTINS|FUNCTIONS|JOBS|VARIABLES` as read-only fallback in Go for compatibility with old snippets.
- New snippets should not export legacy bulk env vars when writing state files.

### State location

Candidate layout:

```text
${XDG_RUNTIME_DIR:-${TMPDIR:-/tmp}}/carapace/builtins/<uid>/<session>/state
```

Where:

- `<uid>` is the numeric user id on Unix; on Windows use an appropriate per-user temp/appdata location with user ACLs.
- `<session>` is derived from shell process identity and shell name, not raw command input. Example inputs: parent shell PID, shell name, and optionally a random boot/session token; hashed to avoid leaking raw paths and to keep filename safe.
- `state` is a directory containing per-key files.

Example:

```text
$XDG_RUNTIME_DIR/carapace/builtins/1000/a1b2c3d4/
  aliases
  builtins
  functions
  jobs
  variables
```

Requirements:

- Create `carapace/builtins/<uid>` as `0700`.
- Create `<session>` as `0700`.
- Create key files as `0600`.
- Do not follow symlinks for directory/file creation or reads.
- If runtime dir is unavailable or unsafe, fall back to a per-user temp dir with the same permission checks.

The exact path helper should be shared by snippets where possible, but snippets are shell-specific strings; the Go side should accept any absolute safe path provided by `CARAPACE_SHELL_STATE`.

### File format

- One key per file: `aliases`, `builtins`, `functions`, `jobs`, `variables`.
- UTF-8 text, newline-delimited entries.
- No header, no shell code, no JSON.
- Ignore empty trailing newline; ignore empty lines when reading.
- Readers must apply a reasonable maximum size/count per key and return partial/empty on malformed oversized input rather than failing completion globally.

### Go read path

In `pkg/actions/shell/shell.go`:

- Add helper like `readShellKey(name string) []string`.
- If `CARAPACE_SHELL_STATE` is set and valid, read `<state>/<name>` lazily.
- Else fall back to `os.Getenv("CARAPACE_SHELL_"+strings.ToUpper(name))`.
- Preserve existing filtering/styling behavior in `ActionFunctions`, `ActionBuiltins`.

### Snippet write path

For shells with builtin capture:

- Create/ensure state directory with private permissions.
- Export `CARAPACE_SHELL_STATE`.
- Write enabled keys to files. Overwrite files on each completion invocation.
- Invoke carapace without exporting bulk `CARAPACE_SHELL_*` variables.
- On shells where cleanup hooks are practical, remove state dir on shell exit; not required for correctness if overwritten and private.

### E2BIG safety net

Even with state files, keep a defensive bridge path:

- bash/zsh/fish: preserve `rc==1` quoting retry. If invocation fails with `126` or non-zero with empty output, retry once with `CARAPACE_SHELL_STATE` unset and legacy shell-internal vars cleared for that invocation.
- bash-ble/oil: add explicit rc handling; they currently have weaker/no retry behavior.
- Log details only behind existing debug/log mechanisms; do not emit noisy stderr during normal completion.

## Compatibility

- New Go + old snippet: works via legacy env fallback.
- Old Go + new snippet: builtin shell-internal completions may be unavailable because old Go does not read `CARAPACE_SHELL_STATE`. This is acceptable if clearly documented for the experimental builtins feature; can be mitigated by dual-write for small payloads, but that weakens the fix.
- Disabled `CARAPACE_BUILTINS`: no state dir, no behavior change.

## Implementation plan

1. **Branch and plan** - done
   - Created `fix/shell-builtin-state-file`.
   - Plan updated as work progresses.

2. **Go state reader** - done
   - Added `pkg/actions/shell/state.go` to read keys from `CARAPACE_SHELL_STATE` with legacy `CARAPACE_SHELL_*` env fallback.
   - Enforced basic path safety: absolute clean dir, `os.OpenRoot`, regular file only, 1 MiB/100k line limits.
   - Updated `pkg/actions/shell/shell.go` to use the helper.

3. **Snippet generators** - in progress
   - Updated bash, bash-ble, oil, zsh, fish to write per-key state files and export `CARAPACE_SHELL_STATE`.
   - Added E2BIG/exec-failure fallback: retry once with `CARAPACE_SHELL_STATE` unset for bash-family shells.
   - Updated elvish, nushell, powershell, xonsh to write state files instead of exporting large env vars.
   - tcsh/cmd-clink currently have no shell-internal builtin capture; no E2BIG state export path to migrate.
   - Still to consider: stronger atomic temp+rename writes, stricter permission checks on existing dirs, Windows ACL behavior, and whether to treat `CARAPACE_SHELL_STATE` as internal.

4. **Fixtures and tests** - in progress
   - Updated snippet fixtures for affected shells.
   - Added Go tests for state-dir read, env fallback, relative-dir rejection, and missing-file fallback.
   - Still to add: symlink/size-limit tests and integration coverage.

5. **Docs** - pending
   - Update `docs/src/setup/environment.md` and related docs if `CARAPACE_SHELL_STATE` becomes user-visible or needs internal mention.
   - Add release note/changelog entry as appropriate.

6. **Validation** - in progress
   - Passed: `go test ./cmd/carapace/cmd/snippet/... ./pkg/actions/shell/... -count=1`.
   - Pending: broader tests/lint as needed.

## Open questions

- Should `CARAPACE_SHELL_STATE` be documented as a user-visible knob or treated as internal?
- Do we accept old-Go/new-snippet incompatibility for experimental builtins, or dual-write small payloads to env as a migration aid?
- Stable per-shell-session directory vs per-invocation temp directory: which is simpler and safer across all supported shells?
- Should state writes be atomic per key (`write temp + rename`) to avoid partial reads on interruption?
- What per-key size limit is appropriate before degrading to empty for that key?
- Do we need Windows-specific ACL handling beyond Go defaults and PowerShell temp-path behavior?
