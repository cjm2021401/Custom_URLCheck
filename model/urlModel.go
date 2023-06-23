package model

type UrlInfoList struct {
	UrlInfos []UrlInfo `json:"urls"`
}

type UrlInfo struct {
	Name    string   `json:"name"`
	Url     string   `json:"url"`
	Timeout int32    `json:"timeout"`
	Alert   []string `json:"alert"`
}

type Response struct {
	Name         string  `json:"name"`
	Url          string  `json:"url"`
	ResponseTime float64 `json:"reponse_name"`
	StatusCode   int32   `json:"status_code`
}
