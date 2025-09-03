package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer);
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	res := ""
    res += strings.Repeat("*", numStarsPerLine) + "\n"
    res += welcomeMsg + "\n"
    res += strings.Repeat("*", numStarsPerLine)
    return res
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    cleaned := strings.ReplaceAll(oldMsg, "*", "")
    cleaned = strings.TrimSpace(cleaned)
	return cleaned
}
