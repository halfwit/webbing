package router

import (
    "crypto/tls"
    "crypto/x509"
    "log"

    "github.com/aws/aws-sdk-go/service/acm"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)

func getTlsConfig() *tls.Config {
	cl := aws.NewConfig().WithRegion("us-east-2")
	sess := session.Must(session.NewSession(cl))
	svc := acm.New(sess)
	arn := "arn:aws:acm:us-east-2:824263434500:certificate/aa0ae6e7-075a-466c-bcb5-8d7874447bcb"
	ca, err := svc.GetCertificate(&acm.GetCertificateInput{
		CertificateArn: &arn,
	})
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
    	if err != nil {
        	log.Fatal(err)
    	}

    	caCertPool := x509.NewCertPool()
    	caCertPool.AppendCertsFromPEM([]byte(*ca.CertificateChain))

    	// Setup HTTPS client
    	tlsConfig := &tls.Config{
        	Certificates: []tls.Certificate{cert},
        	RootCAs:      caCertPool,
    	}
    	tlsConfig.BuildNameToCertificate()
	return tlsConfig
}
