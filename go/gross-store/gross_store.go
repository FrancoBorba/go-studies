package gross



// Units stores the Gross Store unit measurements.
func Units() map[string]int {
		unit := map[string]int{
			"quarter_of_a_dozen" :	3,
			"half_of_a_dozen":	6,
			"dozen":	12,
			"small_gross":	120,
			"gross":	144,
			"great_gross":	1728,
		}
	return  unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int {}
	return  bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	quantity, existsUnit := checkUnit(units, unit)
	if !existsUnit {
		return false
	}

	actualValue, existsItem := bill[item]

	if existsItem {
		bill[item] = actualValue + quantity
	} else {
		bill[item] = quantity
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	newQuantity, existsUnit := checkUnit(units, unit)
	if !existsUnit {
		return false
	}

	actualValue, existsItem := bill[item]
	if !existsItem {
		return  false
	}

	if(actualValue - newQuantity < 0){
		return  false
	}

	if(actualValue - newQuantity == 0){
		delete(bill,item)
		return  true
	}

	bill[item] = actualValue - newQuantity
	return  true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	actualValue, existsItem := bill[item]
	if !existsItem {
		return  0 ,false
	}

	return actualValue, true

}

func checkUnit(units map[string] int , unit string) (int, bool) {

	value , exists := units[unit]

	if(exists){
		return  value , true
	}
	return 0, false
}
