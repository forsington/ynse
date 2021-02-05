# ynse
A CLI written in Golang for importing swedish bank statement files to the budgeting software [You Need A Budget (YNAB)](https://www.youneedabudget.com/).

I created this for personal semi-automated use, since collecting and converting bank statements to QIF for YNAB was just to time-consuming.

## Installation & Docs
* [Download the latest build](https://github.com/forsington/ynse/releases) or [compile it yourself](#Compile)
* Usage documentation can be [found here](https://forsington.github.io/ynse/)

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