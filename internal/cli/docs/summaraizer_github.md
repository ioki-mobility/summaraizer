## summaraizer github

Fetches comments from GitHub

### Synopsis

To summarize a GitHub issue, use the format owner/repo/issue_number.
At GitHub terminology, a pull request is also an issue. Therefore, you can summarize a pull request using the same format.
If the repository is private, you need to provide a token.

```
summaraizer github [flags]
```

### Examples

```
summaraizer github --issue ioki-mobility/summaraizer/1
summaraizer github --issue ioki-mobility/summaraizer/1 --token <token>
```

### Options

```
  -h, --help           help for github
      --issue string   The GitHub issue to summarize. Use the format owner/repo/issue_number.
      --token string   The GitHub token. Only required for private repositories.
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

