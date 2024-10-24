## summaraizer anthropic

Summarizes using Anthropic

### Synopsis

To summarize using Anthropic, you need to provide the API token.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.

```
summaraizer anthropic [flags]
```

### Examples

```
summaraizer anthropic --token <token>
summaraizer anthropic --token <token> --model claude-3-5-sonnet-20241022
```

### Options

```
  -h, --help            help for anthropic
      --model string    The AI model to use (default "claude-3-haiku-20240307")
      --prompt string   The prompt to use for the AI model (default "\nI give you a discussion and you give me a summary.\nEach comment of the discussion is wrapped in a <comment> tag.\nYour summary should not be longer than 1200 chars.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}\n")
      --token string    The API Token for Anthropic
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

