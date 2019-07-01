package goldmachine

import (
	"fmt"
	"regexp"
)

type ExpenseAccount string

const (
	ATM       ExpenseAccount = "ATM"
	TAX       ExpenseAccount = "Tax"
	HOME      ExpenseAccount = "Home"
	ATM_FEE   ExpenseAccount = "ATM Fee"
	HOA       ExpenseAccount = "HOA"
	BIKING    ExpenseAccount = "Biking"
	THERAPY   ExpenseAccount = "Therapy"
	UTILITIES ExpenseAccount = "Utilities"
	TECH      ExpenseAccount = "Tech"
)

func (a ExpenseAccount) AccountName() string {
	return fmt.Sprintf("Expenses:%s", a)
}

type RevenueAccount string

const (
	INTEREST RevenueAccount = "Interest"
	INCOME   RevenueAccount = "Income"
)

func (a RevenueAccount) AccountName() string {
	return fmt.Sprintf("Revenue:%s", a)
}

type LiabilityAccount string

const (
	CREDIT_CARD LiabilityAccount = "Credit Card"
)

func (a LiabilityAccount) AccountName() string {
	return fmt.Sprintf("Liabilities:%s", a)
}

type AssetAccount string

const (
	CHECKING_ACCOUNT AssetAccount = "Checking Account"
	INVESTMENTS      AssetAccount = "Investments"
	SAVINGS          AssetAccount = "Savings"
	VENMO            AssetAccount = "Venmo"
	SQ_CASH          AssetAccount = "Square Cash"
)

func (a AssetAccount) AccountName() string {
	return fmt.Sprintf("Assets:%s", a)
}

type Account interface {
	AccountName() string
}

type patternToAccount struct {
	pattern *regexp.Regexp
	account Account
}

var allPatterns []patternToAccount

func init() {
	allPatterns = []patternToAccount{
		{r(`E\*TRADE ACH TRNSFR`), INVESTMENTS},
		{r(`HANDYMAN HEROES SALE`), HOME},
		{r(`INTEREST PAYMENT`), INTEREST},
		{r(`IRS USATAXPYMT`), TAX},
		{r(`FRANCHISE TAX BD`), TAX},
		{r(`FRANCHISE TAX BO PAYMENTS`), TAX},
		{r(`NON-WELLS FARGO ATM TRANSACTION FEE`), ATM_FEE},
		{r(`ONLINE TRANSFER .*TO VISA SIGNATURE CARD`), CREDIT_CARD},
		{r(`TO ZAEH A`), HOA},
		{r(`NON-WF ATM WITHDRAWAL AUTHORIZED`), ATM},
		{r(`SCOTTCAROLL`), THERAPY},
		{r(`ACERIDEPROD`), BIKING},
		{r(`PGANDE`), UTILITIES},
		{r(`WINDERSON`), HOME},
		{r(`SQC\*CASH APP WINDE`), HOME},
		{r(`SQUARE INC DIRECT DEP`), INCOME},
		{r(`VENMO CASHOUT`), VENMO},
		{r(`Credit Card AUTO PAY`), CREDIT_CARD},
		{r(`GALLEY LLC SRF`), INVESTMENTS},
		{r(`GODADDY.COM`), TECH},
	}
}

func r(s string) *regexp.Regexp {
	return regexp.MustCompile(s)
}

func expenseAccountFromMemo(in string) Account {
	return revenueAccountFromMemo(in)
}

func revenueAccountFromMemo(in string) Account {
	for _, pattern := range allPatterns {
		if pattern.pattern.MatchString(in) {
			return pattern.account
		}
	}
	return ExpenseAccount("UNKNOWN")
}
