package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{
        "quarter_of_a_dozen": 3, 
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
	return units
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := map[string]int{}
    return bill
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    if _, found_unit := units[unit]; found_unit {
    	if i, found_item := bill[item]; found_item {
            quantity := units[unit]
            bill[item] = i + quantity
            return true
        }else{
        	quantity := units[unit]
        	bill[item] = quantity
            return true
        }
	}
	return false
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    if i, found_item := bill[item]; found_item {
		if _, found_unit := units[unit]; found_unit {
			quantity := units[unit]
            if i - quantity > 0 {
            	bill[item] = i - quantity
                return true
            }else if i - quantity == 0{
            	delete(bill, item)
                return true
            }else{
            	return false
            }
        }else{
        	return false
        }
    }else{
    	return false
    }
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    if _, found_item := bill[item]; found_item {
		return bill[item], true
    }
	return 0, false
	panic("Please implement the GetItem() function")
}
