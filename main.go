// utility for determining single origin
// placeholder for http request

package main

reader := strings.NewReader(`{"body":123}`)
request, err := http.NewRequest("GET", "http://localhost:3030/foo", reader)
// TODO: check err
client := &http.Client{}
resp, err := client.Do(request)
// TODO: check err
