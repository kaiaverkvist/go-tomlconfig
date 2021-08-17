package config

import (
	"errors"
	"fmt"
	"github.com/kaiaverkvist/go-tomlconfig/pkg/serializer"
	"io/ioutil"
	"os"
	"reflect"
)

type ConfigSerializer interface {
	Serialize(config interface{}) ([]byte, error)
	Deserialize(path string, out interface{}) (interface{}, error)
}

// Holds a configuration, making it available for use.
type configContainer struct {
	serializer       ConfigSerializer
	configStructType reflect.Type
	defaultConfig    interface{}

	hasLoadedConfig bool

	actualConfig interface{}
}

// Returns a configContainer that uses the default TOML serializer.
func NewConfigContainer(defaultConfig interface{}) *configContainer {
	container := configContainer{}
	container.SetDefaultConfig(defaultConfig)
	container.SetSerializer(serializer.NewTomlSerializer())

	return &container
}

// Returns a config container where you can specify serializer in constructor.
func NewConfigContainerWithSerializer(defaultConfig interface{}, serializer ConfigSerializer) *configContainer {
	container := configContainer{}
	container.SetDefaultConfig(defaultConfig)
	container.SetSerializer(serializer)

	return &container
}

// Sets the serializer used to serialize and deserialize the config file. By default: serializer.TomlSerializer
func (cc *configContainer) SetSerializer(serializer ConfigSerializer) {
	cc.serializer = serializer
}

// Sets the default configuration.
func (cc *configContainer) SetDefaultConfig(defaultConfig interface{}) {
	cc.configStructType = reflect.TypeOf(defaultConfig)
	cc.defaultConfig = defaultConfig
}

// Load looks for a config file at the specified path, and deserializes it using the serializer and returns the config.
func (cc *configContainer) Load(path string) (interface{}, error) {
	if cc.configStructType == nil || cc.defaultConfig == nil {
		return nil, errors.New("must set default config first")
	}

	// If a serializer hasn't been set we can use the default Toml serializer.
	if cc.serializer == nil {
		cc.serializer = serializer.NewTomlSerializer()
	}

	// Here we should return our configuration as a struct of type typeof: cc.defaultConfig
	instance := reflect.New(cc.configStructType)
	actualInterface := instance.Interface()

	deserializedConfig, err := cc.serializer.Deserialize(path, actualInterface)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot deserialize config: %s", err.Error()))
	}

	cc.hasLoadedConfig = true
	cc.actualConfig = deserializedConfig

	return deserializedConfig, nil
}

// Writes serialized config file at specified path.
func (cc *configContainer) Write(data interface{}, path string) error {
	// If a serializer hasn't been set we can use the default Toml serializer.
	if cc.serializer == nil {
		cc.serializer = serializer.NewTomlSerializer()
	}

	bytes, err := cc.serializer.Serialize(data)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to serialize into byte array: %s", err.Error()))
	}

	err = ioutil.WriteFile(path, bytes, os.ModePerm)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to write config file at location <%s>: %s", path, err.Error()))
	}

	cc.actualConfig = data

	return nil
}
