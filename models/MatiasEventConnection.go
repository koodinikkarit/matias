package models

import (
	"time"
)

type MatiasEventConnection struct {
	ID              uint32
	MatiasKey       string
	HostName        string
	ConnectionStart time.Time
	ConnectionEnd   *time.Time
}
