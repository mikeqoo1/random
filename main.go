package main

import (
	"fmt"
	"log"
	"os"
	lib "random/lib"

	"github.com/gotk3/gotk3/gtk"
)

// WindowName is the defined identifier for the main getWindow in the glade template
const WindowName = "window"

// ButtonName is the defined identifier for the list box in the glade template
const ButtonName = "button"

const TextViewName = "text"

// UIMain is the path to our glade file
const UIMain = "glade/test.glade"

func init() {
	lib.InitMyConfig()
}

func main() {

	//解析command line 參數
	gtk.Init(&os.Args)

	bldr, err := getBuilder(UIMain)
	if err != nil {
		fmt.Println("取得glade檔失敗")
		panic(err)
	}

	window, err := getWindow(bldr)
	if err != nil {
		fmt.Println("取得window物件失敗")
		panic(err)
	}

	window.SetTitle("GO GTK3 Glade Example")
	window.SetDefaultSize(300, 300)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		fmt.Println("關閉視窗失敗")
		panic(err)
	}

	textview, err := getTextView(bldr)
	if err != nil {
		panic(err)
	}

	buffer, err := textview.GetBuffer()
	if err != nil {
		log.Fatal("取得textview buffer 失敗")
		panic(err)
	}
	buffer.SetText("請按按鈕")

	button, err := getButton(bldr)
	if err != nil {
		panic(err)
	}

	button.Connect("clicked", func() {
		fmt.Println("Click-Click!")
		nums := lib.CreatRandomNumber(0, 10, 1)
		for x := 0; x < len(nums); x++ {
			//fmt.Printf("index=%d: [%d號座位, 是%s]\n", x, nums[x], lib.MyConfig.Seat[nums[x]])
			fmt.Printf("今天, 是吃%s\n", lib.MyConfig.Eat[nums[x]])
			text := fmt.Sprintf("今天, 是吃%s\n", lib.MyConfig.Eat[nums[x]])
			fmt.Printf("今天, 是喝%s\n", lib.MyConfig.Drink[nums[x]])
			text = fmt.Sprintf("%s今天, 是喝%s\n", text, lib.MyConfig.Drink[nums[x]])
			buffer.SetText(text)
		}
	})
	if err != nil {
		panic(err)
	}

	window.ShowAll()

	gtk.Main()
}

// getBuilder returns *gtk.getBuilder loaded with glade resource (if resource is given)
func getBuilder(filename string) (*gtk.Builder, error) {
	fmt.Println(filename)
	b, err := gtk.BuilderNew()
	if err != nil {
		return nil, err
	}

	if filename != "" {
		err = b.AddFromFile(filename)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

// getWindow returns *gtk.Window object from the glade resource
func getWindow(b *gtk.Builder) (*gtk.Window, error) {

	obj, err := b.GetObject(WindowName)
	if err != nil {
		return nil, err
	}

	window, ok := obj.(*gtk.Window) //斷言判斷是不是gtk物件
	if !ok {
		return nil, err
	}

	return window, nil
}

// getButton returns *gtk.Button object from the glade resource
func getButton(b *gtk.Builder) (*gtk.Button, error) {

	obj, err := b.GetObject(ButtonName)
	if err != nil {
		return nil, err
	}

	button, ok := obj.(*gtk.Button)
	if !ok {
		return nil, err
	}

	return button, nil
}

// getTextView 取得TextView
func getTextView(b *gtk.Builder) (*gtk.TextView, error) {
	obj, err := b.GetObject(TextViewName)
	if err != nil {
		return nil, err
	}

	textview, ok := obj.(*gtk.TextView)
	if !ok {
		return nil, err
	}

	return textview, nil
}
