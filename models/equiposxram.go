package models

// DataMem - Memory data type information
type DataMem struct {
	MemData []*MemData `json:"SPMemoryDataType"` // Memory Slots
}

// MemData - Items Bank memory and global information
type MemData struct {
	Items []*Item `json:"_items"` // Array BANK memory
	//Name     string `json:"_name"`                 // field name Global
	EccState string `json:"global_ecc_state"`      // ECC
	MemUpgr  string `json:"is_memory_upgradeable"` //  Upgradeable Memory
}

// Item - item bank memory information
type Item struct {
	MemoryName string `json:"_name"`              // BANK memory
	Manufact   string `json:"dimm_manufacturer"`  // Manufacturer
	PartNumb   string `json:"dimm_part_number"`   // Part Number
	SerialNumb string `json:"dimm_serial_number"` // Serial Number
	Size       uint64 `json:"dimm_size"`          // Size
	Speed      string `json:"dimm_speed"`         // Speed
	Status     string `json:"dimm_status"`        // Status
	Type       string `json:"dimm_type"`          // Type
}
