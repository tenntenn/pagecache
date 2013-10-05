# pagecache


## usage

```
package main

import (
	"fmt"
	"net/http"
	"pagecache"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now())
}

func main() {
	http.Handle("/", pagecache.CacheHanlder(http.HandlerFunc(handler), 10*time.Second))
	http.ListenAndServe(":8080", nil)
}
```
