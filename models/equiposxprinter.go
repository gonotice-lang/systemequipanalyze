package models

// DataPrinter - Printers data fields
type DataPrinter struct {
	PrinterkData []PrinterData `json:"SPPrintersDataType"` // Printer Devices
}

// PrinterData - Printer data
type PrinterData struct {
	PrinterName    string        `json:"_name"`        // Name Printer
	DateCreate     string        `json:"creationDate"` // Added
	CupsFilters    []CupsFilter  `json:"cups filters,omitempty"`
	CupVer         string        `json:"cupsversion"`     // CUPS Version
	Default        string        `json:"default"`         // Default
	DriverVer      string        `json:"driverversion"`   // Driver Version
	FaxSupport     string        `json:"Fax Support"`     //Fax support
	Ppd            string        `json:"ppd"`             // PPD
	PpdFileVer     string        `json:"ppdfileversion"`  // PPD File Version
	PrinterCmd     string        `json:"printercommands"` // Printer Commands
	PrinterPdes    []PrinterPdes `json:"printerpdes,omitempty"`
	PrinterSharing string        `json:"printersharing"` // System Printer Sharing
	PrintServ      string        `json:"printserver"`    // Printer Server
	PsVersion      string        `json:"psversion"`      // PostScript Version
	Scanner        string        `json:"scanner"`        // Scanning support
	Shared         string        `json:"shared"`         // Shared
	Status         string        `json:"status"`         // Status Printer
	URI            string        `json:"uri"`            // URI
}

// CupsFilter - filters cups data
type CupsFilter struct {
	FilterName string `json:"_name,omitempty"`              // Filters Name
	FilterPath string `json:"filter path,omitempty"`        // Path
	FilterPerm string `json:"filter permissions,omitempty"` // Permissions
}

// PrinterPdes - Printer PDE info
type PrinterPdes struct {
	PDEName string `json:"_name,omitempty"`   // PDE Name
	Sandbox string `json:"sandbox,omitempty"` // Sandbox compliant
}
