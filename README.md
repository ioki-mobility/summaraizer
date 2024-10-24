[![Go Reference](https://pkg.go.dev/badge/github.com/ioki-mobility/summaraizer.svg)](https://pkg.go.dev/github.com/ioki-mobility/summaraizer)
[![Integration Tests](https://github.com/ioki-mobility/summaraizer/actions/workflows/integration_tests.yml/badge.svg)](https://github.com/ioki-mobility/summaraizer/actions/workflows/integration_tests.yml)

# summaraizer

Summarizes comments from a variety of sources, such as GitHub issues, GitLab issues, Slack threads, Reddit threads, and more 
using AI models from different providers, such as Ollama, Anthropic, OpenAI and Google.


## Usage

Check the project website how to use it:
</br>[https://ioki-mobility.github.io/summaraizer](https://ioki-mobility.github.io/summaraizer/)

## Copy & Paste snippets

<details>
    <summary>Fetch from GitHub issue and summarize using Ollama with llama3</summary>

```bash
summaraizer github --issue golang/go/66960 | summaraizer ollama --model llama3
```

</details>

<details>
    <summary>Fetch from Reddit post and summarize using OpenAI with gpt-4o</summary>

```bash
summaraizer reddit --post 'r/ArtificialInteligence/comments/1d16cxl/miss_ai_worlds_first_beauty_contest_with_computer/' | summaraizer openai --model gpt-4o --token SUPER_SECRET_API_TOKEN
```

</details>

<details>
    <summary>Fetch from GitLab issue with custom instance and summarize using Google</summary>

```bash
summaraizer gitlab --issue sre/it-support/203 --token SUPER_SECRET_API_TOKEN --url https://gitlab.url.com | summaraizer google --token SUPER_SECRET_API_TOKEN
```

</details>

<details>
    <summary>Fetch from Slack thread and summarize using Antrhopic</summary>

```bash
summaraizer slack --channel C07ED7YBB1P --ts 1723214080.317439 --token SUPER_SECRET_API_TOKEN | summaraizer anthropic --token SUPER_SECRET_API_TOKEN
```

</details>
