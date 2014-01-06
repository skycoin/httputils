package httputils

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func SendJSON(w http.ResponseWriter, message JSONMessage) error {
    /* Writes JSON to an http response */
    out, err := json.Marshal(message)
    if err == nil {
        _, err := w.Write(out)
        if err != nil {
            return err
        }
    }
    return err
}

func GetJSON(address string, route string, message JSONMessage) error {
    /* GET json content from a url built from address and route */
    return GetJSONParams(message, address, route, nil)
}

func GetJSONParams(message JSONMessage, address string, route string,
    params map[string]string) error {
    /* GET json content from a url built from address, route and params */
    // form the complete url
    url := BuildURL("http", address, route, params)

    // make the request
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // read response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    // unpack the json to our message
    err = json.Unmarshal(body, message)
    if err != nil {
        return err
    }
    return nil
}

func PostJSON(message JSONMessage, outgoing interface{}, address string,
    route string) error {
    /* POST json content to a url built from address and route */
    return PostJSONParams(message, outgoing, address, route, nil)
}

func PostJSONParams(message JSONMessage, outgoing interface{}, address string,
    route string, params map[string]string) error {
    /* POST json content to a url built from address, route and params */
    // marshal our data
    data, err := json.Marshal(outgoing)
    if err != nil {
        return err
    }
    reader := bytes.NewReader(data)

    // form the complete url
    url := BuildURL("http", address, route, params)

    // make request
    resp, err := http.Post(url, "application/json", reader)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // read response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    // unmarshal response
    err = json.Unmarshal(body, message)
    if err != nil {
        return err
    }
    return nil
}

/* Messages */

// Outgoing JSON message
type JSONMessage interface{}

// Incoming JSON response
type JSONResponse struct {
    Success bool
    Message string
}

func (r JSONResponse) String() string {
    /* Returns a JSONResponse as a json encoded string. */
    b, err := json.Marshal(r)
    if err != nil {
        return ""
    }
    return string(b)
}

func FailureResponse(message string) string {
    /* Returns a JSONResponse with Success: false */
    return JSONResponse{Success: false, Message: message}.String()
}

func SuccessResponse(message string) string {
    /* Returns a JSONResponse with Success: true */
    return JSONResponse{Success: true, Message: message}.String()
}
