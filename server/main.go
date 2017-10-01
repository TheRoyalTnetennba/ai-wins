package main

import (
	"log"
	"net/http"
	"time"
	"encoding/gob"

	"golang.org/x/net/context"
	"cloud.google.com/go/datastore"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func init() {
	Store.Options.MaxAge = 86400
	Store.Options.Domain = "localhost"
	gob.Register(time.Time{})
}

func NewClient() *datastore.Client {
	client, err := datastore.NewClient(Ctx, ProjectID)
	if err != nil {
		log.Fatal("Failed to set up DB client", err)
	}
	return client
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
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

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

var (
	Ctx context.Context = context.Background()
	Client *datastore.Client = NewClient()
	Store = sessions.NewCookieStore([]byte(SessionKey))
)

func main() {
	origins := handlers.AllowedOrigins(AllowedOrigins)
	credentials := handlers.AllowCredentials()
	router := NewRouter()
	router.Host(BaseURL)
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(origins, credentials)(router)))
}
