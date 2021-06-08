package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestAuto(t *testing.T) {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{
	//ServerName: "localhost", //与证书 Comman Name 相同

	//	},
	//}
	cli := &http.Client{
		Timeout: 10 * time.Second,
		//Transport: tr,
	}
	req, err := http.NewRequest("GET", "https://cmkj.lunarhalo.cn", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
