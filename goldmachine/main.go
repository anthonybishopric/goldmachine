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
	if *checkingIn != "" {
		journalEntries, err := goldmachine.ParseCheckingCSV(*checkingIn)
		if err != nil {
			panic(err)
		}
		for _, je := range journalEntries {
			fmt.Println(je.ToLedgerCLI())
		}
	}
}
