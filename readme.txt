month: 月薪
monthExcept: 每月五险一金扣除数
monthSpecialExcept: 每月专项扣除数
yearEndBonus: 年终奖额度
yearEndBonusMonth: 年终奖几月份发
calcType: 计算方式，按年还是按月
./CalPersonalTax_mac -month=30000 -monthExcept=4500 -monthSpecialExcept=2000 -yearEndBonus=100000 -yearEndBonusMonth=2 -calcType=month

编译mac:
 go build -o ./CalPersonalTax_mac -i CalPersonalTax.go

编译linux:
GOOS=linux GOARCH=amd64 go build -o ./CalPersonalTax_linux-64 -i ./CalPersonalTax.go

编译windows:
GOOS=windows GOARCH=amd64 go build -o ./CalPersonalTax_x86-64.exe -i ./CalPersonalTax.go
