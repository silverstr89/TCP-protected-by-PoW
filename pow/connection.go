package pow

import "time"

type Connection struct {
	ClientAddress string
	ServerAddress string
	Timestamp     time.Time
}

func NewConnection(client string, server string) *Connection {
	time.Now()
	return &Connection{client, server, timeNow()}
}

//round up to minute
func timeNow() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}
