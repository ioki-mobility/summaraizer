## summaraizer mistral

Summarizes using Mistral AI

### Synopsis

Summarizes using Mistral AI.

```
summaraizer mistral [flags]
```

### Options

```
  -h, --help            help for mistral
      --model string    The AI model to use (default "mistral:7b")
      --prompt string   The prompt to use for the AI model (default "\nI give you a discussion and you give me a summary.\nEach comment of the discussion is wrapped in a <comment> tag.\nYour summary should not be longer than 1200 chars.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}\n")
      --token string    The API Token for Mistral
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

