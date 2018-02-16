# go-env
Simple go lib for environment variables with default values.
## Use

### Basic Use

```go
package main

import env "github.com/Amonsat/go-env"

func main() {
	debug := env.GetBool("DEBUG", false)
	println("Debug mode: ", debug)
}
```