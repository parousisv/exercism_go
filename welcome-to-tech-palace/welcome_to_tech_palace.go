package techpalace
import "fmt"
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    y := strings.Repeat("*", numStarsPerLine)
    return y + "\n" + welcomeMsg+ "\n" + y
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    x := strings.Replace(oldMsg, "*", "", -1)
    y := strings.Replace(x, "\n", "", -1)
    z := strings.TrimSpace(y)
    fmt.Println(z)
    return z
	panic("Please implement the CleanupMessage() function")
}
