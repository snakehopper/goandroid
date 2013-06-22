package goandroid

import (
	"os"
	"testing"
)

var (
	manifest *Manifest

	filepath string

	output string
)

func setup() {
	filepath = "material"
	output = filepath + "/output_AndroidManifest.xml"
	manifest, _ = NewAndroidManifest(filepath)
}

func TestNewAndroidManifest(t *testing.T) {
	var filepath = "material"
	manifest, err := NewAndroidManifest(filepath)
	if err != nil {
		t.Fatal("read manifest", err)
	}

	var want = "com.google.android.apps.iosched"
	if want != manifest.Package {
		t.Fatalf("NewAndroidManifest package want %v, but %v\n", want, manifest.Package)
	}
}

func TestManifest_SetPackageName(t *testing.T) {
	setup()

	var want = "com.example.test.name"

	manifest = manifest.SetPackageName(want)
	if want != manifest.Package {
		t.Fatalf("SetPackageName want %v, but %v\n", want, manifest.Package)
	}
}

func TestManifest_ExportAndroidManifest(t *testing.T) {
	setup()

	if err := manifest.Export(output); err != nil {
		t.Fatal("ExportAndroidManifest error: ", err)
	}

	if om, _ := NewAndroidManifest(output); om.Package != manifest.Package {
		t.Fatalf("Export want %v, but actual %v\n", manifest.Package, om.Package)
	}

	os.Remove(output)
}
