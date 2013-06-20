// testxml.go project testxml.go.go
// references: http://developer.android.com/guide/topics/manifest/manifest-intro.html
package goandroid

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func NewAndroidManifest(filepath string) (*Manifest, error) {
	f, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if f.IsDir() {
		filepath = path.Join(filepath, "AndroidManifest.xml")
	}
	var dir, file = path.Split(filepath)
	b, err := ioutil.ReadFile(path.Join(dir, file))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var manifest Manifest
	err = xml.Unmarshal(b, &manifest)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &manifest, err
}

func (m *Manifest) Export(filepath string) error {
	if f, err := os.Stat(filepath); err == nil {
		if f.IsDir() {
			filepath = path.Join(filepath, "AndroidManifest.xml")
		}
	}

	b, err := xml.MarshalIndent(&m, "  ", "    ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	var dir, file = path.Split(filepath)
	err = ioutil.WriteFile(path.Join(dir, file), b, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

// SetPackageName set the application package name and solve package dependency attributes
// mainfest attributes which contain before package name will convert to new package name
func (m *Manifest) SetPackageName(pname string) *Manifest {
	var oriPackageName = m.Package
	m.Package = pname
	// process defined permission
	m = setDefinedPermission(m, oriPackageName, pname)
	// process uses-permission
	m = setUsesPermission(m, oriPackageName, pname)
	// process GCM receiver category
	m = setReceiverCategory(m, oriPackageName, pname)
	// process activity intent action
	m = setActivityFilterAction(m, oriPackageName, pname)
	// process provider
	m = setProvider(m, oriPackageName, pname)
	return m
}

func setDefinedPermission(m *Manifest, from, to string) *Manifest {
	for index, _ := range m.Permissions {
		var name = m.Permissions[index].Name
		if strings.HasPrefix(name, from) {
			m.Permissions[index].Name = strings.Replace(name, from, to, 1)
		}
	}
	return m
}

func setUsesPermission(m *Manifest, from, to string) *Manifest {
	for index, _ := range m.UsePermissions {
		var name = m.UsePermissions[index].Name
		if strings.HasPrefix(name, from) {
			m.UsePermissions[index].Name = strings.Replace(name, from, to, 1)
		}
	}
	return m
}

func setReceiverCategory(m *Manifest, from, to string) *Manifest {
	for indexReceiver, receiver := range m.Application.Receiver {
		for indexFilter, _ := range receiver.IntentFilters {
			if m.Application.Receiver[indexReceiver].IntentFilters[indexFilter].Category != nil {
				m.Application.Receiver[indexReceiver].IntentFilters[indexFilter].Category.Name = to
				break
			}
		}
	}
	return m
}

func setActivityFilterAction(m *Manifest, from, to string) *Manifest {
	var manifest = m
	for indexActivity, activity := range m.Application.Activities {
		for indexFilter, intentFilter := range activity.IntentFilters {
			for indexAction, action := range intentFilter.Action {
				if strings.HasPrefix(action.Name, from) {
					m.Application.Activities[indexActivity].IntentFilters[indexFilter].Action[indexAction].Name = strings.Replace(action.Name, from, to, 1)
				}
			}
		}
	}
	return manifest
}

func setProvider(m *Manifest, from, to string) *Manifest {
	for indexProvider, _ := range m.Application.Providers {
		m.Application.Providers[indexProvider].Authorities = to
		var readPerm = &m.Application.Providers[indexProvider].ReadPermission
		if strings.HasPrefix(*readPerm, from) {
			*readPerm = strings.Replace(*readPerm, from, to, 1)
		}
		var writePerm = &m.Application.Providers[indexProvider].WritePermission
		if strings.HasPrefix(*writePerm, from) {
			*writePerm = strings.Replace(*writePerm, from, to, 1)
		}
	}
	return m
}
