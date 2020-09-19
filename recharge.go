package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/andanhm/go-prettytime"
	"github.com/fiatjaf/go-lnurl"
	"github.com/gorilla/mux"
	"github.com/imroc/req"
	"github.com/lucsky/cuid"
)

func recharge(w http.ResponseWriter, r *http.Request) {
	subdomain := mux.Vars(r)["subdomain"]

	amount, _ := strconv.ParseInt(r.URL.Query().Get("amount"), 10, 64)
	log.Debug().Str("subdomain", subdomain).Int64("msat", amount).
		Msg("buying subdomain")

	resp, _ := req.Get(s.MicroDBURL + "/" + subdomain)
	var value SubdomainEntry
	resp.ToJSON(&value)

	remaining := value.RemainingHours()
	endTime := time.Now().UTC().Add(time.Hour * time.Duration(remaining))
	canBuy := canBuyMaxHours(remaining)

	jmetadata, _ := json.Marshal([][]string{
		[]string{"text/plain",
			fmt.Sprintf("'%s' has %d hours left (ending %s).\nRecharging for %g sat/h.",
				subdomain, remaining, prettytime.Format(endTime), s.HourlyPriceSat)},
	})

	if amount == 0 {
		// return params
		json.NewEncoder(w).Encode(lnurl.LNURLPayResponse1{
			Tag:             "payRequest",
			Callback:        "https://" + s.Domain + "/" + subdomain,
			MinSendable:     int64(s.HourlyPriceSat * 1000),
			MaxSendable:     int64(s.HourlyPriceSat * 1000 * float64(canBuy)),
			EncodedMetadata: string(jmetadata),
		})
	} else {
		// return invoice
		h := sha256.Sum256(jmetadata)
		label := "subdomainstore/recharge/" + subdomain + "/" + cuid.Slug()

		res, err := spark.Call("invoicewithdescriptionhash",
			amount,
			label,
			hex.EncodeToString(h[:]),
		)

		if err != nil {
			log.Error().Err(err).Str("label", label).Int64("amount", amount).
				Msg("failed to create invoice")
			json.NewEncoder(w).Encode(lnurl.ErrorResponse("failed to create invoice"))
			return
		}

		json.NewEncoder(w).Encode(lnurl.LNURLPayResponse2{
			PR:         res.Get("bolt11").String(),
			Routes:     make([][]lnurl.RouteInfo, 0),
			Disposable: lnurl.FALSE,
		})
	}
}
