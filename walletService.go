package walletapi

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	Router *gin.Engine

	//Move this to a central logger in next version, so can re use across packages and modules, doing this as example of staic init concept
	transactionLog          = "transactionsLog.txt"
	logOutFile, loggerError = os.OpenFile(transactionLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)

	endPointsInitialised bool = false
)

// Initialise end point routes
func InitEndPoints() (success bool) {

	logrus.SetOutput(logOutFile)
	logrus.SetLevel(logrus.DebugLevel)

	Router = gin.Default()

	v1 := Router.Group("/api/v1")
	{
		v1.GET("/wallets/:wid/balance", getWalletBalance)

		v1.POST("/wallets/:wid/credit", postCreditAccount)
		v1.POST("/wallets/:wid/debit", postDebitAccount)
	}

	//In production we would allow specifying the log file and path via settings, this is for simplicity here
	if loggerError != nil {
		logrus.Fatal("Unable to open transactions log file.")
		os.Exit(1) // or hanlde some other way if we want restart logic or someting instead...??
	} else {
		logrus.Info("Endpoints have been initialised, ready for connections.")
		endPointsInitialised = true
	}

	logOutFile.WriteString("File is working fine it's the logger\n")

	return endPointsInitialised
}

// Request balance to be credited with amount in DB for wallet ID
func postCreditAccount(c *gin.Context) {

	//log to to file, just example simple way to do it, could, haev file and terminal, switched fro verbos etc.
	logrus.Info(creditReqText)

	var updateBal updateBal

	// Call BindJSON to bind the received JSON to
	// wallet
	if err := c.BindJSON(&updateBal); err != nil {
		//loggerWrite.Error("Format of JSON data is incorrect...some more meaningfull stuff here...")
		return
	}

	var wid = updateBal.WID
	var amount = updateBal.Amount

	//debit specific account, use a get call or service layer here really!!!!
	logrus.Info("Credit request on wallet Id: " + wid)

	creditDbWallet(wid, amount)
	c.SecureJSON(http.StatusCreated, updateBal)
}

// Request balance to be debited in DB for wallet ID
func postDebitAccount(c *gin.Context) {

	logrus.Info(debitReqText)

	var updateBal updateBal

	// Call BindJSON to bind the received JSON to
	// wallet
	if err := c.BindJSON(&updateBal); err != nil {
		//loggerWrite.Error("Format of JSON data is incorrect / or invalid request error etc. here...")
		return
	}

	var wid = updateBal.WID
	var amount = updateBal.Amount

	//debit specific account, use a get call or service layer here really!!!!
	logrus.Info("Debit request on wallet Id: " + wid)

	debitDbWallet(wid, amount)
	c.SecureJSON(http.StatusCreated, updateBal)
}

// Request current balance from DB for wallet ID
func getWalletBalance(c *gin.Context) {
	wid := c.Param("wid")

	balance := getDbWalletBalance(wid)

	logrus.Info(balanceReqText + wid + ", Balance: " + balance.String())

	//TODO: mm bit hacky but can can think of a better check later
	if !balance.IsNegative() {
		c.SecureJSON(http.StatusOK, balance)
		return
	}

	c.SecureJSON(http.StatusNotFound, gin.H{"message": "wallet with " + wid + " not found!"})

}
