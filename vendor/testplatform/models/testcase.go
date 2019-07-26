package models

type Testcase struct {
	Model

	Module string `json:"module"`
	Describe string `json:"describe"`
	Step string `json:"step"`
	Result string `json:"result"`
}

func GetTestcases(pageSize int, pageCurrent int, sortOrder string, filterModule []string) (testcases []Testcase) {
	var order = "ID"
	if sortOrder == "descend" {
		order = "ID desc"
	}

	if (len(filterModule)) > 0 {
		db.Where("module in (?)", filterModule).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
	} else {
		db.Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
	}
	return
}

func GetTestcaseTotal(filterModule []string) (count int) {
	if (len(filterModule)) > 0 {
		db.Model(&Testcase{}).Where("module in (?)", filterModule).Count(&count)
	} else {
		db.Model(&Testcase{}).Count(&count)
	}

	return
}