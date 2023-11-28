package oop

import (
	"math/rand"
	"strconv"
	"fmt"
	"os/exec"
)

type (
	Mention struct {
		S string `json:"S"`
		E string `json:"E"`
		M string `json:"M"`
	}
	Mentions struct {
		MENTIONEES []Mention `json:"MENTIONEES"`
	}
)

func Contains(ArrList []string, rstr string) bool {
	for _, x := range ArrList {
		if x == rstr {
			return true
		}
	}
	return false
}

func Uncontains(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return false
		}
	}
	return true
}


func ContainsInt(ArrList []string, rstr string) int {
	for z, x := range ArrList {
		if x == rstr {
			return z
		}
	}
	return 404
}

func Clearcache() {
	fmt.Println("CACHE_CLEARED")
	exec.Command("sync;", "echo", "1", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("sync;", "echo", "2", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("sync;", "echo", "3", ">", "/proc/sys/vm/drop_caches").Run()
}

func Remove(arr []string, str string) []string {
	var newArr = []string{}
	for k, v := range arr {
		if v == str {
			newArr = append(arr[:k], arr[k+1:]...)
		}
	}
	return newArr
}

func RemoveInt(arr []int, num int) []int {
	var newArr = []int{}
	for k, v := range arr {
		if v == num {
			newArr = append(arr[:k], arr[k+1:]...)
		}
	}
	return newArr
}

func CheckEqual(list1 []string, list2 []string) bool {
	for _, v := range list1 {
		if Contains(list2, v) {
			return true
		}
	}
	return false
}

func Randint(min int, max int) int {
	return rand.Intn(max-min) + min
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
