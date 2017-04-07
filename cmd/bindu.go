
package cmd

import (
    "net/http"
)

type Status int

const (
    UNCHECKED Status = iota
    DOWN
    UP
)

// The Site struct encapsulates the details about the site being monitored.
type Site struct {
    url         string
    last_status Status
}

// Site.Status makes a GET request to a given URL and checks whether or not the
// resulting status code is 200.
func (s Site) Status() (Status, error) {
    resp, err := http.Get(s.url)
    status := s.last_status

    if (err == nil) && (resp.StatusCode == 200) {
        status = UP
    } else {
        status = DOWN
    }

    return status, err
}
