# proxy-http-status-code-only
A very simple Golang app to reverse proxy an HTTP request but return only the HTTP status code

###### Some apps don't have unauthenticated endpoints, and you want to check their HTTP Status Code without authentication. Put this app up in your private network bypassing authentication, then expose it publically to be able to check the status codes of a GET endpoint publically, without anything else getting through.

## Usage:
```shell
go run main.go -check-url http://localhost:5000/my-app-endpoint -listen-url http://localhost:3000
```
or
```shell
docker run --rm -p 3000:3000 snarlysodboxer/proxy-http-status-code-only:latest -check-url http://localhost:5000/my-app-endpoint -listen-url http://0.0.0.0:3000/status-code
```
This will listen for GET requests on `localhost:3000/status-code`, give a 404 for any path other than `/status-code`, GET `http://localhost:5000/my-api-endpoint`, and respond with only the HTTP status code.
Be sure to set your check-url path to an endpoint on your app that doesn't actually do anything, to avoid attacks.

## Example:
```shell
machine:~$ docker run -d -p 3000:3000 snarlysodboxer/proxy-http-status-code-only:latest -check-url http://www.google.com/robots.txt -listen-url http://0.0.0.0:3000/status-code
machine:~$
machine:~$ curl http://localhost:3000/asdf
404 page not found
machine:~$
machine:~$ curl -I http://localhost:3000/asdf
HTTP/1.1 404 Not Found
Date: Sat, 24 Dec 2016 21:03:45 GMT
Content-Type: text/plain; charset=utf-8
machine:~$
machine:~$ curl http://localhost:3000/status-code
machine:~$
machine:~$ curl -I http://localhost:3000/status-code
HTTP/1.1 200 OK
Date: Sat, 24 Dec 2016 21:03:53 GMT
Content-Type: text/plain; charset=utf-8
machine:~$
```

## Other options
* `-rate-limit` milliseconds at which to limit the rate of requests, defaults to 1000
* `-burst-limit` quantity of requests to allow in bursts, defaults to 10

