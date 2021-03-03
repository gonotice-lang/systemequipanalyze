package models

// NSConn - Running network services information
type NSConn struct {
	Proto       string
	RecvQ       string
	SendQ       string
	LocalAddr   string
	ForeignAddr string
	State       string
}

// NSRoute - Table network routing information
type NSRoute struct {
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
