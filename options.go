package osrm

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
)

type Options map[string][]string

// Set saves a string value by the key
func (opts Options) Set(k, v string) {
	opts[k] = []string{v}
}

// SetBool converts bool to string and set a key
func (opts Options) SetBool(k string, v bool) {
	opts.Set(k, fmt.Sprintf("%t", v))
}

// AddInt converts int to string and appends it to the key
func (opts Options) AddInt(k string, v ...int) {
	for _, n := range v {
		opts[k] = append(opts[k], strconv.Itoa(n))
	}
}

// AddInt64 converts int64 to string and appends it to the key
func (opts Options) AddInt64(k string, v ...int64) {
	for _, n := range v {
		opts[k] = append(opts[k], strconv.FormatInt(n, 10))
	}
}

// AddFloat converts float to string and appends it to the key
func (opts Options) AddFloat(k string, v ...float64) {
	for _, f := range v {
		opts[k] = append(opts[k], strconv.FormatFloat(f, 'f', -1, 64))
	}
}

// Add appends values to the key
func (opts Options) Add(k string, v ...string) {
	opts[k] = append(opts[k], v...)
}

// Encode encodes the options into ``OSRM query`` form
// ({option}={element};{element}[;{element} ... ]) sorted by key
func (opts Options) Encode() string {
	if opts == nil {
		return ""
	}
	var buf []byte
	keys := make([]string, 0, len(opts))
	for k := range opts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vals := opts[k]
		if len(buf) > 0 {
			buf = append(buf, '&')
		}
		buf = append(buf, url.QueryEscape(k)...)
		buf = append(buf, '=')
		for n, val := range vals {
			if n > 0 {
				buf = append(buf, ';')
			}
			buf = append(buf, url.QueryEscape(val)...)
		}
	}
	return string(buf)
}
