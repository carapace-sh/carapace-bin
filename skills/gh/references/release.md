# Releases

How to create, view, download, upload, edit, delete, and verify GitHub releases with `gh release`.

## Overview

`gh release` manages GitHub releases. A release is associated with a git tag. If a matching tag doesn't exist, one is automatically created from the latest state of the default branch (unless `--verify-tag` is used). The `-R, --repo [HOST/]OWNER/REPO` flag targets a specific repository.

## Commands

### `gh release create`

Create a new GitHub release. Aliased as `gh release new`.

```sh
$ gh release create                              # interactive
$ gh release create v1.2.3 --notes "bugfix release"
$ gh release create v1.2.3 --generate-notes      # auto-generated notes
$ gh release create v1.2.3 -F release-notes.md   # notes from file
$ gh release create v1.2.3 --notes-from-tag      # from tag annotation
$ gh release create v1.2.3 ./dist/*.tgz          # upload assets
$ gh release create v1.2.3 '/path/to/asset.zip#My display label'  # with label
$ gh release create v1.2.3 --discussion-category "General"
$ gh release create v1.2.3 --fail-on-no-commits
$ gh release create v1.2.3 --latest=false        # don't mark as latest
```

| Flag | Description |
|------|-------------|
| `--discussion-category string` | Start a discussion in the specified category |
| `-d, --draft` | Save as draft instead of publishing |
| `--fail-on-no-commits` | Fail if no new commits since last release (no effect on first release) |
| `--generate-notes` | Auto-generate title and notes via GitHub Release Notes API |
| `--latest` | Mark as "Latest" (default: automatic by date and version). Use `--latest=false` to explicitly not set |
| `-n, --notes string` | Release notes |
| `-F, --notes-file file` | Read notes from file (`-` for stdin) |
| `--notes-from-tag` | Fetch notes from tag annotation or commit message |
| `--notes-start-tag string` | Tag to use as starting point for generating notes |
| `-p, --prerelease` | Mark as prerelease |
| `--target branch` | Target branch or full commit SHA (default: main branch) |
| `-t, --title string` | Release title |
| `--verify-tag` | Abort if the git tag doesn't already exist in the remote |

**Asset labels**: Append `#display label` after the filename to define a display label for an asset.

**Auto-generated notes**: When using `--generate-notes`, a title is also auto-generated unless `--title` is passed. Additional notes from `--notes` are prepended to the auto-generated content.

**From annotated tag**: First create the tag with git, push it, then run `gh release create`. Use `--notes-from-tag` to get notes from the tag annotation (or commit message if not annotated).

### `gh release list`

List releases in a repository. Aliased as `gh release ls`.

```sh
$ gh release list
$ gh release list --exclude-drafts
$ gh release list --exclude-pre-releases
$ gh release list --limit 50
```

| Flag | Description |
|------|-------------|
| `--exclude-drafts` | Exclude draft releases |
| `--exclude-pre-releases` | Exclude pre-releases |
| `-L, --limit int` | Max items (default 30) |
| `-O, --order string` | Order: `asc`, `desc` (default `desc`) |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `createdAt`, `isDraft`, `isImmutable`, `isLatest`, `isPrerelease`, `name`, `publishedAt`, `tagName`.

### `gh release view`

View information about a release. Without a tag, shows the latest release.

```sh
$ gh release view
$ gh release view v1.2.3
$ gh release view v1.2.3 --web
```

| Flag | Description |
|------|-------------|
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `apiUrl`, `assets`, `author`, `body`, `createdAt`, `databaseId`, `id`, `isDraft`, `isImmutable`, `isPrerelease`, `name`, `publishedAt`, `tagName`, `tarballUrl`, `targetCommitish`, `uploadUrl`, `url`, `zipballUrl`.

### `gh release download`

Download assets from a release. Without a tag, downloads from the latest release (requires `--pattern` or `--archive`).

```sh
$ gh release download v1.2.3
$ gh release download --pattern '*.deb'
$ gh release download -p '*.deb' -p '*.rpm'
$ gh release download v1.2.3 --archive=zip
$ gh release download v1.2.3 -D ./downloads
$ gh release download v1.2.3 -O myfile    # single asset to file
$ gh release download --skip-existing
$ gh release download --clobber
```

| Flag | Description |
|------|-------------|
| `-A, --archive format` | Download source code archive: `zip` or `tar.gz` |
| `--clobber` | Overwrite existing files |
| `-D, --dir directory` | Download directory (default `.`) |
| `-O, --output file` | File for single asset (`-` for stdout) |
| `-p, --pattern stringArray` | Download only assets matching glob patterns |
| `--skip-existing` | Skip when files of same name exist |

### `gh release upload`

Upload asset files to a release. Append `#display label` after filename for a display label.

```sh
$ gh release upload v1.2.3 ./dist/*.tgz
$ gh release upload v1.2.3 asset.zip#My Label
$ gh release upload v1.2.3 *.deb --clobber
```

| Flag | Description |
|------|-------------|
| `--clobber` | Delete and re-upload existing assets of the same name |

> **Caution with `--clobber`**: If the upload fails, the original assets will be lost.

### `gh release delete`

Delete a release.

```sh
$ gh release delete v1.2.3
$ gh release delete v1.2.3 --cleanup-tag    # also delete the tag
$ gh release delete v1.2.3 --yes            # skip confirmation
```

| Flag | Description |
|------|-------------|
| `--cleanup-tag` | Delete the git tag in addition to the release |
| `-y, --yes` | Skip confirmation prompt |

### `gh release edit`

Edit a release.

```sh
$ gh release edit v1.0 --draft=false                    # publish a draft
$ gh release edit v1.0 --notes-file /path/to/notes.md
$ gh release edit v1.0 --prerelease --latest
$ gh release edit v1.0 --target develop
```

| Flag | Description |
|------|-------------|
| `--discussion-category string` | Start a discussion when publishing a draft |
| `--draft` | Save as draft instead of publishing |
| `--latest` | Explicitly mark as "Latest" |
| `-n, --notes string` | Release notes |
| `-F, --notes-file file` | Read notes from file |
| `--prerelease` | Mark as prerelease |
| `--tag string` | New tag name |
| `--target branch` | Target branch or commit SHA |
| `-t, --title string` | Release title |
| `--verify-tag` | Abort if the git tag doesn't exist |

### `gh release verify`

Verify that a release has a valid cryptographically signed attestation. Without a tag, verifies the latest release.

```sh
$ gh release verify
$ gh release verify v1.2.3
$ gh release verify v1.2.3 --format json
```

| Flag | Description |
|------|-------------|
| `--format string` | Output format: `json` |
| `--jq` / `--template` | Filter/format output |

Fetches the attestation and prints metadata about all assets, including their digests.

### `gh release verify-asset`

Verify that an asset file originated from a specific release using cryptographically signed attestations.

```sh
$ gh release verify-asset ./dist/my-asset.zip
$ gh release verify-asset v1.2.3 ./dist/my-asset.zip
$ gh release verify-asset v1.2.3 ./dist/my-asset.zip --format json
```

| Flag | Description |
|------|-------------|
| `--format string` | Output format: `json` |
| `--jq` / `--template` | Filter/format output |

Validates that the asset's digest matches the subject in the attestation and that the attestation is associated with the release.

### `gh release delete-asset`

Delete an asset from a release.

```sh
$ gh release delete-asset v1.2.3 asset.zip
$ gh release delete-asset v1.2.3 asset.zip --yes
```

| Flag | Description |
|------|-------------|
| `-y, --yes` | Skip confirmation prompt |

## Immutable Releases

When release immutability is enabled for a repository:

- Git tags associated with a release **cannot** be modified or deleted
- Release assets **cannot** be modified or deleted
- Immutability is enforced only **after** a release is published
- Draft releases can still be modified/deleted (including their git tags)

When using `create` with assets, separate API calls are made: create as draft → upload assets → publish. Immutability protections apply only after publishing.

## Related

- For repository targeting (`--repo`), see [repo.md](repo.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
- For artifact attestations (`gh attestation`), see [gist-keys.md](gist-keys.md).
