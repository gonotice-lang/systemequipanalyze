package models

// DiskUsageInfo - Information disk usage
type DiskUsageInfo struct {
	FileSys     string
	Size        string
	Used        string
	Avail       string
	Capacity    string
	IUsed       string
	IFree       string
	ProcentUsed string
	Mounted     string
}
