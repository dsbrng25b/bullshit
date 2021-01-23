package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dvob/bullshit"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	var (
		tls  bool
		host string
	)

	flag.BoolVar(&tls, "tls", false, "use acme to configure certificates and serve tls")
	flag.StringVar(&host, "host", "", "hostname for the certificate")
	flag.Parse()

	http.HandleFunc("/", Bullshit)

	s := &http.Server{
		Addr: ":80",
	}

	if tls {
		if host == "" {
			log.Fatal("flag -host required with -tls")
		}
		m := &autocert.Manager{
			Cache:      autocert.DirCache("acme"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(host),
		}
		s.Addr = ":443"
		s.TLSConfig = m.TLSConfig()
		log.Fatal(s.ListenAndServeTLS("", ""))
		return
	}

	log.Fatal(s.ListenAndServe())

}

func Bullshit(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	fmt.Fprintf(w, bullshit.Get())
}
