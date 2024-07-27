# Chainable Error
[![Version](https://img.shields.io/badge/version-1.0.3-blue.svg)]()

Chainable error type allows create chainable and checkable errors.

## Why?
We can create error object, and return it to the caller,
but we can't add error from bottom level in it (for logging for example).

ChainableError resolves it.

### Example:
```go
var ErrMyError = NewChainableError("my error")

func foobar() error {
	err := someFunction()
	if err != nil {
		return ErrMyError.Wrap(err)
	}
	return nil
}

func main() {
	err := foobar()
	if errors.Is(err, ErrMyError) {
		doSomeLogic()
	}
}
```
