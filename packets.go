package main

type TimeInitPacket struct {
	Time int64 `json:"time"`
}

type ServerSelectionPacket struct {
	WsServer string `json:"wsServer"`
}

type Region int
type Polarity int

const (
	Unknown Region = 0
	Europe
	Oceania
	NorthAmerica
	Asia
	SouthAmerica
)

const (
	Negative Polarity = 0
	Positive
)

type Location struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
	Altitude  float32 `json:"alt"`
}

type Signal struct {
	Location
	StationId uint   `json:"sta"`
	Time      uint64 `json:"time"`
	Status    uint   `json:"status"`
}

type StrikePacket struct {
	Location
	Time           uint64   `json:"time"`
	Polarity       Polarity `json:"pol"`
	Deviation      uint     `json:"mds"`
	MaxCircularGap uint     `json:"mcg"`
	Status         uint     `json:"status"`
	Region         Region   `json:"region"`
	Delay          float32  `json:"delay"`
	Signals        []Signal `json:"sig"`
}
