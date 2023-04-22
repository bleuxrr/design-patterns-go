package main

type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "user created"
	}
	return 404, "fail"
}
