package walletapi

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DbConn *gorm.DB
)

// Read stting file specified and openthe DB connection of File or Remote DB
// If tables do not exist in the DB file provided create them.
func DbController(settingsJson string) bool {

	var dbConnInitialised = false

	jsonFile, err := os.Open(settingsJson)

	if err != nil {
		//we could handle this more gracefully than this really, but quick and dirty sorry :()
		panic("Unable to read the settings file, required to conifgure connections etc....")
	}

	byteValue, _ := io.ReadAll(jsonFile)
	fmt.Println("Successfully Opened settings.json")
	var settings settings
	json.Unmarshal(byteValue, &settings)

	defer jsonFile.Close()

	//Could turn this into a switch to handle lots iff different DB conneciton types here
	if settings.RemoteDb {
		db, err := gorm.Open(sqlite.Open(settings.DbLocation), &gorm.Config{})
		DbConn = db

		if err != nil {
			panic("Critical error: unable to open or connect to the Database.")
		}

		dbConnInitialised = true

	} else {
		db, err := gorm.Open(sqlite.Open(settings.DbLocation), &gorm.Config{})
		DbConn = db

		if err != nil {
			panic("Critical error: unable to open or connect to the Database.")
		}

		dbConnInitialised = true
	}

	//create the schema if the table doesnt exist for our file DB
	if !DbConn.Migrator().HasTable("wallets") {
		DbConn.AutoMigrate(&Wallet{})
	} else {
		logrus.Println("DB wallet table and schema already initialised.")
	}

	//Table for auth mocks,
	//TODO: move to separate tesitng logic if time permis!
	if !DbConn.Migrator().HasTable("users") {
		DbConn.AutoMigrate(&user{})
	} else {
		logrus.Println("DB user table and schema already initialised.")
	}

	return dbConnInitialised
}

// Call DB and credit the wallet for {walletId}
func creditDbWallet(walletId string, amount decimal.Decimal) {

	if DbConn == nil {
		logrus.Error("Database connection not open, please check database is initialised and accessible.")
		return
	}

	checkAmount, err := ValidatePositiveAmount(amount)

	if checkAmount && err == nil {

		var currentWallet Wallet

		queryResult := DbConn.First(&currentWallet, "w_id = ?", walletId)

		if queryResult.Error != nil {
			logrus.Errorln(queryResult.Error)
		} else {
			var creditBalance = currentWallet.Balance.Add(amount)
			DbConn.Model(&currentWallet).Update("Balance", creditBalance)
			logrus.Info("Wallet credited, balance updated: ", creditBalance)

		}
	} else {
		logrus.Error("Amount was not a positive value, negative values are note allowed.")
	}
}

// Call DB and debit the wallet for {walletId}
func debitDbWallet(walletId string, amount decimal.Decimal) {

	if DbConn == nil {
		logrus.Error("Database connection not open, please check database is initialised and accessible.")
		return
	}

	checkAmount, err := ValidatePositiveAmount(amount)

	if checkAmount && err == nil {

		var currentWallet Wallet

		queryResult := DbConn.Model(&Wallet{}).Where("w_id = ?", walletId).First(&currentWallet)

		if queryResult.Error != nil {
			//do some error handling and logging
			logrus.Errorln(queryResult.Error)
		} else {

			debitOk, err := ValidateDebitBalance(amount, currentWallet.Balance)
			if debitOk && err == nil {
				var debitBalance = currentWallet.Balance.Sub(amount)
				DbConn.Model(&currentWallet).Update("Balance", debitBalance)
				logrus.Info("Wallet debited, balance updated: ", debitBalance)
			} else {
				logrus.Error("Unable to process transation as it would result in negative balance: ")
			}
		}
	} else {
		logrus.Error("Amount was not a positive value, negative values are note allowed.")
	}
}

// Return specified wallet balance - yeh need error checking here :/
func getDbWalletBalance(walletId string) decimal.Decimal {

	var currentWallet Wallet
	var balance decimal.Decimal

	queryResult := DbConn.Model(&Wallet{}).Where("w_id = ?", walletId).First(&currentWallet)

	//check if we had any error
	if queryResult.Error != nil {
		balance = decimal.New(-1, 0)
		logrus.Errorln(queryResult.Error)
	} else {
		balance = currentWallet.Balance
	}

	return balance
}
