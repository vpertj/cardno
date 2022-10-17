package test

import (
	"fmt"
	"github.com/waketj/cardno"
	"testing"
)

func TestValidate18IdCardNo_1(t *testing.T) {
	idNo := "110101199003074119"
	if !cardno.Validate18CardNo(idNo) {
		t.Fail()
	}
}

func TestValidate18IdCardNo_2(t *testing.T) {
	idNo := "11010119900307411X"
	if !cardno.Validate18CardNo(idNo) {
		t.Fail()
	}
}

func TestParse18IdCardNoInfo(t *testing.T) {
	idNo := "110101199003074119"
	result, info := cardno.Parse18CardNoInfo(idNo)
	if !result {
		t.Fail()
	} else {
		fmt.Println(fmt.Sprintf("%v", *info))
	}
}

func TestAutoCreate18IdCardNo(t *testing.T) {
	idNo := cardno.AutoCreate18CardNo()
	fmt.Println(idNo)
	if !cardno.Validate18CardNo(idNo) {
		t.Fail()
	}
	result, info := cardno.Parse18CardNoInfo(idNo)
	if !result {
		t.Fail()
	} else {
		fmt.Println(fmt.Sprintf("%v", *info))
	}
}
