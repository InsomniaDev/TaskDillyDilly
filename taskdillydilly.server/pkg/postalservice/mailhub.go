package postalservice

// Hub controls mailbag groups
type MailHub struct {
	name string
	// registered mailbags
	mailbags map[string]*Mailbag

	// Inbound messages from the mailbags
	broadcast chan *Envelope

	// register requests from the mailbags
	register chan *Mailbag

	// Unregister requests from the mailbags
	unregister chan *Mailbag
}

// NewHub creates a new hub
func NewHub(name string) *MailHub {
	return &MailHub{
		name:       name,
		broadcast:  make(chan *Envelope),
		register:   make(chan *Mailbag),
		unregister: make(chan *Mailbag),
		mailbags:   make(map[string]*Mailbag),
	}
}

// Run runs the existing hub
func (m *MailHub) Run() {
	for {
		select {
		case mailbag := <-m.register:
			m.mailbags[mailbag.name] = mailbag
		case mailbag := <-m.unregister:
			if _, ok := m.mailbags[mailbag.name]; ok {
				delete(m.mailbags, mailbag.name)
				close(mailbag.delivered)
			}
			// case envelope := <-m.broadcast:
			// 	for mailbag := range m.mailbags {
			// 		select {
			// 		// case mailbag.delivered <- envelope:
			// 		default:
			// 			// close(mailbag.delivered)
			// 			delete(m.mailbags, mailbag)
			// 		}
			// 	}
		}
	}
}
