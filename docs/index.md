# ynse
A CLI written in Golang for importing swedish bank statement files (currently only Handelsbanken .xls files) to the budgeting software [You Need A Budget (YNAB)](https://www.youneedabudget.com/).

I created this for personal semi-automated use, since collecting and converting bank statements to QIF for YNAB was just to time-consuming.

## Installation
* [Download the binary for your platform](https://github.com/forsington/ynse/releases) or [compile it yourself](https://github.com/forsington/ynse#Compile)

## How it works
* Create a YNAB Personal Access Token, [see here](https://api.youneedabudget.com/#authentication)
* Download the transactions statements to be imported to YNAB from your bank.
* Find the ID:s for the budget / account you want to use:

```ynse budgets -a <your-key>```

* Import files either individually or specify a directory. Transactions that already exist will not be imported again.

```ynse import -a <your-access-token> -b <some-id> --account-id <some-id> --bank shb -d path/to/your/files ```

* BAM, your transactions are now in YNAB, ready to be approved. Remember, always double check that balances match up!

## Example workflow
As an example my monthly workflow looks like this:
* Every 5th of the month I dump out the previous months transactions from my bank. I wait this long to make sure all transactions from the previous month have cleared.
* Move the file to the location where I keep the transaction statements.
* Run `ynse import -d ...`
* Categorize the transactions that YNAB can't auto tag.
* Budget as usual.