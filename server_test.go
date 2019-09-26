package server

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var testServer Server
var testText string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
	testText = randSeq(144)
}

func TestGetServer(t *testing.T) {
	testServer = New("8077")
}

func TestEndpointBeforeLoad(t *testing.T) {
	resp, err := http.Get("http://localhost:8077/test")
	if err != nil {
		t.Error(err)
	}
	expect := 404
	if resp.StatusCode != expect {
		t.Errorf("expected response status code of %d but received %d", expect, resp.StatusCode)
	}
}

func TestLoadEndpoint(t *testing.T) {
	testServer.LoadEndpoint("test endpoint", "/test", "GET", testHandlerFunc)
	resp, err := http.Get("http://localhost:8077/test")
	if err != nil {
		t.Error(err)
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	if string(text) != testText {
		t.Errorf("expected \"success\" but received %v", string(text))
	}
}
func TestStopServer(t *testing.T) {
	testServer.Stop()
}

func testHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(testText))
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
