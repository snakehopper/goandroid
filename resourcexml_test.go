package goandroid

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	res *Resources

	res_path string
)

func setup_res() {
	res_path = "material/str.xml"

	b := []byte("<resources>\n<string name=\"app_name\">launcher</string>\n<string name=\"hello\">hello world</string>\n<color name=\"test_color\">#FFFFFF</color>\n</resources>")
	ioutil.WriteFile(res_path, b, 0644)

	res, _ = NewResourceXml(res_path)
}

func teardown_res() {
	os.Remove(res_path)
}

func TestNewResourceXml(t *testing.T) {
	setup_res()

	r, err := NewResourceXml(res_path)
	if err != nil {
		t.Fatal("read resources", err)
	}

	var want = "resources"
	if want != r.XMLName.Local {
		t.Fatalf("NewResourceXml XMLName want %v, but %v\n", want, r.XMLName.Local)
	}

	teardown_res()
}

func TestResources_SetString(t *testing.T) {
	setup_res()

	var want = "com.example.test.name"

	res = res.SetString("app_name", want)

	for _, r := range res.Strings {
		if r.Name == "app_name" && r.Value != want {
			t.Fatalf("SetString want %v, but %v\n", want, r.Value)
		}
	}

	teardown_res()
}

func TestResources_SetColor(t *testing.T) {
	setup_res()

	var want = "#000000"

	res = res.SetColor("test_color", want)

	for _, r := range res.Colors {
		if r.Name == "test_color" && r.Value != want {
			t.Fatalf("SetColor want %v, but %v\n", want, r.Value)
		}
	}

	teardown_res()
}

func TestResources_AddString(t *testing.T) {
	setup_res()

	actual, want := len(res.Strings), len(res.Strings)+1

	res = res.AddString("another_string", "example")
	actual = len(res.Strings)

	if actual != want {
		t.Fatalf("AddString want %v, but %v\n", want, actual)
	}

	teardown_res()
}

func TestManifest_ExportResourceXml(t *testing.T) {
	setup_res()

	output := "material/str_o.xml"
	if err := res.ExportResourceXml(output); err != nil {
		t.Fatal("ExportResourceXml error: ", err)
	}

	_, err := os.Stat(output)
	if err != nil {
		t.Fatal("ExportResourceXml error:", err)
	}
	os.Remove(output)
	teardown_res()
}
