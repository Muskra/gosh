package std

//import "fmt"

func GenerateCombinations(alphabet string, length int) <-chan string {
	c := make(chan string)

	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan string) {
		defer close(c) // Once the iteration function is finished, we close the channel

		addLetter(c, "", alphabet, length) // We start by feeding it an empty string
	}(c)

	return c // Return the channel to the calling function
}

// AddLetter adds a letter to the combination to create a new combination.
// This new combination is passed on to the channel before we call AddLetter once again
// to add yet another letter to the new combination in case length allows it
func addLetter(c chan string, combo string, alphabet string, length int) {
	// Check if we reached the length limit
	// If so, we just return without adding anything
	if length <= 0 {
		return
	}

	var newCombo string
	for _, ch := range alphabet {
		newCombo = combo + string(ch)
		c <- newCombo
		addLetter(c, newCombo, alphabet, length-1)
	}
}
//
// You need to implement this function to use the GenerateCombinations()
//
/*
func ShellTest() {

	for combination := range generateCombinations(" |&;", 6) {
		fmt.Println(combination) // This is instead of process(combination)
	}

	//fmt.Println("Done!")
}
*/
