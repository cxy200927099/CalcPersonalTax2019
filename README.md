# CalcPersonalTax2019
2019年个人所得税计算工具

应该还有bug，大神别喷！


# 各平台使用方式及参数说明
## 参数说明
- month: 月薪
- monthExcept: 每月五险一金扣除数
- monthSpecialExcept: 每月专项扣除数
- yearEndBonus: 年终奖额度
- yearEndBonusMonth: 年终奖几月份发
- calcType: 计算方式，按年还是按月，**按年计算方式还未实现**

## mac平台计算方式
```./CalPersonalTax_mac -month=30000 -monthExcept=4500 -monthSpecialExcept=2000 -yearEndBonus=100000 -yearEndBonusMonth=2 -calcType=month```

## linux平台运行方式
```./CalPersonalTax_linux-64 -month=30000 -monthExcept=4500 -monthSpecialExcept=2000 -yearEndBonus=100000 -yearEndBonusMonth=2 -calcType=month ```

## windows平台运行方式
打开命令行 win+r
将CalPersonalTax_x86-64.exe拖到命令行后，后面添加如下：
 -month=30000 -monthExcept=4500 -monthSpecialExcept=2000 -yearEndBonus=100000 -yearEndBonusMonth=2 -calcType=month
回车即可

我这边运行情况
```\\192.168.116.89\chenxingyi\work\go\src\goDemo\cxy\CalPersonalTax_x86-64.exe -month=30000 -monthExcept=4500 -monthSpecialExcept=2000 -yearEndBonus=100000 -yearEndBonusMonth=2 -calcType=month```
- 运行结果：
```
输入数据 月收入: 30000  每月五险一金扣除: 4500  每月专项扣除 2000  年终奖: 100000  计算方式: month
第  1  月税后收入: 24945  扣税: 555
第  2  月税后收入: 114875  扣税: 10625
第  3  月税后收入: 22500  扣税: 3000
第  4  月税后收入: 21800  扣税: 3700
第  5  月税后收入: 21800  扣税: 3700
第  6  月税后收入: 21800  扣税: 3700
第  7  月税后收入: 21800  扣税: 3700
第  8  月税后收入: 21800  扣税: 3700
第  9  月税后收入: 21800  扣税: 3700
第  10  月税后收入: 21800  扣税: 3700
第  11  月税后收入: 21625  扣税: 3875
第  12  月税后收入: 20875  扣税: 4625
全年税后收入 357420  全年总的扣税: 48580
```

# 源码编译

## 编译mac:
 ```go build -o ./CalPersonalTax_mac -i CalPersonalTax.go```

## 编译linux:
```GOOS=linux GOARCH=amd64 go build -o ./CalPersonalTax_linux-64 -i ./CalPersonalTax.go```

## 编译windows:
```GOOS=windows GOARCH=amd64 go build -o ./CalPersonalTax_x86-64.exe -i ./CalPersonalTax.go```
