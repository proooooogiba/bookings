package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)


func TestFormValid(t * testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}

}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")
	
	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form doesn't have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	if form.Has("whatever") {
		t.Error("form dosn't have any field, but shows that has")
	}

	postedData = url.Values{}
	postedData.Add("a", "b")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("form has field a, but shows that not ")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existence field")
	}

	postedData = url.Values{}
	postedData.Add("a", "b")
	form = New(postedData)

	form.MinLength("a", 100)

	if form.Valid() {
		t.Error("minimum length is 100, current is 4")
	}
	
	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("there should be an error, but it actually doesn't")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abcd123")
	form = New(postedData)

	form.MinLength("another_field", 1)

	if !form.Valid() {
		t.Error("minimum length is 1 current is bigger, but it is still an error")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("there should not be an error, but it actually does")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)
	form.IsEmail("a")

	if form.Valid() {
		t.Error("form doesn't exist, but it shows that it is email")
	}

	postedData = url.Values{}
	postedData.Add("a", "b")
	form = New(postedData)
	form.IsEmail("a")
	if form.Valid() {
		t.Error("field isn't an email, but shows that is")
	}
	
	postedData = url.Values{}
	postedData.Add("a", "pohibuskaivan@gmail.com")
	form = New(postedData)
	form.IsEmail("a")
	if !form.Valid() {
		t.Error("field is an email, but shows that isn't")
	}
}
