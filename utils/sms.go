package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"github.com/go-resty/resty/v2"
)


func SendSMS(phone, name, message string) {
	// Retrieve API credentials from environment variables
	apiKey := os.Getenv("AFRICA_TALKING_API_KEY")
	username := os.Getenv("AFRICA_TALKING_USERNAME")
	apiURL := os.Getenv("AFRICA_TALKING_API_URL") 
	from := os.Getenv("AFRICA_TALKING_SHORTCODE")  

	// Check if credentials are present
	if apiKey == "" || username == "" || apiURL == "" || from == "" {
		fmt.Print("One or more environment variables (API key, username, API URL, or sender ID) are not set.")
	}

	// Format the message string with the customer's name
	messageFormatted := fmt.Sprintf("Dear %s, %s", name, message)

	// Encode phone number and message for URL parameters
	to := url.QueryEscape(phone)
	msg := url.QueryEscape(messageFormatted)

	// Prepare the request body in x-www-form-urlencoded format
	data := "username=" + username + "&to=" + to + "&message=" + msg + "&from=" + from

	// Create a new HTTP client
	client := resty.New()

	// Make the POST request to the Africa's Talking API
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("apiKey", apiKey).
		SetBody(data).
		Post(apiURL)

	// Check for errors
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return
	}

	// Log the response 
	log.Printf("SMS Response: %v", resp.String())

	// check if the message was successfully sent
	if resp.StatusCode() != 200 {
		log.Printf("Failed to send SMS: %v", resp.String())
	} else {
		log.Printf("SMS sent successfully to %s", phone)
	}
}
