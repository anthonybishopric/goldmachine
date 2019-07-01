package goldmachine

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Rhymond/go-money"
)

// converts Wells Fargo transaction data to Ledger CLi data

func ParseCheckingCSV(in string) ([]JournalEntry, error) {
	return convertToJournalEntries(in, func(in []string) (JournalEntry, error) {
		je := &JournalEntry{
			AccountEntries: []AccountEntry{},
		}
		// "06/28/2019","0.41","*","","INTEREST PAYMENT"
		effective, err := time.Parse("01/02/2006", in[0])
		if err != nil {
			return *je, err
		}
		je.EffectiveAt = effective
		amount, err := strconv.ParseFloat(in[1], 64)
		if err != nil {
			return *je, err
		}
		moneyAmount := money.New(int64(amount*100), "USD")
		memo := in[4]
		if amount < 0 {
			je.AddEntry(NewDebitEntry(moneyAmount.Absolute(), accountFromPatterns(memo, checkingPatterns), memo))
			je.AddEntry(NewCreditEntry(moneyAmount.Absolute(), CHECKING_ACCOUNT, memo))
		} else {
			je.AddEntry(NewCreditEntry(moneyAmount, accountFromPatterns(memo, checkingPatterns), memo))
			je.AddEntry(NewDebitEntry(moneyAmount, CHECKING_ACCOUNT, memo))
		}
		je.Identifier = memo
		return *je, nil
	})
}

func ParseCreditCardCSV(in string) ([]JournalEntry, error) {
	return convertToJournalEntries(in, func(in []string) (JournalEntry, error) {
		// "01/09/2019","-28.13","*","","4505 BURGERS AND B SAN FRANCISCOCA"
		// actually identical to checking accounts but we keep them split in case there's future diversions
		je := &JournalEntry{
			AccountEntries: []AccountEntry{},
		}
		effective, err := time.Parse("01/02/2006", in[0])
		if err != nil {
			return *je, err
		}
		je.EffectiveAt = effective
		amount, err := strconv.ParseFloat(in[1], 64)
		if err != nil {
			return *je, err
		}
		moneyAmount := money.New(int64(amount*100), "USD")
		memo := in[4]
		if amount < 0 {
			je.AddEntry(NewDebitEntry(moneyAmount.Absolute(), accountFromPatterns(memo, creditCardPatterns), memo))
			je.AddEntry(NewCreditEntry(moneyAmount.Absolute(), CREDIT_CARD, memo))
		} else {
			je.AddEntry(NewCreditEntry(moneyAmount, accountFromPatterns(memo, creditCardPatterns), memo))
			je.AddEntry(NewDebitEntry(moneyAmount, CREDIT_CARD, memo))
		}
		je.Identifier = memo
		return *je, nil
	})
}

var Skip error = errors.New("Skip this")

func convertToJournalEntries(in string, convert func([]string) (JournalEntry, error)) ([]JournalEntry, error) {
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
		je, err := convert(record)
		if err != nil && err != Skip {
			return nil, err
		} else if err == Skip {
			continue
		}
		if err = je.Validate(); err != nil {
			return nil, err
		}
		ret = append(ret, je)
	}
	return ret, nil
}
