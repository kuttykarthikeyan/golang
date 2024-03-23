package Server

import "net/http"

type Server struct{}
func New() *Server { return &Server{} }

var index string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Simple HTML Page</title>
</head>
<body>
    <h1>Hello, World!</h1>
    <p>This is a simple HTML page.</p>
</body>
</html>
`
var user = ` {
	"name":"Karthi",
	"age":	18
}`



func (s *Server) HandleIndex(w http.ResponseWriter, r http.Request) {
	w.Header().Add("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(index))
}
func (s *Server) Handleinfos(w http.ResponseWriter, r http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user))
}
