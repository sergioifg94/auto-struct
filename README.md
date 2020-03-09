# auto-struct

auto-struct is a simple, lightweight Go library to obtain complex typed values
from a **flat, key value collection**.

## Packages

| Package | Utility |
| ------- | ------- |
| `pkg/generic` | Generic utility to retrieve structs from a key value collection using reflection |
| `pkg/env` | Utility to retrieve structs from environment variables |
| `pkg/stringmap` | Utility to retrieve structs from `map[string]string`
 
## TODO

- [ ] Add functionality to write struct to key value collection

## Unit tests

To perform the unit tests

```sh
go test -v ./...
```