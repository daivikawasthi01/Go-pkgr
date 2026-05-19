# pkgr

> A fast, minimal package manager built in Go — inspired by npm.

pkgr resolves packages from a registry, tracks them in a JSON lockfile, and gives you a clean CLI interface to manage dependencies. Built as a learning project to explore Go's project structure, file I/O, JSON marshaling, CLI design, and testing patterns.

---

## Table of Contents

- [Installation](#installation)
- [Commands](#commands)
- [Lockfile](#lockfile)
- [Registry](#registry)
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
- [Testing](#testing)
- [Built With](#built-with)

---

## Installation

Requires Go 1.26+.

```bash
git clone https://github.com/daivikawasthi01/Go-pkgr.git
cd Go-pkgr
go build -o pkgr .
```

---

## Commands

### Install a package
```bash
./pkgr install <package>
```
Resolves the package from the registry and adds it to `pkgr.lock`.

```bash
$ ./pkgr install gin
Resolving gin...
Installed gin@1.9.1
```

### Remove a package
```bash
./pkgr remove <package>
```
Removes the package entry from `pkgr.lock`.

```bash
$ ./pkgr remove gin
Removed gin
```

### List installed packages
```bash
./pkgr list
```
Reads `pkgr.lock` and prints all currently installed packages.

```bash
$ ./pkgr list
Installed packages (2):
  cobra@1.10.2
  viper@1.18.0
```

### Version
```bash
./pkgr version
```
```bash
$ ./pkgr version
pkgr v0.1.0
```

### Help
```bash
./pkgr --help
./pkgr [command] --help
```

---

## Lockfile

Every install and remove operation updates `pkgr.lock` — a JSON file that acts as the source of truth for what's installed. It is auto-created on first install and should not be committed to version control.

```json
{
  "packages": {
    "cobra": {
      "name": "cobra",
      "version": "1.10.2",
      "source": "github.com/spf13/cobra"
    },
    "gin": {
      "name": "gin",
      "version": "1.9.1",
      "source": "github.com/gin-gonic/gin"
    }
  }
}
```

---

## Registry

pkgr ships with a built-in package registry. The following packages are currently available:

| Package | Version | Source |
|---------|---------|--------|
| cobra | 1.10.2 | github.com/spf13/cobra |
| gin | 1.9.1 | github.com/gin-gonic/gin |
| viper | 1.18.0 | github.com/spf13/viper |
| zap | 1.27.0 | go.uber.org/zap |
| chi | 5.0.12 | github.com/go-chi/chi |

---

## Project Structure

```
pkgr/
├── main.go                       # Entry point — boots the CLI
├── go.mod                        # Module definition and dependencies
├── go.sum                        # Dependency checksums
├── cmd/
│   ├── root.go                   # Root command, version subcommand
│   ├── install.go                # install command
│   ├── remove.go                 # remove command
│   └── list.go                   # list command
└── internal/
    ├── registry/
    │   └── registry.go           # Package resolution logic
    └── lockfile/
        ├── lockfile.go           # Lockfile read/write/mutate logic
        └── lockfile_test.go      # Unit tests for lockfile package
```

The `internal/` directory is intentionally private — Go enforces that packages inside `internal/` cannot be imported by external modules, keeping core logic encapsulated.

---

## How It Works

1. **Install** — `pkgr install <name>` calls `registry.Resolve()` to look up the package, then loads the current `pkgr.lock`, adds the entry, and saves it back to disk.
2. **Remove** — `pkgr remove <name>` loads the lockfile, deletes the entry by key, and saves.
3. **List** — `pkgr list` loads the lockfile and iterates over the packages map to print each entry.
4. **Lockfile persistence** — the lockfile is read fresh on every command and written back immediately after mutation. No in-memory state is held between commands.

---

## Testing

Unit tests cover the core lockfile package — add, remove, save, and load operations.

```bash
# Run tests
go test ./internal/lockfile/...

# Run with verbose output
go test ./internal/lockfile/... -v
```

```
=== RUN   TestAddAndRetrieve
--- PASS: TestAddAndRetrieve (0.00s)
=== RUN   TestRemovePackage
--- PASS: TestRemovePackage (0.00s)
=== RUN   TestRemoveNonExistent
--- PASS: TestRemoveNonExistent (0.00s)
=== RUN   TestSaveAndLoad
--- PASS: TestSaveAndLoad (0.00s)
PASS
ok      github.com/daivikawasthi01/Go-pkgr/internal/lockfile    0.225s
```

---

## Built With

- [cobra](https://github.com/spf13/cobra) — CLI framework used by kubectl, gh, and hugo
