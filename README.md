# ynse
A CLI for importing swedish bank statements to You Need A Budget (YNAB).

I created this for personal semi-automated use, since using collecting and converting bank statements to QIF for YNAB was just to time consuming.

## How it works
* Create a YNAB API key, docs [here](https://api.youneedabudget.com/#authentication)
* Find the ID:s for the budget / account you want to use

```ynse budgets --api-key <your-key>```

* Import files either individually or specify a directory. Transactions that already exist will not be imported again.

```ynse import shb -f kontrotransactionlist.xls --api-key <your-key> --budget-id <some-id> --account-id <some-id>```

* BAM, your transactions are now in YNAB, ready to be approved. Remember, always double check that the balances match up!

## TODO
- [x] Cobra commands
- [ ] Flag parse
    - [ ] -f --file File to import 
    - [ ] -d --dir Directory to import files from 
    - [ ] -b --bank (shb, seb, nordea, swedbank)
    - [x] --api-key  YNAB API Key (docs [here](https://api.youneedabudget.com/#authentication) 
    - [ ] --budget-id  YNAB Budget ID (see [here](https://api.youneedabudget.com/#quick-start))
    - [ ] --account-id  YNAB Account ID, run `ynse budgets` to list your budgets / accounts 
    - [ ] --force
    - [ ] --dry-run
    - [ ] --allow-duplicates
    - [ ] --cutoff
- [x] makefile
- [x] golangci-lint
- [x] `budgets`
- [ ] unit tests
- [ ] confirmation before import
- [ ] ignore duplicates
- [ ] support more banks
    - [ ] Nordea
    - [ ] SEB
    - [ ] Swedbank
    - [ ] Länsförsäkringar
    - [ ] Danske Bank
- [ ] public readme
