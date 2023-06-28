package main

import (
	"net/http"
	"os"
	"testing"
)


func testMain(m *testing.M) {
	


	os.Exit(m.Run())
}


type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}