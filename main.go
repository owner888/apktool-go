package main

import (
    //"fmt"
    "log"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
)

const preferenceCurrentTutorial = "currentTutorial"

// 最上面那一层窗口，因为手机屏幕小，从栏目点开是一层遮住的，这个时候要把child Windows的句柄赋值给这个，当用户点关闭时，子窗口关闭以后把父句柄重新赋值给这个
var topWindow fyne.Window

func main() {
    a := app.NewWithID("com.apktool.go")
    a.SetIcon(theme.FyneLogo())
    logLifecycle(a)
    w := a.NewWindow("APK Tool")
    topWindow = w

    // 生成菜单
    w.SetMainMenu(makeMenu(a, w))
    w.SetMaster()

    // split box 右侧窗体
    centerBox := container.NewMax()

    leftBox  := makeLeftBox(centerBox, w)
    splitBox := makeSplitBox(centerBox, w)

    w.SetContent(container.NewBorder(nil, nil, leftBox, nil, splitBox))
    w.SetMaster()
    //w.SetPadded(false)
    w.Resize(fyne.NewSize(1125, 700))
    //w.SetFullScreen(true)
    w.ShowAndRun()
}

// 打印生命周期日志
func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

// 键盘快捷键
func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

// HBox     --->   container.NewHBox                    --->   按列摆放
// VBox     --->   container.NewVBox                    --->   按行摆放
// HSplit   --->   container.NewHSplit                  --->   按比例分割
// Colmn    --->   layout.NewGridLayoutWithColumns(4)   --->   按列平均分割
// Rows     --->   layout.NewGridLayoutWithRows(4)      --->   按行平均分割
// NewGridLayoutWithColumns：返回一个 gridLayout 结构体，可以指定列数。如果需要垂直布局，可以替换成 NewGridLayoutWithRows
// NewContainerWithLayout 返回一个 Container 实例，使布局生效
// laboxTop := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(3))


