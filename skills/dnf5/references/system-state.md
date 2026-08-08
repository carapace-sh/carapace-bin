# DNF5 System State

The DNF5 system state consists of TOML files storing package install reasons, installed groups, and installed environments.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/system-state.7.html>.

## Location

Determined by the `system_state_dir` configuration option. Default: `/usr/lib/sysimage/libdnf5`.

## Purposes

DNF5 uses the system state to:

1. **Store package install reasons** — why each installed package was added to the system
2. **Track installed groups** and packages installed by these groups
3. **Track installed environmental groups**

## Package Install Reasons

| Reason | Description |
|--------|-------------|
| `user` | Packages the user explicitly asked DNF5 to install |
| `dependency` | Packages pulled in as dependencies |
| `weak dependency` | Packages pulled in as weak dependencies (Recommends/Supplements) |
| `group` | Packages installed by a group |
| `external` | Packages installed by another tool (e.g., `rpm`) |

## Important Notes

- The way of storing the DNF5 system state is an **internal implementation detail** and may change at any time.
- To modify the state, always use the DNF5 **command-line interface** or DNF5 **API**.
- Do not edit the TOML files directly.

## Recovering from Corrupted System State

If system state files become corrupted:

1. Back up the corrupted TOML file mentioned in the error message
2. Remove the corrupted file
3. It will be **regenerated** during the next successful DNF5 transaction

**Warning**: Regenerated files **may lack some data**, such as:
- The reasons why packages were installed
- The repositories from which packages were installed

## State Files vs. DNF4

DNF4 and DNF5 use different state file locations and formats. Transaction history is not migrated when switching from DNF4 to DNF5. See [migration.md](migration.md).

## Related

- [configuration.md](configuration.md) — `system_state_dir`, `transaction_history_dir`
- [migration.md](migration.md) — State file differences between DNF4 and DNF5
