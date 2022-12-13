package utils

// Checks whether all the elements in a list are different
func AreStringElementsDifferent(s string) bool {

	// Create a map to store the elements that have been seen
	hashMap := make(map[string]bool)
	// Loop through the list and add each element to the map
	for _, char := range s {
		hashMap[string(char)] = true
	}
	return len(s) == len(hashMap)
	// Return whether the length of the map is equal to the length of the list
}
