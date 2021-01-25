package models

// DataHardware - Items hardware information
type DataHardware struct {
	HardwareData []HardwareData `json:"SPHardwareDataType"`
}

// HardwareData - General information hardware item
type HardwareData struct {
	NameFields  string `json:"hardware_overview"`       // Field Name
	RomVer      string `json:"boot_rom_version"`        // System Firmware Version
	CPUType     string `json:"cpu_type"`                // Processor Name
	ProcSpeed   string `json:"current_processor_speed"` // Processor Speed
	CacheCore   string `json:"l2_cache_core"`           // L2 Cache (per Core)
	Cache       string `json:"l3_cache"`                // L3 Cache
	MachModel   string `json:"machine_model"`           // Model Identifier
	MachName    string `json:"machine_name"`            // Model Name
	NumbProc    uint8  `json:"number_processors"`       // Total Number of Cores
	Packages    uint8  `json:"packages"`                // Number of Processors
	PhysicalMem string `json:"physical_memory"`         // Memory
	CPUHtt      string `json:"platform_cpu_htt"`        // Hyper-Threading Technology
	UUIDPlatf   string `json:"platform_UUID"`           // Hardware UUID
	UDIDProv    string `json:"provisioning_UDID"`       // Provisioning UDID
	SerialNumb  string `json:"serial_number"`           // Serial Number (system)
	SMCVerSys   string `json:"SMC_version_system"`      // SMC Version (system)
}
