package render

import (
	"app/orchestration"
	"app/orchestration/parser"
	"encoding/json"
	"net/http"
)

type pageInfo struct {
	Symphony string
	Nodes    string
}

// SymphonyDetailPage generates the json data for graph representation of the symphony
func SymphonyDetailPage(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
	page := pageInfo{}

	if ok {
		var nodes []map[string]interface{}

		for _, dag := range orchestration.Dags {
			if dag.Name == name[0] {
				nodes = parser.NodesToMap(dag)
			}
		}

		page.Nodes = generateCytoscapeJSON(nodes)
		page.Symphony = name[0]
	}

	tmpl := AppendContentPage("content/symphony-detail-page.html")
	tmpl.ExecuteTemplate(w, "layout", page)
}

func generateCytoscapeJSON(nodes []map[string]interface{}) string {
	jsonString, _ := json.Marshal(nodes)
	return string(jsonString)
}
