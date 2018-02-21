# go-env
Simple go lib for environment variables with default values.
## Use

Functions:
- GetBool()
- GetString()
- GetInt()
- GetFloat()

### Instalation
```
	go get github.com/amonsat/go-env
```

### Basic Use

```go
package main

import env "github.com/amonsat/go-env"

func main() {
	debug := env.GetBool("DEBUG", false)
	println("Debug mode: ", debug)
}
```