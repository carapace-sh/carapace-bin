# Comps Groups and Environments

Comps files group packages into functional units. They are stored in repository metadata under the `comps.xml` filename.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/comps.7.html>. For specs, see [specs.md](specs.md).

## Two Types of Structures

| Type | Composed Of |
|------|-------------|
| **Group** | Lists of packages |
| **Environment** | Groups (mandatory and optional) |

### Environment Structure

- Each environment is made of **mandatory** and **optional** groups.
- **Mandatory groups** must all be installed for the environment to be considered installed.
- **Optional groups** are not installed by default. Add with `--with-optional`.

## Group Package Levels

| Level | Description | Default Behavior |
|-------|-------------|------------------|
| **mandatory** | Essential for the group's functionality. Must be installed for group to be considered installed. | Installed |
| **default** | Installed together with mandatory packages. | Installed (can be excluded with `--exclude=PACKAGE`) |
| **optional** | Not installed by default. Can be included with `--with-optional`. | Not installed |
| **conditional** | Brought into transaction if their required package is to be installed. | Conditional |

## Configuration

The `group_package_types` config option controls which package types are installed by default:

```
group_package_types = default,mandatory,conditional
```

## Related Commands

- `dnf5 group` — Manage comps groups
- `dnf5 environment` — Manage comps environments

## Related

- [specs.md](specs.md) — `<group-spec>` and `<environment-spec>` patterns
- [configuration.md](configuration.md) — `group_package_types`, `excludegroups`, `excludeenvs`
