package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abinashphulkonwar/go-latics/db"
	"github.com/abinashphulkonwar/go-latics/handlers"
)

func GetJson(t *testing.T, body *handlers.ViewReqBody) []byte {
	data, err := json.Marshal(body)

	if err != nil {
		t.Errorf("Error adding key value pair")

	}
	return data
}

func ReadBodyAsString(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func TestAddViews(t *testing.T) {
	app := App()
	req := httptest.NewRequest("POST", "/add/views", bytes.NewReader(GetJson(t, &handlers.ViewReqBody{
		Id: "123",
	})))
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	println(res.StatusCode)
	body, err := ReadBodyAsString(res)
	if err != nil {
		t.Fatal(err)
	}

	println(body)
	if db.Session != nil {
		db.Session.Close()
	}
}
