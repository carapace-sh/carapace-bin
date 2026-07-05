# Gists, Keys, Labels, Orgs, and Rulesets

How to manage gists, SSH/GPG keys, labels, organizations, and repository rulesets with `gh`.

## Gists

`gh gist` manages GitHub gists. A gist can be referenced by ID or URL.

### Argument Format

- By ID — e.g. `5b0e0062eb8e9654adad7bb1d81cc75f`
- By URL — e.g. `https://gist.github.com/OWNER/5b0e0062eb8e9654adad7bb1d81cc75f`

### `gh gist create`

Create a new gist. Aliased as `gh gist new`. Gists are secret by default.

```sh
$ gh gist create --public hello.py
$ gh gist create hello.py -d "my Hello-World program in Python"
$ gh gist create file1.py file2.py
$ gh gist create - < myfile.txt          # from stdin
$ gh gist create - -f output.txt         # from stdin with filename
$ gh gist create --web hello.py          # open in browser
```

| Flag | Description |
|------|-------------|
| `-d, --desc string` | Description |
| `-f, --filename string` | Filename when reading from stdin |
| `-p, --public` | List publicly (default: secret) |
| `-w, --web` | Open in browser |

### `gh gist list`

List your gists.

```sh
$ gh gist list
$ gh gist list --limit 50
$ gh gist list --public       # only public
$ gh gist list --secret       # only secret
$ gh gist list --filter octo --include-content  # regex filter including content
```

| Flag | Description |
|------|-------------|
| `--filter expression` | Filter gists using a regular expression |
| `--include-content` | Include gists' file content when filtering |
| `-L, --limit int` | Max gists (default 10) |
| `--public` | Show only public gists |
| `--secret` | Show only secret gists |

### `gh gist view`

View a gist.

```sh
$ gh gist view <gist-id>
$ gh gist view <gist-id> --raw        # raw content, no formatting
$ gh gist view <gist-id> --web        # open in browser
$ gh gist view <gist-id> --filename foo.py  # view specific file
$ gh gist view <gist-id> --files      # list file names only
```

| Flag | Description |
|------|-------------|
| `-f, --filename string` | View a specific file |
| `--files` | List file names |
| `-r, --raw` | Print raw content without rendering |
| `-w, --web` | Open in browser |

### `gh gist clone`

Clone a gist locally.

```sh
$ gh gist clone <gist-id>
$ gh gist clone <gist-id> mydir
```

### `gh gist edit`

Edit one of your gists.

```sh
$ gh gist edit <gist-id>                                   # interactive
$ gh gist edit <gist-id> --filename hello.py               # edit specific file
$ gh gist edit <gist-id> --filename hello.py hello.py      # replace with local file
$ gh gist edit <gist-id> --add newfile.py                  # add a new file
$ gh gist edit <gist-id> --desc "new description"          # change description
$ gh gist edit <gist-id> --remove hello.py                 # remove a file
```

| Flag | Description |
|------|-------------|
| `-a, --add string` | Add a new file to the gist |
| `-d, --desc string` | New description for the gist |
| `-f, --filename string` | Select a file to edit |
| `-r, --remove string` | Remove a file from the gist |

### `gh gist delete`

Delete a gist. To delete interactively, use no arguments.

```sh
$ gh gist delete                    # interactive
$ gh gist delete <gist-id>
$ gh gist delete <gist-id> --yes    # skip confirmation
```

| Flag | Description |
|------|-------------|
| `--yes` | Confirm deletion without prompting |

### `gh gist rename`

Rename a file in a gist.

```sh
$ gh gist rename <gist-id> oldname.py newname.py
```

## SSH Keys

`gh ssh-key` manages SSH keys registered with your GitHub account.

### `gh ssh-key add`

Add an SSH key to your account.

```sh
$ gh ssh-key add ~/.ssh/id_ed25519.pub
$ gh ssh-key add ~/.ssh/id_ed25519.pub --title "Laptop"
$ gh ssh-key add ~/.ssh/signing_key.pub --type signing
```

| Flag | Description |
|------|-------------|
| `-t, --title string` | Title for the new key |
| `--type string` | Key type: `authentication` (default) or `signing` |

### `gh ssh-key list`

List SSH keys in your account. Aliased as `gh ssh-key ls`.

### `gh ssh-key delete`

Delete an SSH key from your account.

```sh
$ gh ssh-key delete <id>
$ gh ssh-key delete <id> --yes    # skip confirmation
```

| Flag | Description |
|------|-------------|
| `-y, --yes` | Skip confirmation prompt |

## GPG Keys

`gh gpg-key` manages GPG keys registered with your GitHub account.

### `gh gpg-key add`

Add a GPG key to your account.

```sh
$ gh gpg-key add ~/.gnupg/pubkey.gpg
$ gh gpg-key add ~/.gnupg/pubkey.gpg --title "Work key"
```

| Flag | Description |
|------|-------------|
| `-t, --title string` | Title for the new key |

### `gh gpg-key list`

List GPG keys in your account. Aliased as `gh gpg-key ls`.

### `gh gpg-key delete`

Delete a GPG key from your account.

```sh
$ gh gpg-key delete <key-id>
$ gh gpg-key delete <key-id> --yes
```

| Flag | Description |
|------|-------------|
| `-y, --yes` | Skip confirmation prompt |

## Labels

`gh label` manages GitHub labels. Inherits `-R, --repo`.

### `gh label list`

List labels in a repository. Aliased as `gh label ls`.

```sh
$ gh label list
$ gh label list --sort name
$ gh label list --search bug
$ gh label list --web
```

| Flag | Description |
|------|-------------|
| `-L, --limit int` | Max labels (default 30) |
| `--order string` | Order: `asc`, `desc` (default `asc`) |
| `-S, --search string` | Search label names and descriptions |
| `--sort string` | Sort: `created`, `name` (default `created`) |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `color`, `createdAt`, `description`, `id`, `isDefault`, `name`, `updatedAt`, `url`.

### `gh label create`

Create a new label.

```sh
$ gh label create bug --description "Something isn't working" --color E99695
$ gh label create bug --force    # update if exists
```

| Flag | Description |
|------|-------------|
| `-c, --color string` | Label color (6-char hex) |
| `-d, --description string` | Description |
| `-f, --force` | Update if label already exists |

### `gh label edit`

Edit a label.

### `gh label delete`

Delete a label from a repository.

### `gh label clone`

Clone labels from one repository to another.

```sh
$ gh label clone source-owner/source-repo
$ gh label clone source-owner/source-repo --force    # overwrite existing
```

## Organizations

`gh org` manages GitHub organizations.

### `gh org list`

List organizations for the authenticated user. Aliased as `gh org ls`.

```sh
$ gh org list
$ gh org list --limit 100
```

| Flag | Description |
|------|-------------|
| `-L, --limit int` | Max organizations (default 30) |

## Rulesets

`gh ruleset` views repository rulesets — rules that apply to a repository. Aliased as `gh rs`. Inherits `-R, --repo`.

### `gh ruleset list`

List rulesets for a repository or organization. Aliased as `gh rs ls`, `gh ruleset ls`.

```sh
$ gh ruleset list
$ gh ruleset list --repo owner/repo --parents
$ gh ruleset list --org org-name
```

| Flag | Description |
|------|-------------|
| `-L, --limit int` | Max rulesets (default 30) |
| `-o, --org string` | List organization-wide rulesets (requires `admin:org` scope) |
| `-p, --parents` | Include rulesets from higher levels (default true) |
| `-w, --web` | Open in browser |

### `gh ruleset view`

View information about a ruleset.

```sh
$ gh ruleset view                    # interactive
$ gh ruleset view 43
$ gh ruleset view 23 --repo owner/repo
$ gh ruleset view 23 --org my-org
$ gh ruleset view --no-parents       # only rulesets configured in this repo
```

| Flag | Description |
|------|-------------|
| `-o, --org string` | Organization name for org-level ruleset |
| `-p, --parents` | Include rulesets from higher levels (default true) |
| `-w, --web` | Open in browser |

### `gh ruleset check`

View rules that would apply to a given branch. The branch does not need to exist.

```sh
$ gh ruleset check                   # current branch
$ gh ruleset check my-branch
$ gh ruleset check --default         # default branch
$ gh ruleset check my-branch --repo owner/repo
$ gh ruleset check --web
```

| Flag | Description |
|------|-------------|
| `--default` | Check rules on default branch |
| `-w, --web` | Open in browser |

## Artifact Attestations

`gh attestation` (aliased `gh at`) works with artifact attestations.

### `gh attestation download`

Download an artifact's attestations for offline use.

```sh
$ gh attestation download ./artifact.zip
$ gh attestation download oci://image-uri
$ gh attestation download ./artifact.zip --owner org-name
$ gh attestation download ./artifact.zip --repo owner/repo
```

| Flag | Description |
|------|-------------|
| `-d, --digest-alg string` | Digest algorithm: `sha256`, `sha512` (default `sha256`) |
| `--hostname string` | Configure host |
| `-L, --limit int` | Max attestations (default 30) |
| `-o, --owner string` | GitHub organization to scope lookup |
| `--predicate-type string` | Filter by predicate type |
| `-R, --repo string` | Repository in `<owner>/<repo>` format |

### `gh attestation verify`

Verify an artifact attestation.

### `gh attestation trusted-root`

Output `trusted_root.jsonl` contents for offline verification.

| Flag | Description |
|------|-------------|
| `--hostname string` | Configure host |
| `--tuf-root string` | Path to TUF root.json file |

## Related

- For repository targeting (`--repo`), see [repo.md](repo.md).
- For release attestations, see [release.md](release.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
