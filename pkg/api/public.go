package api

// Information contains publicly available information returned from the Bosch Smart Home Controller
type Information struct {
	APIVersions         []string `json:"apiVersions"`
	SoftwareUpdateState struct {
		Type                     string `json:"@type"`
		SwUpdateState            string `json:"swUpdateState"`
		SwUpdateLastResult       string `json:"swUpdateLastResult"`
		SwUpdateAvailableVersion string `json:"swUpdateAvailableVersion"`
		SwInstalledVersion       string `json:"swInstalledVersion"`
		SwActivationDate         struct {
			Type    string `json:"@type"`
			Timeout int    `json:"timeout"`
		} `json:"swActivationDate"`
	} `json:"softwareUpdateState"`
	Claimed        bool     `json:"claimed"`
	Country        string   `json:"country"`
	TacVersion     string   `json:"tacVersion"`
	ShcIPAddress   string   `json:"shcIpAddress"`
	ClientIds      []string `json:"clientIds"`
	FeatureToggles struct {
		AppStoreRatingIos      bool `json:"app-store-rating.ios"`
		AppLogging             bool `json:"app-logging"`
		SiriShortcuts          bool `json:"siri-shortcuts"`
		ShadingSlatsValues     bool `json:"shading.slats.values"`
		AlarmGapsNotifications bool `json:"alarm.gaps.notifications"`
		PsmPcPairing           bool `json:"psm.pc.pairing"`
		SmartLightPairing      bool `json:"smart-light.pairing"`
		WlsPairing             bool `json:"wls.pairing"`
		CamerasExtension       bool `json:"cameras.extension"`
		ShadingAdvanceMenu     bool `json:"shading.advance.menu"`
		AppStoreRatingAndroid  bool `json:"app-store-rating.android"`
		AnalyticsToggle        bool `json:"analytics.toggle"`
		WhitegoodsPairing      bool `json:"whitegoods.pairing"`
		CloudTokenvalidation   bool `json:"cloud.tokenvalidation"`
	} `json:"featureToggles"`
	ConnectivityVersions []struct {
		Name       string `json:"name"`
		MinVersion int    `json:"minVersion"`
		MaxVersion int    `json:"maxVersion"`
	} `json:"connectivityVersions"`
}

const uriInformation = ":8446/smarthome/public/information"

// Information returns publicly available information from the Bosch Smart Home Controller
func (b *boschShcAPI) Information() (i Information, e error) {
	e = b.get(uriInformation, i)

	return i, e
}
