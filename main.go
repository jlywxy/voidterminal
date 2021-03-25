package main

import (
	"crypto/tls"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	var c net.Conn
	var e error
	//os.Args=append(os.Args,"tcp:la.jlywxy.de:9000")
	//os.Args=append(os.Args,"--tls")
	if len(os.Args)<2{
		println("invalid arguments")
		println(`usage:
voidterminal network:address
	network: unix, tcp, tls
	address: /path/to/socket, IP:Port
`)
		return
	}
	println("connecting to "+os.Args[1]+" ...")

	addr_set:=strings.Split(os.Args[1],":")
	nett:=addr_set[0]
	addr:=strings.Join(addr_set[1:],":")
	if nett=="tls"{
		conf := &tls.Config{
			InsecureSkipVerify: true,
		}
		if strings.Index(addr,":")==-1{
			addr+=":443"
		}
		tlsc,tlse:=tls.Dial("tcp",addr,conf)
		if tlse!=nil{
			log.Fatal(tlse)
		}
		//for _,cert := range tlsc.ConnectionState().PeerCertificates{
		cert:=tlsc.ConnectionState().PeerCertificates[0]
		println("---------TLS INFORMATION---------")
		println("Server Name:",tlsc.ConnectionState().ServerName)
		printTLSInfo(cert)
		if cert.Subject.CommonName!=tlsc.ConnectionState().ServerName{
			println(" * WARNING: server name and subject common name not match.")
		}
		println("------------------------------")
		//}
		c=tlsc; e=tlse
	}else{
		c,e=net.Dial(nett,addr)
	}

	if e!=nil{
		log.Fatal(e)
		println("connection failed.")
		return
	}
	s,_:=terminal.GetState(0)
	terminal.MakeRaw(0)
	go func(){
		io.Copy(os.Stdout,c)
		terminal.Restore(0,s)
		println("disconnected")
		os.Exit(0)
	}()
	io.Copy(c,os.Stdin)
	terminal.Restore(0,s)
}

//package main
//
//import (
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/elliptic"
//	"crypto/rand"
//	"encoding/hex"
//	"github.com/aead/ecdh"
//	"github.com/decred/dcrd/dcrec/secp256k1"
//	"golang.org/x/crypto/ssh/terminal"
//	"io"
//	"log"
//	"net"
//	"os"
//)
//
//func main(){
//	s,_:=terminal.GetState(0)
//
//	//sock:="/Users/jlywxy/voidshell/vssock1"//os.Args[1]
//	sock:="127.0.0.1:9000"
//	println("connecting to "+sock+" ...")
//	//c,e:=net.Dial("unix",sock)
//	c,e:=net.Dial("tcp",sock)
//	if e!=nil{
//		log.Fatal(e)
//	}
//	p256 := ecdh.Generic(secp256k1.S256())
//
//
//	privkey,pubkey,_:=p256.GenerateKey(rand.Reader)
//	ppubkey:=pubkey.(ecdh.Point)
//	pub_b:=elliptic.Marshal(secp256k1.S256(),ppubkey.X,ppubkey.Y)
//	println("this pubkey "+hex.EncodeToString(pub_b))
//
//	ser_pub_b:=make([]byte,65)
//	c.Write(pub_b)
//	c.Read(ser_pub_b)
//	println("get server pubkey "+hex.EncodeToString(ser_pub_b))
//	x,y:=elliptic.Unmarshal(secp256k1.S256(),ser_pub_b)
//	ser_pub:=ecdh.Point{x,y}
//	key:=p256.ComputeSecret(privkey,ser_pub)
//
//	terminal.MakeRaw(0)
//	//key, _ := hex.DecodeString("6368616e676520746869732070617373")
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//	var iv [aes.BlockSize]byte
//	stream := cipher.NewCTR(block, iv[:])
//
//
//
//	go io.Copy(os.Stdout,cipher.StreamReader{
//		S: stream,
//		R: c,
//	})
//	io.Copy(cipher.StreamWriter{
//		S: stream,
//		W: c,
//	},os.Stdin)
//	terminal.Restore(0,s)
//}
//
