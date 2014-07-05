package libgomessage

// yes, this file is togo!
// ba-dum-tss

import(

)

type To struct {
	To []string
}

func NewTo() *To {
	return &To{ }
}

func (to *To) GetAttachementName() string {
	return "To"
}

func (to *To) GetTos() []string {
	return to.To
}

func (to *To) AddAddress(addresses ...string) error {
	for _, address := range addresses {
		to.To = append(to.To, address)
	}
	return nil
}

func (to *To) IsValid(address string) bool {
	return true
}
