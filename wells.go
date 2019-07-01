package goldmachine

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Rhymond/go-money"
)

// converts Wells Fargo transaction data to Ledger CLi data

func ParseCheckingCSV(in string) ([]JournalEntry, error) {
	return convertToJournalEntries(in, func(in []string) JournalEntry {
		je := &JournalEntry{
			AccountEntries: []AccountEntry{},
		}
		// "06/28/2019","0.41","*","","INTEREST PAYMENT"
		effective, err := time.Parse("01/02/2006", in[0])
		if err != nil {
			panic(err)
		}
		je.EffectiveAt = effective
		amount, err := strconv.ParseFloat(in[1], 64)
		if err != nil {
			panic(err)
		}
		moneyAmount := money.New(int64(amount*100), "USD")
		memo := in[4]
		if amount < 0 {
			je.AddEntry(NewDebitEntry(moneyAmount.Absolute(), expenseAccountFromMemo(memo), memo))
			je.AddEntry(NewCreditEntry(moneyAmount.Absolute(), CHECKING_ACCOUNT, memo))
		} else {
			je.AddEntry(NewCreditEntry(moneyAmount, revenueAccountFromMemo(memo), memo))
			je.AddEntry(NewDebitEntry(moneyAmount, CHECKING_ACCOUNT, memo))
		}
		je.Identifier = memo
		return *je
	})
}

func convertToJournalEntries(in string, convert func([]string) JournalEntry) ([]JournalEntry, error) {
	byteContent, err := ioutil.ReadFile(in)
	if err != nil {
		return nil, err
	}
	ret := []JournalEntry{}
	r := csv.NewReader(bytes.NewBuffer(byteContent))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		je := convert(record)
		if err = je.Validate(); err != nil {
			return nil, err
		}
		ret = append(ret, je)
	}
	return ret, nil
}
