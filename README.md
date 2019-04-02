# go-env

Simple go lib for environment variables with default values.

## Use

Functions:

- GetBool()
- GetString()
- GetInt()
- GetInt64()
- GetFloat32()
- GetFloat64()
- GetUint()
- GetUint64()

### Instalation

```go
    go get github.com/amonsat/go-env
```

### Basic Use

```go
package main

import env "github.com/amonsat/go-env"

var (
    debug := env.GetBool("DEBUG", false)
    mongoAddr := env.GetString("MONGO_ADDR", "mongodb://127.0.0.1:27017")
)

func main() {
    println("Debug mode: ", debug)
    println("Mongo addr: ", mongoAddr)
}
```