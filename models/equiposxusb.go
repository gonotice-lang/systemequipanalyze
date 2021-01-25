package models

// DataUSB - Items USB port information
type DataUSB struct {
	USBData []USBData `json:"SPUSBDataType"`
}

// USBData - Items array usb
type USBData struct {
	USBItems    []Items `json:"_items"`
	NameUSB     string  `json:"_name"`           // USB Name
	HostCtrl    string  `json:"host_controller"` // Host Controller Driver
	PCIDevice   string  `json:"pci_device"`      // PCI Device ID
	PCIRevision string  `json:"pci_revision"`    // PCI Revision ID
	PCIVendor   string  `json:"pci_vendor"`      // PCI Vendor ID
}

// Items - items usb
type Items struct {
	Item      []UItem `json:"_items"`
	NameU     string  `json:"_name"`        //Hub Name
	BcdDevice string  `json:"bcd_device"`   // Version
	LocalID   string  `json:"location_id"`  // Location ID
	Manufact  string  `json:"manufacturer"` // Manufacturer
	ProductID string  `json:"product_id"`   // Product ID
	VendorID  string  `json:"vendor_id"`    // Vendor ID
}

// UItem - item usb
type UItem struct {
	DeviceName string `json:"_name"`        // Controller Name
	BcdDevice  string `json:"bcd_device"`   // Version
	LocalID    string `json:"location_id"`  // Location ID
	Manufact   string `json:"manufacturer"` // Manufacturer
	ProductID  string `json:"product_id"`   // Product ID
	VendorID   string `json:"vendor_id"`    // Vendor ID
}
