package keys

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func GetServerCertificate() tls.Certificate {
	cert, err := tls.LoadX509KeyPair("../../keys/server.crt", "../../keys/server.key")
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func GetClientCertificate() tls.Certificate {
	cert, err := tls.LoadX509KeyPair("../../keys/client.crt", "../../keys/client.key")
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func GetCA() *x509.CertPool {
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../keys/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	return certPool
}
