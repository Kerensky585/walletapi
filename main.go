package walletapi

// TODO this  would not be required in a library module!!!! used for testing here to init endpoints, open the DB connection using settings.
/*
//Example of instantiation of the Module end points and connection to the DB:
func main() {

	//Initialise the Web end points and if succssful, start it up
	if InitEndPoints() {

		//Create DB connection and Control, see appsetting.json
		if DbFileConnection("appsettings.json") {
			Router.Run("localhost:8080")
		}
	}
}
*/
