package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

// 不验证服务端证书
func TestSkipVerify(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	cli := &http.Client{
		Timeout:   3 * time.Second,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", "https://127.0.0.1:8787/hello", nil)
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

// 验证服务端证书（服务端自签名证书）
func TestServerVerificationSelf(t *testing.T) {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("../keys/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			// InsecureSkipVerify: true,
			RootCAs:    pool,
			ServerName: "localhost", // 与证书 Comman Name 相同
		},
	}
	cli := &http.Client{
		Timeout:   3 * time.Second,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", "https://127.0.0.1:8787/hello", nil)
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

//验证服务端证书
func TestServerVerification(t *testing.T) {
	pool := x509.NewCertPool()
	cacrt, err := ioutil.ReadFile("../keys/ca/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	pool.AppendCertsFromPEM(cacrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			//InsecureSkipVerify: true,
			RootCAs:    pool,
			ServerName: "localhost",
		},
	}
	cli := &http.Client{
		Timeout:   3 * time.Second,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", "https://127.0.0.1:8787/hello", nil)
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

// 双向认证
func TestClientVerification(t *testing.T) {
	pool := x509.NewCertPool()
	cacrt, err := ioutil.ReadFile("../keys/ca/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	pool.AppendCertsFromPEM(cacrt)

	cliCrt, err := tls.LoadX509KeyPair("../keys/ca/client.crt", "../keys/ca/client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			//InsecureSkipVerify: true,
			RootCAs:      pool,
			ServerName:   "localhost",
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	cli := &http.Client{
		Timeout:   3 * time.Second,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", "https://127.0.0.1:8787/hello", nil)
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
