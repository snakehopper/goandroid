package goandroid

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	sync *SyncAdapter

	sync_path string
)

func setup_sync() {
	sync_path = "material/sync.xml"

	b := []byte("<sync-adapter xmlns:android=\"http://schemas.android.com/apk/res/android\" android:contentAuthority=\"com.google.android.apps.iosched\" android:accountType=\"com.google\" android:supportsUploading=\"false\" android:userVisible=\"true\"/>")
	ioutil.WriteFile(sync_path, b, 0644)

	sync, _ = NewSyncAdapter(sync_path)
}

func teardown_sync() {
	os.Remove(sync_path)
}

func TestNewSyncAdapter(t *testing.T) {
	setup_sync()

	r, err := NewSyncAdapter(sync_path)
	if err != nil {
		t.Fatal("read resources", err)
	}

	var want = "sync-adapter"
	if want != r.XMLName.Local {
		t.Fatalf("NewSyncAdapter XMLName want %v, but %v\n", want, r.XMLName.Local)
	}

	teardown_sync()
}

func TestSyncAdapter_Export(t *testing.T) {
	setup_sync()

	output := "material/sync_o.xml"
	if err := sync.Export(output); err != nil {
		t.Fatal("Export error: ", err)
	}

	_, err := os.Stat(output)
	if err != nil {
		t.Fatal("Export error:", err)
	}
	os.Remove(output)
	teardown_sync()
}
