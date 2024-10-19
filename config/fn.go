
// Code generated with https://github.com/mozey/config DO NOT EDIT

package config

import (
	"fmt"
	"strconv"
	"strings"
)

type Fn struct {
	input string
	// output of the last function,
	// might be useful when chaining multiple functions?
	output string
}

// .............................................................................
// Methods to set function input


// FnIh2024Amount sets the function input to the value of APP_IH2024_AMOUNT
func (c *Config) FnIh2024Amount() *Fn {
	fn := Fn{}
	fn.input = c.ih2024Amount
	fn.output = ""
	return &fn
}

// FnIh2024InWalletAddressUrl sets the function input to the value of APP_IH2024_IN_WALLET_ADDRESS_URL
func (c *Config) FnIh2024InWalletAddressUrl() *Fn {
	fn := Fn{}
	fn.input = c.ih2024InWalletAddressUrl
	fn.output = ""
	return &fn
}

// FnIh2024KeyId sets the function input to the value of APP_IH2024_KEY_ID
func (c *Config) FnIh2024KeyId() *Fn {
	fn := Fn{}
	fn.input = c.ih2024KeyId
	fn.output = ""
	return &fn
}

// FnIh2024Nonce sets the function input to the value of APP_IH2024_NONCE
func (c *Config) FnIh2024Nonce() *Fn {
	fn := Fn{}
	fn.input = c.ih2024Nonce
	fn.output = ""
	return &fn
}

// FnIh2024OutWalletAddressUrl sets the function input to the value of APP_IH2024_OUT_WALLET_ADDRESS_URL
func (c *Config) FnIh2024OutWalletAddressUrl() *Fn {
	fn := Fn{}
	fn.input = c.ih2024OutWalletAddressUrl
	fn.output = ""
	return &fn
}

// FnIh2024PrivateKey sets the function input to the value of APP_IH2024_PRIVATE_KEY
func (c *Config) FnIh2024PrivateKey() *Fn {
	fn := Fn{}
	fn.input = c.ih2024PrivateKey
	fn.output = ""
	return &fn
}

// FnIh2024SuccessUrl sets the function input to the value of APP_IH2024_SUCCESS_URL
func (c *Config) FnIh2024SuccessUrl() *Fn {
	fn := Fn{}
	fn.input = c.ih2024SuccessUrl
	fn.output = ""
	return &fn
}

// FnDir sets the function input to the value of APP_DIR
func (c *Config) FnDir() *Fn {
	fn := Fn{}
	fn.input = c.dir
	fn.output = ""
	return &fn
}


// .............................................................................
// Type conversion functions

// Bool parses a bool from the value or returns an error.
// Valid values are "1", "0", "true", or "false".
// The value is not case-sensitive
func (fn *Fn) Bool() (bool, error) {
	v := strings.ToLower(fn.input)
	if v == "1" || v == "true" {
		return true, nil
	}
	if v == "0" || v == "false" {
		return false, nil
	}
	return false, fmt.Errorf("invalid value %s", fn.input)
}

// Float64 parses a float64 from the value or returns an error
func (fn *Fn) Float64() (float64, error) {
	f, err := strconv.ParseFloat(fn.input, 64)
	if err != nil {
		return f, err
	}
	return f, nil
}

// Int64 parses an int64 from the value or returns an error
func (fn *Fn) Int64() (int64, error) {
	i, err := strconv.ParseInt(fn.input, 10, 64)
	if err != nil {
		return i, err
	}
	return i, nil
}

// String returns the input as is
func (fn *Fn) String() string {
	return fn.input
}
