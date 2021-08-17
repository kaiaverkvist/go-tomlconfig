package serializer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
)

type TomlSerializer struct{}

func NewTomlSerializer() TomlSerializer {
	return TomlSerializer{}
}

func (ts TomlSerializer) Serialize(config interface{}) ([]byte, error) {
	var buffer bytes.Buffer

	// Create a new encoder with no indenting (because it's ugly!).
	encoder := toml.NewEncoder(&buffer)
	encoder.Indent = ""

	// Run the encoding with default config.
	err := encoder.Encode(config)

	return buffer.Bytes(), err
}

func (ts TomlSerializer) Deserialize(path string, out interface{}) (interface{}, error) {

	// Attempt to deserialize the file into the &conf pointer reference.
	if _, err := toml.DecodeFile(path, out); err != nil {
		return nil, errors.New(fmt.Sprintf("cannot deserialize config: %s", err.Error()))
	}

	return out, nil
}
