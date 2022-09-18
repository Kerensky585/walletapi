package walletapi

type settings struct {
	RemoteDb   bool   `json:"remoteDb"`
	DbLocation string `json:"dblocation"`
	DbUser     string `json:"dbUser"`
	DbPass     string `json:"dbPass"`
	OtherStuff string `json:"otherStuff"`
}
