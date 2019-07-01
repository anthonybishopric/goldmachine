package goldmachine

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Rhymond/go-money"
)

func ParseCashAppCSV(in string) ([]JournalEntry, error) {
	cashAppHeader := regexp.MustCompile(`Transaction ID`)

	cashAppPatterns := ps(
		p("🧼|✨|🧹", HOME),
	)

	return convertToJournalEntries(in, func(line []string) (JournalEntry, error) {
		//"Transaction ID","Name","Notes","Date","Status","Account","Amount","Fee","Net Amount","Currency"
		//"f69na6n","'Winderson'","'🧼'","2019-06-20 09:19:02 PDT","PAYMENT SENT","Visa Debit 3167","-$140","$0","-$140","USD"
		je := &JournalEntry{}
		if cashAppHeader.MatchString(line[0]) {
			return *je, Skip
		}

		stamp, err := time.Parse("2006-01-02 15:04:05 MST", line[3])
		if err != nil {
			return *je, err
		}

		je.EffectiveAt = stamp
		je.Identifier = line[1] + " " + line[2] // identifier is also the payee
		if line[5] != "Your Cash" {
			return *je, Skip
		}

		amountString := strings.Replace(strings.Replace(strings.Replace(line[8], "$", "", -1), "+", "", -1), " ", "", -1)
		if amountString == "" {
			return *je, Skip
		}

		amountFloat, err := strconv.ParseFloat(amountString, 64)
		if err != nil {
			return *je, err
		}
		amount := money.New(int64(amountFloat*100), line[9])

		if amountFloat == 0 {
			return *je, Skip
		}

		if amountFloat < 0 {
			var account Account
			account = accountFromPatterns(line[5], cashAppPatterns)
			je.AddEntry(NewCreditEntry(amount.Absolute(), CASH_APP, line[5]))
			je.AddEntry(NewDebitEntry(amount.Absolute(), account, line[5]))
		} else {
			je.AddEntry(NewCreditEntry(amount.Absolute(), accountFromPatterns(line[5], cashAppPatterns), line[5]))
			je.AddEntry(NewDebitEntry(amount.Absolute(), CASH_APP, line[5]))
		}

		return *je, nil
	})
}
