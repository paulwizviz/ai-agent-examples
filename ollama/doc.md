# Ollama Working Examples

[API Documentation](https://github.com/ollama/ollama/blob/main/docs/api.md)

## Go

* Example 1
  * Text completion
  * Based on the official [Ollama API](https://github.com/ollama/ollama) to interact with the server.
  * Response with no streaming.
  * [Source](./gollama/ex1/main.go).
* Example 2
  * Same as Example 1 with no streaming response.
  * [Source](./gollama/ex2/main.go).
* Example 3
  * Using custom client.
  * Request the server to analyse a picture of a [kitten](../testdata/cat.jpeg).
  * [Source](./gollama/ex3/main.go).
* Example 4
  * Chat scenario where the model is the role of an assistant.
  * A user sends a greeting by saying Hi and then followed by asking the model to tell him about Italy in one sentence.
  * [Source](./gollama/ex4/main.go)
* Example 5
  * Text response
  * Asking for a response in JSON
  * [Source](./gollama/ex5/main.go)
