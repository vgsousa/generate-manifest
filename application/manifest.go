package application

import (
	"encoding/json"
	"generate-manifest/common"
	domain "generate-manifest/domains"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func GenerateManifest(w http.ResponseWriter, r *http.Request) {

	common.ValidateRequestMethod(w, r, http.MethodPost)

	var app domain.Project
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := domain.ValidateProject(app); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Substitui os valores no template YAML
	manifest := renderDocument(app)

	// Retorna o template YAML modificado
	writeReturnManifest(w, manifest)
}

func renderDocument(app domain.Project) string {

	// Cria um mapa com os valores a serem substituídos no template
	data := map[string]interface{}{
		"ProjectName": common.NormalizeString(app.Name),
		"Environment": app.Environment,
		"Namespace":   app.Namespace + "-" + app.Environment,
		"Port":        app.Port,
	}

	document := renderYAML("./templates/namespace_v1.yaml", data)
	document += "\n---\n\n"
	document += renderYAML("./templates/deployment_v1.yaml", data)
	document += "\n---\n\n"
	document += renderYAML("./templates/service_v1.yaml", data)
	document += "\n---\n\n"
	document += renderYAML("./templates/ingress_v1.yaml", data)

	return document
}

func renderYAML(templateFile string, data interface{}) string {

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error parsing YAML template: %v", err)
		return ""
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		log.Printf("Error rendering YAML template: %v", err)
		return ""
	}

	return result.String()
}

func writeReturnManifest(w http.ResponseWriter, manifest string) {
	w.Header().Set("Content-Type", "application/yaml")
	w.Header().Set("Content-Disposition", "attachment; filename=manifest.yaml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(manifest))
}
