package main

import (
	"regexp"
	"strconv"
	"testing"
)

type convertStep struct {
	reg     *regexp.Regexp
	replace string
}

var (
	convertSteps    []convertStep
	sliceUnit       []string
	sliceUnit2      []rune
	upperDigitUnit  map[string]string
	upperDigitUnit2 map[rune]string
	upperDigitUnit3 map[rune]rune
)

func init() {
	convertSteps = make([]convertStep, 0, 10)

	reg, _ := regexp.Compile(`零角零分$`)
	convertSteps = append(convertSteps, convertStep{reg, "整"})

	reg, _ = regexp.Compile(`零角`)
	convertSteps = append(convertSteps, convertStep{reg, "零"})

	reg, _ = regexp.Compile(`零分$`)
	convertSteps = append(convertSteps, convertStep{reg, "整"})

	reg, _ = regexp.Compile(`零[仟佰拾]`)
	convertSteps = append(convertSteps, convertStep{reg, "零"})

	reg, _ = regexp.Compile(`零{2,}`)
	convertSteps = append(convertSteps, convertStep{reg, "零"})

	reg, _ = regexp.Compile(`零亿`)
	convertSteps = append(convertSteps, convertStep{reg, "亿"})

	reg, _ = regexp.Compile(`零万`)
	convertSteps = append(convertSteps, convertStep{reg, "万"})

	reg, _ = regexp.Compile(`零*元`)
	convertSteps = append(convertSteps, convertStep{reg, "元"})

	reg, _ = regexp.Compile(`亿零{0, 3}万`)
	convertSteps = append(convertSteps, convertStep{reg, "^元"})

	reg, _ = regexp.Compile(`零元`)
	convertSteps = append(convertSteps, convertStep{reg, "零"})

	sliceUnit = []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	sliceUnit2 = []rune{'仟', '佰', '拾', '亿', '仟', '佰', '拾', '万', '仟', '佰', '拾', '元', '角', '分'}
	upperDigitUnit = map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	upperDigitUnit2 = map[rune]string{'0': "零", '1': "壹", '2': "贰", '3': "叁", '4': "肆", '5': "伍", '6': "陆", '7': "柒", '8': "捌", '9': "玖"}
	upperDigitUnit3 = map[rune]rune{'0': '零', '1': '壹', '2': '贰', '3': '叁', '4': '肆', '5': '伍', '6': '陆', '7': '柒', '8': '捌', '9': '玖'}
}

// 最初的版本
//
func ConvertNumToCny1(num float64) (string, error) {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum {
		str = str + upperDigitUnit[string(v)] + s[k]
	}

	reg, err := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零{2,}`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零*元`)
	str = reg.ReplaceAllString(str, "元")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^元")
	if err != nil {
		return "", err
	}

	reg, err = regexp.Compile(`零元`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		return "", err
	}

	return str, nil
}

// 把正则编译提取到外部，避免反复初始化
//
func ConvertNumToCny2(num float64) (string, error) {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum {
		str = str + upperDigitUnit[string(v)] + s[k]
	}

	for i := 0; i < len(convertSteps); i++ {
		str = convertSteps[i].reg.ReplaceAllString(str, convertSteps[i].replace)
	}

	return str, nil
}

// 把变量初始化提取到外部，避免不必要的重复创建
//
func ConvertNumToCny3(num float64) (string, error) {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	str := ""
	for k, v := range strnum {
		str = str + upperDigitUnit[string(v)] + s[k]
	}

	for i := 0; i < len(convertSteps); i++ {
		str = convertSteps[i].reg.ReplaceAllString(str, convertSteps[i].replace)
	}

	return str, nil
}

// 把map的key换成rune，避免反复生成字符串
//
func ConvertNumToCny4(num float64) (string, error) {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	str := ""
	for k, v := range strnum {
		str = str + upperDigitUnit2[v] + s[k]
	}

	for i := 0; i < len(convertSteps); i++ {
		str = convertSteps[i].reg.ReplaceAllString(str, convertSteps[i].replace)
	}

	return str, nil
}

// 避免字符串反复拼接，并把类型换成rune
//
func ConvertNumToCny5(num float64) (string, error) {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	s := sliceUnit2[len(sliceUnit2)-len(strnum) : len(sliceUnit2)]
	buff := make([]rune, 0, 2*len(s))
	for k, v := range strnum {
		buff = append(buff, upperDigitUnit3[v])
		buff = append(buff, s[k])
	}
	str := string(buff)

	for i := 0; i < len(convertSteps); i++ {
		str = convertSteps[i].reg.ReplaceAllString(str, convertSteps[i].replace)
	}

	return str, nil
}

func Test_Method1(t *testing.T) {
	if str, _ := ConvertNumToCny1(1001.01); str != "壹仟零壹元零壹分" {
		t.Fatal()
	}
}

func Test_Method2(t *testing.T) {
	if str, _ := ConvertNumToCny2(1001.01); str != "壹仟零壹元零壹分" {
		t.Fatal()
	}
}

func Test_Method3(t *testing.T) {
	if str, _ := ConvertNumToCny3(1001.01); str != "壹仟零壹元零壹分" {
		t.Fatal()
	}
}

func Test_Method4(t *testing.T) {
	if str, _ := ConvertNumToCny4(1001.01); str != "壹仟零壹元零壹分" {
		t.Fatal()
	}
}

func Test_Method5(t *testing.T) {
	if str, _ := ConvertNumToCny5(1001.01); str != "壹仟零壹元零壹分" {
		t.Fatal()
	}
}

func Benchmark_Method1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertNumToCny1(1001.01)
	}
}

func Benchmark_Method2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertNumToCny2(1001.01)
	}
}

func Benchmark_Method3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertNumToCny3(1001.01)
	}
}

func Benchmark_Method4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertNumToCny4(1001.01)
	}
}

func Benchmark_Method5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertNumToCny5(1001.01)
	}
}
