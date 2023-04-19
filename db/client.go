package db

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var Session *gocql.Session

func InitClient() (*gocql.Session, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	astra_uri := os.Getenv("ASTRA_URI")
	astra_username := os.Getenv("ASTRA_USERNAME")
	astra_password := os.Getenv("ASTRA_PASSWORD")

	if astra_uri == "" || astra_username == "" || astra_password == "" {
		log.Fatal("Astra URI is empty")
	}

	cluster := gocql.NewCluster(astra_uri)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: astra_username,
		Password: astra_password,
	}
	certPath, _ := filepath.Abs("./secure-connect-way/cert") //extracted bundle
	caPath, _ := filepath.Abs("./secure-connect-way/ca.crt")
	keyPath, _ := filepath.Abs("./secure-connect-way/key")

	cluster.Keyspace = "way"
	cluster.Port = 29042

	cluster.ProtoVersion = 4
	cluster.CQLVersion = "3.4.5"
	cluster.ConnectTimeout = time.Second * 6
	myRootCAs, err := os.ReadFile("./secure-connect-way/ca.crt")
	if err != nil {
		log.Fatal("Failed to read CA certificate: ", err)
	}

	cert, err := tls.LoadX509KeyPair("./secure-connect-way/cert", "./secure-connect-way/key")
	if err != nil {
		log.Fatal("Failed to read client certificate and key: ", err)
	}

	Config := &tls.Config{
		ServerName:         astra_uri,
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
		Certificates:       []tls.Certificate{cert},

		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		RootCAs:    x509.NewCertPool(),
		ClientAuth: tls.VerifyClientCertIfGiven,
	}

	Config.RootCAs.AppendCertsFromPEM(myRootCAs)

	cluster.SslOpts = &gocql.SslOptions{
		CertPath:               certPath,
		CaPath:                 caPath,
		KeyPath:                keyPath,
		EnableHostVerification: true,
		Config:                 Config,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		println(err.Error(), " 🚀")
		return nil, err
	}
	println("connected 🚀")
	Session = session
	return session, nil

}
