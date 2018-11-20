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

## Contibuting

I would love to have more screen size names included in this. If you need something please raise an issue so we know what to add.

If you want to have a go at adding some your self you are more than welcome. At the bottom of [resolution.go] there is a map[string]struct{x,y int} that contains the known mappings.
Send through a pull request with extras. Please try and keep it in alphabetical order. It makes it slightly easier to maintain. The names must be in upper case. Test cases are also warmly welcomed.

If you find a bug or don't understand something please raise an issue. If it is not a bug we can still probably improve the docs. 

## Thanks

Most of the screen sizes in this project were cribbed from wikipedia.