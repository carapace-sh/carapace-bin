# Git Remotes and Transfer Protocols

How Git communicates with remote repositories — remotes, refspecs, fetch, push, pull, and the wire protocols.

> **Source of truth**: <https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols>, <https://git-scm.com/docs/gitprotocol-pack>. For **reference types and refspecs**, see [refs.md](refs.md). For **object storage and packfiles**, see [objects.md](objects.md).

## Remote Configuration

### Adding and Managing Remotes

```bash
git remote add origin https://github.com/user/repo.git
git remote add upstream https://github.com/original/repo.git
git remote -v                          # List with URLs
git remote set-url origin <new-url>    # Change URL
git remote rename <old> <new>          # Rename
git remote remove <name>               # Remove
git remote show origin                 # Show remote details
```

### Remote Config

Each remote is stored in `.git/config`:

```ini
[remote "origin"]
    url = https://github.com/user/repo.git
    fetch = +refs/heads/*:refs/remotes/origin/*
    pushurl = git@github.com:user/repo.git    # Different push URL
    push = refs/heads/main:refs/heads/main    # Push refspec
    tagOpt = --no-tags                         # Don't auto-fetch tags
```

### Multiple Push URLs

```bash
git remote set-url --add --push origin <url1>
git remote set-url --add --push origin <url2>
```

Pushes to both URLs simultaneously.

## Refspecs

A refspec maps source refs to destination refs during fetch and push:

```
[+]<src>:<dst>
```

- `+` prefix = force update (allow non-fast-forward)
- `<src>` = source ref (remote for fetch, local for push)
- `<dst>` = destination ref

### Default Fetch Refspec

```
+refs/heads/*:refs/remotes/origin/*
```

Fetches all remote branches and stores them as remote-tracking branches under `refs/remotes/origin/`.

### Default Push Refspec

If not explicitly configured, `git push` uses the matching branch name:

```
refs/heads/<branch>:refs/heads/<branch>
```

With `push.default = current`, it pushes the current branch to a same-named branch on the remote.

### Push Default Modes

|| Mode | Behavior |
||------|----------|
|| `nothing` | Do not push anything (must specify explicitly) |
|| `current` | Push current branch to same-named remote branch |
|| `upstream` / `tracking` | Push current branch to its upstream |
|| `simple` | Like `upstream` but refuse if branch names differ (default since Git 2.0) |
|| `matching` | Push all branches that have same name on remote |

## Fetch

`git fetch` downloads objects and updates remote-tracking refs:

```bash
git fetch origin                    # Fetch all refs from origin
git fetch origin main               # Fetch only main
git fetch origin refs/heads/main:refs/heads/main  # Fetch into local branch
git fetch --prune origin            # Delete stale remote-tracking refs
git fetch --all                     # Fetch from all remotes
git fetch --depth=1 origin          # Shallow fetch
git fetch --unshallow               # Convert shallow to full clone
```

### Fetch Head

After a fetch, `.git/FETCH_HEAD` records the fetched tips:

```
abc1234def...    branch 'main' of https://github.com/user/repo
567890abc...     not-for-merge    branch 'feature' of ...
```

Lines marked `not-for-merge` are informational. The unmarked line is what `git pull` would merge.

## Pull

`git pull` = `git fetch` + `git merge` (or `git rebase`):

```bash
git pull origin main                # Fetch + merge
git pull --rebase origin main       # Fetch + rebase
git pull --ff-only                  # Only if fast-forward
```

### Pull Rebase

```bash
git config pull.rebase true         # Always rebase on pull
git config pull.rebase merges      # Rebase with --rebase-merges
git config pull.ff only             # Only fast-forward
```

## Push

`git push` uploads objects and updates remote refs:

```bash
git push origin main                # Push main
git push -u origin main             # Push and set upstream
git push origin --all               # Push all branches
git push origin --tags              # Push all tags
git push origin --follow-tags       # Push annotated tags reachable from pushed refs
git push origin --delete feature    # Delete remote branch
git push origin :feature            # Delete remote branch (old syntax)
```

### Force Push

```bash
git push --force origin main        # Overwrite remote history (dangerous)
git push --force-with-lease origin main  # Force push with safety check
git push --force-with-lease=main:abc123  # Force push only if remote is at abc123
```

`--force-with-lease` checks that the remote ref is at the expected value before overwriting. This prevents overwriting someone else's push.

### Push Options

```bash
git push -o ci.skip                 # Send push option to server
git config receive.advertisePushOptions true  # Server-side: advertise support
```

Push options are passed to `pre-receive` and `post-receive` hooks on the server.

## Transfer Protocols

Git supports several protocols for remote communication:

### Local Protocol

```
git clone /path/to/repo.git
git clone file:///path/to/repo.git
```

- Without `file://`: Uses hard links or local copy (fast, no protocol)
- With `file://`: Uses the smart HTTP protocol over a pipe (slower, but validates like a remote)

### HTTP/HTTPS Protocol

```
https://github.com/user/repo.git
```

Two variants:

|| Variant | Description |
||---------|-------------|
|| **Dumb HTTP** | Static file serving — no server-side Git. Fetches loose objects and packs directly. No push. |
|| **Smart HTTP** | Server runs `git-http-backend`. Supports fetch and push. Negotiation over HTTP. |

Smart HTTP uses `git-http-backend` as a CGI program. The client sends a `GET /info/refs?service=git-upload-pack` (fetch) or `GET /info/refs?service=git-receive-pack` (push) to start negotiation.

### SSH Protocol

```
git@github.com:user/repo.git
ssh://user@host/path/to/repo.git
```

SSH runs `git-upload-pack` (fetch) or `git-receive-pack` (push) on the remote host. The protocol is the same as the smart pack protocol, tunneled over SSH.

### Git Protocol

```
git://host/path/to/repo.git
```

A custom, unauthenticated protocol using `git-daemon`. Read-only by default. Rarely used today due to lack of authentication.

## Pack Protocol (Smart Protocol)

The smart protocol is used by SSH, smart HTTP, and the git protocol. It has two phases:

### Fetch (Upload Pack)

```
Client                              Server
  │                                   │
  ├── GET /info/refs?service=git-upload-pack ──►
  │                                   │
  │◄── ref advertisement (refs + capabilities) ──┤
  │                                   │
  ├── want <OID> (of each ref to fetch) ─────────►
  ├── have <OID> (of objects client has) ────────►
  │                                   │
  │◄── ACK/NAK (negotiation) ────────────────────┤
  │◄── packfile data ────────────────────────────┤
```

1. **Reference advertisement**: Server lists all refs and their OIDs
2. **Negotiation**: Client sends `want` lines for desired refs and `have` lines for objects it already has
3. **Pack transmission**: Server sends a thin packfile containing only the objects the client needs

### Push (Receive Pack)

```
Client                              Server
  │                                   │
  ├── GET /info/refs?service=git-receive-pack ──►
  │                                   │
  │◄── ref advertisement + capabilities ────────┤
  │                                   │
  ├── packfile data ─────────────────►│
  ├── update ref commands ───────────►│
  │                                   │
  │◄── report status ────────────────┤
```

1. **Reference advertisement**: Server lists refs and capabilities
2. **Pack transmission**: Client sends a thin packfile with new objects
3. **Ref update**: Client sends ref update commands (`<old-oid> <new-oid> <ref>`)
4. **Report**: Server reports success/failure for each ref update

## Credential Helpers

Git delegates credential storage to external helpers:

```bash
git config credential.helper cache          # In-memory cache (15 min default)
git config credential.helper cache --timeout=3600  # 1 hour
git config credential.helper store          # Store in ~/.git-credentials (plaintext)
git config credential.helper osxkeychain   # macOS Keychain
git config credential.helper manager       # Windows Credential Manager
git config credential.helper libsecret     # Linux libsecret
```

### Custom Credential Helpers

```bash
git credential fill    # Prompt for credentials (stdin protocol)
git credential approve # Store credentials
git credential reject  # Discard credentials
```

The credential helper protocol uses key-value pairs on stdin/stdout:

```
protocol=https
host=github.com
username=user
password=secret
```

## Shallow Clones

Shallow clones truncate history at a certain depth:

```bash
git clone --depth 1 <url>           # Only the latest commit
git clone --depth 1 --single-branch <url>  # Only one branch
git clone --shallow-submodules <url>  # Shallow submodules too
```

### Shallow File

`.git/shallow` lists the commit OIDs at the shallow boundary:

```
abc1234def567890...
567890abcdef1234...
```

Commits below the shallow boundary are not present. Some operations are limited:

- `git log` cannot show history beyond the shallow point
- `git push` from a shallow repo may be rejected (server config `receive.denyShallowCommits`)
- `git fetch --unshallow` converts to a full clone

## Mirror Clones

```bash
git clone --mirror <url>
```

A bare clone that:
- Fetches all refs (including refs outside `refs/heads/` and `refs/tags/`)
- Sets `remote.origin.fetch = +refs/*:refs/*`
- Sets `remote.origin.mirror = true`
- Sets `remote.origin.prune = true`

## Edge Cases and Known Issues

- **Force push safety**: `--force-with-lease` is not foolproof — if you fetched but didn't integrate, the expected value may be stale.
- **Shallow clone limitations**: Many operations (blame, log, bisect) are limited or broken with shallow clones. Use `--unshallow` to convert.
- **Credential storage**: `credential.helper store` stores passwords in **plaintext**. Use OS-native helpers instead.
- **Push to non-bare**: Pushing to a non-bare repository can cause the working tree to diverge from HEAD. The receiving repo should be bare or use `receive.denyCurrentBranch = refuse`.
- **Refspec ordering**: When multiple refspecs match the same destination, the last one wins.
- **Large repos**: For very large repositories, partial clone (`--filter=blob:none`) and sparse-checkout can dramatically reduce initial download size.

## References

- <https://git-scm.com/docs/git-remote>
- <https://git-scm.com/docs/gitprotocol-pack>
- <https://git-scm.com/book/en/v2/Git-on-the-Server-The-Protocols>

## Related Skills

- For reference types and refspecs, see [refs.md](refs.md)
- For packfile format, see [objects.md](objects.md)
- For partial clone and sparse-checkout, see [worktree.md](worktree.md)
