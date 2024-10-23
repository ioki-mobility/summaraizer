## summaraizer ollama

Summarizes using Ollama

### Synopsis

To summarize using Ollama, you *can* provide the URL where Ollama is accessible.
If you are running Ollama locally, you can use the default URL. There is no need to provide the URL.
Optional flags are the AI model and the prompt. The prompt can make use of Go template functions.

```
summaraizer ollama [flags]
```

### Examples

```
summaraizer ollama
summaraizer ollama --model llama3.2:3b
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

