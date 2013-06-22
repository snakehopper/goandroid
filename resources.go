package goandroid

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Resources struct {
	XMLName xml.Name `xml:"resources"`
	Strings []String `xml:"string"`
	Colors  []Color  `xml:"color"`
}

type String struct {
	Value  string `xml:",innerxml"`
	Name   string `xml:"name,attr"`
	Format string `xml:"formatted,attr,omitempty"`
}

type Color struct {
	Value string `xml:",innerxml"`
	Name  string `xml:"name,attr"`
}

func NewResources(filepath string) (*Resources, error) {
	f, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if f.IsDir() {
		log.Fatal(filepath, "is not a valid file")
	}
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var res Resources
	err = xml.Unmarshal(b, &res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, err
}

func (res *Resources) SetString(name, value string) *Resources {
	for index, _ := range res.Strings {
		if name == res.Strings[index].Name {
			m := new(bytes.Buffer)
			xml.EscapeText(m, []byte(value))
			res.Strings[index].Value = m.String()
			//res.Strings[index].Value = value
			break
		}
	}
	return res
}

func (res *Resources) AddString(name, value string) *Resources {
	res.Strings = append(res.Strings, String{Value: value, Name: name})
	return res
}

func (res *Resources) SetColor(name, value string) *Resources {
	for index, _ := range res.Colors {
		if name == res.Colors[index].Name {
			m := new(bytes.Buffer)
			xml.EscapeText(m, []byte(value))
			res.Colors[index].Value = m.String()
			break
		}
	}
	return res
}

func (res *Resources) Export(filepath string) error {
	if f, err := os.Stat(filepath); err == nil {
		if f.IsDir() {
			return fmt.Errorf("%s is not a valid file", filepath)
		}
	}

	b, err := xml.MarshalIndent(&res, "", "    ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = ioutil.WriteFile(filepath, b, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
