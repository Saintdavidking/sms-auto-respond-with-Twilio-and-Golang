package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the route for incoming messages
	http.HandleFunc("/sms", smsHandler)

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// smsHandler receives and responds to SMS messages from Twilio
func smsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost, http.MethodGet:
		// Read SMS data (GET will only work if params are in the URL)
		from := r.FormValue("From")
		body := r.FormValue("Body")
		fmt.Printf("Received message from %s: %s\n", from, body)

		// Send TwiML response
		response := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
	<Message>Thanks for your message David. We’ll get back to you soon!</Message>
</Response>`

		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprint(w, response)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
