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

func main() {
    debug := env.GetBool("DEBUG", false)
    println("Debug mode: ", debug)
}
```