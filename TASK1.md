
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
$ curl -v localhost/
2020-08-10 11:43:19.155007 -0700 PDT m=+3.639003257
$ curl localhost/kitchen
11:43AM
$ curl localhost/unix
Mon Aug 10 11:43:57 PDT 2020
```

### Yay!! all working correcly!
