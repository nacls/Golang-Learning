package banking

type SavingsAccount struct {
	balance     int
	monthlyRate int
}

func (s *SavingsAccount) MonthlyInterest() int {
	return (s.monthlyRate * s.balance / 100) / 12
}

func (s *SavingsAccount) Transfer(receiver Account, amount int) string {
	return DoTransfer(s, receiver, amount)
}

func (s *SavingsAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	s.balance += amount
	return "Success"
}

func (s *SavingsAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	if amount > s.balance {
		return "Account balance is not enough"
	}

	s.balance -= amount
	return "Success"
}

func (s *SavingsAccount) CheckBalance() int {
	return s.balance
}

type CheckingAccount struct {
	balance     int
	monthlyRate int
}

func (c *CheckingAccount) MonthlyInterest() int {
	return (c.monthlyRate * c.balance / 100) / 12
}

func (c *CheckingAccount) Transfer(receiver Account, amount int) string {
	return DoTransfer(c, receiver, amount)
}

func (c *CheckingAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	c.balance += amount
	return "Success"
}

func (c *CheckingAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	if amount > c.balance {
		return "Account balance is not enough"
	}

	c.balance -= amount
	return "Success"
}

func (c *CheckingAccount) CheckBalance() int {
	return c.balance
}

type InvestmentAccount struct {
	balance     int
	monthlyRate int
}

func (i *InvestmentAccount) MonthlyInterest() int {
	return (i.monthlyRate * i.balance / 100) / 12
}

func (i *InvestmentAccount) Transfer(receiver Account, amount int) string {
	return DoTransfer(i, receiver, amount)
}

func (i *InvestmentAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	i.balance += amount
	return "Success"
}

func (i *InvestmentAccount) Withdraw(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	if amount > i.balance {
		return "Account balance is not enough"
	}

	i.balance -= amount
	return "Success"
}

func (i *InvestmentAccount) CheckBalance() int {
	return i.balance
}

func DoTransfer(sender Account, receiver Account, amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	if amount > sender.CheckBalance() {
		return "Account balance is not enough"
	}

	isAccOk := checkAccountType(receiver)
	if !isAccOk {
		return "Invalid receiver account"
	}

	sender.Withdraw(amount)
	receiver.Deposit(amount)
	return "Success"
}

func checkAccountType(account Account) bool {
	switch account.(type) {
	case *SavingsAccount, *CheckingAccount, *InvestmentAccount:
		return true
	default:
		return false
	}
}

func NewSavingsAccount() *SavingsAccount {
	return &SavingsAccount{monthlyRate: 5}
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{monthlyRate: 1}
}

func NewInvestmentAccount() *InvestmentAccount {
	return &InvestmentAccount{monthlyRate: 2}
}
