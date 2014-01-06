package httputils

import (
    "errors"
    "net"
)

/* Common IP address functions */

func LocalIP() (net.IP, error) {
    /* Returns the local machine's IP address */
    tt, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    for _, t := range tt {
        aa, err := t.Addrs()
        if err != nil {
            return nil, err
        }
        for _, a := range aa {
            ipnet, ok := a.(*net.IPNet)
            if !ok {
                continue
            }
            v4 := ipnet.IP.To4()
            if v4 == nil || v4[0] == 127 { // loopback address
                continue
            }
            return v4, nil
        }
    }
    return nil, errors.New("cannot find local IP address")
}

func LocalIPString() (string, error) {
    /* Returns the local machine's IP address as a string */
    _ip, err := LocalIP()
    var ip string = ""
    if err == nil {
        ip = _ip.String()
    }
    return ip, err
}
