package env

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment map[string]interface{}

func init() {
	_ = godotenv.Load(".env")
}

func NewEnvironment() Environment {
	return Environment{}
}

func (e Environment) SetEnv(name string, value interface{}) Environment {
	e[name] = value
	return e
}

func (e Environment) GetFloat64(key string) float64 {
	value := e[key]

	valueAsString := value.(string)
	valueAsFloat, err := strconv.ParseFloat(valueAsString, 8)

	if err != nil {
		log.Fatal("couldn't parse value as string: ", err.Error())
		return 0
	}
	return valueAsFloat
}

func (e Environment) GetAsString(key string) string {
	value := e[key]

	valueAsString, ok := value.(string)
	if !ok {
		return ""
	}

	return valueAsString
}

func (e Environment) GetAsBytes(key string) []byte {
	value := e[key]

	valueAsString, ok := value.(string)
	if !ok {
		return nil
	}

	return []byte(valueAsString)
}
