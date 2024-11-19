package utils

import (
	"os"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	
)

//  initializes the OAuth provider with necessary credentials
func SetupGoth() {
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:3000/auth/google/callback",
			"profile",
			"email",
		),
	)
}
