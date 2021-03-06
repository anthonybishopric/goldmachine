package goldmachine

import (
	"fmt"
	"regexp"
)

type ExpenseAccount string

const (
	ACLU           ExpenseAccount = "ACLU"
	ATM            ExpenseAccount = "ATM"
	ATM_FEE        ExpenseAccount = "ATM Fee"
	BARS           ExpenseAccount = "Bars"
	BIKING         ExpenseAccount = "Biking:Other"
	BIKING_FOOD    ExpenseAccount = "Biking:Food"
	BIKING_APPAREL ExpenseAccount = "Biking:Apparel"
	BIKING_SERVICE ExpenseAccount = "Biking:Service"
	BIKING_EQUIP   ExpenseAccount = "Biking:Equipment"
	BIKING_EVENTS  ExpenseAccount = "Biking:Events"
	BIRB           ExpenseAccount = "Birb"
	CAFES          ExpenseAccount = "Food:Cafes"
	DELIVERY       ExpenseAccount = "Food:Delivery"
	FATTY          ExpenseAccount = "Food:Fatty"
	GROCERIES      ExpenseAccount = "Food:Groceries"
	LUNCH          ExpenseAccount = "Food:Lunch"
	RESTAURANTS    ExpenseAccount = "Food:Restaurants"
	DONATION       ExpenseAccount = "Donation"
	ENTERTAINMENT  ExpenseAccount = "Entertainment"
	HEALTH         ExpenseAccount = "Health"
	HOA            ExpenseAccount = "HOA"
	HOME           ExpenseAccount = "Home"
	SHOPPING       ExpenseAccount = "Shopping"
	STYLE          ExpenseAccount = "Style"
	TAX            ExpenseAccount = "Tax"
	TAXIS          ExpenseAccount = "Taxis"
	TECH           ExpenseAccount = "Tech"
	THERAPY        ExpenseAccount = "Therapy"
	TRANSIT        ExpenseAccount = "Transit"
	TRAVEL         ExpenseAccount = "Travel"
	UTILITIES      ExpenseAccount = "Utilities"
	REIMBURSABLE   ExpenseAccount = "Reimbursable"
)

func (a ExpenseAccount) AccountName() string {
	return fmt.Sprintf("Expenses:%s", a)
}

type RevenueAccount string

const (
	INTEREST RevenueAccount = "Interest"
	INCOME   RevenueAccount = "Income"
) /**/

func (a RevenueAccount) AccountName() string {
	return fmt.Sprintf("Revenue:%s", a)
}

type LiabilityAccount string

const (
	CREDIT_CARD       LiabilityAccount = "Credit Card"
	CREDIT_CARD_RECON LiabilityAccount = "Credit Card Recon"
)

func (a LiabilityAccount) AccountName() string {
	return fmt.Sprintf("Liabilities:%s", a)
}

type AssetAccount string

const (
	CHECKING_ACCOUNT AssetAccount = "Checking Account"
	INVESTMENTS      AssetAccount = "Investments"
	IRA              AssetAccount = "IRA"
	SAVINGS          AssetAccount = "Savings"
	VENMO            AssetAccount = "Venmo"
	VENMO_RECON      AssetAccount = "Venmo Recon"
	CASH_APP         AssetAccount = "Cash App"
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

var checkingPatterns []patternToAccount
var creditCardPatterns []patternToAccount

func p(patternToCompile string, acc Account) patternToAccount {
	return patternToAccount{r(patternToCompile), acc}
}

func init() {
	checkingPatterns = ps(
		p(`ACERIDEPROD`, BIKING),
		p(`Credit Card AUTO PAY`, CREDIT_CARD_RECON),
		p(`E\*TRADE ACH TRNSFR`, INVESTMENTS),
		p(`FRANCHISE TAX BD`, TAX),
		p(`FRANCHISE TAX BO PAYMENTS`, TAX),
		p(`GALLEY LLC SRF`, INVESTMENTS),
		p(`GODADDY.COM`, TECH),
		p(`HANDYMAN HEROES SALE`, HOME),
		p(`INTEREST PAYMENT`, INTEREST),
		p(`IRS USATAXPYMT`, TAX),
		p(`NON-WELLS FARGO ATM TRANSACTION FEE`, ATM_FEE),
		p(`NON-WF ATM WITHDRAWAL AUTHORIZED`, ATM),
		p(`ONLINE TRANSFER .*TO VISA SIGNATURE CARD`, CREDIT_CARD_RECON),
		p(`PGANDE`, UTILITIES),
		p(`SCOTTCAROLL`, THERAPY),
		p(`SQC\*CASH APP WINDE`, HOME),
		p(`SQUARE INC DIRECT DEP`, INCOME),
		p(`TO ZAEH A`, HOA),
		p(`VENMO CASHOUT`, VENMO_RECON),
		p(`WINDERSON`, HOME),
	)

	creditCardPatterns = ps(
		p(`BURGERS AND B`, RESTAURANTS),
		p(`A MANO`, RESTAURANTS),
		p(`ACLU`, ACLU),
		p(`AIDSLIFECYCLE`, BIKING_EVENTS),
		p(`AIRBNB`, TRAVEL),
		p(`ALAMO NEW MISSION`, ENTERTAINMENT),
		p(`ALASKA AIR`, TRAVEL),
		p(`Amazon Prime`, SHOPPING),
		p(`AMZN MKTP US`, SHOPPING),
		p(`ARIZMENDI`, CAFES),
		p(`ATM WITHDRAWAL`, ATM),
		p(`ITUNES\.COM`, ENTERTAINMENT),
		p(`AUTOMATIC PAYMENT - THANK YOU`, CREDIT_CARD_RECON),
		p(`BANDCAMP`, ENTERTAINMENT),
		p(`BASQUE BOULANGERIE`, RESTAURANTS),
		p(`BATH \& BODY WORKS`, HOME),
		p(`BED BATH \& BEYOND`, HOME),
		p(`BERKELEY REP BOX OFFICE`, ENTERTAINMENT),
		p(`BI-RITE CREAMERY`, FATTY),
		p(`BI-RITE (DIVIS|MARKET)`, GROCERIES),
		p(`BIKE MONKEY`, BIKING),
		p(`BLACKBIRD`, BARS),
		p(`BOOGALOOS`, RESTAURANTS),
		p(`BOOT \& SHOE SERVICE`, RESTAURANTS),
		p(`BOUCHON BAKERY`, CAFES),
		p(`BUFFALO WHOLE FOODS`, GROCERIES),
		p(`C CASA NAPA`, RESTAURANTS),
		p(`CALA SAN FRANCISCO`, RESTAURANTS),
		p(`CALIFORNIA DONUTS`, CAFES),
		p(`CB2`, HOME),
		p(`CLIFFS VARIETY`, HOME),
		p(`CLIPPER SERVICE`, TRANSIT),
		p(`CO TRATTORIA`, RESTAURANTS),
		p(`COMCAST`, UTILITIES),
		p(`COMMON SAGE`, RESTAURANTS),
		p(`COMMUNITY THRIFT STORE`, STYLE),
		p(`COOK SHOPPE`, RESTAURANTS),
		p(`COWGIRL CREAMERY`, RESTAURANTS),
		p(`CROWNE PLAZA LAX`, TRAVEL),
		p(`GODADDY\.COM`, TECH),
		p(`DOBBS FERRY`, RESTAURANTS),
		p(`DUC LOI SUPERMARKET`, GROCERIES),
		p(`VERVE COFFEE`, BIKING_FOOD),
		p(`EASY\-BREEZY\_`, FATTY),
		p(`easyrentcars.com`, TRAVEL),
		p(`EB THE ENCHANTED FORE`, ENTERTAINMENT),
		p(`EMIL VILLA'S HICKORY`, RESTAURANTS),
		p(`EMMYS SPAGHETTI SHACK`, RESTAURANTS),
		p(`Equator Coffees`, BIKING_FOOD),
		p(`EXPEDIA`, TRAVEL),
		p(`Experian`, UTILITIES),
		p(`FAIRFAX COFFEE`, BIKING_FOOD),
		p(`FENTONS`, FATTY),
		p(`FOREIGN CINEMA`, RESTAURANTS),
		p(`FOSTERS FREEZE`, RESTAURANTS),
		p(`GITHUB`, TECH),
		p(`GOKU RAMEN IZAKAYA`, RESTAURANTS),
		p(`GOOGLE ?\*(CLOUD|Domains)`, TECH),
		p(`Amazon web services`, TECH),
		p(`HAND JOB NAILS`, STYLE),
		p(`HAWKER FARE SF`, RESTAURANTS),
		p(`HEROKU`, TECH),
		p(`HIWAY`, RESTAURANTS),
		p(`HOTELS.COM`, TRAVEL),
		p(`HQ FUELS`, TRAVEL),
		p(`IMATHLETE.COM`, BIKING_EVENTS),
		p(`INCYCLE BICYCLES`, BIKING_EVENTS),
		p(`INTUIT`, REIMBURSABLE),
		p(`QuickBooks`, REIMBURSABLE),
		p(`IPPUDO`, RESTAURANTS),
		p(`(LYFT|UBER)`, TAXIS),
		p(`JUMPBIKESHARESANFRACA`, TRANSIT),
		p(`KATE'S KITCHEN`, RESTAURANTS),
		p(`KQED Public Media`, DONATION),
		p(`LABCORP`, HEALTH),
		p(`LA BOULANGERIE`, CAFES),
		p(`LA CITY PARKING METER`, TRANSIT),
		p(`LA FOLIE`, RESTAURANTS),
		p(`LA TORTILLA`, RESTAURANTS),
		p(`LE GARAGE BISTRO`, RESTAURANTS),
		p(`Le Marais Bakery`, CAFES),
		p(`LERS ROS`, RESTAURANTS),
		p(`LIMON\s`, RESTAURANTS),
		p(`MAMA JI'S`, RESTAURANTS),
		p(`MARTINS TAVERN WASHINGTON DC`, RESTAURANTS),
		p(`MIKES BIKES`, BIKING),
		p(`MISSION MANAGEMENT GROUP`, ENTERTAINMENT),
		p(`MIXT VALENCIA`, RESTAURANTS),
		p(`MOLLIE STONES`, GROCERIES),
		p(`NAMU STONEPOT`, RESTAURANTS),
		p(`Netflix\.com`, ENTERTAINMENT),
		p(`NEW YORK TIMES DIGITAL`, ENTERTAINMENT),
		p(`NOB HILL CAFE`, RESTAURANTS),
		p(`NOE HILL MARKET`, GROCERIES),
		p(`NOPALITO`, RESTAURANTS),
		p(`NOVY RESTAURANT`, RESTAURANTS),
		p(`RESTAURANT`, RESTAURANTS),
		p(`Old West Cinnamon Rolls`, BIKING_FOOD),
		p(`ONE MED\s`, HEALTH),
		p(`ONLINE PAYMENT`, CREDIT_CARD_RECON),
		p(`ORENCHI RAMEN`, RESTAURANTS),
		p(`PADRECITO`, RESTAURANTS),
		p(`PEETS`, CAFES),
		p(`PlaystationNetwork`, ENTERTAINMENT),
		p(`CREAMSANFRA`, FATTY),
		p(`PRESS RESTAURANT`, RESTAURANTS),
		p(`REGISTRATION INSURANCE`, BIKING),
		p(`ROCK HARD SAN FRANCISCO`, ENTERTAINMENT),
		p(`SAFECO INSURANCE`, HOME),
		p(`SAN FRANCYCLO`, BIKING),
		p(`Scotch Soda San`, STYLE),
		p(`SEPHORA`, STYLE),
		p(`SMITTEN ICE CREAM`, FATTY),
		p(`SOLAGE`, TRAVEL),
		p(`HILLKILLER`, BIKING_APPAREL),
		p(`SPARROWS LODGE`, TRAVEL),
		p(`SPORTS BASEMENT`, BIKING),
		p(`Spotify`, ENTERTAINMENT),
		p(`SQ \*9`, LUNCH),
		p(`SQ \*AIDS/LIFECYCLE`, BIKING),
		p(`ANDYTOWN`, CAFES),
		p(`ARSICAULT`, FATTY),
		p(`SQ \*AVERY`, RESTAURANTS),
		p(`BARZOTTO`, RESTAURANTS),
		p(`BLUE BOTTLE`, CAFES),
		p(`BOB'S DONUTS`, FATTY),
		p(`BOVINE BAKERY`, BIKING_FOOD),
		p(`BROWN BUTTER COOKIE`, BIKING_FOOD),
		p(`CAVIAR`, DELIVERY),
		p(`COMIX`, ENTERTAINMENT),
		p(`CORRIDOR`, BARS),
		p(`CRAFTSMAN AND WOLVES`, FATTY),
		p(`DANDELION`, FATTY),
		p(`DEVIL'S TEETH`, BIKING_FOOD),
		p(`DOG EARED BOOKS`, ENTERTAINMENT),
		p(`DRIP! MOBILE ESPRESSO`, CAFES),
		p(`FLYWHEEL COFFEE`, BIKING_FOOD),
		p(`FOUR BARREL COFFEE`, CAFES),
		p(`GIAN PERRONE`, REIMBURSABLE),
		p(`THISTLE.CO`, DELIVERY),
		p(`GIDDY CANDY`, FATTY),
		p(`GINO D'AGOSTINO`, STYLE),
		p(`GRANOLA'S COFFE`, BIKING_FOOD), // bakery in HMB
		p(`HUMPHRY SLOCOMB`, FATTY),
		p(`INSTACART`, GROCERIES),
		p(`JUICEY LUCY'S`, GROCERIES),
		p(`KAGAWA\-YA`, LUNCH),
		p(`LEV SAN FRANCISCO`, LUNCH),
		p(`MAUER PARK`, CAFES),
		p(`NEIGHBORS CORNE`, BIKING_FOOD),
		p(`NOE VALLEY BAKERY`, FATTY),
		p(`OZ PIZZA`, RESTAURANTS),
		p(`P\-FITS/FITTED`, BIKING),
		p(`PENTACLE COFFEE`, CAFES),
		p(`PLAYSTATION NETWORK`, ENTERTAINMENT),
		p(`PHILZ`, CAFES),
		p(`QUALITEA`, FATTY),
		p(`R. FARMS SANTA ROSA`, GROCERIES),
		p(`TRANSFER TO STERLING BANK & TRUST`, HOME),
		p(`RAINBOW GROCERY`, GROCERIES),
		p(`REVEILLE COFFEE`, CAFES),
		p(`RITUAL COFFEE`, CAFES),
		p(`SEE'S CANDIES`, FATTY),
		p(`SEE\.SAW\.SEEN`, HEALTH),
		p(`SIDEWALK JUICE`, CAFES),
		p(`SENOR SISIG`, LUNCH),
		p(`SMART MART BERKELEY`, GROCERIES),
		p(`SNOWBIRD COFFEE`, CAFES),
		p(`SOUVLA`, RESTAURANTS),
		p(`STORIES BOOKS`, ENTERTAINMENT),
		p(`SUPPORTING ORCU CASMALIA`, DONATION),
		p(`TACORGASMICO`, RESTAURANTS),
		p(`THE CASTRO FOUN`, FATTY),
		p(`THE CHOKE COACH`, RESTAURANTS),
		p(`THE LITTLE CHIH`, RESTAURANTS),
		p(`THE MARKET CAFE`, CAFES),
		p(`\bCAFE\b`, CAFES),
		p(`THE MILL San Francisco`, CAFES),
		p(`THE SOCIAL STUD`, BARS),
		p(`TRIPLE DELIGHT`, RESTAURANTS),
		p(`WISE SONS`, LUNCH),
		p(`STEAMGAMES`, ENTERTAINMENT),
		p(`StravaCOM`, BIKING),
		p(`SUPER DUPER BURGER`, RESTAURANTS),
		p(`TARGET\b`, HOME),
		p(`TICKETFLY EVENTS`, ENTERTAINMENT),
		p(`THAI HOUSE EXPRESS`, RESTAURANTS),
		p(`THE ANIMAL COMPANY`, BIRB),
		p(`BAY AREA BIRD HOSPITAL`, BIRB),
		p(`THE CHEESE BOARD`, RESTAURANTS),
		p(`THE KITCHEN T`, RESTAURANTS),
		p(`THE MODEL BAKERY`, BIKING_FOOD),
		p(`NATURAL SISTERS CAFE`, CAFES),
		p(`THE SANDWHICH PALM`, RESTAURANTS),
		p(`CAFFE ACRI BELVEDERE`, BIKING),
		p(`DELLA FATTORIA - BAKPETALUMA`, BIKING_FOOD),
		p(`TST\* KASA`, RESTAURANTS),
		p(`STATE BIRD PROVISIONS`, RESTAURANTS),
		p(`TARTINE BAKERY`, FATTY),
		p(`TURO INC\.`, TRAVEL),
		p(`URBAN TORTILLA`, RESTAURANTS),
		p(`UVA ENOTECA`, RESTAURANTS),
		p(`VALENCIA CYCLERY`, BIKING_EQUIP),
		p(`VENICE GOURMET`, BIKING_FOOD),
		p(`VITAMINSHOPPE`, HEALTH),
		p(`VZWRLSS`, UTILITIES),
		p(`WALGREENS`, HOME),
		p(`WHOLEFDS`, GROCERIES),
		p(`Grizzly Peak Cyclists`, BIKING_EVENTS),
		p(`WWW.TOPMAN.COM`, STYLE),
		p(`YOGA TREE`, HEALTH),
		p(`YOGA WORKS`, HEALTH),
		p(`ZARA USA`, STYLE),
		p(`LEMONADE PALO ALTO`, BIKING_FOOD),
		p(`PANAMABAYCO`, BIKING_FOOD),
	)

}

func ps(ps ...patternToAccount) []patternToAccount {
	return ps
}

func r(s string) *regexp.Regexp {

	return regexp.MustCompile(s)
}

func accountFromPatterns(in string, patterns []patternToAccount) Account {
	for _, pattern := range patterns {
		if pattern.pattern.MatchString(in) {
			return pattern.account
		}
	}
	return ExpenseAccount("UNKNOWN")
}
