package processor

// Estructura del mensaje AIS
type AISMessage struct {
	MsgType   int     `json:"msg_type"`
	IMO       int     `json:"imo"`
	MMSI      int     `json:"mmsi"`
	CALLSIGN  string  `json:"callsign"`
	SHIPNAME  string  `json:"shipname"`
	SHIP_TYPE string  `json:"ship_type"`
	STATUS    string  `json:"status"`
	SPEED     float64 `json:"speed"`
	LON       float64 `json:"lon"`
	LAT       float64 `json:"lat"`
}

type AISMessageType5 struct {
	MsgType   int    `json:"msg_type"`
	IMO       int    `json:"imo"`
	MMSI      int    `json:"mmsi"`
	CALLSIGN  string `json:"callsign"`
	SHIPNAME  string `json:"shipname"`
	SHIP_TYPE string `json:"ship_type"`
}
