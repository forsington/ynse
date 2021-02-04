# ynse
A CLI written in Golang for importing swedish bank statement files to You Need A Budget (YNAB).

I created this for personal semi-automated use, since using collecting and converting bank statements to QIF for YNAB was just to time-consuming.

## Installation & Docs
* [Download the latest build](https://ynse.f0.rs) or [compile it yourself](#Compile)
* Usage documentation can be [found here](https://ynse.f0.rs)

## How it works
* Create a YNAB Personal Access Token, docs [here](https://api.youneedabudget.com/#authentication)
* Find the ID:s for the budget / account you want to use

```ynse budgets -a <your-key>```

* Import files either individually or specify a directory. Transactions that already exist will not be imported again.

```ynse import -a <your-key> -b <some-id> --account-id <some-id> --bank shb -d path/to/your/files ```

* BAM, your transactions are now in YNAB, ready to be approved. Remember, always double check that balances match up!
    
## Compile
* [Install the latest version of Golang](https://golang.org/doc/install)
* Clone this repo `git clone git@github.com:forsington/ynse.git`
* `make`, or `go build -o ./bin/ynse`

## TODO
Support for more Banks:
- [x] Handelsbanken
- [ ] Nordea
- [ ] SEB
- [ ] Swedbank
- [ ] Danske Bank
- [ ] Länsförsäkringar