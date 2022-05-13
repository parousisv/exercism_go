package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    if balance < 0{
        return 3.213
    }else if 0 <= balance && balance <1000{
    	return 0.5
    }else if 1000 <= balance && balance <5000{
    	return 1.621
    }else if 5000 <= balance{
    	return 2.475
    }
	return 0
	panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    if balance < 0{
        return balance * 0.03213
    }else if 0 <= balance && balance <1000{
    	return balance * 0.005
    }else if 1000 <= balance && balance <5000{
    	return balance * 0.01621
    }else if 5000 <= balance{
    	return balance * 0.02475
    }
	return 0
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return balance + Interest(balance)
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    i := 0
    for{
        if balance < targetBalance{
        	balance = AnnualBalanceUpdate(balance)
            i++
        }else{
        	break
        }
    }
	return i
	panic("Please implement the YearsBeforeDesiredBalance function")
}
