package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	map1 := map[string]int{}
    map1["quarter_of_a_dozen"] = 3
    map1["half_of_a_dozen"] = 6
    map1["dozen"] = 12
    map1["small_gross"] = 120
	map1["gross"] = 144
	map1["great_gross"] = 1728
    return map1
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := map[string]int{}
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
    if !ok {
        return false
    }
	_, exists := bill[item]
    if exists {
        bill[item] += units[unit]
    } else {
    	bill[item] = units[unit]
    }
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitQty, unitExists := units[unit]
    itemQty, itemExists := bill[item]
    qty := itemQty - unitQty
    if !itemExists || !unitExists || qty < 0 {
        return false
    } else if qty == 0 {
    	delete(bill, item)
    	return true
    } else {
		bill[item] = qty
    	return true
    }
    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, itemExists := bill[item]
    return qty, itemExists

}
