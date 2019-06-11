package goldmachine

import (
	"bytes"
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
	ledgerCLITemplate, err = template.New("Template for LedgerCLI").Parse(`
{{.EntryDate}} {{.Identifier}}
{{range $accountEntry := .AccountEntries}}
	{{$accountEntry.AccountName}} {{$accountEntry.Value -}}
{{end}}

`)
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

type AccountEntry struct {
	Account      string
	DebitAmount  money.Money
	CreditAmount money.Money
	Memo         string
}

func (ae AccountEntry) Validate() error {
	if ae.CreditAmount.IsZero() && ae.DebitAmount.IsZero() {
		return fmt.Errorf("Neither money field set: %+v", ae)
	}
	if ae.CreditAmount.IsNegative() || ae.DebitAmount.IsNegative() {
		return fmt.Errorf("Cannot set a negative value on an account entry: %+v", ae)
	}
	if !ae.CreditAmount.IsZero() && !ae.DebitAmount.IsZero() {
		return fmt.Errorf("Cannot set both credit and debit amounts: %+v", ae)
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
