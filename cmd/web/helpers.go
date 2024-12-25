package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, status int, data templateData) {
	tmpl, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buff := new(bytes.Buffer)

	err := tmpl.ExecuteTemplate(buff, "layout", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buff.WriteTo(w)
}

func (app *application) newTemplateData() templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}
