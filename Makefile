#主程式
Main=main.go

#執行檔
Bin_Out=-o bin/David.exe

main:
	go install
	go build $(Bin_Out) $(Main)