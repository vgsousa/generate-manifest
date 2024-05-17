package domain

type Resources struct {
	Requests struct {
		Memory string `json:"memory"`
		CPU    string `json:"cpu"`
	} `json:"requests"`
	Limits struct {
		Memory string `json:"memory"`
		CPU    string `json:"cpu"`
	} `json:"limits"`
}

type Deployment struct {
	Replicas           int       `json:"replicas"`
	Port               int       `json:"port"`
	Resources          Resources `json:"resources"`
	LivenessProbePath  string    `json:"livenessProbePath"`
	ReadinessProbePath string    `json:"readinessProbePath"`
}

type HPA struct {
	Min     int `json:"min"`
	Max     int `json:"max"`
	Metrics struct {
		CPU    *int `json:"cpu,omitempty"`
		Memory *int `json:"memory,omitempty"`
	} `json:"metrics"`
}

type Config struct {
	Application        string     `json:"application"`
	Environment        string     `json:"environment"`
	ClusterVersion     float64    `json:"clusterVersion"`
	InternalAccessOnly bool       `json:"internalAccessOnly"`
	Registry           string     `json:"registry"`
	Namespace          string     `json:"namespace"`
	Deployment         Deployment `json:"deployment"`
	HPA                HPA        `json:"hpa"`
	IPs                string
}

func IsHPAEmpty(hpa HPA) bool {
	// Verificar se MinReplicas e MaxReplicas são nil
	return (hpa.Min > 0 && hpa.Max > 0)
}
