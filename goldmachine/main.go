package main

import (
	"fmt"
	"sort"

	"github.com/anthonybishopric/goldmachine"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	creditCardIn *string = kingpin.Flag("credit-card-in", "CSV exported from Wells Fargo of Credit Card Transactions").String()
	checkingIn   *string = kingpin.Flag("checking-in", "CSV exported from Wells Fargo of Checking Account Transactions").String()
	venmoIn      *string = kingpin.Flag("venmo-in", "CSV exported from Venmo").String()
)

type ByEffective []goldmachine.JournalEntry

func (a ByEffective) Len() int           { return len(a) }
func (a ByEffective) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEffective) Less(i, j int) bool { return a[i].EffectiveAt.Before(a[j].EffectiveAt) }

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
	if *venmoIn != "" {
		venmoEntries, err := goldmachine.ParseVenmoCSV(*venmoIn)
		if err != nil {
			panic(err)
		}
		jes = append(jes, venmoEntries...)
	}
	sort.Sort(ByEffective(jes))
	for _, je := range jes {
		fmt.Println(je.ToLedgerCLI())
	}
}
