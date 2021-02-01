# ynse
A CLI for importing swedish bank statements to You Need A Budget (YNAB).

I created this for personal semi-automated use, since using collecting and converting bank statements to QIF for YNAB was just to time consuming.

## How it works
* Create a YNAB API key, docs [here](https://api.youneedabudget.com/#authentication)
* Find the ID:s for the budget / account you want to use

```ynse budgets -a <your-key>```

* Import files either individually or specify a directory. Transactions that already exist will not be imported again.

```ynse import -a <your-key> -b <some-id> --account-id <some-id> --bank shb -d path/to/your/files ```

* BAM, your transactions are now in YNAB, ready to be approved. Remember, always double check that the balances match up!

## TODO
- [x] Cobra commands
- [ ] Flag parse
    - [x] -f --file File to import 
    - [x] -d --dir Directory to import files from 
    - [x] -b --bank (shb, seb, nordea, swedbank)
    - [x] --api-key  YNAB API Key (docs [here](https://api.youneedabudget.com/#authentication) 
    - [x] --budget-id  YNAB Budget ID (see [here](https://api.youneedabudget.com/#quick-start))
    - [x] --account-id  YNAB Account ID, run `ynse budgets` to list your budgets / accounts 
    - [ ] --dry-run
    - [x] --allow-duplicates
- [x] makefile
- [x] golangci-lint
- [x] `budgets`
- [ ] unit tests
- [x] ignore duplicates
- [ ] banks drivers
    - [x] Handelsbanken
    - [ ] Nordea
    - [ ] SEB
    - [ ] Swedbank
    - [ ] Danske Bank
    - [ ] Länsförsäkringar
- [ ] public readme
