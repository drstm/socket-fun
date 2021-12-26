package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {

	// create listener (server side)
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Starting tcp listener on 8888")
	defer l.Close()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world!\n"))
	})

	// some basic file descriptor tom foolery
	file, err := l.(*net.TCPListener).File() // net.Listener typecasted to net.TCPListener
	if err != nil {
		panic(err)
	}
	fmt.Println("Socket file descriptor: ", file.Fd())

	if err := http.Serve(l, nil); err != nil {
		panic(err)
	}

}

/* helpful resources
* https://iximiuz.com/en/posts/go-net-http-setsockopt-example/
 */
