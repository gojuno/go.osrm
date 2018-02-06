package osrm

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
)

// options represents OSRM query params to be encoded in URL
type options map[string][]string

// Set saves a string value by the key
func (opts options) set(k, v string) options {
	if v != "" {
		opts[k] = []string{v}
	}
	return opts
}

func (opts options) setStringer(k string, v fmt.Stringer) options {
	return opts.set(k, v.String())
}

// SetBool converts bool to string and set a key
func (opts options) setBool(k string, v bool) options {
	return opts.set(k, fmt.Sprintf("%t", v))
}

// AddInt converts int to string and appends it to the key
func (opts options) addInt(k string, v ...int) options {
	for _, n := range v {
		opts[k] = append(opts[k], strconv.Itoa(n))
	}
	return opts
}

// AddInt64 converts int64 to string and appends it to the key
func (opts options) addInt64(k string, v ...int64) options {
	for _, n := range v {
		opts[k] = append(opts[k], strconv.FormatInt(n, 10))
	}
	return opts
}

// AddFloat converts float to string and appends it to the key
func (opts options) addFloat(k string, v ...float64) options {
	for _, f := range v {
		opts[k] = append(opts[k], strconv.FormatFloat(f, 'f', -1, 64))
	}
	return opts
}

// Add appends values to the key
func (opts options) add(k string, v ...string) options {
	opts[k] = append(opts[k], v...)
	return opts
}

// Encode encodes the options into OSRM query form
// ({option}={element};{element}[;{element} ... ]) sorted by key
func (opts options) encode() string {
	if opts == nil {
		return ""
	}

	keys := make([]string, 0, len(opts))
	for k := range opts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf []byte
	for _, k := range keys {
		if len(buf) > 0 {
			buf = append(buf, '&')
		}
		buf = append(buf, url.QueryEscape(k)...)
		buf = append(buf, '=')
		for n, val := range opts[k] {
			if n > 0 {
				buf = append(buf, ';')
			}
			buf = append(buf, url.QueryEscape(val)...)
		}
	}
	return string(buf)
}
