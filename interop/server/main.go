package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/lucas-clemente/quic-go/internal/testdata"
)

func main() {
	testcase := os.Getenv("TESTCASE")
	var quicConf *quic.Config
	switch testcase {
	case "transfer":
		quicConf = &quic.Config{
			AcceptToken: func(_ net.Addr, _ *quic.Token) bool { return true },
		}
	case "retry":
	default:
		fmt.Printf("unsupported test case: %s\n", testcase)
		os.Exit(127)
	}

	if err := runServer(quicConf); err != nil {
		panic(err)
	}
}

func runServer(quicConf *quic.Config) error {
	server := http3.Server{
		Server:     &http.Server{Addr: "0.0.0.0:443"},
		QuicConfig: quicConf,
	}
	http.DefaultServeMux.Handle("/", http.FileServer(http.Dir("/www")))
	return server.ListenAndServeTLS(testdata.GetCertificatePaths())
}
