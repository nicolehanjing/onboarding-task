
# Task 1
### Building a HTTP Server in Go
First create a new folder and open that from VSCode 

make it a Git folder and cd into it

write Go code main.go
```Golang
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, t.String())
}

func unix(w http.ResponseWriter, req *http.Request) {
	t := time.Now().Format(time.UnixDate)
	fmt.Fprintln(w, t)
}

func kitchen(w http.ResponseWriter, req *http.Request) {
	t := time.Now().Format(time.Kitchen)
	fmt.Fprintln(w, t)
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/kitchen", kitchen)
	http.HandleFunc("/unix", unix)

	http.ListenAndServe(":8083", nil)
}

```

Then open another terminal
In the first terminal:
```
$ go build
$ ./webserver
// OR
$ go run main.go
```

In the second terminal:
```
$ curl localhost:8083/
2021-07-28 23:17:13.06243 +0000 UTC
TIMEZONE:
$ curl localhost:8083/kitchen
11:16PM
TIMEZONE:
$ curl localhost:8083/unix
Wed Jul 28 23:17:47 UTC 2021
TIMEZONE:
```

### Yay!! all working correcly!
