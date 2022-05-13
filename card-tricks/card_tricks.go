package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    favorite := [] int{2, 6, 9}
    return favorite
	panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index > -1 && index < len(slice){
    	return slice[index]
    }
	return -1
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index > -1 && index < len(slice){
    	slice[index] = value
	}else{
		slice = append(slice, value)
    }
    return slice
	panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
    // s1 := value[:len(value)]
    s2 := append(value, slice...)
    // copy(slice[:0], slice[:len(slice)+len(value)])
    // slice[0:] = value
    // newSlice := []int{value}
    // newSlice = append(newSlice, slice)
    return s2
	panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index > -1 && index < len(slice){
        copy(slice[index:], slice[index+1:]) // Shift a[i+1:] left one index.
    	slice[len(slice)-1] = 0     // Erase last element (write zero value).
    	slice = slice[:len(slice)-1]     // Truncate slice.
    }
    return slice
	panic("Please implement the RemoveItem function")
}
