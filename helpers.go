package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/imroc/req"
)

func canBuyMaxHours(current int64) (max int64) {
	now := time.Now().UTC()
	max = int64(s.MaxDurationEnding.Sub(now).Hours())

	if s.MaxHours < max {
		max = s.MaxHours
	}

	return
}

func checkAuthorization(r *http.Request, subdomain string) bool {
	resp, _ := req.Get(s.MicroDBURL + "/" + subdomain)
	var value SubdomainEntry
	resp.ToJSON(&value)

	spl := strings.Split(r.Header.Get("Authorization"), " ")
	preimage := spl[len(spl)-1]
	bpreimage, _ := hex.DecodeString(preimage)
	bhash := sha256.Sum256(bpreimage)
	if hex.EncodeToString(bhash[:]) != value.Hash {
		log.Warn().Str("subdomain", subdomain).Msg("wrong auth token")
		return false
	}

	return true
}
