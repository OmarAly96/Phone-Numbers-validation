package pkg

import (
	"fmt"
	"strings"
)

var (
	CodeToCountry = map[string]string{
		"237": "Cameroon",
		"251": "Ethiopia",
		"212": "Morocco",
		"258": "Mozambique",
		"256": "Uganda",
	}
	CodeToExpression = map[string]string{
		"237": "\\(237\\)\\ ?[2368]\\d{7,8}$",
		"251": "\\(251\\)\\ ?[1-59]\\d{8}$",
		"212": "\\(212\\)\\ ?[5-9]\\d{8}$",
		"258": "\\(258\\)\\ ?[28]\\d{7,8}$",
		"256": "\\(256\\)\\ ?\\d{9}$",
	}
)

type CountryExp struct {
	Country string
	Exp     string
}

func SegregateCode(number string) (val string) {
	var (
		idx1 = strings.Index(number, "(")
		idx2 = strings.Index(number, ")")
	)
	if idx1 != -1 && idx2 != -1 {
		val = number[idx1+1 : idx2]
	}
	return
}

func CodeExists(code string) error {
	if _, ok := CodeToCountry[code]; ok {
		return nil
	}
	return fmt.Errorf("country code does not exist")
}

func CodeCountryExpression(code string) CountryExp {
	return CountryExp{
		Country: CodeToCountry[code],
		Exp:     CodeToExpression[code],
	}
}
