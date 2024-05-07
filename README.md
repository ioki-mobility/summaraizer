# summaraizer

Summarize GitHub issue (or pull request) comments.

## Requirements

* Go installed (version `1.22.2`, `brew install go`)
* Ollama installed (`brew install ollama`)
* Any model installed via Ollama
    * `ollama serve`
    * `ollama pull mistral:7b`

## Usage

```bash
go run cmd/summaraizer/summaraizer.go <aiProvider> <customAiProviderParams> --owner <owner> --repo <repo> --issue-number <issueNumber> [--token <token>]
```

You can also run `--help` to get the list of available `commands`, `arguments` and `flags`.

| Command                  | Description                                                                        |
|--------------------------|------------------------------------------------------------------------------------| 
| `aiProvider`             | One of `ollama`, `mistral`, `openai`. Defaults to `ollama`.                        |
| `customAiProviderParams` | The custom parameters for the AI provider. Depending on the selected `aiProvider`. |
| `--owner`                | The owner of the GitHub repository.                                                |
| `--repo`                 | The GitHub repository name.                                                        |
| `--issue-number`         | The GitHub issue number.                                                           |
| `--token`                | (Optional) The GitHub API token. It is only required for private repos.            |

### Common AIProvider Parameters

There are some common parameters that can be used with all AI providers, namely, `--ai-model` and `--ai-prompt`.
The following sections describe these parameters.

#### AiModel

The `--ai-model` parameter is used to specify the model to be used by the AI provider.
You have to check the documentation of the AI provider to know which models are available.

#### AiPrompt

The `--ai-prompt` parameter is used to specify the prompt to be used by the AI provider.
The prompt make use of Go's [`text/template` package](https://pkg.go.dev/text/template) to render the prompt.
The prompt will receive a `struct` in form of `Comments`, with the following structure:
```go
type Comments []Comment

type Comment struct {
    Author string
    Body   string
}
```
You could, for example, inject the following prompt:
```go
var myAwesomeTemplate = `
Please summarize the following discussions between different Authors.
{{ range $comment := . }}
Author {{ $comment.Author }} said: {{ $comment.Body }}
{{ end }}
`
```

## Example

```bash 
go run cmd/summaraizer/summaraizer.go ollama --url http://localhost:11434 --ai-model llama3 --owner golang --repo go --issue-number 66960
```

This will run the CLI with the `ollama` AI provider, pointing to a local `ollama` instance, 
using the `llama3` model, for [`golang/go/issues/66960`](https://github.com/golang/go/issues/66960).

<details>

<summary><b>The example above produced the following output:</b></summary>

Here is a summary of the discussion:

The proposal is for a new type `atomic.Chan` that allows for atomic access to a channel. 
The motivation behind this proposal is to improve existing code that uses `atomic.Value` and `*hchan` to create lazy channels,
which have 488 matches in GitHub search results.

The current implementation of lazy channels requires the use of `sync.Once`, 
which has some drawbacks such as increasing the footprint and making on-demand channel swapping more complicated. 
The proposed `atomic.Chan` type would enable no new code but improve existing code by providing a more efficient 
and clear way to implement lazy channels.

Some examples of where this proposal could be used include
[pkg context](https://github.com/golang/go/blob/go1.22.2/src/context/context.go#L425), 
[gRPC-Go](https://github.com/grpc/grpc-go/blob/v1.63.2/internal/transport/controlbuf.go#L311), and others.

The discussion also touched on the idea that if this proposal is approved, 
it could set a precedent for approving similar proposals for other Go types
that are secretly just pointers to a special data structure or involve pointers, such as `string`, `slice`, and `map`.

</details>