package urltype

import "fmt"

type DataURL struct {
	Prefix string
	Data   string
}

func (u *DataURL) Url() string {
	return fmt.Sprintf("%s,%s", u.Prefix, u.Data)
}
