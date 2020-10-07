package webserver

import "app/orchestration"

// MainPageData contains the data needed to show main page information
type MainPageData struct {
	Title string
	Dags  []DagPreview
}

// DagPreview ..
type DagPreview struct {
	Name string
	Cron string
}

// GenerateMainMainPageData will join the information needed to rendering the main page
func GenerateMainMainPageData() MainPageData {
	dags := make([]DagPreview, len(orchestration.Dags))

	for idx, dag := range orchestration.Dags {
		cron := orchestration.Triggers[dag.Name].Info
		dags[idx] = DagPreview{Name: dag.Name, Cron: cron}
	}

	return MainPageData{Title: "FROM OBJECT", Dags: dags}
}
