package bank

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"

	"github.com/forsington/ynse/budget"
)

const (
	// TokensPerTransaction is the amount of tokens per Transaction
	TokensPerTransaction = 21

	// DateToken is the Token that contains the Date
	DateToken = 1

	// PayeeToken is the Token that contains the Payee
	PayeeToken = 13

	// AmountToken is the Token that contains the Amount
	AmountToken = 19

	// AmountMultiplier is the multiplier to match a Transactions int format
	AmountMultiplier = 1000
)

// Handelsbanken is the File Parser implementations for Handelsbanken .xls files
type Handelsbanken struct {
}

// Bank returns the name of the bank
func (h *Handelsbanken) Bank() string {
	return SHB
}

// ReadFile parsers the file and extracts transactions
func (h *Handelsbanken) ReadFile(b []byte) ([]*budget.Transaction, error) {
	// Convert the latin1 to utf-8
	var body = strings.NewReader(string(b))
	r, err := charset.NewReader(body, "latin1")
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cleanBody := strings.NewReader(string(buf))
	tokenizer := html.NewTokenizer(cleanBody)

	var trans []*budget.Transaction
	i := 0
	// While the </html> tag has not been hit, iterate through the files tokens
	for tokenizer.Token().Data != "html" {
		tt := tokenizer.Next()
		i++
		if tt == html.StartTagToken {
			t := tokenizer.Token()
			if t.Data == "tr" && i > 100 {
				inner := tokenizer.Next()
				var transaction budget.Transaction

				j := 0
				// A Transaction row is 21 tokens long
				for j < TokensPerTransaction {
					if inner == html.TextToken {
						h.parseTagToTransaction(j, tokenizer.Text(), &transaction)
					}
					j++
					inner = tokenizer.Next()
				}

				if transaction.Date.IsZero() {
					fmt.Println("transaction has invalid date, skipping:", transaction.PayeeName)
				} else {
					trans = append(trans, &transaction)
				}
			}
		}
	}
	return trans, nil
}

func (h *Handelsbanken) parseTagToTransaction(tokenLocation int, b []byte, transaction *budget.Transaction) {
	text := string(b)
	text = strings.TrimSpace(text)
	if text != "" {
		switch tokenLocation {
		case DateToken:
			transaction.Date, _ = time.Parse("2006-01-02", text)
		case PayeeToken:
			transaction.PayeeName = text
		case AmountToken:
			text = strings.Replace(text, " ", "", -1)
			text = strings.Replace(text, ",", ".", -1)
			amount, _ := strconv.ParseFloat(text, 64)
			transaction.Amount = int64(amount * AmountMultiplier)
		}
	}
}