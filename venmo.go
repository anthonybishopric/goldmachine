package goldmachine

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Rhymond/go-money"
)

func ParseVenmoCSV(in string) ([]JournalEntry, error) {
	venmoHeader := regexp.MustCompile(` Username`)

	venmoPatterns := ps(
		p("Standard Transfer", VENMO_RECON),
		p("Dinner", RESTAURANTS),
	)

	return convertToJournalEntries(in, func(line []string) (JournalEntry, error) {
		//  Username,ID,Datetime,Type,Status,Note,From,To,Amount (total),Amount (fee),Funding Source,Destination,Beginning Balance,Ending Balance,Statement Period Venmo Fees,Year to Date Venmo Fees,Disclaimer
		// ,2729439306803315572,2019-04-24T21:48:40,Payment,Complete,🍜🖤,Arianna Kellogg,Anthony Bishopric,+ $29.25,,,Venmo balance,,,,,
		je := &JournalEntry{}
		if venmoHeader.MatchString(line[0]) {
			return *je, Skip
		}
		if line[1] == "" {
			// first or last lines of the file
			return *je, Skip
		}
		stamp, err := time.Parse("2006-01-02T15:04:05", line[2])
		if err != nil {
			return *je, err
		}
		je.EffectiveAt = stamp
		je.Identifier = line[6] + " " + line[5] // identifier is also the payee
		if je.Identifier == "" {
			je.Identifier = line[1] + " " + line[5]
		}

		amountString := strings.Replace(strings.Replace(strings.Replace(line[8], "$", "", -1), "+", "", -1), " ", "", -1)
		if amountString == "" {
			return *je, Skip
		}

		amountFloat, err := strconv.ParseFloat(amountString, 64)
		if err != nil {
			return *je, err
		}
		amount := money.New(int64(amountFloat*100), "USD")

		if amountFloat == 0 {
			return *je, Skip
		}

		if amountFloat < 0 {
			var account Account
			if line[3] == "Standard Transfer" {
				account = VENMO_RECON
			} else {
				account = accountFromPatterns(line[5], venmoPatterns)
			}
			je.AddEntry(NewCreditEntry(amount.Absolute(), VENMO, line[5]))
			je.AddEntry(NewDebitEntry(amount.Absolute(), account, line[5]))
		} else {
			je.AddEntry(NewCreditEntry(amount.Absolute(), accountFromPatterns(line[5], venmoPatterns), line[5]))
			je.AddEntry(NewDebitEntry(amount.Absolute(), VENMO, line[5]))
		}

		return *je, nil
	})
}
