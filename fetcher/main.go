package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type Price struct {
    Bitcoin struct {
        USD float64 `json:"usd"`
    } `json:"bitcoin"`
}

func main() {
    url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var price Price
    json.Unmarshal(body, &price)

    output, _ := json.MarshalIndent(price, "", "  ")
    ioutil.WriteFile("../analyzer/latest.json", output, 0644)

    fmt.Printf("✅ Saved latest Bitcoin price (USD) to analyzer/latest.json\n")
}

