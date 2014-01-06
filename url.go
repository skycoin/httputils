package httputils

import (
    "fmt"
    "strings"
)

func BuildURL(protocol string, address string, route string,
    params map[string]string) string {
    /* Construct a URL from protocol, address, route and params.
       protocol is e.g. http, https.  If a protocol is detected in address,
       this value is ignored.
       address is the domain name or ip address.
       route is the resource path for the destination address
       params are optional and are used for the GET query string
    */
    if strings.HasPrefix(address, "http://") ||
        strings.HasPrefix(address, "https://") {
        protocol = ""
    }
    query := ""
    if params != nil {
        var _params []string
        for key, value := range params {
            _params = append(_params, fmt.Sprintf("%s=%s", key, value))
        }
        query = "?" + strings.Join(_params, "&")
    }

    if protocol != "" && !strings.HasSuffix(protocol, "://") {
        protocol += "://"
    }
    return fmt.Sprintf("%s%s%s%s", protocol, address, route, query)
}
