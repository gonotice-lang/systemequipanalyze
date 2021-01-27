package models

// DataEthernet - Ethernet data fields
type DataEthernet struct {
	EthernetData []EthernetData `json:"SPEthernetDataType,omitempty"`
}

// EthernetData - Ethernet information field
type EthernetData struct {
	NameDevice     string `json:"_name"`                          // Name Controller
	BSDName        string `json:"spethernet_BSD_Name"`            // BSD Name
	BundleID       string `json:"spethernet_BUNDLE_IDentifier"`   // Kext name
	Bus            string `json:"spethernet_bus"`                 // Bus
	DeviceType     string `json:"spethernet_device_type"`         // Type
	DeviecID       string `json:"spethernet_device-id"`           // Devcie ID
	KextPath       string `json:"spethernet_kext_path"`           // Location
	RevisionID     string `json:"spethernet_revision-id"`         // Revision ID
	SubSysID       string `json:"spethernet_subsystem-id"`        // Subsystem ID
	SubSysVendorID string `json:"spethernet_subsystem-vendor-id"` // Subsystem Vendor ID
	VendorID       string `json:"spethernet_vendor-id"`           // Vendor
	Ver            string `json:"spethernet_version"`             // Version
}
