package models

// DataEthernet - Ethernet data fields
type DataEthernet struct {
	EthernetData []EthernetData `json:"SPEthernetDataType"`
}

// EthernetData - Ethernet information field
type EthernetData struct{}
