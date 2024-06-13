package subdomains

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func SetupRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		if dataForwading {
			ForwardDataHandler(w, r)
		}
	})

	http.HandleFunc("/subdomain", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}

		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		setWebSocketConnection(websocket)
		MessageAccepterHandler(websocket)
	})
}

var dataForwading = false

func MessageAccepterHandler(conn *websocket.Conn) {
	var timer *time.Timer

	go func() {
		for {
			_, encodedMessage, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			message := string(encodedMessage)
			fmt.Print("Received the encoded message.\n")

			if message == "afjhfa793n&%$$kjhfah8H&h88h78uHY&99yhyfauhh8YUIH98Hh9``hhhre9rfhh93%4&" {
				dataForwading = true

				if timer != nil {
					timer.Stop()
				}

				timer = time.AfterFunc(1*time.Hour, func() {
					log.Println("Timer expired, disconnecting...")
					conn.Close()
				})
			}
		}
	}()
}

func ForwardDataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	go func() {
		conn, err := waitForConnection()
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(body)); err != nil {
			log.Println("Error sending webhook to whtest server", err)
			return
		}
		fmt.Print("Message received and forwaded to CLI.\n")
	}()
}
