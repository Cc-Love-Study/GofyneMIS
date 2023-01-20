package server

import (
	"binhaiControl/model"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xuri/excelize/v2"
)

type PersonService struct {
	Persons   []*model.Person
	PersonNum int
	PersonSum int
	DataPath  string
}

func NewPersonService() *PersonService {
	return &PersonService{PersonNum: 1, PersonSum: 0, DataPath: "data.json"}
}

func CreateOrReadData() (ps *PersonService) {
	ps = NewPersonService()
	data, _ := ioutil.ReadFile("data.json")
	_ = json.Unmarshal(data, ps)
	return ps
}

/* 数据存储 */
func (this *PersonService) Save() error {
	data, err := json.Marshal(this)
	if err != nil {
		return err
	} else {
		err = ioutil.WriteFile(this.DataPath, data, 0666)
		if err != nil {
			return err
		}
		return nil
	}
}

/* 添加一个人员信息 */
func (this *PersonService) Add(p *model.Person) {
	p.Id = this.PersonNum
	this.PersonNum += 1
	this.PersonSum += 1
	this.Persons = append(this.Persons, p)
}

/* 查找id的索引 */
func (this *PersonService) FindId(id int) int {
	for i, v := range this.Persons {
		if v.Id == id {
			return i
		}
	}
	return -1
}

/* 删除一个人员信息 */
func (this *PersonService) Detel(id int) bool {
	index := this.FindId(id)
	if index == -1 {
		return false
	} else {
		this.Persons = append(this.Persons[:index], this.Persons[index+1:]...)
		this.PersonSum -= 1
		return true
	}
}

/* 修改一个人员的信息 */
func (this *PersonService) Change(id int, name string, sex string, age int, phoneone string, phonetwo string, phonethree string, idcard string, content string, job int, in string, out string) bool {
	index := this.FindId(id)
	if index == -1 {
		return false
	} else {
		this.Persons[index].Name = name
		this.Persons[index].Sex = sex
		this.Persons[index].Age = age
		this.Persons[index].Phoneone = phoneone
		this.Persons[index].Phonetwo = phonetwo
		this.Persons[index].Phonethree = phonethree
		this.Persons[index].IdCard = idcard
		this.Persons[index].Content = content
		this.Persons[index].Job = job
		this.Persons[index].InDate = in
		this.Persons[index].OutDate = out

		return true
	}
}

/* 重置索引号 */
func (this *PersonService) ReId() {
	for i, v := range this.Persons {
		v.Id = i + 1
	}
	this.PersonNum = this.PersonSum + 1
}

func (this *PersonService) ShowAllInfo() {
	for _, v := range this.Persons {
		v.PrintInfo()
	}
}

func (this *PersonService) ImportInfo(sel int) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "序号")
	f.SetCellValue("Sheet1", "B1", "姓名")
	f.SetCellValue("Sheet1", "C1", "性别")
	f.SetCellValue("Sheet1", "D1", "年龄")
	f.SetCellValue("Sheet1", "E1", "家属电话1")
	f.SetCellValue("Sheet1", "F1", "家属电话2")
	f.SetCellValue("Sheet1", "G1", "家属电话3")
	f.SetCellValue("Sheet1", "H1", "身份证号")
	f.SetCellValue("Sheet1", "I1", "身份")
	f.SetCellValue("Sheet1", "J1", "入住时间")
	f.SetCellValue("Sheet1", "K1", "出院时间")
	f.SetCellValue("Sheet1", "L1", "详细描述")
	// 设置列宽
	// 设置列宽
	f.SetColWidth("Sheet1", "A", "L", 20)
	line := 1
	// 循环写入数据
	if sel == 0 {
		for _, v := range this.Persons {
			line++
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.Id)
			f.SetCellStr("Sheet1", fmt.Sprintf("B%d", line), v.Name)
			f.SetCellStr("Sheet1", fmt.Sprintf("C%d", line), v.Sex)
			f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.Age)
			f.SetCellStr("Sheet1", fmt.Sprintf("E%d", line), v.Phoneone)
			f.SetCellStr("Sheet1", fmt.Sprintf("F%d", line), v.Phonetwo)
			f.SetCellStr("Sheet1", fmt.Sprintf("G%d", line), v.Phonethree)
			f.SetCellStr("Sheet1", fmt.Sprintf("H%d", line), v.IdCard)
			if v.Job == 1 {
				f.SetCellStr("Sheet1", fmt.Sprintf("I%d", line), "入院老人")
			} else {
				f.SetCellStr("Sheet1", fmt.Sprintf("I%d", line), "服务员")
			}
			f.SetCellStr("Sheet1", fmt.Sprintf("J%d", line), v.InDate)
			f.SetCellStr("Sheet1", fmt.Sprintf("K%d", line), v.OutDate)
			f.SetCellStr("Sheet1", fmt.Sprintf("L%d", line), v.Content)
		}
	} else if sel == 1 {
		for _, v := range this.Persons {
			if v.Job == 1 {
				line++
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.Id)
				f.SetCellStr("Sheet1", fmt.Sprintf("B%d", line), v.Name)
				f.SetCellStr("Sheet1", fmt.Sprintf("C%d", line), v.Sex)
				f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.Age)
				f.SetCellStr("Sheet1", fmt.Sprintf("E%d", line), v.Phoneone)
				f.SetCellStr("Sheet1", fmt.Sprintf("F%d", line), v.Phonetwo)
				f.SetCellStr("Sheet1", fmt.Sprintf("G%d", line), v.Phonethree)
				f.SetCellStr("Sheet1", fmt.Sprintf("H%d", line), v.IdCard)
				if v.Job == 1 {
					f.SetCellStr("Sheet1", fmt.Sprintf("I%d", line), "入院老人")
				} else {
					f.SetCellStr("Sheet1", fmt.Sprintf("I%d", line), "服务员")
				}
				f.SetCellStr("Sheet1", fmt.Sprintf("J%d", line), v.InDate)
				f.SetCellStr("Sheet1", fmt.Sprintf("K%d", line), v.OutDate)
				f.SetCellStr("Sheet1", fmt.Sprintf("L%d", line), v.Content)
			}
		}
	} else if sel == 2 {
		for _, v := range this.Persons {
			if v.Job == 2 {
				line++
				f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.Id)
				f.SetCellStr("Sheet1", fmt.Sprintf("B%d", line), v.Name)
				f.SetCellStr("Sheet1", fmt.Sprintf("C%d", line), v.Sex)
				f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.Age)
				f.SetCellStr("Sheet1", fmt.Sprintf("E%d", line), v.Phoneone)
				f.SetCellStr("Sheet1", fmt.Sprintf("F%d", line), v.Phonetwo)
				f.SetCellStr("Sheet1", fmt.Sprintf("G%d", line), v.Phonethree)
				f.SetCellStr("Sheet1", fmt.Sprintf("H%d", line), v.IdCard)
				if v.Job == 1 {
					f.SetCellStr("Sheet1", fmt.Sprintf("I%d", line), "入院老人")
				} else {
					f.SetCellStr("Sheet1", fmt.Sprintf("I%d", line), "服务员")
				}
				f.SetCellStr("Sheet1", fmt.Sprintf("J%d", line), v.InDate)
				f.SetCellStr("Sheet1", fmt.Sprintf("K%d", line), v.OutDate)
				f.SetCellStr("Sheet1", fmt.Sprintf("L%d", line), v.Content)
			}
		}
	}
	if err := f.SaveAs("滨海人员信息.xlsx"); err != nil {
		fmt.Println(err)
	}
}
