package stt

// https://cloud.google.com/speech-to-text/v2/docs/reference/rest#service-endpoint
type ENDPOINT int8

const (
	ENDPOINTS_US_CENTRAL1_SPEECH_GOOGLEAPIS_COM ENDPOINT = iota
	ENDPOINTS_US_EAST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_US_WEST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_EUROPE_WEST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_EUROPE_WEST2_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_EUROPE_WEST3_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_EUROPE_WEST4_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_ASIA_SOUTH1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_ASIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_ASIA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_AUSTRALIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_NORTHAMERICA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_US_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_EU_SPEECH_GOOGLEAPIS_COM
	ENDPOINTS_SPEECH_GOOGLEAPIS_COM
)

var ENDPOINTS_STRING = map[ENDPOINT]string{
	ENDPOINTS_US_CENTRAL1_SPEECH_GOOGLEAPIS_COM:             "us-central1-speech.googleapis.com",
	ENDPOINTS_US_EAST1_SPEECH_GOOGLEAPIS_COM:                "us-east1-speech.googleapis.com",
	ENDPOINTS_US_WEST1_SPEECH_GOOGLEAPIS_COM:                "us-west1-speech.googleapis.com",
	ENDPOINTS_EUROPE_WEST1_SPEECH_GOOGLEAPIS_COM:            "europe-west1-speech.googleapis.com",
	ENDPOINTS_EUROPE_WEST2_SPEECH_GOOGLEAPIS_COM:            "europe-west2-speech.googleapis.com",
	ENDPOINTS_EUROPE_WEST3_SPEECH_GOOGLEAPIS_COM:            "europe-west3-speech.googleapis.com",
	ENDPOINTS_EUROPE_WEST4_SPEECH_GOOGLEAPIS_COM:            "europe-west4-speech.googleapis.com",
	ENDPOINTS_ASIA_SOUTH1_SPEECH_GOOGLEAPIS_COM:             "asia-south1-speech.googleapis.com",
	ENDPOINTS_ASIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM:         "asia-southeast1-speech.googleapis.com",
	ENDPOINTS_ASIA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM:         "asia-northeast1-speech.googleapis.com",
	ENDPOINTS_AUSTRALIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM:    "australia-southeast1-speech.googleapis.com",
	ENDPOINTS_NORTHAMERICA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM: "northamerica-northeast1-speech.googleapis.com",
	ENDPOINTS_US_SPEECH_GOOGLEAPIS_COM:                      "us-speech.googleapis.com",
	ENDPOINTS_EU_SPEECH_GOOGLEAPIS_COM:                      "eu-speech.googleapis.com",
	ENDPOINTS_SPEECH_GOOGLEAPIS_COM:                         "speech.googleapis.com",
}

var ENDPOINTS_ENDPOINT = map[string]ENDPOINT{
	"us-central1-speech.googleapis.com":             ENDPOINTS_US_CENTRAL1_SPEECH_GOOGLEAPIS_COM,
	"us-east1-speech.googleapis.com":                ENDPOINTS_US_EAST1_SPEECH_GOOGLEAPIS_COM,
	"us-west1-speech.googleapis.com":                ENDPOINTS_US_WEST1_SPEECH_GOOGLEAPIS_COM,
	"europe-west1-speech.googleapis.com":            ENDPOINTS_EUROPE_WEST1_SPEECH_GOOGLEAPIS_COM,
	"europe-west2-speech.googleapis.com":            ENDPOINTS_EUROPE_WEST2_SPEECH_GOOGLEAPIS_COM,
	"europe-west3-speech.googleapis.com":            ENDPOINTS_EUROPE_WEST3_SPEECH_GOOGLEAPIS_COM,
	"europe-west4-speech.googleapis.com":            ENDPOINTS_EUROPE_WEST4_SPEECH_GOOGLEAPIS_COM,
	"asia-south1-speech.googleapis.com":             ENDPOINTS_ASIA_SOUTH1_SPEECH_GOOGLEAPIS_COM,
	"asia-southeast1-speech.googleapis.com":         ENDPOINTS_ASIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM,
	"asia-northeast1-speech.googleapis.com":         ENDPOINTS_ASIA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM,
	"australia-southeast1-speech.googleapis.com":    ENDPOINTS_AUSTRALIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM,
	"northamerica-northeast1-speech.googleapis.com": ENDPOINTS_NORTHAMERICA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM,
	"us-speech.googleapis.com":                      ENDPOINTS_US_SPEECH_GOOGLEAPIS_COM,
	"eu-speech.googleapis.com":                      ENDPOINTS_EU_SPEECH_GOOGLEAPIS_COM,
	"speech.googleapis.com":                         ENDPOINTS_SPEECH_GOOGLEAPIS_COM,
}

var ENDPOINTS_PARENT = map[ENDPOINT]string{
	ENDPOINTS_US_CENTRAL1_SPEECH_GOOGLEAPIS_COM:             "us-central1",
	ENDPOINTS_US_EAST1_SPEECH_GOOGLEAPIS_COM:                "us-east1",
	ENDPOINTS_US_WEST1_SPEECH_GOOGLEAPIS_COM:                "us-west1",
	ENDPOINTS_EUROPE_WEST1_SPEECH_GOOGLEAPIS_COM:            "europe-west1",
	ENDPOINTS_EUROPE_WEST2_SPEECH_GOOGLEAPIS_COM:            "europe-west2",
	ENDPOINTS_EUROPE_WEST3_SPEECH_GOOGLEAPIS_COM:            "europe-west3",
	ENDPOINTS_EUROPE_WEST4_SPEECH_GOOGLEAPIS_COM:            "europe-west4",
	ENDPOINTS_ASIA_SOUTH1_SPEECH_GOOGLEAPIS_COM:             "asia-south1",
	ENDPOINTS_ASIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM:         "asia-southeast1",
	ENDPOINTS_ASIA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM:         "asia-northeast1",
	ENDPOINTS_AUSTRALIA_SOUTHEAST1_SPEECH_GOOGLEAPIS_COM:    "australia-southeast1",
	ENDPOINTS_NORTHAMERICA_NORTHEAST1_SPEECH_GOOGLEAPIS_COM: "northamerica-northeast1",
	ENDPOINTS_US_SPEECH_GOOGLEAPIS_COM:                      "us",
	ENDPOINTS_EU_SPEECH_GOOGLEAPIS_COM:                      "eu",
	ENDPOINTS_SPEECH_GOOGLEAPIS_COM:                         "global",
}
