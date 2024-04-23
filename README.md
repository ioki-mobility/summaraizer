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
go run cli/cli.go -owner <owner> -repo <repo> -issue-number <issue-number> [-ai-provider <ai-provider>] [-ai-model <ai-model>] [-token <token>]
```

The `ai-provider` can be changed to one of our supported providers. 
Right now only `ollama` is supported. 
Defaults to `ollama`.

The `ai-model` can be changed to any model supported by the `ai-provider`.
Defaults to `mistral:7b`.

The GitHub `token` is only required for private repositories.

**Example:**

```bash 
go run cli/cli.go -owner golang -repo go -issue-number 66960
```

**This resulted with `Ollama` & `mistral:7b` to the following output:**

The proposed `sync.LazyChannel` implementation using `sync.OnceValue` has some issues. Here's why:

1. It requires the user to initialize the channel value with an anonymous function, which is not a common idiom and may lead to confusion for new Go programmers.
2. The initialization of the channel happens outside the scope of the `sync.LazyChannel` type, making it harder to reason about when the channel is actually initialized and ready to use.
3. There is no way to support buffered channels using this approach because there is no way to pass a buffer size as a type parameter in Go.
4. The implementation of `sync.OnceValue` uses an internal mutex for synchronization during initialization, which may add unnecessary contention in multithreaded scenarios.

An alternative approach that addresses some of these issues is using a struct with a single field (the channel) and a method to initialize it on demand using `sync.Once`. Here's the updated implementation:

```go
type LazyChannel[T] struct {
        ch chan T
        once sync.Once
}

func (lc *LazyChannel[T]) Ch() chan T {
        if lc.ch == nil {
                lc.once.Do(func () { lc.ch = make(chan T) })
        }
        return lc.ch
}
```

This implementation uses a struct with a single field (the channel) and a method `Ch()` to initialize the channel on demand using `sync.Once`. 
The initialization only happens once, making the zero value of the `LazyChannel` type ready to use.

This approach does not require the user to initialize the channel value with an anonymous function and keeps all initialization logic within the `LazyChannel` type. 
It also supports buffered channels by simply passing the desired buffer size when creating the `make(chan T, bufferSize)`. 
The implementation of `sync.Once` uses an internal mutex for synchronization during initialization, but only once per goroutine, making it more efficient than the earlier proposed approach.

This implementation still requires some boilerplate code to set up the type and method, 
but it's a more straightforward solution that better fits the Go idiom of composing types to build more complex data structures.

**End of output**