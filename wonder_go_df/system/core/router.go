package core
 
import (
    "github.com/gorilla/mux"
	"net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
 
type Routes []Route
 
func NewRouter(routesVal Routes) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routesVal {
        var handler http.Handler
        handler = route.HandlerFunc 
        handler = Logger(handler, route.Name)
 
        router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
    }
    return router
}
