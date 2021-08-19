package server

type DbData struct {
	Room  string `json:"room"`
	Uid   string `json:"admin_uid"`
	Token string `json:"token"`
}
