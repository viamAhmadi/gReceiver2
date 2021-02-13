package conn

type Message struct {
	Destination string
	Id          int    // 5 bytes
	ConnId      string // 20 char
	Content     string // 8 kilo
}
