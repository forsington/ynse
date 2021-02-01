package bank

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/forsington/ynse/budget"
)

// TokensPerTransaction is the amount of tokens per Transaction
const TokensPerTransaction = 21

// DateToken is the Token that contains the Date
const DateToken = 1

// PayeeToken is the Token that contains the Payee
const PayeeToken = 13

// AmountToken is the Token that contains the Amount
const AmountToken = 19

type Handelsbanken struct {
}

func (h *Handelsbanken) ReadFile(f *os.File) ([]*budget.Transaction, error) {
	fmt.Println("parsing Handelsbanken file:", f.Name())

	if !strings.HasSuffix(f.Name(), ".xls") {
		return []*budget.Transaction{}, nil
	}


	byteValue, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Convert the latin1 to utf-8
	var body = strings.NewReader(string(byteValue))
	r, err := charset.NewReader(body, "latin1")
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	cleanBody := strings.NewReader(string(buf))
	z := html.NewTokenizer(cleanBody)

	var trans []*budget.Transaction
	i := 0
	// While the </html> tag has not been hit
	for z.Token().Data != "html" {
		tt := z.Next()
		i++
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "tr" && i > 100 {
				inner := z.Next()
				var transaction budget.Transaction

				j := 0
				// A Transaction row is 21 tokens long
				for j < TokensPerTransaction {
					if inner == html.TextToken {
						text := (string)(z.Text())
						text = strings.TrimSpace(text)
						if text != "" {
							switch j {
							case DateToken:
								transaction.Date, _ = time.Parse("2006-01-02", text)
							case PayeeToken:
								transaction.PayeeName = text
							case AmountToken:
								text = strings.Replace(text, " ", "", -1)
								text = strings.Replace(text, ",", ".", -1)
								floatT, _ := strconv.ParseFloat(text, 64)
								transaction.Amount = int64(floatT * 1000)
							}
						}
					}
					j++
					inner = z.Next()
				}
				if transaction.Date.IsZero() {
					fmt.Println("transaction has invalid date, skipping:", transaction.PayeeName)
				} else {
					fmt.Println("added trans", transaction.Date, transaction.PayeeName, transaction.Amount, f.Name())
					trans = append(trans, &transaction)
				}
			}
		}
	}
	return trans, nil
}
