package models

// BluetoothDataInfo - bluetooth informations data
type BluetoothDataInfo struct {
	BluetoothInfo []*BluetoothInfo `json:"SPBluetoothDataType"`
}

// BluetoothInfo - bluetooth informations
type BluetoothInfo struct {
	BlueVer          string                       `json:"apple_bluetooth_version"`     // Apple Bluetooth Software Version
	DeviceTitle      []map[string]DeviceTitle     `json:"device_title"`                // Name Device
	IncPortsTitle    []map[string]SetsSerialPorts `json:"incoming_serial_ports_title"` // Incoming Serial Ports
	LocalDeviceTitle *LocalDeviceTitle            `json:"local_device_title"`          // Local Device Name
	OutPortsTitle    []map[string]SetsSerialPorts `json:"outgoing_serial_ports_title"` // Outgoing Serial Ports
	ServicesTitle    []map[string]ServicesTitle   `json:"services_title"`              // Services
}

// DeviceTitle - Device Title structure
type DeviceTitle struct {
	DeviceAddr         string `json:"device_addr"`                      // Address Device
	ClassDevice        string `json:"device_classOfDevice"`             // Class of Device
	CoreSpec           string `json:"device_core_spec"`                 // Bluetooth Core Spec
	FmVersion          string `json:"device_fw_version"`                // Firmware Version
	IsConfigured       string `json:"device_isconfigured"`              // Configured:
	IsConnect          string `json:"device_isconnected"`               // Connected
	IsPaired           string `json:"device_ispaired"`                  // Paired
	ClassOfDeviceMajor string `json:"device_majorClassOfDevice_string"` // Major Type
	Manufacturer       string `json:"device_manufacturer"`              // Manufacturer
	ClassOfDeviceMinor string `json:"device_minorClassOfDevice_string"` // Minor Type
	ProductID          string `json:"device_productID"`                 // Product ID
	DeviceServices     string `json:"device_services"`                  // Services
	SupportsEDR        string `json:"device_supportsEDR"`               // EDR Supported
	SupportsESCO       string `json:"device_supportsESCO"`              // eSCO Supported
	SupportsSSP        string `json:"device_supportsSSP"`               // SSP Supported
	VendorID           string `json:"device_vendorID"`                  // Vendor ID
}

// SetsSerialPorts - serial ports maps
type SetsSerialPorts struct {
	DeviceAuth    string `json:"device_authentication,omitempty"` // Requires Authentication
	DeviceChannel string `json:"device_channel,omitempty"`        // RFCOMM Channel
	GenAddr       string `json:"general_address,omitempty"`       // Address
}

// LocalDeviceTitle - localdevice title
type LocalDeviceTitle struct {
	GeneralAddress string `json:"general_address"`                 // Address
	AutoseekKey    string `json:"general_autoseek_keyboard"`       // Auto Seek Keyboard:
	AutoSeekPoint  string `json:"general_autoseek_pointing"`       // Auto Seek Pointing
	Chipset        string `json:"general_chipset"`                 // Chipset
	Connectable    string `json:"general_connectable"`             // Connectable
	ClassComposite string `json:"general_device_class_composite"`  // Composite Class Of Device
	ClassMajor     string `json:"general_device_class_major"`      // Device Class (Major)
	ClassMinor     string `json:"general_device_class_minor"`      // Device Class (Minor)
	Discover       string `json:"general_discoverable"`            // Discoverable
	FwVersion      string `json:"general_fw_version"`              // Firmware Version
	HardTransport  string `json:"general_hardware_transport"`      // Transport
	HciRevision    string `json:"general_hci_revision"`            // HCI Revision
	HciVersion     string `json:"general_hci_version"`             // Bluetooth Core Spec
	LmpSubversion  string `json:"general_lmp_subversion"`          // LMP Subversion
	LmpVersion     string `json:"general_lmp_version"`             // LMP Version
	GenMfg         string `json:"general_mfg"`                     // Manufacturer
	GenName        string `json:"general_name"`                    // Name
	GenPower       string `json:"general_power"`                   // Bluetooth Power
	GenProductID   string `json:"general_productID"`               // Product ID
	RemoteWake     string `json:"general_remoteWake"`              // Remote wake
	ServiceClass   string `json:"general_service_class"`           // Service Class
	Handoff        string `json:"general_supports_handoff"`        // Handoff Supported
	InstHotspot    string `json:"general_supports_instantHotspot"` // Instant Hot Spot Supported
	LowEnergy      string `json:"general_supports_lowEnergy"`      // Bluetooth Low Energy Supported
	CompleteStr    string `json:"general_type_complete_string"`    // Device Type (Complete)
	MajorStr       string `json:"general_type_major_string"`       // Device Type (Major)
	GenVendorID    string `json:"general_vendorID"`                // Vendor ID
}

// ServicesTitle - title service
type ServicesTitle struct {
	FTPRootFolder string `json:"service_FTPRootFolder,omitempty"`    // Folder other devices can browse
	OBEXFolder    string `json:"service_OBEXFolder,omitempty"`       // Folder for accepted items
	OBEXHandle    string `json:"service_OBEXHandle_Other,omitempty"` // When other items are accepted
	OBEXReceive   string `json:"service_OBEXReceive,omitempty"`      // When receiving items
	ServiceState  string `json:"service_state"`                      // State
}
