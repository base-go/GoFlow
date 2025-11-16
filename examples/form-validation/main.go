package main

import (
	"fmt"
	"regexp"

	"github.com/base-go/GoFlow/signals"
)

func main() {
	fmt.Println("=== Form Validation Example ===")

	// Form fields
	email := signals.New("")
	password := signals.New("")

	// Email validation
	emailValid := signals.NewComputed(func() bool {
		e := email.Get()
		if e == "" {
			return false
		}
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		return emailRegex.MatchString(e)
	})

	emailError := signals.NewComputed(func() string {
		e := email.Get()
		if e == "" {
			return "Email is required"
		}
		if !emailValid.Get() {
			return "Invalid email format"
		}
		return ""
	})

	// Password validation
	passwordValid := signals.NewComputed(func() bool {
		p := password.Get()
		return len(p) >= 8
	})

	passwordError := signals.NewComputed(func() string {
		p := password.Get()
		if p == "" {
			return "Password is required"
		}
		if len(p) < 8 {
			return "Password must be at least 8 characters"
		}
		return ""
	})

	// Form validity
	formValid := signals.NewComputed(func() bool {
		return emailValid.Get() && passwordValid.Get()
	})

	// Display validation status
	dispose := signals.NewEffect(func() {
		fmt.Println("\n--- Form Status ---")

		// Use Untracked to avoid creating dependencies on error messages
		// when we just want to check validity
		isValid := formValid.Get()

		signals.Untracked(func() struct{} {
			if err := emailError.Get(); err != "" {
				fmt.Printf("Email Error: %s\n", err)
			} else {
				fmt.Println("Email: ✓ Valid")
			}

			if err := passwordError.Get(); err != "" {
				fmt.Printf("Password Error: %s\n", err)
			} else {
				fmt.Println("Password: ✓ Valid")
			}
			return struct{}{}
		})

		if isValid {
			fmt.Println("Form: ✓ Ready to submit")
		} else {
			fmt.Println("Form: ✗ Please fix errors")
		}
	})
	defer dispose()

	// Simulate user input
	fmt.Println("\n1. Empty form:")
	// (already displayed by effect)

	fmt.Println("\n2. Invalid email:")
	email.Set("invalid-email")

	fmt.Println("\n3. Valid email, short password:")
	signals.Batch(func() {
		email.Set("user@example.com")
		password.Set("short")
	})

	fmt.Println("\n4. Valid form:")
	password.Set("securePassword123")
}
