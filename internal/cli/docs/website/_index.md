# summaraizer
Summarizes comments from a variety of sources, such as GitHub issues, Slack threads, Reddit posts, and more
using AI models from different providers, such as Ollama, OpenAI, and more.

## CLI
### Installation

- Download pre-built binaries from the [releases](https://github.com/ioki-mobility/summaraizer/releases/latest) page
- Install via go toolchain:
```shell
go install github.com/ioki-mobility/summaraizer/cmd/summaraizer@latest
```

### Source vs Summarization

The CLI is split up into two parts.
The **source** part and the **summarization** part.
*First* you need to fetch comments from a source, *then* you can summarize the comments.

Of course, you can also do both independently.

#### Source aka. fetching comments

The general usage is:
```bash
summaraizer [SOURCE] [flags] 
```

Example sources are:
* `github`
* `gitlab`
* `slack`
* `reddit`

The source writes to `Stdout` and writes the output in a JSON format:
```json
[
    {
        "author": "Author1",
        "body": "Body1"
    },
    {
        "author": "Author2",
        "body": "Body2"
    }
]
```

### Summarization aka. calling providers

The general usage is:
```bash
summaraizer [PROVIDER] [FLAGS]
```

Example providers are:
* `ollama`
* `anthropic`
* `openai`
* `google`

The provider reads from `Stdin` and requires a special JSON format as input:
```json
[
    {
        "author": "Author1",
        "body": "Body1"
    },
    {
        "author": "Author2",
        "body": "Body2"
    }
]
```

### Configuration and flags

Each source and provider has a unique set of flags. 
Some flags are required, while others are optional. 
If you do not provide the required flags, you will be prompted to enter them.

Alternatively, you can specify app flags (both required and optional) 
in a configuration file named `config.json`.
The JSON format uses the same key names as the flag names,
and each provider or source can be included as an optional object in the JSON file.

Example configuration file:
```json
{
  "github": {
    "token": "[TOKEN]"
  },
  "ollama": {
    "model": "llama3.2:3b"
  },
  "openai": {
    "token": "[TOKEN]",
    "model": "gpt-4o-mini"
  },
  "anthropic": {
    "prompt": "Count the comments in the following discussion.\nEach comment is divided into a <comment> tag.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}"
  }
}
```

You can also provide flags via environment variables.
The env. variables should have a prefix of `SUMMARAIZER_` and the provider/source name in uppercase,
followed by the flag name in uppercase and separated by an underscore.

For example
```bash
export SUMMARAIZER_GITHUB_TOKEN=[TOKEN]
export SUMMARAIZER_OPENAI_TOKEN=[API-TOKEN]

// The following command will use the environment variables
summaraizer github | summaraizer openai
```

A combination of both (config file and env. variables) is also possible of course:
```bash
cat config.json
{
  "gitlab": {
    "url": "https://gitlab.example.com",
  },
  "google": {
    "model": "gemini-1.5-pro"
  }
}

export SUMMARAIZER_GOOGLE_TOKEN=[API-TOKEN]

summaraizer gitlab --issue ioki-mobility/issue/1 | summaraizer google
```

The CLi searches in the following directories for the configuration file:
* `.` (current directory)
* `$HOME/.config/summaraizer/config.json`
* `$HOME/.summaraizer/config.json` 

### Tips and tricks

#### Custom source

Using a provider to summarize from a custom¹ source:
```bash 
echo '[{"author": "Author1", "body": "I like to"}, {"author": "Author2", "body": "Move it!"}]' | summaraizer ollama
```

¹: Custom source means that you can provide your own comments in the JSON format.

#### Custom prompt

Customize the prompt to fit your needs, using Go templates:
```bash
summaraizer github --issue ioki-mobility/summaraizer/1 | summaraizer ollama --prompt 'Please count the comments in the following discussion.\nEach comment is divided into a <comment> tag.\nHere is the discussion:\n{{ range $comment := . }}\n<comment>{{ $comment.Body }}</comment>\n{{end}}'
```

The prompt will receive a `struct` in form of `Comments`, with the following structure:
```go
type Comments []Comments

type Comment struct {
    Author string
    Body   string
}
```

## Go module
### Installation

```bash
go get github.com/ioki-mobility/summaraizer@latest
```

### Example

```go
package main

import (
	"bytes"
	"fmt"
	"log"
	
	"github.com/ioki-mobility/summaraizer"
)

func main() {
	buffer := bytes.Buffer{}
	gh := summaraizer.GitHub{
		RepoOwner:   "ioki-mobility",
		RepoName:    "summaraizer",
		IssueNumber: "1",
	}
	err := gh.Fetch(&buffer)
	if err != nil {
		log.Fatal(err)
	}

	openAi := summaraizer.OpenAi{
		Model:    "gpt-4o-mini",
		Prompt:   "A prompt that can make use of templates. See the Comments type",
		ApiToken: "API-TOKEN",
	}
	summarization, err := openAi.Summarize(&buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(summarization)
}
```
