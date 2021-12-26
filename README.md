# socket-fun

Quick exercise in trying to see how sockets work in Go. This is not a _true_ socket implementation as I ended up having an http listener - but we can swap that part out for an infinite loop and connect it to a small client or telnet to it. 

Some interesting points:
* In python if you open a socket and send it a binary string, there is some magic that happens under the hood that understands the protocol to be used. For example if you open a socket and send a string with HTTP headers you can connect to that port and receive an HTTP response. Not sure how this is handled under the hood
* Go is way more straightforward. Sockets have listeners and listeners can be pased to a connection obj. The connection object accepts a Writer interface to which you can send any byte data (I think? I just tried strings, but don't see why we can't send other types). 
* If you want an HTTP response, don't write to the socket connection object, which makes sense (python doesn't seem to care, YOLO), get an HTTP listener and a ResponseWriter obj to write out. IMO this is cleaner because it has a strong separation of concern. Writing to a socket direcrtly should be a telnet/socket client only business. 
* `http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request)` reminded me of node.js. How does `rw` get initialized? Not sure, but I think during compile time Go keeps an `http.ResponseWriter` object ready for us?
