package application

import (
	"bufio"
	"encoding/json"
	"generate-manifest/common"
	domain "generate-manifest/domain"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GenerateManifest(w http.ResponseWriter, r *http.Request) {

	common.ValidateRequestMethod(w, r, http.MethodPost)

	var params domain.Config
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	params.Namespace = params.Namespace + "-" + params.Environment
	params.Application = common.NormalizeString(params.Application)

	// Ler o conteúdo do arquivo

	ips := loadIPList()
	if len(ips) > 0 {
		params.IPs = ips
	}

	manifest := renderDocument(params)

	// Retorna o template YAML modificado
	writeReturnManifest(w, manifest)
}

func loadIPList() string {

	fileContent, err := os.Open("./config/IPList.txt")
	if err != nil {
		log.Fatal("Error parsing YAML template: ", err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(fileContent)
	fileScanner.Split(bufio.ScanLines)
	var ips string
	for fileScanner.Scan() {
		ips += fileScanner.Text()
	}
	defer fileContent.Close()

	return ips
}

func renderDocument(params domain.Config) string {

	folderVersion := "./templates/v" + strconv.FormatFloat(params.ClusterVersion, 'f', 2, 64)
	_, err := os.Stat(folderVersion)

	if os.IsNotExist(err) {
		log.Fatal("Cluster version not found", err)
	}

	document := renderYAML(folderVersion+"/namespace.yaml", params)
	document += "\n---\n\n"
	document += renderYAML(folderVersion+"/deployment.yaml", params)
	document += "\n---\n\n"
	document += renderYAML(folderVersion+"/service.yaml", params)
	document += "\n---\n\n"
	document += renderYAML(folderVersion+"/ingress.yaml", params)

	if domain.IsHPAEmpty(params.HPA) {
		document += "\n---\n\n"
		document += renderYAML(folderVersion+"/hpa.yaml", params)
	}

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
