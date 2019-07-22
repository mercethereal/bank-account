/*
Functional requirements for this package are listed in the
bank_account_test.go test file

Some of the design decisions in this package are required by
the authors of this test problem. The bank_account_test file
was written entirely by exercism (exercism.io) unaltered by myself
The bank_account.go file is entirely my work.

Some implementation notes:

I'm surrounding any read or write to an account object instance
with a mutex lock/unlock.

*/
package account

import "sync"

/*
simple bank account structure protected from concurrent reads/writes
Protected reads and writes are needed to prevent:
-Negative account balances
-incorrect balance.
-operations on a closed account
*/

type Account interface {
	Deposit(amount int64) (newBalance int64, ok bool)
	Close() (payout int64, ok bool)
	Balance() (balance int64, ok bool)
}

type bankAccount struct {
	sync.RWMutex
	balance int64
	IsOpen  bool
}

/*
Open an account and make an initial deposit.
Prevent the account from opening if an attempt
is made to deposit a negative amount.
*/
func Open(initialDeposit int64) Account {
	if initialDeposit < 0 {
		return nil
	} else {
		i := new(bankAccount)
		i.balance = initialDeposit
		i.IsOpen = true
		return i
	}
}

/*
Close the account, return any balance, and prevent any future operations
*/
func (i *bankAccount) Close() (payout int64, ok bool) {
	i.Lock()
	defer i.Unlock()
	if i.IsOpen {
		x := i.balance
		i.balance = 0
		i.IsOpen = false
		return x, true
	} else {
		return 0, false
	}
}

/*
Report the account balance
*/
func (i *bankAccount) Balance() (balance int64, ok bool) {
	i.Lock()
	defer i.Unlock()
	if i.IsOpen {
		return i.balance, i.IsOpen
	} else {
		return 0, false
	}
}

/*
Make a deposit (or withdrawal). Negative balance transaction is prevented.
*/
func (i *bankAccount) Deposit(amount int64) (newBalance int64, ok bool) {
	i.Lock()
	defer i.Unlock()
	x := i.balance
	if i.balance+amount < 0 || i.IsOpen == false {
		return i.balance, false
	} else {
		x += amount
		i.balance = x
		return x, true
	}
}
