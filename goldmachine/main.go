package main

import (
	"fmt"

	"github.com/anthonybishopric/goldmachine"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	creditCardIn *string = kingpin.Flag("credit-card-in", "CSV exported from Wells Fargo of Credit Card Transactions").String()
	checkingIn   *string = kingpin.Flag("checking-in", "CSV exported from Wells Fargo of Checking Account Transactions").String()
)

func main() {
	kingpin.Parse()
	jes := []goldmachine.JournalEntry{}
	if *checkingIn != "" {
		checkingEntries, err := goldmachine.ParseCheckingCSV(*checkingIn)
		if err != nil {
			panic(err)
		}
		jes = append(jes, checkingEntries...)
	}
	if *creditCardIn != "" {
		creditCardEntries, err := goldmachine.ParseCreditCardCSV(*creditCardIn)
		if err != nil {
			panic(err)
		}
		jes = append(jes, creditCardEntries...)
	}
	for _, je := range jes {
		fmt.Println(je.ToLedgerCLI())
	}
}
