# random

隨機幹嘛!? 吃飯 位子...

## 安裝glade

```bash
sudo apt-get install glade
```

## 需要套件

```bash
sudo apt-get install libgtk-3-dev libcairo2-dev libglib2.0-dev
```

## 執行

需要加上gtk3的版本, 怎麼查?

```bash
apt-cache policy libgtk-3-0 | grep Installed(已安裝)
libgtk-3-0:
  已安裝：3.24.20-0ubuntu1.1
  候選： 3.24.20-0ubuntu1.1
  版本列表：
 *** 3.24.20-0ubuntu1.1 500
     3.24.18-1ubuntu1 500

go run     -v -tags gtk_3_24 main.go
go install -v -tags gtk_3_24
go build   -v -tags gtk_3_24 -o main.out
```
