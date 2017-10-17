package tpatterns

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	const s, sep, want = "chicken", "ken", 4
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Expected %d but got index as: %d\n", want, got)
	}
}

/**
The following test cases will test the http server by mocking the http server
responses through an implementation of httptest.NewServer(http.HandlerFunc())
*/

func TestHttpServer(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello golang test server")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", greeting)
}

type statusHandler int

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(*h))
}

func TestIntegration(t *testing.T) {
	status := statusHandler(404)
	ts := httptest.NewServer(&status)
	defer ts.Close()

	s := NewServer("1.x", ts.URL, 1*time.Millisecond)
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)

	if b := w.Body.String(); !strings.Contains(b, "No.") {
		t.Errorf("body =  %q, wanted no.", b)
	}
	status = 200
	time.Sleep(3 * time.Millisecond)

	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "Yes.") {
		t.Errorf("body =  %q, wanted Yes.", b)
	}

}

func TestTaggedFunc(t *testing.T) {
	status := statusHandler(404)
	ts := httptest.NewServer(&status)
	defer ts.Close()
	if isTagged(ts.URL) {
		t.Error("Expected false but received true from isTageed function")
	}
	status = 200
	if !isTagged(ts.URL) {
		t.Error("Expected true but received false from isTageed function")
	}
}

// writing test to make a responsewriter and request to a server
func TestResponseRecorder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	req, err := http.NewRequest("GET", "http://abcd.com/foo", nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler(w, req)
	fmt.Printf("%d - %s", w.Code, w.Body.String())
}

// Testing the sub process
func TestCrasher(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Crasher()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCrasher")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with error %v, want exit status 1", err)
}
