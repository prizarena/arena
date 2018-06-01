package arena

import (
		"strings"
)

type CommaSeparatedUniqueValuesList string

func (csv CommaSeparatedUniqueValuesList) Add(v string) CommaSeparatedUniqueValuesList {
	if strings.Contains(v, ",") {
		panic("contains comma")
	}
	if csv == "" {
		return CommaSeparatedUniqueValuesList(v)
	}
	vals := strings.Split(string(csv), ",")
	for _, val := range vals {
		if val == v {
			return csv
		}
	}
	return CommaSeparatedUniqueValuesList(strings.Join(append(vals, v), ","))
}

func (csv CommaSeparatedUniqueValuesList) Contains(v string) bool {
	s := string(csv)
	vLen := len(v)
	return (strings.HasPrefix(s, v) && (len(s) == vLen || s[vLen:vLen+1] == ",")) ||
		strings.HasSuffix(s, ","+v) || strings.Contains(s, ","+v+",")
}

func (csv CommaSeparatedUniqueValuesList) Remove(v string) CommaSeparatedUniqueValuesList {
	return CommaSeparatedUniqueValuesList(removeValFromCSV(string(csv), v))
}

func (csv CommaSeparatedUniqueValuesList) Strings() []string {
	return strings.Split(string(csv), ",")
}

func (csv CommaSeparatedUniqueValuesList) String() string {
	return string(csv)
}
