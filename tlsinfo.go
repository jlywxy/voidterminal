package main

import (
	"crypto/x509"
	"strings"
)

func printTLSInfo(cert x509.Certificate){
	println("Subject Info:",
		strings.Join(cert.Subject.Country,","),
		strings.Join(cert.Subject.Province,","),
		strings.Join(cert.Subject.Locality,","),
		strings.Join(cert.Subject.Organization,","),
		cert.Subject.CommonName)
	println("Issuer Info:",
		strings.Join(cert.Issuer.Country,","),
		strings.Join(cert.Issuer.Province,","),
		strings.Join(cert.Issuer.Locality,","),
		strings.Join(cert.Issuer.Organization,","),
		cert.Issuer.CommonName)

	println("Public Key Algorithm:",
		cert.PublicKeyAlgorithm.String(),
		"\nSignature Algorithm:",
		cert.SignatureAlgorithm.String())
}
