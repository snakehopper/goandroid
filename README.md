goandroid
================

A convenient struct interface for unmarshal AndroidMainifest.xml with standard golang pkg "encoding/xml"

simple usage:
	b, _ := ioutil.ReadFile("AndroidManifest.xml")
        var m Manifest
        err = xml.Unmarshal(b, &manifest)
	// check error & and can use now
	fmt.Println("version name:", m.VersionName)
	fmt.Println("package name:", m.Package)
	fmt.Println("an action name of activity:", m.Application.Activities[index].IntentFilters[index].Action.Name)

