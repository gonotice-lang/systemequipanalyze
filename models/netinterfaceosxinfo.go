package models

// NetIntInfo - Information network interfaces
type NetIntInfo struct {
	NameInterface string
	Flags         string
	Mtu           string   `json:"mtu,omitempty"`
	Options       string   `json:"options,omitempty"`
	Ether         []string `json:"ether,omitempty"`
	ConfigMember  []string `json:"configmember,omitempty"`
	Inet          *Inet    `json:"inet,omitempty"`
	Inet6         []*Inet6 `json:"inet6,omitempty"`
	Nd6Options    string   `json:"nd6options,omitempty"`
	Media         string   `json:"media,omitempty"`
	Status        string   `json:"status,omitempty"`
}

// Inet - Information Inet Interface
type Inet struct {
	InetAddr  string `json:"inetaddr,omitempty"`
	Netmask   string `json:"netmask,omitemty"`
	Broadcast string `json:"broadcast,omitempty"`
}

// Inet6 - Information Inet6 Interface
type Inet6 struct {
	Inet6Addr string `json:"inet6addr,omitempty"`
	Prefixlen string `json:"prefixlen,omitempty"`
	ScopeID   string `json:"scopeid,omitempty"`
}
