package main

import "fmt"

func main() {
	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	fmt.Println()
	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}
