package models

// AppsInfo - All applications information structure
type AppsInfo struct {
	Apps []*App `json:"SPApplicationsDataType"`
}

// App - Application information structure
type App struct {
	AppName    string   `json:"_name"`
	ArchKind   string   `json:"arch_kind"`
	LastMod    string   `json:"lastModified"`
	ObtainFrom string   `json:"obtained_from"`
	AppPath    string   `json:"path"`
	SignedBy   []string `json:"signed_by,omitempty"`
	Version    string   `json:"version"`
}
