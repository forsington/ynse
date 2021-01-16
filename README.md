# ynse
A CLI for importing swedish bank statements to You Need A Budget (YNAB).

I created this for personal use, since using collecting and converting bank statements to QIF for YNAB was just to time consuming.

## How it works
* Create a YNAB API key, docs [here](https://api.youneedabudget.com/#authentication)
* Run `ynse budgets` to list your budgets / accounts
* Import files either individually or specify a directory. Transactions that already exist will not be imported again.
* BAM, your transactions are now in YNAB, ready to be approved. Remember, always double check that the balances match up!

## Examples 
`ynse budgets --api-key your-key-here`

`ynse import shb -f kontrotransactionlist.xls --api-key your-key-here --budget-id some-budget-id-here`

`ynse import shb -d /dir/where/you/keep/this/accounts/statements --api-key your-key-here --budget-id some-budget-id-here`

## flags

## TODO
- [x] Cobra commands
- [ ] Flag parse
    - [ ] -f --file File to import 
    - [ ] -d --dir Directory to import files from 
    - [ ] --api-key  YNAB API Key (docs [here](https://api.youneedabudget.com/#authentication) 
    - [ ] --budget-id  YNAB Budget ID (see [here](https://api.youneedabudget.com/#quick-start))
    - [ ] --account-id  YNAB Account ID, run `ynse budgets` to list your budgets / accounts 
    - [ ] --force
    - [ ] --dry-run
    - [ ] --allow-duplicates
- [ ] makefile
- [ ] golangci-lint
- [ ] unit tests
- [ ] bank + supported list
- [ ] confirmation before import
- [ ] dry run
- [ ] ignore duplicates
- [ ] support more banks
    - [ ] Nordea
    - [ ] SEB
    - [ ] Swedbank
    - [ ] Länsförsäkringar
    - [ ] Danske Bank
- [ ] file location
- [ ] public readme
