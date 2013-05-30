package goandroid

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

//http://developer.android.com/reference/android/content/AbstractThreadedSyncAdapter.html
/*
<sync-adapter xmlns:android="http://schemas.android.com/apk/res/android"
    android:contentAuthority="authority"
    android:accountType="accountType"
    android:userVisible="true|false"
    android:supportsUploading="true|false"
    android:allowParallelSyncs="true|false"
    android:isAlwaysSyncable="true|false"
    android:syncAdapterSettingsAction="ACTION_OF_SETTINGS_ACTIVITY"
 />
*/
type SyncAdapter struct {
	XMLName                   xml.Name `xml:"sync-adapter"`
	ContentAuthority          string   `xml:"http://schemas.android.com/apk/res/android contentAuthority,attr,omitempty"`
	AccountType               string   `xml:"http://schemas.android.com/apk/res/android accountType,attr,omitempty"`
	UserVisible               string   `xml:"http://schemas.android.com/apk/res/android userVisible,attr,omitempty"`
	SupportsUploading         string   `xml:"http://schemas.android.com/apk/res/android supportsUploading,attr,omitempty"`
	AllowParallelSyncs        string   `xml:"http://schemas.android.com/apk/res/android allowParallelSyncs,attr,omitempty"`
	IsAlwaysSyncable          string   `xml:"http://schemas.android.com/apk/res/android isAlwaysSyncable,attr,omitempty"`
	SyncAdapterSettingsAction string   `xml:"http://schemas.android.com/apk/res/android syncAdapterSettingsAction,attr,omitempty"`
}

func NewSyncAdapter(filepath string) (*SyncAdapter, error) {
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
	var res SyncAdapter
	err = xml.Unmarshal(b, &res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, err
}

func (res *SyncAdapter) ExportSyncAdapter(filepath string) error {
	f, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if f.IsDir() {
		log.Fatal(filepath, " is not a valid file")
	}

	b, err := xml.MarshalIndent(&res, "  ", "    ")
	//b, err := xml.Marshal(&res)
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
