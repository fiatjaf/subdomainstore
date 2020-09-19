package main

import "time"

type SubdomainEntry struct {
	Hash    string `json:"hash"`
	Started int64  `json:"started"`
	Hours   int64  `json:"hours"`
}

func (se SubdomainEntry) RemainingHours() int64 {
	started := time.Unix(se.Started, 0)
	ending := started.Add(time.Hour * time.Duration(se.Hours))
	now := time.Now().UTC()
	return int64(ending.Sub(now).Hours())
}

func (se SubdomainEntry) HasEnded() bool {
	return se.RemainingHours() < 0
}

type DNSRecord struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Target   string `json:"target"`
	Priority int    `json:"priority"`
	Proxied  bool   `json:"proxied"`
}
