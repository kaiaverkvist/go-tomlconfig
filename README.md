# go-tomlconfig
    Generic library to serialize and deserialize structs intended for config files.

I created this project because I found myself copy pasting a "config.go" file (*which only served to
serialize and deserialize some toml from/into structs*) between all my projects.

## How to use 101:

### 🧩 Get the library
`go get -u github.com/kaiaverkvist/go-tomlconfig`

### 🖨️ Copy this code into your project or whatever
```go
// Set up a default config type:
type Config struct{
	Username string
	Password string
}

var defaultConfig = Config{
	Username: "test",
	Password: "test",
}

func main() {
    // This either loads the config itself 
    err := config.LoadOrCreateConfig("example.toml", defaultConfig)
    if err != nil {
        // In this case, you might want to use your default config instead, or keep the application from starting up!
        log.Println("cannot load or create config:", err.Error())
        return
    }
    
    // Unfortunately a type conversion is required due to the lack of generics in Go at this time.
    conf := config.GetConfig().(*Config)
}
```

## 🏗️ Contributing
Contributions are very much welcome. The thing I want is a clean and useful PR formatted using `go fmt`.

## 🧪 Tests
Run tests with `go test -v ./...`