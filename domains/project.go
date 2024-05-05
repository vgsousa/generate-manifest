package domain

import "errors"

type Project struct {
	Name        string `json:"project"`
	Team        string `json:"team"`
	Environment string `json:"environment"`
	Namespace   string `json:"namespace"`
	Port        int64  `json:"port"`
}

func ValidateProject(domain Project) error {
	// Valida se todos os campos da struct Application foram preenchidos
	if domain.Name == " " {
		return errors.New("Project field is required")
	}

	return nil
}
