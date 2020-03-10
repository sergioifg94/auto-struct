# auto-struct

auto-struct is a simple, lightweight Go library to obtain complex typed values
from a **flat, key value collection**.

> Inspired by the [.NET Core environment config provider.](https://docs.microsoft.com/en-us/aspnet/core/fundamentals/configuration/?view=aspnetcore-3.1#environment-variables-configuration-provider)

## Example

Having the following structs:

```go
type innerTestType struct {
	Value2 int64
	Value3 bool
}

type testType struct {
	Value1     string
	InnerValue innerTestType
}
```

An instance of the `testType` struct can be retrieved with the following key-value
collection, using `.` as field separator:

| Key | Value |
| --- | ----- |
| `"testType.Value1"` | `"hello"` |
| `"testType.InnerValue.Value2"` | `"12"` |
| `"testType.InnerValue.Value3"` | `"true"` |

## Packages

| Package | Utility |
| ------- | ------- |
| `pkg/generic` | Generic utility to retrieve structs from a key value collection using reflection |
| `pkg/env` | Utility to retrieve structs from environment variables |
| `pkg/stringmap` | Utility to retrieve structs from `map[string]string`
 
## Unit tests

To execute the unit tests

```sh
go test -v ./...
```

## TODO

- [x] Add support for slices
- [ ] Add support for `map[string]xxxx`
- [ ] Add functionality to write struct to key value collection