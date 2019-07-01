package goldmachine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/Rhymond/go-money"
)

type JournalEntry struct {
	AccountEntries []AccountEntry
	EffectiveAt    time.Time
	Identifier     string
}

var ledgerCLITemplate *template.Template

func init() {
	var err error
	ledgerCLITemplate, err = template.New("Template for LedgerCLI").Parse(`{{.EntryDate}} {{.Identifier}}
{{- range $accountEntry := .AccountEntries}}
	{{$accountEntry.Account}}			{{$accountEntry.ValueString -}}
{{end}}`)
	if err != nil {
		panic("invalid template " + err.Error())
	}
}

func (je JournalEntry) ToLedgerCLI() string {
	var b []byte
	buffer := bytes.NewBuffer(b)
	err := ledgerCLITemplate.Execute(buffer, je)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (je JournalEntry) Validate() error {
	if je.Identifier == "" {
		return fmt.Errorf("No identifier on journal entry")
	}
	if len(je.AccountEntries) == 0 {
		return fmt.Errorf("No account entries present on JE")
	}
	if len(je.AccountEntries) == 1 {
		return fmt.Errorf("Only one account entry present on JE")
	}
	var checksum int64
	currency := ""
	for _, ae := range je.AccountEntries {
		if err := ae.Validate(); err != nil {
			return err
		}
		if currency == "" {
			currency = ae.CurrencyCode()
		} else if currency != ae.CurrencyCode() {
			return fmt.Errorf("Mismatched currencies found %s and %s", currency, ae.CurrencyCode())
		}
		checksum += ae.DebitAmount.Amount() - ae.CreditAmount.Amount()
	}
	if checksum != 0 {
		return fmt.Errorf("Account entries did not sum to 0. %+v", je)
	}
	return nil
}

func (je *JournalEntry) AddEntry(ae *AccountEntry) {
	je.AccountEntries = append(je.AccountEntries, *ae)
}

func (je JournalEntry) EntryDate() string {
	return je.EffectiveAt.Format("2006/01/02")
}

type AccountEntry struct {
	Account      string
	DebitAmount  *money.Money
	CreditAmount *money.Money
	Memo         string
}

func NewDebitEntry(amount *money.Money, account Account, memo string) *AccountEntry {
	return &AccountEntry{
		DebitAmount:  amount,
		CreditAmount: money.New(0, "USD"),
		Account:      account.AccountName(),
		Memo:         memo,
	}
}

func NewCreditEntry(amount *money.Money, account Account, memo string) *AccountEntry {
	return &AccountEntry{
		DebitAmount:  money.New(0, "USD"),
		CreditAmount: amount,
		Account:      account.AccountName(),
		Memo:         memo,
	}
}

func (ae AccountEntry) ValueString() string {
	if ae.DebitAmount.IsZero() {
		return fmt.Sprintf("$-%.2f", float64(ae.CreditAmount.Amount())/100)
	} else {
		return fmt.Sprintf("$%.2f", float64(ae.DebitAmount.Amount())/100)
	}
}

func (ae AccountEntry) String() string {
	out, _ := json.MarshalIndent(ae, "\t", "\t")
	return string(out)
}

func (ae *AccountEntry) Validate() error {
	if ae.CreditAmount.IsZero() && ae.DebitAmount.IsZero() {
		return fmt.Errorf("Neither money field set: %+v", ae)
	}
	if !ae.CreditAmount.IsZero() && !ae.DebitAmount.IsZero() {
		return fmt.Errorf("Cannot set both credit and debit amounts: %s", ae)
	}
	if ae.CreditAmount.IsNegative() || ae.DebitAmount.IsNegative() {
		return fmt.Errorf("Cannot set a negative value on an account entry: %s", ae)
	}
	return nil
}

func (ae AccountEntry) CurrencyCode() string {
	if ae.DebitAmount.IsZero() {
		return ae.CreditAmount.Currency().Code
	} else {
		return ae.DebitAmount.Currency().Code
	}
}
