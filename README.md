# ass

`ass` is a tiny testing lib making error assertions easier. It basically consists of a single function:

```go
ass.Err(t, ...)
```

As you can see, it spells `assErrt`, which is convenient and conveys the intent very clearly.

`ass` can be used with many different inputs, making it very practical when doing table driven testing:

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
