package main

import (
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/cloudflare/cloudflare-go"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	lightning "github.com/fiatjaf/lightningd-gjson-rpc"
	"github.com/gorilla/mux"
	"github.com/imroc/req"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

var err error
var s Settings
var cf *cloudflare.API
var log = zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr})
var spark *lightning.Client
var httpPublic = &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: ""}
var router = mux.NewRouter()

type Settings struct {
	Host               string `envconfig:"HOST" default:"0.0.0.0"`
	Port               string `envconfig:"PORT" required:"true"`
	Domain             string `envconfig:"DOMAIN" required:"true"`
	MicroDBURL         string `envconfig:"MICRODB_URL" required:"true"`
	CloudflareAPIToken string `envconfig:"CLOUDFLARE_API_TOKEN" required:"true"`
	CloudflareZoneId   string `envconfig:"CLOUDFLARE_ZONE_ID" required:"true"`
	SparkURL           string `envconfig:"SPARK_URL" required:"true"`
	SparkToken         string `envconfig:"SPARK_TOKEN" required:"true"`

	HourlyPriceSat    float64   `envconfig:"HOURLY_PRICE_SAT" default:"1"`
	MaxHours          int64     `envconfig:"MAX_HOURS" default:"8760"`
	MaxDurationEnding time.Time `envconfig:"MAX_DURATION_ENDING" default:"2026-01-02T15:04:05Z"`
}

func main() {
	err = envconfig.Process("", &s)
	if err != nil {
		log.Fatal().Err(err).Msg("couldn't process envconfig.")
	}

	// check microdb bucket status
	r, _ := req.Get(regexp.MustCompile(`(\w)/(\w)`).
		ReplaceAllString(s.MicroDBURL, "$1/status/$2"))
	var resp struct {
		Bucket string `json:"bucket"`
		Funds  int64  `json:"funds"`
		Size   int64  `json:"size"`
	}
	err = r.ToJSON(&resp)
	if err != nil {
		log.Panic().Err(err).Str("url", s.MicroDBURL).Msg("failed to parse microdb")
	}
	if resp.Funds <= 0 {
		log.Panic().Err(err).Str("url", s.MicroDBURL).Msg("microdb funds are low")
	}

	// check cloudflare api key
	cf, err = cloudflare.NewWithAPIToken(s.CloudflareAPIToken)
	if resp.Funds <= 0 {
		log.Panic().Err(err).Msg("couldn't connect to cloudflare api")
	}

	// spark caller and listener
	spark = &lightning.Client{
		SparkURL:    s.SparkURL,
		SparkToken:  s.SparkToken,
		CallTimeout: time.Second * 3,
	}
	go sparkoListener()

	// routes
	router.PathPrefix("/static/").Methods("GET").Handler(http.FileServer(httpPublic))
	router.Path("/favicon.ico").Methods("GET").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/png")
			iconf, _ := httpPublic.Open("static/icon.png")
			fstat, _ := iconf.Stat()
			http.ServeContent(w, r, "static/icon.png", fstat.ModTime(), iconf)
			return
		})

	router.Path("/buy/{subdomain}").Methods("GET").HandlerFunc(buy)
	router.Path("/{subdomain}").Methods("GET").HandlerFunc(recharge)

	router.Path("/{subdomain}/records").Methods("GET").
		HandlerFunc(listRecords)
	router.Path("/{subdomain}/{type}/{name}/{target}").Methods("POST").
		HandlerFunc(setRecord)

	router.Path("/{subdomain}").Methods("DELETE").
		HandlerFunc(deleteRecord)
	router.Path("/{subdomain}/{type}").Methods("DELETE").
		HandlerFunc(deleteRecord)
	router.Path("/{subdomain}/{type}/{name}").Methods("DELETE").
		HandlerFunc(deleteRecord)

	router.PathPrefix("/").Methods("GET").HandlerFunc(serveClient)

	// start http server
	log.Info().Str("host", s.Host).Str("port", s.Port).Msg("listening")
	srv := &http.Server{
		Handler:      router,
		Addr:         s.Host + ":" + s.Port,
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  300 * time.Second,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Error().Err(err).Msg("error serving http")
	}
}

func serveClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	indexf, err := httpPublic.Open("static/index.html")
	if err != nil {
		log.Error().Err(err).Str("file", "static/index.html").
			Msg("make sure you generated bindata.go without -debug")
		return
	}
	fstat, _ := indexf.Stat()
	http.ServeContent(w, r, "static/index.html", fstat.ModTime(), indexf)
	return
}
