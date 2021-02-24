package models

// NetStatConn - Running network services information
type NetStatConn struct {
	Proto       string
	RecvQ       string
	SendQ       string
	LocalAddr   string
	ForeignAddr string
	State       string
}

// NetStatRoute - Table network routing information
type NetStatRoute struct {
	VerIP     string
	RouteInfo []*RouteInfo
}

// RouteInfo - routing information
type RouteInfo struct {
	Dst     string
	Gateway string
	Flags   string
	Netif   string
	Expire  string
}
