package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fiatjaf/go-lnurl"
	"github.com/gorilla/mux"
	"github.com/imroc/req"
	"github.com/lucsky/cuid"
)

func buy(w http.ResponseWriter, r *http.Request) {
	subdomain := mux.Vars(r)["subdomain"]

	if strings.Index(subdomain, ".") != -1 {
		json.NewEncoder(w).Encode(lnurl.ErrorResponse("name can't contain a dot"))
		return
	}

	amount, _ := strconv.ParseInt(r.URL.Query().Get("amount"), 10, 64)
	log.Debug().Str("subdomain", subdomain).Int64("msat", amount).
		Msg("buying subdomain")

	resp, _ := req.Get(s.MicroDBURL + "/" + subdomain)
	var value struct {
		Hash  string `json:"hash"`
		Start int64  `json:"started"`
		Hours int64  `json:"hours"`
	}
	resp.ToJSON(&value)

	if value.Hours > 0 {
		json.NewEncoder(w).Encode(
			lnurl.ErrorResponse(subdomain + " already has an owner."))
		return
	}

	canBuy := canBuyMaxHours(0)

	jmetadata, _ := json.Marshal([][]string{
		[]string{"text/plain",
			fmt.Sprintf("Buy %s.%s for some time.", subdomain, s.Domain)},
	})

	if amount == 0 {
		// return params
		json.NewEncoder(w).Encode(lnurl.LNURLPayResponse1{
			Tag:             "payRequest",
			Callback:        "https://" + s.Domain + "/buy/" + subdomain,
			MinSendable:     int64(s.HourlyPriceSat * 1000),
			MaxSendable:     int64(s.HourlyPriceSat * 1000 * float64(canBuy)),
			EncodedMetadata: string(jmetadata),
		})
	} else {
		// return invoice
		h := sha256.Sum256(jmetadata)
		label := "subdomainstore/buy/" + subdomain + "/" + cuid.Slug()

		res, err := spark.Call("invoicewithdescriptionhash",
			amount,
			label,
			hex.EncodeToString(h[:]),
		)

		if err != nil {
			log.Error().Err(err).Str("label", label).Int64("amount", amount).
				Msg("failed to create invoice")
			json.NewEncoder(w).Encode(lnurl.ErrorResponse("failed to create invoice!"))
			return
		}

		preimage := res.Get("preimage").String()
		bpreimage, _ := hex.DecodeString(preimage)
		aesAction, err := lnurl.AESAction("The key for editing DNS records on this subdomain is the preimage of this payment.", bpreimage, preimage)
		if err != nil {
			log.Error().Err(err).Str("label", label).Int64("amount", amount).
				Str("preimage", preimage).
				Msg("failed to encode aes action")
			json.NewEncoder(w).Encode(lnurl.ErrorResponse("failed to create invoice!"))
			return
		}

		json.NewEncoder(w).Encode(lnurl.LNURLPayResponse2{
			PR:            res.Get("bolt11").String(),
			Routes:        make([][]lnurl.RouteInfo, 0),
			Disposable:    lnurl.TRUE,
			SuccessAction: aesAction,
		})
	}
}
