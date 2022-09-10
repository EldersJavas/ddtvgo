package ddtvgo

import "encoding/json"

func UnmarshalRpSystemInfo(data []byte) (RpSystemInfo, error) {
	var r RpSystemInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RpSystemInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RpSystemInfo struct {
	Code    int64  `json:"code"`
	Cmd     string `json:"cmd"`
	Massage string `json:"massage"`
	Data    Data   `json:"data"`
}

type Data struct {
	DDTVCoreVer  string       `json:"DDTVCore_Ver"`
	RoomQuantity int64        `json:"Room_Quantity"`
	ServerName   string       `json:"ServerName"`
	ServerAID    string       `json:"ServerAID"`
	OSInfo       OSInfo       `json:"os_Info"`
	DownloadInfo DownloadInfo `json:"download_Info"`
}

type DownloadInfo struct {
	Downloading        int64 `json:"Downloading"`
	CompletedDownloads int64 `json:"Completed_Downloads"`
}

type OSInfo struct {
	OSVer            string `json:"OS_Ver"`
	OSTpye           string `json:"OS_Tpye"`
	MemoryUsage      int64  `json:"Memory_Usage"`
	RuntimeVer       string `json:"Runtime_Ver"`
	UserInteractive  bool   `json:"UserInteractive"`
	AssociatedUsers  string `json:"Associated_Users"`
	CurrentDirectory string `json:"Current_Directory"`
	AppCoreVer       string `json:"AppCore_Ver"`
	WebCoreVer       string `json:"WebCore_Ver"`
}
