package main

import (
	"crypto/x509"
	"strings"
)

func printTLSInfo(cert *x509.Certificate){
	println("Subject:",
		strings.Join(cert.Subject.Country,","),
		strings.Join(cert.Subject.Province,","),
		strings.Join(cert.Subject.Locality,","),
		strings.Join(cert.Subject.Organization,","),
		cert.Subject.CommonName)
	println("Issuer:",
		strings.Join(cert.Issuer.Country,","),
		strings.Join(cert.Issuer.Province,","),
		strings.Join(cert.Issuer.Locality,","),
		strings.Join(cert.Issuer.Organization,","),
		cert.Issuer.CommonName)

	println("Public Key & Signature Algorithm:",
		cert.PublicKeyAlgorithm.String(),",",
		cert.SignatureAlgorithm.String())
}
