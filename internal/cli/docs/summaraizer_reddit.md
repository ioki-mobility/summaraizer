## summaraizer reddit

Fetches comments from a Reddit post

### Synopsis

To summarize a Reddit post, use the URL path. Everything after reddit.com without the leading slash.
Note that we only fetch the top-level comments. Nested comments are ignored.

```
summaraizer reddit [flags]
```

### Examples

```
summaraizer reddit --post r/ArtificialInteligence/comments/1d16cxl/miss_ai_worlds_first_beauty_contest_with_computer/
```

### Options

```
  -h, --help          help for reddit
      --post string   The Reddit post to summarize. Use the URL path.
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

