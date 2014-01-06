package httputils

import (
    "html/template"
    "net/http"
    "strings"
)

/* HTTP Server Utilities */

func ListenAndServeBackground(address string, mux *http.ServeMux,
    errchan chan<- error) {
    /* Runs an http server in the background */
    err := http.ListenAndServe(address, mux)
    errchan <- err
}

/* HTTP Error Responses */

func HttpError(w http.ResponseWriter, status int, messages ...string) {
    /* Emits an HTTP status code to an http response */
    message := strings.Join(messages, "<br>")
    http.Error(w, message, status)
}

func Error400(w http.ResponseWriter, messages ...string) {
    /* Emits an HTTP 400 error to an http response */
    messages = createDefaultMessage("Bad request", messages)
    HttpError(w, http.StatusBadRequest, messages...)
}

func Error404(w http.ResponseWriter, messages ...string) {
    /* Emits an HTTP 404 error to an http response */
    messages = createDefaultMessage("Not found", messages)
    HttpError(w, http.StatusNotFound, messages...)
}

func Error405(w http.ResponseWriter, messages ...string) {
    /* Emits an HTTP 405 error to an http response */
    messages = createDefaultMessage("Method not allowed", messages)
    HttpError(w, http.StatusMethodNotAllowed, messages...)
}

func Error500(w http.ResponseWriter, messages ...string) {
    /* Emits an HTTP 500 error to an http response */
    messages = createDefaultMessage("Internal server error", messages)
    HttpError(w, http.StatusInternalServerError, messages...)
}

func Error501(w http.ResponseWriter, messages ...string) {
    /* Emits an HTTP 501 error to an http response */
    messages = createDefaultMessage("Not implemented", messages)
    HttpError(w, http.StatusNotImplemented, messages...)
}

func createDefaultMessage(message string, messages []string) []string {
    /* Returns the message as the sole element in a string slice if messages
       is empty, otherwise returns messages unmodified */
    if len(messages) == 0 {
        messages = make([]string, 1)
        messages[0] = message
    }
    return messages
}

/* Template helpers */

func ShowTemplate(w http.ResponseWriter, t *template.Template,
    p interface{}) error {
    /* Renders a template to an http response */
    err := t.Execute(w, p)
    if err != nil {
        Error500(w, err.Error())
        return err
    }
    return nil
}
