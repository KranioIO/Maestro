package render

import (
	"path/filepath"
	"text/template"
)

const webFolder = "../www/"

// AppendContentPage makes all imports to the template and adds the content template required
func AppendContentPage(pageName string) *template.Template {
	layoutPath := filepath.Join(webFolder+"templates", "dark-material-layout.html")
	sidebarPath := filepath.Join(webFolder+"templates", "sidebar-menu.html")
	navbarPath := filepath.Join(webFolder+"templates", "navbar.html")
	contentPath := filepath.Join(webFolder+"templates", pageName)

	// fp := filepath.Join("./src/web/templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(layoutPath, sidebarPath, navbarPath, contentPath)

	return tmpl
}
