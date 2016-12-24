# proxy-http-status-code-only
A very simple Golang app to proxy an HTTP request and return only the HTTP status code

###### Some apps don't have unauthenticated endpoints, and you want to check their HTTP Status Code without authentication. Put this app up in your private network, then expose it publically to be able to check the status codes of those endpoints publically, without anything else getting through.

## Usage:
```shell
go run main.go -url http://localhost:5000 -serveaddress localhost:3000
```
This will listen for requests on `localhost:3000` and proxy the request through to `http://localhost:5000` including the full URI, and respond with only the HTTP status code.

## Example:
```shell
machine:~$ go run main.go -url http://www.google.com -serveaddress localhost:3000 &
machine:~$ curl http://localhost:3000/asdf
machine:~$
machine:~$ curl -I http://localhost:3000/asdf
HTTP/1.1 404 Not Found
Date: Sat, 24 Dec 2016 21:03:45 GMT
Content-Type: text/plain; charset=utf-8
machine:~$ curl http://localhost:3000/robots.txt
machine:~$
machine:~$ curl -I http://localhost:3000/robots.txt
HTTP/1.1 200 OK
Date: Sat, 24 Dec 2016 21:03:53 GMT
Content-Type: text/plain; charset=utf-8
machine:~$
```
