Browser types localhost:8080/about
        ↓
TCP connection opens
        ↓
Go reads raw HTTP request:
"GET /about HTTP/1.1
Host: localhost:8080"
        ↓
ServeMux looks up "/about"
        ↓
Calls your handler function
        ↓
w.Write sends back:
"HTTP/1.1 200 OK
Content-Type: text/plain
WELCOME TO ABOUT SECTION"
        ↓
Browser renders response