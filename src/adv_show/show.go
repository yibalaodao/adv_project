package adv_show

import (
	_ "database/sql"
	"fmt"
	_ "github/lib/pg"
)

type AdvInfo struct {
	Id      int
	Url     string
	Image   string
	DitchId int
}
