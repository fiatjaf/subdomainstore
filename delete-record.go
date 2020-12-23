package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gorilla/mux"
)

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	subdomain := mux.Vars(r)["subdomain"]
	typ, _ := mux.Vars(r)["type"]
	name, _ := mux.Vars(r)["name"]
	target, _ := mux.Vars(r)["target"]
	log.Debug().Str("subdomain", subdomain).Str("target", target).Msg("deleting record")

	if !checkAuthorization(r, subdomain) {
		http.Error(w, "invalid authorization token. use the preimage of the payment sent to buy '"+subdomain+"'.", 401)
		return
	}

	matching, err := cf.DNSRecords(s.CloudflareZoneId, cloudflare.DNSRecord{
		Name:    name,
		Type:    typ,
		Content: target,
	})
	if err != nil {
		log.Warn().Err(err).Str("subdomain", subdomain).
			Msg("failed to list records on delete")
		http.Error(w, "cloudflare error", 500)
	}

	deleted := 0
	failed := 0
	for _, cfr := range matching {
		if !strings.HasSuffix(cfr.Name, subdomain+"."+cfr.ZoneName) {
			continue
		}

		err := cf.DeleteDNSRecord(s.CloudflareZoneId, cfr.ID)
		if err != nil {
			log.Warn().Err(err).Str("subdomain", subdomain).
				Interface("record", cfr).
				Msg("failed to delete")
			failed++
		} else {
			deleted++
		}
	}

	fmt.Fprintf(w, "records deleted: %d; failures: %d.", deleted, failed)
}
