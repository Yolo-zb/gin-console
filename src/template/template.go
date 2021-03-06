package template

import (
	"bytes"
	"fmt"
	"github.com/Yolo-zb/gin-console/helper"
	"github.com/Yolo-zb/gin-console/src/gorm"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)
var ModelTemplate = `
{{- /* delete empty line */ -}}
package model

const (
	// Dictionary
)

type {{ .BigCamel }} struct {
{{- /* delete empty line */ -}}
{{range $index, $value := .Column}} 
	{{$value.ColumnName}} {{$value.DataType}} {{end}}
}

func (*{{ .BigCamel }}) TableName() string {
	return "{{ .TableName }}"
}
`
var DaoTemplate = `
{{- /* delete empty line */ -}}
package dao

import (
	localGorm "github.com/Yolo-zb/gin-console/src/gorm"
	"github.com/jinzhu/gorm"
	"{{ .Module }}/common/model"
)

type {{ .BigCamel }} struct {
	gorm *gorm.DB
}

func New{{ .BigCamel }}() {{ .BigCamel }} {
	return {{ .BigCamel }}{
		gorm:localGorm.GetGorm("localhost"),
	}
}

func (ctl *{{ .BigCamel }}) GetById(id int) model.{{ .BigCamel }} {
	{{ .Camel }} := model.{{ .BigCamel }}{}
	ctl.gorm.Where("id = ?", id).First(&{{ .Camel }})
	return model.{{ .BigCamel }}{}
}
`
var ServiceTemplate = `
{{- /* delete empty line */ -}}
package service

import (
	"{{ .Module }}/common/dao"
	"{{ .Module }}/common/model"
)

type {{ .Camel }}Service struct {
	data dao.{{ .BigCamel }}
}

var {{ .BigCamel }} = {{ .Camel }}Service{
	data: dao.New{{ .BigCamel }}(),
}

func (ctl *{{ .Camel }}Service) GetById(id int) model.{{ .BigCamel }} {
	return ctl.data.GetById(id)
}
`
var sqlTypeMap = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"bigint":             "int64",
	"int unsigned":       "uint",
	"integer unsigned":   "uint",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint32",
	"bigint unsigned":    "uint64",
	"bit":                "byte",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}
type Model struct {
	TableName string // ??????
	Module 	  string // ????????????
	Column	  []TableColumn
	PathTemplate	map[string]string // ?????????????????????????????????
	Camel	  string // ???????????????
	BigCamel  string // ???????????????
	Gorm 	  gorm.A // ??????????????????
}
type TableExist struct {
	TableName string `gorm:"column:TABLE_NAME"`
}
var tableExistRes TableExist
type TableColumn struct {
	ColumnName string `gorm:"column:COLUMN_NAME"` // ????????????
	DataType string `gorm:"column:DATA_TYPE"` // ????????????
}
var tableColumnRes []TableColumn
func (s *Model) Execute() {
	fmt.Println(s.Camel)
	if (s.Gorm.GetGorm("localhost").Raw("select TABLE_NAME from INFORMATION_SCHEMA.TABLES WHERE `TABLE_NAME` = '"+ s.TableName +"' AND `TABLE_SCHEMA` = 'zulin'").Scan(&tableExistRes).RecordNotFound()){
		log.Fatal("??????????????????")
	}
	s.Gorm.GetGorm("localhost").Raw("select COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT from information_schema.COLUMNS where table_name = '"+ s.TableName +"' AND `TABLE_SCHEMA` = 'zulin'  ORDER BY ORDINAL_POSITION ASC").Scan(&tableColumnRes)
	s.Column = tableColumnRes
	for key, value := range s.Column{
		s.Column[key].ColumnName = TranCamel(value.ColumnName)
		s.Column[key].DataType = sqlTypeMap[value.DataType]
	}
	for pathName, templateString := range s.PathTemplate{
		dirPathName, filePathName := s.getPath(pathName)
		s.createFile(dirPathName, filePathName, templateString)
	}
	//fmt.Println(service.User.GetById(1))
	//fmt.Println(service.User.GetById(2337204))
	//gorm.Close("localhost")
}

func (s *Model) getPath(dir string) (string, string) {
	wd, _ := os.Getwd()
	return path.Join(wd, "/common/" + dir), path.Join(wd, "/common/" + dir + "/" + s.Camel + ".go")
}

func (s *Model) createFile(dirPathName string, filePathName string, templateString string) {
	exist := helper.PathExists(dirPathName)
	if !exist {
		err := os.MkdirAll(dirPathName, 0777)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Println("mkdir" + dirPathName + " success!")
		}
	}
	buf := new(bytes.Buffer)
	tmpl, _ := template.New("name").Parse(templateString)
	tmpl.Execute(buf, s)
	ioutil.WriteFile(filePathName, buf.Bytes(), 0777)
}

// ?????????????????????????????????
func TranCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
