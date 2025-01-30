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
	MsgType        int     `json:"msg_type"`
	IMO            int     `gorm:"primaryKey" json:"imo"`
	MMSI           int     `gorm:"uniqueIndex" json:"mmsi"`
	Callsign       string  `json:"callsign"`
	Shipname       string  `json:"shipname"`
	ShipType       string  `json:"ship_type"`
	BuiltYear      *string `gorm:"column:built_year" json:"Built"`
	Shipyard       *string `gorm:"column:shipyard" json:"shipyard"`
	HullNumber     *string `gorm:"column:hull_number" json:"Hull-No."`
	KeelLaying     *string `gorm:"column:keel_laying" json:"Keel Laying"`
	LaunchDate     *string `gorm:"column:launch_date" json:"Launch"`
	DeliveryDate   *string `gorm:"column:delivery_date" json:"Delivery"`
	GT             *string `gorm:"column:gt" json:"gt"`
	NT             *string `gorm:"column:nt" json:"nt"`
	CarryingCapTDW *string `gorm:"column:carrying_capacity_tdw" json:"Carrying capacity (tdw)"`
	LengthOverall  *string `gorm:"column:length_overall" json:"Length overall (m)"`
	Breadth        *string `gorm:"column:breadth" json:"Breadth (m)"`
	Depth          *string `gorm:"column:depth" json:"Depth (m)"`
	Propulsion     *string `gorm:"column:propulsion" json:"propulsion"`
	Power          *string `gorm:"column:power" json:"power"`
	Screws         *string `gorm:"column:screws" json:"screws"`
	Speed          *string `gorm:"column:speed" json:"speed"`
}
