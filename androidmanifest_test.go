package goandroid

import (
	"log"
	"testing"
)

func TestNewAndroidManifest(t *testing.T) {
	var filepath = "material"
	manifest, err := NewAndroidManifest(filepath)
	if err != nil {
		t.Fatal("read manifest", err)
	}

	var pname = "com.example.test.name"

	manifest = manifest.SetPackageName(pname, "35")
	if manifest.Package != pname {
		t.Fatal("setPackageName failed.")
	}
	log.Println("Permissions:\n", manifest.Permissions)
	log.Println("UsePermissions:\n", manifest.UsePermissions)
	log.Println("Receiver:\n", manifest.Application.Receiver)
	log.Println("Activity:\n", manifest.Application.Activities)
	log.Println("Provider:\n", manifest.Application.Providers)

	err = manifest.ExportAndroidManifest("material/output_AndroidManifest.xml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewResourceXml(t *testing.T) {
	var stringPath = "material/strings.xml"
	var appName = "呵呵>App"
	resString, err := NewResourceXml(stringPath)
	if err != nil {
		t.Fatal(err)
	}
	var isWedding = false
	resString = resString.SetString("app_name", appName)
	resString = resString.SetString("content_authority", "com.example.name")
	if isWedding {
		resString = resString.SetString("picture_dir", "享幸福")
	} else {
		resString = resString.SetString("picture_dir", "Miiroad")
	}
	err = resString.ExportResourceXml(stringPath)
	if err != nil {
		t.Fatal(err)
	}
}
