package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	Blockchain "github.com/01Joseph-Hwang10/josephcoin/blockchain"
	"github.com/01Joseph-Hwang10/josephcoin/utils"
)

const (
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*Blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", Blockchain.Blockchain().Blocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		Blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	handler.HandleFunc("/", utils.WithLog(home))
	handler.HandleFunc("/add", utils.WithLog(add))
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
