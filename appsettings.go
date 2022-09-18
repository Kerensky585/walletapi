package walletapi

type settings struct {
	RemoteDb   bool   `json:"RemoteDb"`
	DbLocation string `json:"Dblocation"`
	DbUser     string `json:"DbUser"`
	DbPass     string `json:"DbPass"`
	OtherStuff string `json:"OtherStuff"`
}
