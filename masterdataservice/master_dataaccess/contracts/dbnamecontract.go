package contracts

type ClientInfo struct {
	UserName    string
	Password    string
	Client      string
	HostName    string
	Port        int
	IsMigration bool
}
