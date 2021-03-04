package models

// DataPower - Power settings data type information
type DataPower struct {
	PowerDataType []*PowerDataType `json:"SPPowerDataType"` // Power Settings
}

// PowerDataType - Power settings information
type PowerDataType struct {
	NameSetting string   `json:"_name"`
	ACPower     *ACPower `json:"AC Power,omitempty"`
	UPSInstall  string   `json:"sppower_ups_installed"`
}

// ACPower - Power AC information
type ACPower struct {
	PowerSrc          string `json:"Current Power Source,omitempty"`
	DiskSleepTimer    uint8  `json:"Disk Sleep Timer,omitempty"`
	DisplaySleepTimer uint8  `json:"Display Sleep Timer,omitempty"`
	HibMode           uint8  `json:"Hibernate Mode,omitempty"`
	PowerButton       string `json:"Sleep On Power Button,omitempty"`
	SysSleepTimer     uint8  `json:"System Sleep Timer,omitempty"`
}
