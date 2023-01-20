package main

import (
	"binhaiControl/model"
	"binhaiControl/server"
	"fmt"
	"image/color"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*初始化配置*/
func init() {
	os.Setenv("FYNE_FONT", "Alibaba-PuHuiTi-Medium.ttf")
}

/* 一条信息 里面应该有各个输入框方便获取/填入输入框内信息 */
type InfoWidget struct {
	id          *widget.Label
	dbutton     *widget.Button
	wname       *widget.Entry
	wsex        *widget.RadioGroup
	wage        *widget.Entry
	wphoneone   *widget.Entry
	wphonetwo   *widget.Entry
	wphonethree *widget.Entry
	widcard     *widget.Entry
	wcontent    *widget.Entry
	wjob        *widget.RadioGroup
	win         *widget.Entry
	wout        *widget.Entry
}

/*
	peopleList是显示容器 放入peopleList用于显示
	InfoList是组件数组 用于各个Info的各个Widget的信息获取
*/
type OsView struct {
	w          float32
	h          float32
	app        fyne.App
	win        fyne.Window
	ps         *server.PersonService
	peopleList *fyne.Container
	infoScroll *container.Scroll
	infoList   []*InfoWidget
	window     fyne.Canvas
	showFlag   int
	findW      string
}

/* 工厂函数 */
func NewView() *OsView {
	return &OsView{w: 1000, h: 800, showFlag: 0}
}

/* 返还系统界面上部分 标题 搜索框 按键等内容 */
func (this *OsView) NewTop() *fyne.Container {
	Title := canvas.NewText("滨海信息管理系统", color.Black)
	Title.TextSize = 32
	Titlec := container.New(layout.NewCenterLayout(), Title)
	newButton := widget.NewButton("新建", func() {
		fmt.Println("新建")
		/* 这里放入新建函数 这里只是在界面层创建 在server层 还没有保存 */
		this.NewInfo(this.ps.PersonNum, "", "男", 0, "", "", "", "", "", 1, "", "")
		person := model.NewPerson("", "男", 0, "", "", "", "", "", 1)
		this.ps.Add(person)
		this.ps.Save()
	})
	importButton := widget.NewButton("导出", func() {
		fmt.Println("导出")
		/* 这里放入导出函数 */
		this.ps.ImportInfo(this.showFlag)
	})
	saveButton := widget.NewButton("保存", func() {
		fmt.Println("保存")
		/* 这里放入导出函数 */
		for _, v := range this.infoList {
			person := model.NewPerson("", "", 99, "", "", "", "", "", 1)
			person.Id, _ = strconv.Atoi(v.id.Text)
			person.Name = v.wname.Text
			person.Sex = v.wsex.Selected
			person.Age, _ = strconv.Atoi(v.wage.Text)
			person.Phoneone = v.wphoneone.Text
			person.Phonetwo = v.wphonetwo.Text
			person.Phonethree = v.wphonethree.Text
			person.IdCard = v.widcard.Text
			person.Content = v.wcontent.Text
			if v.wjob.Selected == "入住老人" {
				person.Job = 1
			} else {
				person.Job = 2
			}
			person.InDate = v.win.Text
			person.OutDate = v.wout.Text
			this.ps.Change(person.Id, person.Name, person.Sex, person.Age, person.Phoneone, person.Phonetwo,
				person.Phonethree, person.IdCard, person.Content, person.Job, person.InDate, person.OutDate)
		}

		this.ps.ReId()
		this.ps.Save()
		this.peopleList.RemoveAll()
		this.NewShow()
		widget.ShowPopUpAtPosition(widget.NewLabel("保存成功!"), this.window, fyne.NewPos(350, 30))
	})
	selectButton := widget.NewSelect([]string{"全部", "入住老人", "服务员"}, func(value string) {
		fmt.Println("筛选:", value)
		/* 这里放入筛选函数 */
		if value == "全部" {
			this.showFlag = 0
		} else if value == "入住老人" {
			this.showFlag = 1
		} else if value == "服务员" {
			this.showFlag = 2
		}
		this.peopleList.RemoveAll()
		this.NewShow()
	})
	selectButton.Selected = "全部"
	findEnt := widget.NewEntry()
	findEnt.PlaceHolder = "输入姓名"
	findButton := widget.NewButton("查询", func() {
		fmt.Println("查询:", findEnt.Text)
		/* 这里放入查询函数 */
		if len(findEnt.Text) != 0 {
			this.showFlag = 3
			this.findW = findEnt.Text
		} else {
			this.showFlag = 0
		}

		this.peopleList.RemoveAll()
		this.NewShow()
	})
	line := canvas.NewLine(color.Black)
	// left
	funcToolsone := container.New(layout.NewHBoxLayout(), newButton, importButton, saveButton, selectButton)
	// 搜索框
	funcToolstwo := container.New(layout.NewFormLayout(), findButton, findEnt)
	funcTools := container.New(layout.NewGridLayout(2), funcToolsone, funcToolstwo)

	attr1 := widget.NewLabel("ID")
	attr2 := widget.NewLabel("姓名")
	attr3 := widget.NewLabel("性别")
	attr4 := widget.NewLabel("年龄")
	attr3_4 := container.New(layout.NewHBoxLayout(), attr3, attr4)
	attr5 := widget.NewLabel("家属电话1")
	attr6 := widget.NewLabel("家属电话2")
	attr7 := widget.NewLabel("家属电话3")
	attr8 := widget.NewLabel("身份证号")
	attr10 := widget.NewLabel("职位")
	attr11 := widget.NewLabel("入院日期")
	attr12 := widget.NewLabel("出院日期")

	attriContent := container.New(layout.NewGridLayout(10), attr1, attr2, attr3_4, attr5, attr6, attr7,
		attr8, attr10, attr11, attr12)
	topContent := container.New(layout.NewVBoxLayout(), Titlec, funcTools, line, attriContent)
	return topContent
}

func (this *OsView) infoDetel(id int) {
	var index int
	for i, v := range this.infoList {
		if v.id.Text == fmt.Sprint(id) {
			index = i
		}
	}
	this.infoList = append(this.infoList[:index], this.infoList[index+1:]...)
}

func (this *OsView) NewInfo(id int, name string, sex string, age int, phoneone string, phonetwo string, phonethree string, idcard string,
	content string, job int, indate string, outdate string) {
	// 创建一个 容器
	info := container.New(layout.NewGridLayout(10))

	wcontent := widget.NewEntry()
	wcontent.SetText(content)
	wcontent.MultiLine = true

	cont2 := container.New(layout.NewFormLayout(), widget.NewLabel("老人\n详细\n描述"), wcontent)
	info2 := container.New(layout.NewVBoxLayout(), info, cont2)
	deletButton := widget.NewButton("删除", func() {
		fmt.Println("删除这条信息")
		this.infoDetel(id)
		this.ps.Detel(id)
		this.peopleList.Remove(info2)
	})
	wid := widget.NewLabel(fmt.Sprint(id))
	widadet := container.New(layout.NewHBoxLayout(), wid, deletButton)

	wname := widget.NewEntry()
	wname.SetText(name)

	wsex := widget.NewRadioGroup([]string{"男", "女"}, func(s string) {})
	wsex.SetSelected(sex)

	wage := widget.NewEntry()
	wage.SetText(fmt.Sprint(age))

	wsexAwage := container.New(layout.NewHBoxLayout(), wsex, wage)

	wphoneone := widget.NewEntry()
	wphoneone.SetText(phoneone)

	wphonetwo := widget.NewEntry()
	wphonetwo.SetText(phonetwo)

	wphonethree := widget.NewEntry()
	wphonethree.SetText(phonethree)

	widcard := widget.NewEntry()
	widcard.SetText(idcard)

	wjob := widget.NewRadioGroup([]string{"入住老人", "服务员"}, func(value string) {
	})
	if job == 1 {
		wjob.SetSelected("入住老人")
	} else {
		wjob.SetSelected("服务员")
	}

	win := widget.NewEntry()
	win.SetText(indate)

	wout := widget.NewEntry()
	wout.SetText(outdate)

	iw := InfoWidget{wid, deletButton, wname, wsex, wage, wphoneone, wphonetwo, wphonethree, widcard, wcontent, wjob, win, wout}
	this.infoList = append(this.infoList, &iw)
	info.Add(widadet)
	info.Add(wname)
	info.Add(wsexAwage)
	info.Add(wphoneone)
	info.Add(wphonetwo)
	info.Add(wphonethree)
	info.Add(widcard)
	info.Add(wjob)
	info.Add(win)
	info.Add(wout)
	this.peopleList.Add(info2)
}

/* 显示全部的 已存在的 */
func (this *OsView) NewShow() {
	if this.showFlag == 0 {
		for _, v := range this.ps.Persons {
			this.NewInfo(v.Id, v.Name, v.Sex, v.Age, v.Phoneone, v.Phonetwo, v.Phonethree, v.IdCard, v.Content, v.Job, v.InDate, v.OutDate)
		}
	} else if this.showFlag == 1 {
		for _, v := range this.ps.Persons {
			if v.Job == 1 {
				this.NewInfo(v.Id, v.Name, v.Sex, v.Age, v.Phoneone, v.Phonetwo, v.Phonethree, v.IdCard, v.Content, v.Job, v.InDate, v.OutDate)
			}
		}
	} else if this.showFlag == 2 {
		for _, v := range this.ps.Persons {
			if v.Job == 2 {
				this.NewInfo(v.Id, v.Name, v.Sex, v.Age, v.Phoneone, v.Phonetwo, v.Phonethree, v.IdCard, v.Content, v.Job, v.InDate, v.OutDate)
			}
		}
	} else if this.showFlag == 3 {
		for _, v := range this.ps.Persons {
			if v.Name == this.findW {
				this.NewInfo(v.Id, v.Name, v.Sex, v.Age, v.Phoneone, v.Phonetwo, v.Phonethree, v.IdCard, v.Content, v.Job, v.InDate, v.OutDate)
			}
		}
	}
}

/* 放入全部的内容 */
func (this *OsView) NewContent() *fyne.Container {
	// 这里是固定的
	topContent := this.NewTop()
	// 这里放入 List
	this.ps = server.CreateOrReadData()
	this.NewShow()
	content := container.New(layout.NewBorderLayout(topContent, nil, nil, nil), topContent, this.infoScroll)
	return content
}

/* 界面运行函数 */
func (this *OsView) MainView() {

	this.win.Resize(fyne.NewSize(this.w, this.h))
	this.window = this.win.Canvas()
	this.peopleList = container.New(layout.NewVBoxLayout())
	this.infoScroll = container.NewVScroll(this.peopleList)
	content := this.NewContent()
	this.win.SetContent(content)
}

func main() {
	defer os.Unsetenv("FYNE_FONT")
	// 创建一个界面结构体
	view := NewView()
	// 创建一个服务结构体
	MyPersonServer := server.NewPersonService()
	view.ps = MyPersonServer
	myApp := app.New()
	myWindow := myApp.NewWindow("滨海信息管理系统")
	view.app = myApp
	view.win = myWindow
	// 运行界面
	view.MainView()
	view.win.CenterOnScreen()
	view.win.SetCloseIntercept(func() {

		tipc := container.New(layout.NewVBoxLayout())
		pop := widget.NewPopUp(tipc, view.window)
		cancel := func() {
			pop.Hide()
		}
		ok := func() {
			view.app.Quit()
		}
		okw := widget.NewButton("确定", ok)
		cancelw := widget.NewButton("取消", cancel)
		tipc.Add(widget.NewLabel("您是否要退出?请检查是否保存"))
		tipc.Add(okw)
		tipc.Add(cancelw)
		pop.ShowAtPosition(fyne.NewPos(view.w/2-100, view.h/2-100))
	})
	ico, _ := fyne.LoadResourceFromPath("./myicon.png")
	view.win.SetIcon(ico)
	view.win.Show()
	view.app.Run()
}
