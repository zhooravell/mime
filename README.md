GoLang MIME Types Utilities
===========================
> This library provides functions to work with MIME types (also known as "media types").

[![License][license-image]][license-link]
[![Build][build-image]][build-link]

# Examples

```go
package main

import (
	"github.com/zhooravell/mime"
	"log"
)

func main() {
	ext, _ := mime.GetExtensions("application/javascript")

	log.Println(ext)
	// Output: {"js", "jsm", "mjs"}

	mt, _ := mime.GetMimeTypes("yml")

	log.Println(mt)
	// Output: {"application/x-yaml", "text/yaml", "text/x-yaml"}
}
```

# Source(s)

* [mime-db](https://github.com/jshttp/mime-db)
* [Wikipedia](https://en.wikipedia.org/wiki/MIME)
* [FreeDesktop](https://github.com/freedesktop/xdg-shared-mime-info)

[license-link]: https://github.com/zhooravell/mime/blob/master/LICENSE

[license-image]: https://img.shields.io/dub/l/vibe-d.svg

[build-image]: https://github.com/zhooravell/mime/actions/workflows/go.yml/badge.svg

[build-link]: https://github.com/zhooravell/mime/actions