package models

// DataNetwork - Network data fields
type DataNetwork struct {
	NetworkData []NetworkData `json:"SPNetworkDataType"`
}

// NetworkData - Field Network data
type NetworkData struct {
	AdapterName     string       `json:"_name"`                   // Name device
	DHCPData        DHCPData     `json:"dhcp,omitempty"`          // DHCP Server Responses
	DNSData         DNSData      `json:"DNS,omitempty"`           // DNS
	EthernetInfo    EthernetInfo `json:"Ethernet,omitempty"`      // Ethernet
	Hardware        string       `json:"hardware"`                // Hardware Name
	Interface       string       `json:"interface"`               // Interface Name:
	IPaddr          []string     `json:"ip_address,omitempty"`    // IPv4 Addresses
	IPv4Data        IPv4Data     `json:"IPv4"`                    // IPv4
	IPv6Data        IPv6Data     `json:"IPv6"`                    // IPv6
	ProxiesData     ProxiesData  `json:"Proxies"`                 // Proxies
	SubNetSrvcOrder uint8        `json:"spnetwork_service_order"` // Service Order
	Type            string       `json:"type"`
}

// DHCPData - DHCP information
type DHCPData struct {
	ServerDomain  string `json:"dhcp_domain_name_servers,omitempty"` // Domain Name Servers
	DurationLease uint8  `json:"dhcp_lease_duration,omitempty"`      // Lease Duration (seconds)
	MessageeType  string `json:"dhcp_message_type,omitempty"`        // DHCP Message Type
	Routes        string `json:"dhcp_routers,omitempty"`             // Routers
	ServerIdent   string `json:"dhcp_server_identifier,omitempty"`   // Server Identifier
	SubnetMask    string `json:"dhcp_subnet_mask,omitempty"`         // Subnet Mask
}

// DNSData - DNS information
type DNSData struct {
	ServerAddr []string `json:"ServerAddresses"` // Server Addresses
}

// EthernetInfo - Ethernet information
type EthernetInfo struct {
	MACAddr      string   `json:"MAC Address,omitempty"`  // MAC Address
	MediaOpt     []string `json:"MediaOptions,omitempty"` // Media Options
	MediaSubType string   `json:"MediaSubType,omitempty"` // Media Subtype
}

// IPv4Data - IPv4 information
type IPv4Data struct {
	AddRoute      []AddRoute `json:"AdditionalRoutes,omitempty"`           // Additional Routes
	Addresses     []string   `json:"Addresses,omitempty"`                  // Addresses
	ARPHardResolv string     `json:"ARPResolvedHardwareAddress,omitempty"` // ARP Resolved Hardware Address
	ARPIPResolv   string     `json:"RPResolvedIPAddress,omitempty"`        // ARP Resolved IP Address
	ConfMethod    string     `json:"ConfigMethod"`                         // Configuration Method
	ConfIntfName  string     `json:"ConfirmedInterfaceName,omitempty"`     // Confirmed Interface Name
	IntfName      string     `json:"InterfaceName,omitempty"`              // Interface Name
	NetSign       string     `json:"NetworkSignature,omitempty"`           // Network Signature
	Router        string     `json:"Router,omitempty"`                     // Router
	SubnetMasks   []string   `json:"SubnetMasks,omitempty"`                // Subnet Masks
}

// AddRoute - Route add information
type AddRoute struct {
	DstAddr    string `json:"DestinationAddress,omitempty"` // Destination Address
	SubnetMask string `json:"SubnetMask,omitempty"`         // Subnet Mask
}

// IPv6Data - IPv6 information
type IPv6Data struct {
	Addresses    []string `json:"Addresses,omitempty"`              // Addresses
	ConfMethod   string   `json:"ConfigMethod"`                     // Configuration Method
	ConfIntfName string   `json:"ConfirmedInterfaceName,omitempty"` // Confirmed Interface Name
	IntfName     string   `json:"InterfaceName,omitempty"`          // Interface Name
	PrefixLeght  []string `json:"PrefixLength,omitempty"`           // Prefix Length
}

// ProxiesData - Proxies data information
type ProxiesData struct {
	ExcpList []string `json:"ExceptionsList"` // Exceptions List
	FTPPass  string   `json:"FTPPassive"`     // FTP Passive Mode
}
