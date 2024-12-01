package auth

import (
	"fmt"
)

// ValidateTicket checks if the provided ticket is valid (for demo purposes).
func ValidateTicket(ticket string) bool {
	// Simulate ticket validation
	// You can replace this with actual validation logic later
	if ticket == "sampleTicket" {
		return true
	}
	return false
}