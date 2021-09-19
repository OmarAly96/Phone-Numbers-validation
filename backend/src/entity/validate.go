package entity

import (
	"regexp"
)

func (p *PhoneNumber) ValidateState(exp string) {
	if match, _ := regexp.MatchString(exp, p.Number); match {
		p.State = true
	}
}
