## summaraizer google

Summarizes using Google

### Synopsis

To summarize using Google, you need to provide the API token.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.

```
summaraizer google [flags]
```

### Examples

```
summaraizer google --token <token>
summaraizer google --token <token> --model gemini-1.5-pro-002
```

### Options

```
  -h, --help            help for google
      --model string    The AI model to use (default "gemini-1.5-flash-8b")
      --prompt string   The prompt to use for the AI model (default "\nI give you a discussion and you give me a summary.\nEach comment of the discussion is wrapped in a <comment> tag.\nYour summary should not be longer than 1200 chars.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}\n")
      --token string    The API Token for Google
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

