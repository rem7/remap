# Change Log

**ATTN**: This project uses [semantic versioning](http://semver.org/).

## [Unreleased]

## [1.15.0] - 2016-04-30
### Added
- This file!
- Support for placeholders in flag usage strings
- `App.Metadata` map for arbitrary data/state management
- `Set` and `GlobalSet` methods on `*cli.Context` for altering values after
parsing.
- Support for nested lookup of dot-delimited keys in structures loaded from
YAML.

### Changed
- The `App.Action` and `Command.Action` now prefer a return signature of
`func(*cli.Context) error`, as defined by `cli.ActionFunc`.  If a non-nil
`error` is returned, there may be two outcomes:
    - If the error fulfills `cli.ExitCoder`, then `os.Exit` will be called
    automatically
    - Else the error is bubbled up and returned from `App.Run`
- Specifying an `Action` with the legacy return signature of
`func(*cli.Context)` will produce a deprecation message to stderr
- Specifying an `Action` that is not a `func` type will produce a non-zero exit
from `App.Run`
- Specifying an `Action` func that has an invalid (input) signature will
produce a non-zero exit from `App.Run`

### Deprecated
- <a name="deprecated-cli-app-runandexitonerror"></a>
`cli.App.RunAndExitOnError`, which should now be done by returning an error
that fulfills `cli.ExitCoder` to `cli.App.Run`.
- <a name="deprecated-cli-app-action-signature"></a> the legacy signature for
`cli.App.Action` of `func(*cli.Context)`, which should now have a return
signature of `func(*cli.Context) error`, as defined by `cli.ActionFunc`.

### Fixed
- Added missing `*cli.Context.GlobalFloat64` method

## [1.14.0] - 2016-04-03 (backfilled 2016-04-25)
### Added
- Codebeat badge
- Support for categorization via `CategorizedHelp` and `Categories` on app.

### Changed
- Use `filepath.Base` instead of `path.Base` in `Name` and `HelpName`.

### Fixed
- Ensure version is not shown in help text when `HideVersion` set.

## [1.13.0] - 2016-03-06 (backfilled 2016-04-25)
### Added
- YAML file input support.
- `NArg` method on context.

## [1.12.0] - 2016-02-17 (backfilled 2016-04-25)
### Added
- Custom usage error handling.
- Custom text support in `USAGE` section of help output.
- Improved help messages for empty strings.
- AppVeyor CI configuration.

### Changed
- Removed `panic` from default help printer func.
- De-duping and optimizations.

### Fixed
- Correctly handle `Before`/`After` at command level when no subcommands.
- Case of literal `-` argument causing flag reordering.
- Environment variable hints on Windows.
- Docs updates.

## [1.11.1] - 2015-12-21 (backfilled 2016-04-25)
### Changed
- Use `path.Base` in `Name` and `HelpName`
- Export `GetName` on flag types.

### Fixed
- Flag parsing when skipping is enabled.
- Test output cleanup.
- Move completion check to account for empty input case.

## [1.11.0] - 2015-11-15 (backfilled 2016-04-25)
### Added
- Destination scan support for flags.
- Testing against `tip` in Travis CI config.

### Changed
- Go version in Travis CI config.

### Fixed
- Removed redundant tests.
- Use correct example naming in tests.

## [1.10.2] - 2015-10-29 (backfilled 2016-04-25)
### Fixed
- Remove unused var in bash completion.

## [1.10.1] - 2015-10-21 (backfilled 2016-04-25)
### Added
- Coverage and reference logos in README.

### Fixed
- Use specified values in help and version parsing.
- Only display app version and help message once.

## [1.10.0] - 2015-10-06 (backfilled 2016-04-25)
### Added
- More tests for existing functionality.
- `ArgsUsage` at app and command level for help text flexibility.

### Fixed
- Honor `HideHelp` and `HideVersion` in `App.Run`.
- Remove juvenile word from README.

## [1.9.0] - 2015-09-08 (backfilled 2016-04-25)
### Added
- `FullName` on command with accompanying help output update.
- Set default `$PROG` in bash completion.

### Changed
- Docs formatting.

### Fixed
- Removed self-referential imports in tests.

## [1.8.0] - 2015-06-30 (backfilled 2016-04-25)
### Added
- Support for `Copyright` at app level.
- `Parent` func at context level to walk up context lineage.

### Fixed
- Global flag processing at top level.

## [1.7.1] - 2015-06-11 (backfilled 2016-04-25)
### Added
- Aggregate errors from `Before`/`After` funcs.
- Doc comments on flag structs.
- Include non-global flags when checking version and help.
- Travis CI config updates.

### Fixed
- Ensure slice type flags have non-nil values.
- Collect global flags from the full command hierarchy.
- Docs prose.

## [1.7.0] - 2015-05-03 (backfilled 2016-04-25)
### Changed
- `HelpPrinter` signature includes output writer.

### Fixed
- Specify go 1.1+ in docs.
- Set `Writer` when running command as app.

## [1.6.0] - 2015-03-23 (backfilled 2016-04-25)
### Added
- Multiple author support.
- `NumFlags` at context level.
- `Aliases` at command level.

### Deprecated
- `ShortName` at command level.

### Fixed
- Subcommand help output.
- Backward compatible support for deprecated `Author` and `Email` fields.
- Docs regarding `Names`/`Aliases`.

## [1.5.0] - 2015-02-20 (backfilled 2016-04-25)
### Added
- `After` hook func support at app and command level.

### Fixed
- Use parsed context when running command as subcommand.
- Docs prose.

## [1.4.1] - 2015-01-09 (backfilled 2016-04-25)
### Added
- Support for hiding `-h / --help` flags, but not `help` subcommand.
- Stop flag parsing after `--`.

### Fixed
- Help text for generic flags to specify single value.
- Use double quotes in output for defaults.
- Use `ParseInt` instead of `ParseUint` for int environment var values.
- Use `0` as base when parsing int environment var values.

## [1.4.0] - 2014-12-12 (backfilled 2016-04-25)
### Added
- Support for environment variable lookup "cascade".
- Support for `Stdout` on app for output redirection.

### Fixed
- Print command help instead of app help in `ShowCommandHelp`.

## [1.3.1] - 2014-11-13 (backfilled 2016-04-25)
### Added
- Docs and example code updates.

### Changed
- Default `-v / --version` flag made optional.

## [1.3.0] - 2014-08-10 (backfilled 2016-04-25)
### Added
- `FlagNames` at context level.
- Exposed `VersionPrinter` var for more control over version output.
- Zsh completion hook.
- `AUTHOR` section in default app help template.
- Contribution guidelines.
- `DurationFlag` type.

## [1.2.0] - 2014-08-02
### Added
- Support for environment variable defaults on flags plus tests.

## [1.1.0] - 2014-07-15
### Added
- Bash completion.
- Optional hiding of built-in help command.
- Optional skipping of flag parsing at command level.
- `Author`, `Email`, and `Compiled` metadata on app.
- `Before` hook func support at app and command level.
- `CommandNotFound` func support at app level.
- Command reference available on context.
- `GenericFlag` type.
- `Float64Flag` type.
- `BoolTFlag` type.
- `IsSet` flag helper on context.
- More flag lookup funcs at context level.
- More tests &amp; docs.

### Changed
- Help template updates to account for presence/absence of flags.
- Separated subcommand help template.
- Exposed `HelpPrinter` var for more control over help output.

## [1.0.0] - 2013-11-01
### Added
- `help` flag in default app flag set and each command flag set.
- Custom handling of argument parsing errors.
- Command lookup by name at app level.
- `StringSliceFlag` type and supporting `StringSlice` type.
- `IntSliceFlag` type and supporting `IntSlice` type.
- Slice type flag lookups by name at context level.
- Export of app and command help functions.
- More tests &amp; docs.

## 0.1.0 - 2013-07-22
### Added
- Initial implementation.

[Unreleased]: https://github.com/codegangsta/cli/compare/v1.15.0...HEAD
[1.15.0]: https://github.com/codegangsta/cli/compare/v1.14.0...v1.15.0
[1.14.0]: https://github.com/codegangsta/cli/compare/v1.13.0...v1.14.0
[1.13.0]: https://github.com/codegangsta/cli/compare/v1.12.0...v1.13.0
[1.12.0]: https://github.com/codegangsta/cli/compare/v1.11.1...v1.12.0
[1.11.1]: https://github.com/codegangsta/cli/compare/v1.11.0...v1.11.1
[1.11.0]: https://github.com/codegangsta/cli/compare/v1.10.2...v1.11.0
[1.10.2]: https://github.com/codegangsta/cli/compare/v1.10.1...v1.10.2
[1.10.1]: https://github.com/codegangsta/cli/compare/v1.10.0...v1.10.1
[1.10.0]: https://github.com/codegangsta/cli/compare/v1.9.0...v1.10.0
[1.9.0]: https://github.com/codegangsta/cli/compare/v1.8.0...v1.9.0
[1.8.0]: https://github.com/codegangsta/cli/compare/v1.7.1...v1.8.0
[1.7.1]: https://github.com/codegangsta/cli/compare/v1.7.0...v1.7.1
[1.7.0]: https://github.com/codegangsta/cli/compare/v1.6.0...v1.7.0
[1.6.0]: https://github.com/codegangsta/cli/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/codegangsta/cli/compare/v1.4.1...v1.5.0
[1.4.1]: https://github.com/codegangsta/cli/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/codegangsta/cli/compare/v1.3.1...v1.4.0
[1.3.1]: https://github.com/codegangsta/cli/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/codegangsta/cli/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/codegangsta/cli/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/codegangsta/cli/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/codegangsta/cli/compare/v0.1.0...v1.0.0
