package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func SendRequest(path string) {

	const (
		proto = "http://"
		id    = "localhost"
		dir   = "/appointment"
		port  = ":2088"

		//cookieName  = "flag"
		//cookieValue = "vka{example_flag}"
	)

	url := proto + id + port + dir + path
	url = url[:len(url)-1]

	// jar, err := cookiejar.New(nil)
	// if err != nil {
	// 	fmt.Println("Error creating cookie jar:", err)
	// 	return
	// }

	client := &http.Client{
		//Jar:     jar,
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// req.AddCookie(&http.Cookie{
	// 	Name:  cookieName,
	// 	Value: cookieValue,
	// })

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		fmt.Println("Request:", req.URL.String())
		return
	}

	defer resp.Body.Close()

}
