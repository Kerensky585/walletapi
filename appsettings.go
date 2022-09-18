package walletapi

type settings struct {
	remoteDb   bool   `json:"remoteDb"`
	dbLocation string `json:"dblocation"`
	bbUser     string `json:"dbUser"`
	bbPass     string `json:"dbPass"`
	otherStuff string `json:"otherStuff"`
}
