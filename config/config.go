
// Code generated with https://github.com/mozey/config DO NOT EDIT

package config

import (
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/mozey/config/pkg/share"
	"github.com/pkg/errors"
)

// KeyPrefix is not made publicly available on this package,
// users must use the getter or setter methods.
// This package must not change the config file


// APP_IH2024_AMOUNT
var ih2024Amount string
// APP_IH2024_IN_WALLET_ADDRESS_URL
var ih2024InWalletAddressUrl string
// APP_IH2024_KEY_ID
var ih2024KeyId string
// APP_IH2024_NONCE
var ih2024Nonce string
// APP_IH2024_OUT_WALLET_ADDRESS_URL
var ih2024OutWalletAddressUrl string
// APP_IH2024_PRIVATE_KEY
var ih2024PrivateKey string
// APP_IH2024_SUCCESS_URL
var ih2024SuccessUrl string
// APP_DIR
var dir string

// Config fields correspond to config file keys less the prefix
type Config struct {
	
	ih2024Amount string // APP_IH2024_AMOUNT
	ih2024InWalletAddressUrl string // APP_IH2024_IN_WALLET_ADDRESS_URL
	ih2024KeyId string // APP_IH2024_KEY_ID
	ih2024Nonce string // APP_IH2024_NONCE
	ih2024OutWalletAddressUrl string // APP_IH2024_OUT_WALLET_ADDRESS_URL
	ih2024PrivateKey string // APP_IH2024_PRIVATE_KEY
	ih2024SuccessUrl string // APP_IH2024_SUCCESS_URL
	dir string // APP_DIR
}


// Ih2024Amount is APP_IH2024_AMOUNT
func (c *Config) Ih2024Amount() string {
	return c.ih2024Amount
}
// Ih2024InWalletAddressUrl is APP_IH2024_IN_WALLET_ADDRESS_URL
func (c *Config) Ih2024InWalletAddressUrl() string {
	return c.ih2024InWalletAddressUrl
}
// Ih2024KeyId is APP_IH2024_KEY_ID
func (c *Config) Ih2024KeyId() string {
	return c.ih2024KeyId
}
// Ih2024Nonce is APP_IH2024_NONCE
func (c *Config) Ih2024Nonce() string {
	return c.ih2024Nonce
}
// Ih2024OutWalletAddressUrl is APP_IH2024_OUT_WALLET_ADDRESS_URL
func (c *Config) Ih2024OutWalletAddressUrl() string {
	return c.ih2024OutWalletAddressUrl
}
// Ih2024PrivateKey is APP_IH2024_PRIVATE_KEY
func (c *Config) Ih2024PrivateKey() string {
	return c.ih2024PrivateKey
}
// Ih2024SuccessUrl is APP_IH2024_SUCCESS_URL
func (c *Config) Ih2024SuccessUrl() string {
	return c.ih2024SuccessUrl
}
// Dir is APP_DIR
func (c *Config) Dir() string {
	return c.dir
}


// SetIh2024Amount overrides the value of ih2024Amount
func (c *Config) SetIh2024Amount(v string) {
	c.ih2024Amount = v
}

// SetIh2024InWalletAddressUrl overrides the value of ih2024InWalletAddressUrl
func (c *Config) SetIh2024InWalletAddressUrl(v string) {
	c.ih2024InWalletAddressUrl = v
}

// SetIh2024KeyId overrides the value of ih2024KeyId
func (c *Config) SetIh2024KeyId(v string) {
	c.ih2024KeyId = v
}

// SetIh2024Nonce overrides the value of ih2024Nonce
func (c *Config) SetIh2024Nonce(v string) {
	c.ih2024Nonce = v
}

// SetIh2024OutWalletAddressUrl overrides the value of ih2024OutWalletAddressUrl
func (c *Config) SetIh2024OutWalletAddressUrl(v string) {
	c.ih2024OutWalletAddressUrl = v
}

// SetIh2024PrivateKey overrides the value of ih2024PrivateKey
func (c *Config) SetIh2024PrivateKey(v string) {
	c.ih2024PrivateKey = v
}

// SetIh2024SuccessUrl overrides the value of ih2024SuccessUrl
func (c *Config) SetIh2024SuccessUrl(v string) {
	c.ih2024SuccessUrl = v
}

// SetDir overrides the value of dir
func (c *Config) SetDir(v string) {
	c.dir = v
}


// New creates an instance of Config.
// Build with ldflags to set the package vars.
// Env overrides package vars.
// Fields correspond to the config file keys less the prefix.
// The config file must have a flat structure
func New() *Config {
	conf := &Config{}
	SetVars(conf)
	SetEnv(conf)
	return conf
}

// SetVars sets non-empty package vars on Config
func SetVars(conf *Config) {
	
	if ih2024Amount != "" {
		conf.ih2024Amount = ih2024Amount
	}
	
	if ih2024InWalletAddressUrl != "" {
		conf.ih2024InWalletAddressUrl = ih2024InWalletAddressUrl
	}
	
	if ih2024KeyId != "" {
		conf.ih2024KeyId = ih2024KeyId
	}
	
	if ih2024Nonce != "" {
		conf.ih2024Nonce = ih2024Nonce
	}
	
	if ih2024OutWalletAddressUrl != "" {
		conf.ih2024OutWalletAddressUrl = ih2024OutWalletAddressUrl
	}
	
	if ih2024PrivateKey != "" {
		conf.ih2024PrivateKey = ih2024PrivateKey
	}
	
	if ih2024SuccessUrl != "" {
		conf.ih2024SuccessUrl = ih2024SuccessUrl
	}
	
	if dir != "" {
		conf.dir = dir
	}
	
}

// SetEnv sets non-empty env vars on Config
func SetEnv(conf *Config) {
	var v string

	
	v = os.Getenv("APP_IH2024_AMOUNT")
	if v != "" {
		conf.ih2024Amount = v
	}
	
	v = os.Getenv("APP_IH2024_IN_WALLET_ADDRESS_URL")
	if v != "" {
		conf.ih2024InWalletAddressUrl = v
	}
	
	v = os.Getenv("APP_IH2024_KEY_ID")
	if v != "" {
		conf.ih2024KeyId = v
	}
	
	v = os.Getenv("APP_IH2024_NONCE")
	if v != "" {
		conf.ih2024Nonce = v
	}
	
	v = os.Getenv("APP_IH2024_OUT_WALLET_ADDRESS_URL")
	if v != "" {
		conf.ih2024OutWalletAddressUrl = v
	}
	
	v = os.Getenv("APP_IH2024_PRIVATE_KEY")
	if v != "" {
		conf.ih2024PrivateKey = v
	}
	
	v = os.Getenv("APP_IH2024_SUCCESS_URL")
	if v != "" {
		conf.ih2024SuccessUrl = v
	}
	
	v = os.Getenv("APP_DIR")
	if v != "" {
		conf.dir = v
	}
	
}

// GetMap of all env vars
func (c *Config) GetMap() map[string]string {
	m := make(map[string]string)
	
	m["APP_IH2024_AMOUNT"] = c.ih2024Amount
	
	m["APP_IH2024_IN_WALLET_ADDRESS_URL"] = c.ih2024InWalletAddressUrl
	
	m["APP_IH2024_KEY_ID"] = c.ih2024KeyId
	
	m["APP_IH2024_NONCE"] = c.ih2024Nonce
	
	m["APP_IH2024_OUT_WALLET_ADDRESS_URL"] = c.ih2024OutWalletAddressUrl
	
	m["APP_IH2024_PRIVATE_KEY"] = c.ih2024PrivateKey
	
	m["APP_IH2024_SUCCESS_URL"] = c.ih2024SuccessUrl
	
	m["APP_DIR"] = c.dir
	
	return m
}

// LoadMap sets the env from a map and returns a new instance of Config
func LoadMap(configMap map[string]string) (conf *Config)  {
	for key, val := range configMap {
		_ = os.Setenv(key, val)
	}
	return New()
}

// SetEnvBase64 decodes and sets env from the given base64 string
func SetEnvBase64(configBase64 string) (err error) {
	// Decode base64
	decoded, err := base64.StdEncoding.DecodeString(configBase64)
	if err != nil {
		return errors.WithStack(err)
	}
	// UnMarshall json
	configMap := make(map[string]string)
	err = json.Unmarshal(decoded, &configMap)
	if err != nil {
		return errors.WithStack(err)
	}
	// Set config
	for key, value := range configMap {
		err = os.Setenv(key, value)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// LoadFile sets the env from file and returns a new instance of Config
func LoadFile(env string) (conf *Config, err error) {
	appDir := os.Getenv("APP_DIR")
	if appDir == "" {
		// Use current working dir
		appDir, err = os.Getwd()
		if err != nil {
			return conf, errors.WithStack(err)
		}
	}

	var configPath string
	filePaths, err := share.GetConfigFilePaths(appDir, env)
	if err != nil {
		return conf, err
	}
	for _, configPath = range filePaths {
		_, err := os.Stat(configPath)
		if err != nil {
			if os.IsNotExist(err) {
				// Path does not exist
				continue
			}
			return conf, errors.WithStack(err)
		}
		// Path exists
		break
	}
	if configPath == "" {
		return conf, errors.Errorf("config file not found in %s", appDir)
	}

	b, err := os.ReadFile(configPath)
	if err != nil {
		return conf, errors.WithStack(err)
	}

	configMap, err := share.UnmarshalConfig(configPath, b)
	if err != nil {
		return conf, err
	}
	for key, val := range configMap {
		_ = os.Setenv(key, val)
	}
	return New(), nil
}
