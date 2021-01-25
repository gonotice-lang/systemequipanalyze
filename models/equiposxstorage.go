package models

// DataStorage - Storage information
type DataStorage struct {
	StorageData []StorageData `json:"SPStorageDataType"` // Storage equipments
}

// StorageData - Storage more information
type StorageData struct {
	StorageName   string        `json:"_name"`               // Name storage equip
	BsdName       string        `json:"bsd_name"`            // BSD name
	FileSys       string        `json:"file_system"`         // File System
	FreeSpace     uint64        `json:"free_space_in_bytes"` // Free
	IgnoreOwn     string        `json:"ignore_ownership"`    // Ignore Ownership
	MountPoint    string        `json:"mount_point"`         // Mount Point
	PhysicalDrive PhysicalDrive `json:"physical_drive"`      // Physical Drive
	Size          uint64        `json:"size_in_bytes"`       // Size
	UUIDVol       string        `json:"volume_uuid"`         // Volume UUID
	Write         string        `json:"writable"`            // Writable
}

// PhysicalDrive - Physical drive information
type PhysicalDrive struct {
	NameDevice   string `json:"device_name"`        // Device Name
	InternalDisk string `json:"is_internal_disk"`   // Internal
	MediaName    string `json:"media_name"`         // Media Name
	MediumType   string `json:"medium_type"`        // Medium Type
	PartMap      string `json:"partition_map_type"` // Partition Map Type
	Proto        string `json:"protocol"`           // Protocol
	SmartStatus  string `json:"smart_status"`       // S.M.A.R.T. Status
}
