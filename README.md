# blablo
TS/JS/Bash/Go logger

## How to deploy manually

### Go

- before commit/merge changes to `main`, bump `latest-version.txt` version
- commit/merge changes to `main`
- create tag = `latest-version.txt`, e.g. `v0.3.0`
- `git push --tags`
- `make go-publish`
