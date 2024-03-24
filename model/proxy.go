package model

type GetProxyResp struct {
	Anonymous  string `json:"anonymous"`
	CheckCount int    `json:"check_count"`
	FailCount  int    `json:"fail_count"`
	HTTPS      bool   `json:"https"`
	LastStatus bool   `json:"last_status"`
	LastTime   string `json:"last_time"`
	Proxy      string `json:"proxy"`
	Region     string `json:"region"`
	Source     string `json:"source"`
}

type GetProxyCountResp struct {
	Count    int `json:"count"`
	HTTPType struct {
		HTTP  int `json:"http"`
		HTTPS int `json:"https"`
	} `json:"http_type"`
	Source struct { //todo
		FreeProxy03 int `json:"freeProxy03"`
		FreeProxy07 int `json:"freeProxy07"`
		FreeProxy11 int `json:"freeProxy11"`
	} `json:"source"`
}
