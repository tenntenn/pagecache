# pagecache


## usage

```
package main

import (
	"fmt"
	"net/http"
	"time"

    "github.com/tenntenn/pagecache"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now())
}

func main() {
	http.Handle("/", pagecache.CacheHandlerFunc(handler, 10*time.Second))
	http.ListenAndServe(":8080", nil)
}
```
