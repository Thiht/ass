# ass [![Go Reference](https://pkg.go.dev/badge/github.com/Thiht/ass.svg)](https://pkg.go.dev/github.com/Thiht/ass)

`ass` is a non-inflated testing lib making error assertions easier. It basically does one thing:

```go
ass.Err(t, ...)
```

As you can see, it spells `assErrt`, which is convenient and conveys the intent very clearly. Writing `ass` is also 50% more efficient than writing `assert`.

`ass` accepts many different inputs (one at a time though, a `bbl` companion lib might exist in the future if some expansion is needed), making it very practical when doing table driven testing:

```go
err := f()

// Assert no error
ass.Err(t, err, nil)
ass.Err(t, err, "")

// Substring matching
ass.Err(t, err, "an error occurred")

// errors.Is matching
ass.Err(t, err, ErrValue)

// errors.As matching
ass.Err(t, err, reflect.TypeOf(&ErrType{}))
```

Please star this repo if you like `ass`.
