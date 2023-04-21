package tower

type CombinedPayload struct {
	Workspace struct {
		Name       string `json:"name"`
		FullName   string `json:"fullName"`
		Visibility string `json:"visibility"`
	} `json:"workspaces"`
	Organization struct {
		Name     string `json:"name"`
		FullName string `json:"fullName"`
	} `json:"organization"`
}
