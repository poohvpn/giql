# GiQL

Elegant [GraphiQL](https://github.com/graphql/graphiql) server handler for Go.

## Usage

```go
package main

import (
	"github.com/poohvpn/giql"
	"net/http"
)

func main() {
	http.HandleFunc("/", giql.New())
	// or http.HandleFunc("/", giql.New("http://example.com/graphql"))
	http.ListenAndServe(":8080", nil)
}
```

## License

Apache License 2.0

Copyright 2021 PoohVPN
