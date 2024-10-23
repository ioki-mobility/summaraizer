## summaraizer gitlab

Fetches comments from GitLab issues

### Synopsis

To summarize a GitLab issue, use the format owner/repo/issue_number. 
You always have to provide a token.
If you have a custom GitLab instance, you can provide the URL with the --url flag.
Note that we only fetch the top-level comments. Nested comments are ignored.

```
summaraizer gitlab [flags]
```

### Examples

```
summaraizer gitlab --issue ioki-mobility/summaraizer/1 --token <token>
summaraizer gitlab --issue ioki-mobility/summaraizer/1 --token <token> --url https://gitlab.url.com
```

### Options

```
  -h, --help           help for gitlab
      --issue string   The GitLab issue to summarize. Use the format owner/repo/issue_number.
      --token string   The GitLab token.
      --url string     The URL of the GitLab instance. (default "https://gitlab.com")
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

