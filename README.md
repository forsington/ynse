# ynse
A CLI for importing swedish bank statements to You Need A Budget (YNAB).

## How it works

## Examples 
`ynse import shb -f kontrotransactionlist.xls --api-key your-key-here --budget-id some-budget-id-here`

`ynse import shb -d /dir/where/you/keep/this/accounts/statements --api-key your-key-here --budget-id some-budget-id-here`

## flags
-f --file
-d --dir
--api-key
--budget-id
--account-id
--force
--dry-run
--allow-duplicates

## TODO
- [x] Cobra commands
- [ ] Flag parse
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
