```
curl -v localhost:8080/curl
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /curl HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< content-type: text/plain; charset=utf-8
< Cache-Control: no-cache
< Expires: Fri, 01 Jan 1990 00:00:00 GMT
< Content-Length: 920
< Server: Development/2.0
< Date: Fri, 11 Mar 2016 01:40:52 GMT
< 
* Connection #0 to host localhost left intact
http.Header{"X-Appengine-User-Email":[]string{""}, "User-Agent":[]string{"AppEngine-Google; (+http://code.google.com/appengine; appid: dev~test)"}, "Accept-Encoding":[]string{"identity"}, "X-Appengine-Server-Name":[]string{"localhost"}, "X-Appengine-Remote-Addr":[]string{"::1"}, "X-Appengine-Default-Version-Hostname":[]string{"localhost:8080"}, "X-Appengine-Server-Software":[]string{"Development/2.0"}, "X-Appengine-Request-Id-Hash":[]string{"xxx"}, "X-Appengine-Country":[]string{"ZZ"}, "X-Appengine-User-Is-Admin":[]string{"0"}, "X-Appengine-Request-Log-Id":[]string{"xxx"}, "X-Appengine-Server-Protocol":[]string{"HTTP/1.1"}, "X-Appengine-Server-Port":[]string{"8080"}, "X-Appengine-User-Organization":[]string{""}, "X-Appengine-User-Nickname":[]string{""}, "X-Test":[]string{"Hello3"}, "X-Appengine-User-Id":[]string{""}}% 
```
