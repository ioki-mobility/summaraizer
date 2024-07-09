## summaraizer ollama

Summarizes using Ollama AI

### Synopsis

Summarizes using Ollama AI.

```
summaraizer ollama [flags]
```

### Options

```
  -h, --help            help for ollama
      --model string    The AI model to use (default "gemma:2b")
      --prompt string   The prompt to use for the AI model (default "\nI give you a discussion and you give me a summary.\nEach comment of the discussion is wrapped in a <comment> tag.\nYour summary should not be longer than 1200 chars.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}\n")
      --url string      The URl where ollama is accessible (default "http://localhost:11434")
```

### SEE ALSO

* [summaraizer](summaraizer.md)	 - Summarizes comments

