Run:
```go build main.go
nm main | grep Example
```

Observe:
InlinedOrProbablySkippedExample and InlinedExample are inlined and not shown by _nm_, NotInlinedExample is _not_ inlined, because Cgo is involved.
