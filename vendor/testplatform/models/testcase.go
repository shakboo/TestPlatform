package models

import (
	"github.com/jinzhu/gorm"
)

type Testcase struct {
	Model

	Id int `json:"id"`
	Module string `json:"module"`
	Importance string `json:"importance"`
	Describe string `json:"describe"`
	Step string `json:"step"`
	Result string `json:"result"`

	Item string `json:"item"`
}


func GetAllTestcases(item string) (testcases []Testcase) {
	db.Where("item = ?", item).Order("ID").Find(&testcases)

	return
}

func GetTestcases(item string, pageSize int, pageCurrent int, sortOrder string, filterImportance, filterModule []string) (testcases []Testcase) {
	var order = "ID"
	if sortOrder == "descend" {
		order = "ID desc"
	}

	if (len(filterImportance) > 0) {
		if (len(filterModule) > 0) {
			db.Where("item = ?", item).Where("importance in (?)", filterImportance).Where("module in (?)", filterModule).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
		} else {
			db.Where("item = ?", item).Where("importance in (?)", filterImportance).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
		}	
	} else {
		if (len(filterModule) > 0) {
			db.Where("item = ?", item).Where("module in (?)", filterModule).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
		} else {
			db.Where("item = ?", item).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
		}
	}
	return
}

func GetTestcasesTotal(item string, filterImportance, filterModule []string) (count int) {
	if (len(filterImportance) > 0) {
		if (len(filterModule) > 0) {
			db.Model(&Testcase{}).Where("item = ?", item).Where("importance in (?)", filterImportance).Where("module in (?)", filterModule).Count(&count)
		} else {
			db.Model(&Testcase{}).Where("item = ?", item).Where("importance in (?)", filterImportance).Count(&count)
		}
	} else {
		if (len(filterModule) > 0) {
			db.Model(&Testcase{}).Where("item = ?", item).Where("module in (?)", filterModule).Count(&count)
		} else {
			db.Model(&Testcase{}).Where("item = ?", item).Count(&count)
		}	
	}

	return
}

func GetTestcasesModule(item string) (testcases []Testcase) {
	db.Select("DISTINCT(module)").Find(&testcases)
	return
}

func PutTestcases(item string, id int, module string, importance string, describe string, step string, result string) {
	db.Model(&Testcase{}).Where("item = ? AND id = ?", item, id).Assign(Testcase{Item: item, Id: id, Module: module, Importance: importance, Describe: describe, Step: step, Result: result}).FirstOrCreate(&Testcase{})
}

func PostTestcases(item string, preId int, module string, importance string, describe string, step string, result string) {
	db.Model(&Testcase{}).Where("item = ?", item).Where("id > ?", preId).Order("ID desc").UpdateColumn("id", gorm.Expr("id + ?", 1))
	var testcase = Testcase{Id: preId + 1, Module: module, Importance: importance, Describe: describe, Step: step, Result: result, Item: item}
	db.Create(&testcase)
}

func DeleteTestcases(item string, id int) {
	db.Where("item = ?", item).Where("id = ?", id).Delete(&Testcase{})
	db.Where("item = ?", item).Model(&Testcase{}).Where("id > ?", id).Order("ID").UpdateColumn("id", gorm.Expr("id - ?", 1))
}