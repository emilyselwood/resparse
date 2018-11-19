# ResParse

Simple library to parse screen resolutions or image sizes.

Copes with resolutions like "800x600" and things "UHD" or "1080p"

## Example

```go
import "github.com/wselwood/resparse"

x, y, err := resparse.ParseResolution("HD")
if err != nil {
    log.Fatal(err)
}
fmt.Println("x:", x, "y:", y)
```