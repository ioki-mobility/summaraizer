## summaraizer mistral

Summarizes using Mistral

### Synopsis

To summarize using Mistral, you need to provide the API token.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.

```
summaraizer mistral [flags]
```

### Examples

```
summaraizer mistral --token <token>
summaraizer mistral --token <token> --model pixtral-12b-2409
```

### Options

```
  -h, --help            help for mistral
      --model string    The AI model to use (default "pixtral-12b-2409")
      --prompt string   The prompt to use for the AI model (default "\nI give you a discussion and you give me a summary.\nEach comment of the discussion is wrapped in a <comment> tag.\nYour summary should not be longer than 1200 chars.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}\n")
      --token string    The API Token for Mistral
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

