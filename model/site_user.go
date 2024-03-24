package model

type SiteUser struct {
	ID       uint64 `json:"id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Ctime    int64  `json:"ctime"`
	Mtime    int64  `json:"mtime"`
	IsBan    bool   `json:"is_ban"`
}
