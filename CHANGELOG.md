# CHANGELOG
All notable changes to this project will be documented in this file. 

## [0.1.0] - Initial Release 
### Added 
- `devlocal init` command to initialize configuration
- `devlocal apply` command to apply overlook rules
- `devlocal revert` command to remove overlook rules
- `devlocal cleanup` command to clean up Git state
- `devlocal status` command to inspect current state
-  YAML-based configuration (`devlocal.yaml`)
-  Support for Git `skip-worktree` to ignore local changes Basic CLI interface for managing overlooked files


### Notes 
- First usable version of DevLocal
- Focused on preventing accidental commits of local-only changes
- Designed to be simple, Git-native, and workflow-friendly
