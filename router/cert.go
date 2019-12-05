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
    	caCertPool := x509.NewCertPool()
    	caCertPool.AppendCertsFromPEM([]byte(*ca.CertificateChain))

    	// Setup HTTPS client
    	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        	RootCAs:      caCertPool,
    	}
    	tlsConfig.BuildNameToCertificate()
	return tlsConfig
}
