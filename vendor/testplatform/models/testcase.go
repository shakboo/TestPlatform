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

func GetTestcases(item string, pageSize int, pageCurrent int, sortOrder string, filterImportance []string) (testcases []Testcase) {
	var order = "ID"
	if sortOrder == "descend" {
		order = "ID desc"
	}

	if (len(filterImportance) > 0) {
		db.Where("item = ?", item).Where("importance in (?)", filterImportance).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
	} else {
		db.Where("item = ?", item).Offset((pageCurrent-1) * 10).Limit(pageSize).Order(order).Find(&testcases)
	}
	return
}

func GetTestcasesTotal(item string, filterImportance []string) (count int) {
	if (len(filterImportance) > 0) {
		db.Model(&Testcase{}).Where("item = ?", item).Where("importance in (?)", filterImportance).Count(&count)
	} else {
		db.Model(&Testcase{}).Where("item = ?", item).Count(&count)
	}

	return
}

func PutTestcases(item string, id int, module string, importance string, describe string, step string, result string) {
	db.Model(&Testcase{}).Where("item = ?", item).Where("id = ?", id).Updates(map[string]interface{}{"module": module, "importance": importance, "describe": describe, "step": step, "result": result})
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