package postalservice

// Envelope hods the data and the routing information for the message
type Envelope struct {
	// transaction id
	trackingnumber string
	// event name
	postalcode string
	// contents/data of the event
	letter []byte
}
