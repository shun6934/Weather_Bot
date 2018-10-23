package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    values := url.Values{}
    base_url := "https://map.yahooapis.jp/weather/V1/place?"

    values.Add("coordinates", "36.530213,136.627226")
    values.Add("appid", "dj00aiZpPUd0aTBqdms3dW1sOCZzPWNvbnN1bWVyc2VjcmV0Jng9YzU-")
    values.Add("output", "json")
    resp, err := http.Get(base_url + values.Encode())
    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()

    execute(resp)
}

func execute(response *http.Response) {
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}