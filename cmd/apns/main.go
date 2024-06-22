package main

import "fmt"

const (
	// You can also use port 2197 (instead of port 443) on either server when communicating with APNs.
	//You might use this port to allow APNs traffic through your firewall but to block other HTTPS traffic.
	apnsDev  = "api.sandbox.push.apple.com:443"
	apnsProd = "api.push.apple.com:443"
)

func main() {

	// https://developer.apple.com/documentation/usernotifications/sending-notification-requests-to-apns
	fmt.Println("Hello, APNS!")
}
