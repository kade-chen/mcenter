package domain

import "github.com/kade-chen/library/tools/format"

func (d *Domain) ToJson() string {
	return format.ToJSON(d)
}
