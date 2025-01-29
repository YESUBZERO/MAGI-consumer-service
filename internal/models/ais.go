package models

// AISMessageType5 representa los datos AIS estáticos
type AISMessageType5 struct {
	MsgType  int    `json:"msg_type"`
	IMO      int    `json:"imo"`
	MMSI     int    `json:"mmsi"`
	Callsign string `json:"callsign"`
	Shipname string `json:"shipname"`
	ShipType string `json:"ship_type"`
}

// Ship representa la estructura almacenada en PostgreSQL
type Ship struct {
	IMO            int      `gorm:"primaryKey" json:"imo"`
	MMSI           int      `gorm:"uniqueIndex" json:"mmsi,omitempty"`
	Callsign       string   `json:"callsign,omitempty"`
	Shipname       string   `json:"shipname,omitempty"`
	ShipType       string   `json:"ship_type,omitempty"`
	BuiltYear      *int     `gorm:"column:built_year" json:"built_year,omitempty"`
	Shipyard       *string  `gorm:"column:shipyard" json:"shipyard,omitempty"`
	HullNumber     *string  `gorm:"column:hull_number" json:"hull_number,omitempty"`
	KeelLaying     *string  `gorm:"column:keel_laying" json:"keel_laying,omitempty"`
	LaunchDate     *string  `gorm:"column:launch_date" json:"launch_date,omitempty"`
	DeliveryDate   *string  `gorm:"column:delivery_date" json:"delivery_date,omitempty"`
	GT             *int     `gorm:"column:gt" json:"gt,omitempty"`
	NT             *int     `gorm:"column:nt" json:"nt,omitempty"`
	CarryingCapTDW *int     `gorm:"column:carrying_capacity_tdw" json:"carrying_capacity_tdw,omitempty"`
	LengthOverall  *float64 `gorm:"column:length_overall" json:"length_overall,omitempty"`
	Breadth        *float64 `gorm:"column:breadth" json:"breadth,omitempty"`
	Depth          *float64 `gorm:"column:depth" json:"depth,omitempty"`
	Propulsion     *string  `gorm:"column:propulsion" json:"propulsion,omitempty"`
	Power          *string  `gorm:"column:power" json:"power,omitempty"`
	Screws         *string  `gorm:"column:screws" json:"screws,omitempty"`
	Speed          *float64 `gorm:"column:speed" json:"speed,omitempty"`
}
