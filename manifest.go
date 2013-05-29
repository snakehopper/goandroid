package go-android-tools

import (
	"encoding/xml"
)

type Manifest struct {
	XMLName         xml.Name `xml:"manifest"`
	Package         string   `xml:"package,attr"`
	SharedUserId    string   `xml:"http://schemas.android.com/apk/res/android sharedUserId,attr,omitempty"`
	SharedUserLabel string   `xml:"http://schemas.android.com/apk/res/android sharedUserLabel,attr,omitempty"`
	VersionCode     string   `xml:"http://schemas.android.com/apk/res/android versionCode,attr"`
	VersionName     string   `xml:"http://schemas.android.com/apk/res/android versionName,attr"`
	InstallLocation string   `xml:"http://schemas.android.com/apk/res/android installLocation,attr,omitempty"`

	UsesFeatures      []UsesFeature       `xml:"uses-feature,omitempty"`
	SupportsScreens   []SupportsScreens   `xml:"supports-screens,omitempty"`
	CompatibleScreens []CompatibleScreens `xml:"compatible-screensm,omitempty"`
	SupportGlTexture  []SupportGlTexture  `xml:"supports-gl-texture,omitempty"`

	Application       Application         `xml:"application"`
	Instrumentation   []Instrumentation   `xml:"instrumentation,omitempty"`
	Permissions       []Permission        `xml:"permission,omitempty"`
	PermissionGroup   []PermissionGroup   `xml:"permission-group,omitempty"`
	PermissionTree    []PermissionTree    `xml:"permission-tree,omitempty"`
	UsesConfiguration []UsesConfiguration `xml:"uses-configuration,omitempty"`
	UsePermissions    []UsePermission     `xml:"uses-permission,omitempty"`
	UsesSdk           UsesSdk             `xml:"uses-sdk"`
}

type UsePermission struct {
	Name string `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

type Permission struct {
	Description     string `xml:"http://schemas.android.com/apk/res/android description,attr,omitempty"`
	Icon            string `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label           string `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name            string `xml:"http://schemas.android.com/apk/res/android name,attr"`
	PermissionGroup string `xml:"http://schemas.android.com/apk/res/android permissionGroup,attr,omitempty"`
	ProtectionLevel string `xml:"http://schemas.android.com/apk/res/android protectionLevel,attr,omitempty"`
}

type PermissionTree struct {
	Icon  string `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label string `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name  string `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

type PermissionGroup struct {
	Description string `xml:"http://schemas.android.com/apk/res/android description,attr,omitempty"`
	Icon        string `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label       string `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name        string `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

type Instrumentation struct {
	FunctionalTest  string `xml:"http://schemas.android.com/apk/res/android functionalTest,attr,omitempty"`
	HandleProfiling string `xml:"http://schemas.android.com/apk/res/android handleProfiling,attr,omitempty"`
	Icon            string `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label           string `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name            string `xml:"http://schemas.android.com/apk/res/android name,attr"`
	TargetPackage   string `xml:"http://schemas.android.com/apk/res/android targetPackage,attr,omitempty"`
}

type UsesSdk struct {
	MaxSdkVersion    string `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr,omitempty"`
	MinSdkVersion    string `xml:"http://schemas.android.com/apk/res/android minSdkVersion,attr,omitempty"`
	TargetSdkVersion string `xml:"http://schemas.android.com/apk/res/android targetSdkVersion,attr,omitempty"`
}

type UsesConfiguration struct {
	Resizeable              string `xml:"http://schemas.android.com/apk/res/android resizeable,attr,omitempty"`
	SmallScreens            string `xml:"http://schemas.android.com/apk/res/android smallScreens,attr,omitempty"`
	NormalScreens           string `xml:"http://schemas.android.com/apk/res/android normalScreens,attr,omitempty"`
	LargeScreens            string `xml:"http://schemas.android.com/apk/res/android largeScreens,attr,omitempty"`
	XlargeScreens           string `xml:"http://schemas.android.com/apk/res/android xlargeScreens,attr,omitempty"`
	AnyDensity              string `xml:"http://schemas.android.com/apk/res/android anyDensity,attr,omitempty"`
	RequiresSmallestWidthDp string `xml:"http://schemas.android.com/apk/res/android requiresSmallestWidthDp,attr,omitempty"`
	CompatibleWidthLimitDp  string `xml:"http://schemas.android.com/apk/res/android compatibleWidthLimitDp,attr,omitempty"`
	LargestWidthLimitDp     string `xml:"http://schemas.android.com/apk/res/android largestWidthLimitDp,attr,omitempty"`
}

type UsesFeature struct {
	GlEsVersion string `xml:"http://schemas.android.com/apk/res/android glEsVersion,attr,omitempty"`
	Name        string `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Required    string `xml:"http://schemas.android.com/apk/res/android required,attr,omitempty"`
}

type SupportsScreens struct {
	Resizeable              string `xml:"http://schemas.android.com/apk/res/android resizeable,attr,omitempty"`
	SmallScreens            string `xml:"http://schemas.android.com/apk/res/android smallScreens,attr,omitempty"`
	NormalScreens           string `xml:"http://schemas.android.com/apk/res/android normalScreens,attr,omitempty"`
	LargeScreens            string `xml:"http://schemas.android.com/apk/res/android largeScreens,attr,omitempty"`
	XlargeScreens           string `xml:"http://schemas.android.com/apk/res/android xlargeScreens,attr,omitempty"`
	AnyDensity              string `xml:"http://schemas.android.com/apk/res/android anyDensity,attr,omitempty"`
	RequiresSmallestWidthDp string `xml:"http://schemas.android.com/apk/res/android requiresSmallestWidthDp,attr,omitempty"`
	CompatibleWidthLimitDp  string `xml:"http://schemas.android.com/apk/res/android compatibleWidthLimitDp,attr,omitempty"`
	LargestWidthLimitDp     string `xml:"http://schemas.android.com/apk/res/android largestWidthLimitDp,attr,omitempty"`
}

type CompatibleScreens struct {
	screen Screen `xml:"screen,omitempty"`
}

type Screen struct {
	screenSize    string `xml:"http://schemas.android.com/apk/res/android screenSize,attr,omitempty"`
	screenDensity string `xml:"http://schemas.android.com/apk/res/android screenDensity,attr,omitempty"`
}

type SupportGlTexture struct {
	name string `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}

type Application struct {
	AllowBackup             string          `xml:"http://schemas.android.com/apk/res/android allowBackup,attr,omitempty"`
	AllowClearUserData      string          `xml:"http://schemas.android.com/apk/res/android allowClearUserData,attr,omitempty"`
	AllowTaskReparenting    string          `xml:"http://schemas.android.com/apk/res/android allowTaskReparenting,attr,omitempty"`
	BackupAgent             string          `xml:"http://schemas.android.com/apk/res/android backupAgent,attr,omitempty"`
	Debuggable              string          `xml:"http://schemas.android.com/apk/res/android debuggable,attr,omitempty"`
	Description             string          `xml:"http://schemas.android.com/apk/res/android description,attr,omitempty"`
	Enabled                 string          `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	HardwareAccelerated     string          `xml:"http://schemas.android.com/apk/res/android hardwareAccelerated,attr,omitempty"`
	HasCode                 string          `xml:"http://schemas.android.com/apk/res/android hasCode,attr,omitempty"`
	Icon                    string          `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	KillAfterRestore        string          `xml:"http://schemas.android.com/apk/res/android killAfterRestore,attr,omitempty"`
	Label                   string          `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	LargeHeap               string          `xml:"http://schemas.android.com/apk/res/android largeHeap,attr,omitempty"`
	Logo                    string          `xml:"http://schemas.android.com/apk/res/android logo,attr,omitempty"`
	ManageSpaceActivity     string          `xml:"http://schemas.android.com/apk/res/android manageSpaceActivity,attr,omitempty"`
	Name                    string          `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Permission              string          `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	Persistent              string          `xml:"http://schemas.android.com/apk/res/android persistent,attr,omitempty"`
	Process                 string          `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	RestoreAnyVersion       string          `xml:"http://schemas.android.com/apk/res/android restoreAnyVersion,attr,omitempty"`
	RestoreNeedsApplication string          `xml:"http://schemas.android.com/apk/res/android restoreNeedsApplication,attr,omitempty"`
	SupportsRtl             string          `xml:"http://schemas.android.com/apk/res/android supportsRtl,attr,omitempty"`
	TaskAffinity            string          `xml:"http://schemas.android.com/apk/res/android taskAffinity,attr,omitempty"`
	TestOnly                string          `xml:"http://schemas.android.com/apk/res/android testOnly,attr,omitempty"`
	Theme                   string          `xml:"http://schemas.android.com/apk/res/android theme,attr,omitempty"`
	UiOptions               string          `xml:"http://schemas.android.com/apk/res/android uiOptions,attr,omitempty"`
	VmSafeMode              string          `xml:"http://schemas.android.com/apk/res/android vmSafeMode,attr,omitempty"`
	Activities              []Activity      `xml:"activity,omitempty"`
	ActivityAlias           []ActivityAlias `xml:"activity-alias,omitempty"`
	Service                 []Service       `xml:"service,omitempty"`
	Receiver                []Receiver      `xml:"receiver,omitempty"`
	Providers               []Provider      `xml:"provider,omitempty"`
	UsesLibrary             []UsesLibrary   `xml:"uses-library,omitempty"`
	MetaData                []MetaData      `xml:"meta-data,omitempty"`
}

type Activity struct {
	AllowTaskReparenting       string         `xml:"http://schemas.android.com/apk/res/android allowTaskReparenting,attr,omitempty"`
	AlwaysRetainTaskState      string         `xml:"http://schemas.android.com/apk/res/android alwaysRetainTaskState,attr,omitempty"`
	ClearTaskOnLaunch          string         `xml:"http://schemas.android.com/apk/res/android clearTaskOnLaunch,attr,omitempty"`
	ConfigChanges              string         `xml:"http://schemas.android.com/apk/res/android configChanges,attr,omitempty"`
	Description                string         `xml:"http://schemas.android.com/apk/res/android description,attr,omitempty"`
	Enabled                    string         `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	ExcludeFromRecents         string         `xml:"http://schemas.android.com/apk/res/android excludeFromRecents,attr,omitempty"`
	Exported                   string         `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	FinishOnCloseSystemDialogs string         `xml:"http://schemas.android.com/apk/res/android finishOnCloseSystemDialogs,attr,omitempty"`
	FinishOnTaskLaunch         string         `xml:"http://schemas.android.com/apk/res/android finishOnTaskLaunch,attr,omitempty"`
	HardwareAccelerated        string         `xml:"http://schemas.android.com/apk/res/android hardwareAccelerated,attr,omitempty"`
	Icon                       string         `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Immersive                  string         `xml:"http://schemas.android.com/apk/res/android immersive,attr,omitempty"`
	Label                      string         `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	LaunchMode                 string         `xml:"http://schemas.android.com/apk/res/android launchMode,attr,omitempty"`
	Logo                       string         `xml:"http://schemas.android.com/apk/res/android logo,attr,omitempty"`
	Multiprocess               string         `xml:"http://schemas.android.com/apk/res/android multiprocess,attr,omitempty"`
	Name                       string         `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	NoHistory                  string         `xml:"http://schemas.android.com/apk/res/android noHistory,attr,omitempty"`
	ParentActivityName         string         `xml:"http://schemas.android.com/apk/res/android parentActivityName,attr,omitempty"`
	Permission                 string         `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	Process                    string         `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	ScreenOrientation          string         `xml:"http://schemas.android.com/apk/res/android screenOrientation,attr,omitempty"`
	ShowOnLockScreen           string         `xml:"http://schemas.android.com/apk/res/android showOnLockScreen,attr,omitempty"`
	SingleUser                 string         `xml:"http://schemas.android.com/apk/res/android singleUser,attr,omitempty"`
	StateNotNeeded             string         `xml:"http://schemas.android.com/apk/res/android stateNotNeeded,attr,omitempty"`
	TaskAffinity               string         `xml:"http://schemas.android.com/apk/res/android taskAffinity,attr,omitempty"`
	Theme                      string         `xml:"http://schemas.android.com/apk/res/android theme,attr,omitempty"`
	UiOptions                  string         `xml:"http://schemas.android.com/apk/res/android uiOptions,attr,omitempty"`
	WindowSoftInputMode        string         `xml:"http://schemas.android.com/apk/res/android windowSoftInputMode,attr,omitempty"`
	IntentFilters              []IntentFilter `xml:"intent-filter,omitempty"`
	MetaDatas                  []MetaData     `xml:"meta-data,omitempty"`
}

type ActivityAlias struct {
	Enabled        string         `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	Exported       string         `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	Icon           string         `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label          string         `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name           string         `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Permission     string         `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	TargetActivity string         `xml:"http://schemas.android.com/apk/res/android targetActivity,attr,omitempty"`
	IntentFilters  []IntentFilter `xml:"intent-filter,omitempty"`
	MetaDatas      []MetaData     `xml:"meta-data,omitempty"`
}

type Service struct {
	Enabled         string         `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	Exported        string         `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	Icon            string         `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	IsolatedProcess string         `xml:"http://schemas.android.com/apk/res/android isolatedProcess,attr,omitempty"`
	Label           string         `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name            string         `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Permission      string         `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	Process         string         `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	IntentFilters   []IntentFilter `xml:"intent-filter,omitempty"`
	MetaDatas       []MetaData     `xml:"meta-data,omitempty"`
}

type Receiver struct {
	Enabled       string         `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	Exported      string         `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	Icon          string         `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label         string         `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Name          string         `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Permission    string         `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	Process       string         `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	IntentFilters []IntentFilter `xml:"intent-filter,omitempty"`
	MetaDatas     []MetaData     `xml:"meta-data,omitempty"`
}

type Provider struct {
	Authorities         string              `xml:"http://schemas.android.com/apk/res/android authorities,attr,omitempty"`
	Enabled             string              `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	Exported            string              `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	GrantUriPermissions string              `xml:"http://schemas.android.com/apk/res/android grantUriPermissions,attr,omitempty"`
	Icon                string              `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	InitOrder           string              `xml:"http://schemas.android.com/apk/res/android initOrder,attr,omitempty"`
	Label               string              `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Multiprocess        string              `xml:"http://schemas.android.com/apk/res/android multiprocess,attr,omitempty"`
	Name                string              `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Permission          string              `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	Process             string              `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	ReadPermission      string              `xml:"http://schemas.android.com/apk/res/android readPermission,attr,omitempty"`
	Syncable            string              `xml:"http://schemas.android.com/apk/res/android syncable,attr,omitempty"`
	WritePermission     string              `xml:"http://schemas.android.com/apk/res/android writePermission,attr,omitempty"`
	MetaDatas           *MetaData           `xml:"meta-data,omitempty,omitempty"`
	GrantUriPermission  *GrantUriPermission `xml:"grant-uri-permission,omitempty"`
	PathPermission      *PathPermission     `xml:"path-permission,omitempty"`
}

type UsesLibrary struct {
	Name     string `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Required string `xml:"http://schemas.android.com/apk/res/android required,attr,omitempty"`
}

type IntentFilter struct {
	Icon     string    `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label    string    `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	Priority string    `xml:"http://schemas.android.com/apk/res/android priority,attr,omitempty"`
	Action   Action    `xml:"action"`
	Category *Category `xml:"category,omitempty"`
	Data     *Data     `xml:"data,omitempty"`
}

type Action struct {
	Name string `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}

type Category struct {
	Name string `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}

type Data struct {
	Host        string `xml:"http://schemas.android.com/apk/res/android host,attr,omitempty"`
	MimeType    string `xml:"http://schemas.android.com/apk/res/android mimeType,attr,omitempty"`
	Path        string `xml:"http://schemas.android.com/apk/res/android path,attr,omitempty"`
	PathPattern string `xml:"http://schemas.android.com/apk/res/android pathPattern,attr,omitempty"`
	PathPrefix  string `xml:"http://schemas.android.com/apk/res/android pathPrefix,attr,omitempty"`
	Port        string `xml:"http://schemas.android.com/apk/res/android port,attr,omitempty"`
	Scheme      string `xml:"http://schemas.android.com/apk/res/android scheme,attr,omitempty"`
}

type GrantUriPermission struct {
	Path        string `xml:"http://schemas.android.com/apk/res/android path,attr,omitempty"`
	PathPattern string `xml:"http://schemas.android.com/apk/res/android pathPattern,attr,omitempty"`
	PathPrefix  string `xml:"http://schemas.android.com/apk/res/android pathPrefix,attr,omitempty"`
}

type PathPermission struct {
	Path            string `xml:"http://schemas.android.com/apk/res/android path,attr,omitempty"`
	PathPattern     string `xml:"http://schemas.android.com/apk/res/android pathPattern,attr,omitempty"`
	PathPrefix      string `xml:"http://schemas.android.com/apk/res/android pathPrefix,attr,omitempty"`
	Permission      string `xml:"http://schemas.android.com/apk/res/android permission,attr,omitempty"`
	ReadPermission  string `xml:"http://schemas.android.com/apk/res/android readPermission,attr,omitempty"`
	WritePermission string `xml:"http://schemas.android.com/apk/res/android writePermission,attr,omitempty"`
}

type MetaData struct {
	Name     string `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Resource string `xml:"http://schemas.android.com/apk/res/android resource,attr,omitempty"`
	Value    string `xml:"http://schemas.android.com/apk/res/android value,attr,omitempty"`
}
