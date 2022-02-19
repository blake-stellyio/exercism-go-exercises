package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32
    switch {

        case balance < 0:
    	interest = 3.213

        case balance >= 0 && balance < 1000:
    	interest = .5

        case balance >= 1000 && balance < 5000:
    	interest = 1.621

        case balance >= 5000:
    	interest =2.475
    }
	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return (float64(InterestRate(balance)) / float64(100)) * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
    var amount float64 = balance
    for {
        if (amount / targetBalance) < 1 {
            amount = AnnualBalanceUpdate(amount)
            years = years + 1
        } else {
        return years
        }
    }  
	return years
}
