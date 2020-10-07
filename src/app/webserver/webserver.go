package webserver

import (
	"app/webserver/render"

	"log"
	"net/http"
)

const webFolder = "../www/"

// StartServer ..
func StartServer() {
	fs := http.FileServer(http.Dir(webFolder + "static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", showMainPage)
	http.HandleFunc("/symphony", render.SymphonyDetailPage)
	http.HandleFunc("/connections", showConnectionsPage)
	http.HandleFunc("/update_pipelines", updatePipelines)

	log.Println("Listening on :3000...")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func showMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := render.AppendContentPage("content-main-page.html")
	tmpl.ExecuteTemplate(w, "layout", GenerateMainMainPageData())
}

func showConnectionsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := render.AppendContentPage("content-connections-page.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func updatePipelines(w http.ResponseWriter, r *http.Request) {
	log.Println("update pipelines")
}
