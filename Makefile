#主程式
Main=main.go

#執行檔
Bin_Out=-o bin/David.out

main:
	go install
	go build $(Bin_Out) $(Main)