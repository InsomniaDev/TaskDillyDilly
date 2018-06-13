package postalservice

// Mailbag is the queue for the event messages
type Mailbag struct {
	name      string
	envelopes []Envelope
	items     int
	packed    chan *Envelope
	delivered chan *Envelope
}

// New creates a new mailbag
func New(name string) *Mailbag {
	return &Mailbag{
		name:      name,
		items:     0,
		packed:    make(chan *Envelope),
		delivered: make(chan *Envelope),
	}
}

func (m *Mailbag) beginDelivery(mailbagName string) {
	// for {
	// 	select {
	// 	case envelope := <-m.packed:
	// 		// envelope.sendOutForDelivery()
	// 	case envelope := <-m.delivered:
	// 		// envelope.deliverMessage()
	// 	}
	// }
}

func (h *MailHub) deliverMessage(e Envelope) {
	// send the message out to the hub for delivery to recipient
	// return e
}

func (e *Envelope) sendOutForDelivery(m *MailHub) {
	// check to see if the delivery has begun for the mailbag using its name
	// if m.mailbags[]
	// if it has not started delivery go ahead and start it

	//  else add the message to the mailbag for delivery

}
