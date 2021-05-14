# GoLaget 🍺
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lnus/golaget)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/lnus/golaget)
[![Go Report Card](https://goreportcard.com/badge/github.com/lnus/golaget)](https://goreportcard.com/report/github.com/lnus/golaget)
![GitHub issues](https://img.shields.io/github/issues/lnus/golaget)
![GitHub](https://img.shields.io/github/license/lnus/golaget)

GoLaget is a wrapper for [Systembolagets V2 API](https://api-portal.systembolaget.se/). It currently includes every function listed under V2 on their open API. **No plans to include V1 right now**, since I don't personally need it. But feel free to PR in case you write it out. ​

## Installation

```bash
go get github.com/lnus/golaget
```
For running the tests, the .env file should be populated with
```bash
PRIMARY_KEY=YOUR_PRIMARY_KEY
```
Please refer to the .env.example file for a safe configuration.
## Usage

```golang
import "github.com/lnus/golaget"

func main() {
    instance = golaget.New(API_KEY, false) // API_KEY & verbose mode
    result, error := instance.GetAgents() // Returns a list of all agents
    for k, v := range *result {
        fmt.Println(v.AgentID, v.IsOpen, v.Phone, v.Address)
    }
}
```
## Contributing ❤
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License ⚖
[MIT](https://choosealicense.com/licenses/mit/)
