```go
import "github/jgengo/otac"
```

**add to your main:**

```go
const AppName string = ""
const AppVersion string = ""
const OTAUrl string = "http://url/check"
```

**in your main:**

```go
if err := OTACheck(); err != nil {
    fmt.Println(err)
}
```
