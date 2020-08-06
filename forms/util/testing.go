package util

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"

	"golang.org/x/text/message"
)

func BuildMultiRequest(mpw *multipart.Writer, buff *bytes.Buffer) *http.Request {
	mpw.Close()

	req := httptest.NewRequest("POST", "/", buff)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+mpw.Boundary())

	return req
}

func WriteFile(mpw *multipart.Writer, key, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	fw, err := mpw.CreateFormFile(key, path)
	if err != nil {
		return err
	}

	_, err = io.Copy(fw, file)
	return err
}

func TestValues(values url.Values, fn func(*http.Request, *message.Printer) []string) error {
	req := httptest.NewRequest("GET", "/", nil)

	req.PostForm = values
	req.Header.Set("Content-Type", "form-urlencoded")

	printer := message.NewPrinter(message.MatchLanguage("en"))
	for _, err := range fn(req, printer) {
		return errors.New(err)
	}

	return nil
}
