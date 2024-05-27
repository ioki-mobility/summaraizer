# summaraizer

Summarizes comments from a variety of sources, such as GitHub issues, Reddit threads, and more 
using AI models from different providers, such as Ollama, Mistral, and OpenAI, and more.

## Installation

You can install the CLI by running the following command:
```bash
go install github.com/ioki-mobility/summaraizer/cmd/summaraizer@latest
```

## Usage

### CLI

The minimum required Go version can be found in the [go.mod](go.mod) file.

The code is split up into two parts.
The `source` part and the `summaraization` part.
First you need to fetch comments from a source, then you can summarize the comments.

Obviously, you can also do both independently.

#### Fetch comments

The general usage is:
```bash
summaraizer [SOURCE] [ARGUMENTS] 
```

Example sources are:
* `github`
* `gitlab`
* `reddit`

#### Summarize comments

The general usage is:
```bash
summaraizer [PROVIDER] [ARGUMENTS]
```

Example providers are:
* `ollama`
* `openai`
* `mistral`

Please note that the provider reads from `Stdin` as well as require a special JSON format as input:
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

#### Examples (summarize comments from a source)

Each provided sources already respects the JSON format required by the summarization providers.
Being said, you can use of pipe to chain the commands together.

```bash
summaraizer github --issue golang/go/66960 | summaraizer ollama --model llama3
```

This command fetches the comments of [this GitHub issue](https://github.com/golang/go/issues/66960)
and summarizes them using the `llama3` model via Ollama.

```bash
summaraizer reddit --post 'r/ArtificialInteligence/comments/1d16cxl/miss_ai_worlds_first_beauty_contest_with_computer/' | summaraizer openai --model gpt-4o --token SUPER_SECRE_API_TOKEN
```

This command fetches the comments of [this Reddit post](https://www.reddit.com/r/ArtificialInteligence/comments/1d16cxl/miss_ai_worlds_first_beauty_contest_with_computer/)
and summarizes them using the `gpt-4o` model via OpenAI.
