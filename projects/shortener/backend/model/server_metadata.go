package model

type ServerMetadata struct {
	Version   string `json:"version"`
	DBType    string `json:"db_type"`
	DBVersion string `json:"db_version"`
	OSType    string `json:"os_type"`
	OSArch    string `json:"os_arch"`
}
