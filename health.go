package healthchecker

import "os"

func main() {
	urls := getEnv()
	for _, url := range urls {
		go checkHealth(url)
	}
}

func checkHealth(url string) {

}

func getEnv() []string {
	var res []string
	url := os.Getenv("URL1")
	res = append(res, url)
	url = os.Getenv("URL2")
	res = append(res, url)
	url = os.Getenv("URL3")
	res = append(res, url)
	url = os.Getenv("URL4")
	res = append(res, url)
	url = os.Getenv("URL5")
	res = append(res, url)

	return res
}
