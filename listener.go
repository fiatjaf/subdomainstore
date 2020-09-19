package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/donovanhide/eventsource"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

func sparkoListener() {
	r, _ := http.NewRequest("GET", s.SparkURL+"/stream", nil)
	r.Header.Set("X-Access", s.SparkToken)
	stream, err := eventsource.SubscribeWithRequest("", r)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to sparko.")
	}

	defer stream.Close()

	go func() {
		for err := range stream.Errors {
			log.Debug().Err(err).Msg("spark stream error")
		}
	}()

	for ev := range stream.Events {
		switch ev.Event() {
		case "invoice_payment":
			data := gjson.Parse(ev.Data())
			label := data.Get("invoice_payment.label").String()
			if !strings.HasPrefix(label, "subdomainstore/") {
				continue
			}
			preimage := data.Get("invoice_payment.preimage").String()
			msat := data.Get("invoice_payment.msat").String()
			log.Debug().Str("label", label).Str("r", preimage).Str("msat", msat).
				Msg("got payment")

			amount, err := strconv.ParseInt(strings.TrimSuffix(msat, "msat"), 10, 64)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse msat on invoice_payment")
				return
			}

			p := strings.Split(label, "/")
			if len(p) != 4 {
				log.Error().Err(err).Msg("failed to parse subdomain on invoice_payment")
				return
			}
			op := p[1]
			subdomain := p[2]

			resp, _ := req.Get(s.MicroDBURL + "/" + subdomain)
			var value SubdomainEntry
			resp.ToJSON(&value)

			hoursBought := int64(float64(amount) / (s.HourlyPriceSat * 1000))

			switch op {
			case "buy":
				log.Info().Str("subdomain", subdomain).Str("preimage", preimage).
					Msg("subdomain bought")

				bpreimage, _ := hex.DecodeString(preimage)
				bhash := sha256.Sum256(bpreimage)
				hash := hex.EncodeToString(bhash[:])

				value := SubdomainEntry{
					Hash:    hash,
					Started: time.Now().UTC().Unix(),
					Hours:   hoursBought,
				}
				resp, _ := req.Put(s.MicroDBURL+"/"+subdomain, req.BodyJSON(value))
				if resp.Response().StatusCode >= 300 {
					log.Warn().Str("subdomain", subdomain).Str("resp", resp.String()).
						Interface("value", value).
						Msg("failed to mark subdomain as bought")
				}
			case "recharge":
				log.Info().Str("subdomain", subdomain).Msg("subdomain recharged")

				value.Hours += hoursBought
				resp, _ := req.Put(s.MicroDBURL+"/"+subdomain, req.BodyJSON(value))
				if resp.Response().StatusCode >= 300 {
					log.Warn().Str("subdomain", subdomain).Str("resp", resp.String()).
						Interface("value", value).
						Msg("failed to mark subdomain recharge")
				}
			}
		}
	}
}
