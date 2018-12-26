package main

import (
	"flag"
	"fmt"
)

const (
	TOTAL_YEAR_36000 = 0.03
	TOTAL_YEAR_144000 = 0.1
	TOTAL_YEAR_300000 = 0.20
	TOTAL_YEAR_420000 = 0.25
	TOTAL_YEAR_660000 = 0.3
	TOTAL_YEAR_960000 = 0.35
	TOTAL_YEAR_960000_UPER = 0.45

	EXCEPTE_36000 = 0
	EXCEPTE_144000 = 2520
	EXCEPTE_300000 = 16920
	EXCEPTE_420000 = 31920
	EXCEPTE_660000 = 52920
	EXCEPTE_960000 = 85920
	EXCEPTE_960000_UPER = 181920
)

var monthSalary = flag.Float64("month",0, "月薪")
var monthExcept = flag.Float64("monthExcept", 0, "每月五险一金的扣除")
var monthSpecialExcept = flag.Float64("monthSpecialExcept", 0, "每月专项扣除")
var yearEndBonus= flag.Float64("yearEndBonus",0, "年终奖")
var yearEndBonusMonth= flag.Int("yearEndBonusMonth",0, "年终奖计算的月份")
var calcType = flag.String("calcType", "month", "选择扣除方式 month:按月; year:按年")

func main(){
	flag.Parse()

	fmt.Println("输入数据 月收入:",*monthSalary," 每月五险一金扣除:",*monthExcept," 每月专项扣除",*monthSpecialExcept," 年终奖:",*yearEndBonus," 计算方式:",*calcType)
	var realPay float64
	mPay := *monthSalary
	mEPay := *monthExcept + 5000 + *monthSpecialExcept
	yearEndPay := *yearEndBonus
	var totalExceptTax float64
	var curMonthTax float64
	var totalRealPay float64
	for i := 1; i <= 12; i++ {

		tmp := mPay * float64(i) - mEPay * float64(i)
		if i >= *yearEndBonusMonth {
			tmp = mPay * float64(i) + yearEndPay - mEPay * float64(i)
		}
		if tmp < 0{
			tmp = 0
		}

		if tmp < 36000{
			curMonthTax = tmp * TOTAL_YEAR_36000 - EXCEPTE_36000 - totalExceptTax
		}else if tmp < 144000{
			curMonthTax = tmp * TOTAL_YEAR_144000 - EXCEPTE_144000 - totalExceptTax
		}else if tmp < 300000{
			curMonthTax = tmp * TOTAL_YEAR_300000 - EXCEPTE_300000 - totalExceptTax
		}else if tmp < 420000{
			curMonthTax = tmp * TOTAL_YEAR_420000 - EXCEPTE_420000 - totalExceptTax
		}else if tmp < 660000{
			curMonthTax = tmp * TOTAL_YEAR_660000 - EXCEPTE_660000 - totalExceptTax
		}else if tmp < 960000{
			curMonthTax = tmp * TOTAL_YEAR_960000 - EXCEPTE_960000 - totalExceptTax
		}else{
			curMonthTax = tmp * TOTAL_YEAR_960000_UPER - EXCEPTE_960000_UPER - totalExceptTax
		}
		if curMonthTax < 0{
			curMonthTax = 0
		}


		realPay = mPay - *monthExcept - curMonthTax
		if i == *yearEndBonusMonth{
			realPay = mPay - *monthExcept - curMonthTax + yearEndPay
		}
		if tmp == 0{
			realPay = mPay
		}

		totalExceptTax += curMonthTax
		totalRealPay += realPay
		fmt.Println("第 ",i," 月税后收入:",realPay," 扣税:",curMonthTax)
	}
	fmt.Println("全年税后收入",totalRealPay," 全年总的扣税:",totalExceptTax)
}
