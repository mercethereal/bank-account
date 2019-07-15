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
In some circumstances, I'm required to return an in instance
variable, which cannot be inside the lock/unlock. In that
case, I'm assigning temporary variables (specifically for the return)
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
type Account struct {
	sync.RWMutex
	balance int64
	IsOpen  bool
}

/*
Open an account and make an initial deposit.
Prevent the account from opening if an attempt
is made to deposit a negative amount.
*/
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	} else {
		i := new(Account)
		i.Lock()
		i.balance = initialDeposit
		i.IsOpen = true
		i.Unlock()
		return i
	}
}

/*
Close the account, return any balance, and prevent any future operations
*/
func (i *Account) Close() (payout int64, ok bool) {
	i.Lock()
	if i.IsOpen {
		x := i.balance
		i.balance = 0
		i.IsOpen = false
		i.Unlock()
		return x, true
	} else {
		i.Unlock()
		return 0, false
	}
}

/*
Report the account balance
*/
func (i *Account) Balance() (balance int64, ok bool) {
	i.Lock()
	if i.IsOpen {
		x := i.balance
		y := i.IsOpen
		i.Unlock()
		return x, y
	} else {
		i.Unlock()
		return 0, false
	}
}

/*
Make a deposit (or withdrawal). Negative balance transaction is prevented.
*/
func (i *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	i.Lock()
	x := i.balance
	if i.balance+amount < 0 || i.IsOpen == false {
		i.Unlock()
		return x, false
	} else {
		x += amount
		i.balance = x
		i.Unlock()
		return x, true
	}
}
