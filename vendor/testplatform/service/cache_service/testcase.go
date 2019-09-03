package cache_service

import(
	"strconv"
	"strings"

	"testplatform/pkg/e"
)

type Testcase struct {
	ID         int
	Item       string
	Module     string
	Importance string
	Describe   string
	Step       string
	Result     string

	PageNum  int
	PageSize int
}

func (t *Testcase) GetTestcasesKey() string {
	keys := []string{
		e.CACHE_TESTCASE,
		"LIST",
	}

	if t.ID > 0 {
		keys = append(keys, strconv.Itoa(t.ID))
	}
	if t.Item != "" {
		keys = append(keys, t.Item)
	}
	if t.Module != "" {
		keys = append(keys, t.Module)
	}
	if t.Importance != "" {
		keys = append(keys, t.Importance)
	}
	if t.Describe != "" {
		keys = append(keys, t.Describe)
	}
	if t.Step != "" {
		keys = append(keys, t.Step)
	}
	if t.Result != "" {
		keys = append(keys, t.Result)
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}