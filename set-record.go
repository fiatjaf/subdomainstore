package main

import (
	"net/http"
	"strconv"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gorilla/mux"
)

func setRecord(w http.ResponseWriter, r *http.Request) {
	subdomain := mux.Vars(r)["subdomain"]
	typ := mux.Vars(r)["type"]
	name := mux.Vars(r)["name"]
	target := mux.Vars(r)["target"]
	log.Debug().Str("subdomain", subdomain).Str("target", target).Msg("setting record")

	if !checkAuthorization(r, subdomain) {
		http.Error(w, "invalid authorization token. use the preimage of the payment sent to buy '"+subdomain+"'.", 401)
		return
	}

	priority, _ := strconv.Atoi(r.URL.Query().Get("priority"))
	proxy := r.URL.Query().Get("proxy") != ""
	if name == "@" {
		name = subdomain
	} else {
		name = name + "." + subdomain
	}

	resp, err := cf.CreateDNSRecord(s.CloudflareZoneId, cloudflare.DNSRecord{
		Type:     typ,
		Name:     name,
		Content:  target,
		TTL:      1,
		Priority: priority,
		Proxied:  proxy,
	})
	if err != nil || resp.Response.Success != true {
		log.Warn().Err(err).Interface("response", resp.Response).
			Msg("error setting dns record")
		http.Error(w, "error setting dns record", 400)
	}
}
