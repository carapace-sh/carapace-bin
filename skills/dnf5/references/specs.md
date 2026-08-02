# DNF5 Package Specs

Pattern matching rules for package, group, environment, module, and transaction specifications in DNF5.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/specs.7.html>.

## Globs

DNF5 supports the same glob pattern matching as the shell:

| Pattern | Matches |
|---------|---------|
| `*` | Any number of characters |
| `?` | Any single character |
| `[abc]` | Any one of the enclosed characters |
| `[a-z]` | Any character in the range |
| `[!abc]` or `[^abc]` | Any character NOT enclosed |

**Note**: Curly brackets `{}` are **not supported** by DNF5. Shells that support them can expand them before DNF5 sees the argument.

## Package Specs

Many commands take a `<package-spec>` parameter. Matching is tried sequentially against:

1. **NEVRA** [N] or case-insensitive NEVRA [NI]
2. **Provides** [P]
3. **File provides** [F]
4. **Binaries** [B]

For `<package-spec-NP>`, matching is first against NEVRAs, then Provides (if no NEVRA match).

### NEVRA Matching

Each package is uniquely identified by NEVRA:

| Component | Description |
|-----------|-------------|
| **Name** | Package name |
| **Epoch** | Epoch number (not always included). Overrides other version checking. |
| **Version** | Version string (not strictly numeric, matches upstream software version) |
| **Release** | Edition string (particular package build, usually a number) |
| **Architecture** | Target processor type. Can be `src` (source) or `noarch` (architecture-independent). |

### Partial Matching

DNF5 tries the spec against these forms in decreasing priority:

1. `NAME-[EPOCH:]VERSION-RELEASE.ARCH`
2. `NAME.ARCH`
3. `NAME`
4. `NAME-[EPOCH:]VERSION-RELEASE`
5. `NAME-[EPOCH:]VERSION`

The first form that matches any packages is used; remaining forms are not tried. If none match, an attempt is made to match against full package NEVRAs (relevant if globs are present).

Globs can be specified as part of any NEVRA component, or across multiple NEVRA components (matching across separators like `-` and `.`). In the latter case, the spec must match against full package NEVRAs.

### Provides Matching

Users can specify version comparisons for provides:

```
<provide>
<provide> = <version>
<provide> > <version>
<provide> >= <version>
<provide> < <version>
<provide> <= <version>
```

Rich dependency expressions (boolean operators, nesting) are supported:

```
(<provide1> or (<provide2> and <provide3>))
```

### File Provides Matching

If a spec starts with `/` or `*/`, it is considered a potential file provide. DNF5 checks if any package provides that file.

### Binaries Matching

DNF5 checks if the given spec is a binary in `/usr/bin/` or `/usr/sbin/`.

## Comps Specs

`<group-spec>` and `<environment-spec>` are case-insensitive strings (supporting globs) matched against:

- Group's/environment's **ID**
- Group's/environment's **canonical name**
- Group's/environment's **name** translated into the current `LC_MESSAGES` locale

Comps specs are prefixed by `@` for commands that also accept package specs. Group and environment commands prefer their corresponding type; other commands affect both types.

## Module Specs

`<module-spec>` identifies modules/profiles via the **NSVCA** format:

```
NAME:STREAM:VERSION:CONTEXT:ARCH/PROFILE
```

### Supported Partial Forms

- `NAME`
- `NAME:STREAM`
- `NAME:STREAM:VERSION`
- `NAME:STREAM:VERSION:CONTEXT`
- All above with `::ARCH` (e.g., `NAME::ARCH`)
- `NAME:STREAM:VERSION:CONTEXT:ARCH`
- All above with `/PROFILE` (e.g., `NAME/PROFILE`)

### Defaults

- If **stream** is not specified, the enabled or default stream is used (in that order).
- If **profile** is not specified, the system default profile or 'default' profile is used.

## Transaction Specs

`<transaction-spec>` can be:

| Spec | Description |
|------|-------------|
| Integer | Specifies a transaction ID |
| `last` | Same as the ID of the most recent transaction |
| `last-<offset>` | Where `<offset>` is a positive integer, specifying the offset-th transaction preceding the most recent |

### Examples

- `dnf5 history info 42` — Show info for transaction #42
- `dnf5 history info last` — Show info for the most recent transaction
- `dnf5 history info last-1` — Show info for the transaction before the most recent
