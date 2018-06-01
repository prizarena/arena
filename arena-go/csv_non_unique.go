package arena

import (
	"bytes"
	"strings"
)

type CommaSeparatedValuesList string

func (csv CommaSeparatedValuesList) Add(v string) CommaSeparatedValuesList {
	if strings.Contains(v, ",") {
		panic("contains comma")
	}
	if csv == "" {
		return CommaSeparatedValuesList(v)
	}
	return CommaSeparatedValuesList(string(csv) +"," + v)
}

func (csv CommaSeparatedValuesList) Count() int {
	if csv == "" {
		return 0
	}
	return strings.Count(string(csv), ",") + 1
}

func (csv CommaSeparatedValuesList) Set(i int, v string) CommaSeparatedValuesList {
	vals := strings.Split(string(csv), ",")
	if i >= len(vals) {
		panic("out of range")
	}
	vals[i] = v
	return CommaSeparatedValuesList(strings.Join(vals, ","))
}

func (csv CommaSeparatedValuesList) Contains(v string) bool {
	s := string(csv)
	vLen := len(v)
	return (strings.HasPrefix(s, v) && (len(s) == vLen || s[vLen:vLen+1] == ",")) ||
		strings.HasSuffix(s, ","+v) || strings.Contains(s, ","+v+",")
}

func (csv CommaSeparatedValuesList) Remove(v string) CommaSeparatedValuesList {
	return CommaSeparatedValuesList(removeValFromCSV(string(csv), v))
}

func (csv CommaSeparatedValuesList) Strings() []string {
	return strings.Split(string(csv), ",")
}

func (csv CommaSeparatedValuesList) String() string {
	return string(csv)
}

func removeValFromCSV(s, v string) string {
	if len(s) == 0 {
		return s
	}
	var buf bytes.Buffer
	vals := strings.Split(s, ",")
	for _, val := range vals {
		if val != v {
			buf.WriteString(val)
			buf.WriteString(",")
		}
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	return string(buf.Bytes())
}