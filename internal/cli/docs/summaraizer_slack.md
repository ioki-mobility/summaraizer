## summaraizer slack

Fetches comments from a Slack thread

### Synopsis

To summarize a Slack thread, you need to provide the token, the channel ID, and the timestamp of the thread.
You can get the channel ID and the timestamp from the URL of the thread.

```
summaraizer slack [flags]
```

### Examples

```
summaraizer slack --token <token> --channel <channel_id> --ts <timestamp>
```

### Options

```
      --channel string   The channel ID of the Slack thread.
  -h, --help             help for slack
      --token string     The Slack token.
      --ts string        The timestamp of the Slack thread.
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

