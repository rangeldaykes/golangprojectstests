package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/login/", rootHandler(loginHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ClientError is an error whose details to be shared with client.
type ClientError interface {
	Error() string

	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)

	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

// HTTPError implements ClientError interface.
type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

// Use rootHandler as a wrapper around handler.
type rootHandler func(http.ResponseWriter, *http.Request) error

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function.
	if err == nil {
		return
	}

	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.)

	clientError, ok := err.(ClientError) // check if it is a ClientError.
	if !ok {
		// if the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := clientError.ResponseBody() // Try to get response body of ClientError.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}

	status, headers := clientError.ResponseHeaders() // Get http status code and header
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	w.Write(body)
}

type loginSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginUser(username string, password string) (bool, error) {
	if username == "padrao" && password == "1234" {
		return true, nil
	} else {
		return false, nil
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return NewHTTPError(nil, http.StatusMethodNotAllowed, "Method not allowed.")
	}

	// Read request body.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Request body read error : %v", err)
	}

	// Parse body as json.
	var schema loginSchema
	if err = json.Unmarshal(body, &schema); err != nil {
		return NewHTTPError(nil, http.StatusBadRequest, "Bad request : invalid JSON,")
	}

	ok, err := loginUser(schema.Username, schema.Password)
	if err != nil {
		return fmt.Errorf("Login user DB error, %v", err)
	}

	if !ok {
		return NewHTTPError(nil, http.StatusUnauthorized, "Wrong password or username.")
	}

	w.WriteHeader(http.StatusOK) // Successfully logged in.
	return nil
}
