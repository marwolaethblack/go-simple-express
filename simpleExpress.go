package express

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

//Routes is the type in which your types are stored
//It is a map where the key is string and the value is another map whose key is a string and value is a slice of functions
type Routes map[string]map[string][]func(w http.ResponseWriter, req *http.Request, stop func())

//MiddlewareHandler is the type your middleware functions passed to the app methods should be
type MiddlewareHandler func(w http.ResponseWriter, req *http.Request, stop func())

type appmethods interface {
	Get(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func()))
	Post(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func()))
	Put(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func()))
	Delete(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func()))
	Patch(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func()))
}

/*App is the type of the application with the relevant methods
You need to make a new variable of this type first then use the relevant methodss
*/
type App struct {
	routes Routes
}

//Get handles a GET http request to the given path
func (a *App) Get(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func())) {
	if a.routes == nil {
		a.routes = make(map[string]map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))
	}
	if a.routes[path] == nil {
		a.routes[path] = make(map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))

	}
	if a.routes[path]["GET"] == nil {
		a.routes[path]["GET"] = make([]func(w http.ResponseWriter, req *http.Request, stop func()), 5)
	}
	a.routes[path]["GET"] = handlers
}

//Post handles a POST http request to the given path
func (a *App) Post(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func())) {
	if a.routes == nil {
		a.routes = make(map[string]map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))
	}
	if a.routes[path] == nil {
		a.routes[path] = make(map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))

	}
	if a.routes[path]["POST"] == nil {
		a.routes[path]["POST"] = make([]func(w http.ResponseWriter, req *http.Request, stop func()), 5)
	}
	a.routes[path]["POST"] = handlers
}

//Put handles a PUT http request to the given path
func (a *App) Put(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func())) {
	if a.routes == nil {
		a.routes = make(map[string]map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))
	}
	if a.routes[path] == nil {
		a.routes[path] = make(map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))

	}
	if a.routes[path]["PUT"] == nil {
		a.routes[path]["PUT"] = make([]func(w http.ResponseWriter, req *http.Request, stop func()), 5)
	}
	a.routes[path]["PUT"] = handlers
}

//Delete handles a DELETE http request to the given path
func (a *App) Delete(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func())) {
	if a.routes == nil {
		a.routes = make(map[string]map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))
	}
	if a.routes[path] == nil {
		a.routes[path] = make(map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))

	}
	if a.routes[path]["DELETE"] == nil {
		a.routes[path]["DELETE"] = make([]func(w http.ResponseWriter, req *http.Request, stop func()), 5)
	}
	a.routes[path]["DELETE"] = handlers
}

//Patch handles a PATCH http request to the given path
func (a *App) Patch(path string, handlers ...func(w http.ResponseWriter, req *http.Request, stop func())) {
	if a.routes == nil {
		a.routes = make(map[string]map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))
	}
	if a.routes[path] == nil {
		a.routes[path] = make(map[string][]func(w http.ResponseWriter, req *http.Request, stop func()))

	}
	if a.routes[path]["PATCH"] == nil {
		a.routes[path]["PATCH"] = make([]func(w http.ResponseWriter, req *http.Request, stop func()), 5)
	}
	a.routes[path]["PATCH"] = handlers
}

func (a App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	handlers := a.routes[path][method]

	stopflag := false
	stop := func() {
		stopflag = true
	}

	if handlers != nil {
		for _, fn := range handlers {
			if stopflag {
				break
			}
			fn(w, req, stop)
		}
	} else {
		fmt.Fprintf(w, "Cannot %v %v", method, path)
	}

}

//Run starts the http server with the given port e.g ":8080", : is required
func (a App) Run(port string) {
	fmt.Printf("Server is running on localhost%s\n", port)
	http.ListenAndServe(port, a)
}

//GzipJSON gzips your data, converts it to json and sends it to the writer
func GzipJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	gz := gzip.NewWriter(w)
	json.NewEncoder(gz).Encode(data)
	gz.Close()
}
