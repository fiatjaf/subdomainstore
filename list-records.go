package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gorilla/mux"
)

func listRecords(w http.ResponseWriter, r *http.Request) {
	subdomain := mux.Vars(r)["subdomain"]
	log.Debug().Str("subdomain", subdomain).Msg("listing records")

	if !checkAuthorization(r, subdomain) {
		http.Error(w, "invalid authorization token. use the preimage of the payment sent to buy '"+subdomain+"'.", 401)
		return
	}

	cfrecords, err := cf.DNSRecords(s.CloudflareZoneId, cloudflare.DNSRecord{})
	if err != nil {
		log.Warn().Err(err).Str("subdomain", subdomain).Msg("failed to list records")
		http.Error(w, "cloudflare error", 500)
	}

	records := make([]DNSRecord, 0, len(cfrecords))
	for _, cfr := range cfrecords {
		if !strings.HasSuffix(cfr.Name, "."+subdomain+"."+s.Domain) {
			continue
		}

		records = append(records, DNSRecord{
			cfr.Name,
			cfr.Type,
			cfr.Content,
			cfr.Priority,
			cfr.Proxied,
		})
	}

	json.NewEncoder(w).Encode(records)
}
