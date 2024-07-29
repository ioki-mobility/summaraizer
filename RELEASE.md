# Release

The release process is fully automated via GitHub Actions.
To dispatch a release, go to the Actions tab and select the ["Release" workflow](https://github.com/ioki-mobility/summaraizer/actions/workflows/release.yml).
Click on the "Run workflow" button and fill in the tag name.
The workflow will then do the following:
* Run integration tests
* Build the CLI binaries for Linux (amd64), macOS (amd64, arm64), and Windows (amd64)
* Create a git tag with the specified tag name
* Create a GitHub release and publish it with an automatically generated changelog

It will create a draft release by default, so you can review the changelog before publishing it.

## Tag name

Please follow the [Go module versioning scheme](https://go.dev/doc/modules/version-numbers) for the tag name.