package main
import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	serveWeb()
}
func serveWeb() {
	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{page_alias}", serveContent)  // Dynamic URL values

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8181", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request){
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]

	w.Write([]byte("Hello World! " + page_alias))
}

