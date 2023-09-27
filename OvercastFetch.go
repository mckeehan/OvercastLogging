package main

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

func GetOpml(cfg OvercastConfig) string {
	loginUrl, err := url.Parse( cfg.LoginUrl )
	if err != nil {
		log.Fatal(err)
	}
	dataUrl, err := url.Parse( cfg.DataUrl )
	if err != nil {
		log.Fatal(err)
	}

	// All users of cookiejar should import "golang.org/x/net/publicsuffix"
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	client.PostForm(loginUrl.String(), url.Values{
		"email": {cfg.Email},
		"password": {cfg.Password},
	})
	if _, err = client.Get(loginUrl.String()); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("After 1st request:")
	// for _, cookie := range jar.Cookies(loginUrl) {
	// 	fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	// }

	resp, err := client.Get(dataUrl.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.Status != "200 OK" {
		log.Fatal("Invalid response: ", resp.Status, "\n", string(body))
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
