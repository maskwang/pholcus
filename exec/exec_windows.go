// +build windows
package exec

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/maskwang/pholcus/config"

	"github.com/maskwang/pholcus/cmd" // cmd版
	"github.com/maskwang/pholcus/gui" // gui版
	"github.com/maskwang/pholcus/web" // web版
)

func run(which string) {
	exec.Command("cmd.exe", "/c", "title", config.FULL_NAME).Start()

	// 选择运行界面
	switch which {
	case "gui":
		gui.Run()

	case "cmd":
		cmd.Run()

	case "web":
		fallthrough
	default:
		ctrl := make(chan os.Signal, 1)
		signal.Notify(ctrl, os.Interrupt, os.Kill)
		go web.Run()
		<-ctrl
	}
}
