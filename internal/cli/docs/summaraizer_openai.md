## summaraizer openai

Summarizes using OpenAI AI

### Synopsis

Summarizes using OpenAI AI.

```
summaraizer openai [flags]
```

### Options

```
  -h, --help            help for openai
      --model string    The AI model to use (default "gpt-3.5-turbo")
      --prompt string   The prompt to use for the AI model (default "\nI give you a discussion and you give me a summary.\nEach comment of the discussion is wrapped in a <comment> tag.\nYour summary should not be longer than 1200 chars.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}\n")
      --token string    The API Token for OpenAI
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

