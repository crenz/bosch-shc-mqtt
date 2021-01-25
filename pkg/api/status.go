package api

// Rooms describes the rooms configured by the user
type Rooms []struct {
	Type   string `json:"@type"`
	ID     string `json:"id"`
	IconID string `json:"iconId"`
	Name   string `json:"name"`
}

// Devices describes available devices
type Devices []struct {
	Type             string        `json:"@type"`
	RootDeviceID     string        `json:"rootDeviceId"`
	ID               string        `json:"id"`
	DeviceServiceIds []string      `json:"deviceServiceIds"`
	Manufacturer     string        `json:"manufacturer"`
	RoomID           string        `json:"roomId,omitempty"`
	DeviceModel      string        `json:"deviceModel"`
	Serial           string        `json:"serial,omitempty"`
	Profile          string        `json:"profile,omitempty"`
	Name             string        `json:"name"`
	Status           string        `json:"status"`
	ChildDeviceIds   []interface{} `json:"childDeviceIds"`
	IconID           string        `json:"iconId,omitempty"`
	ParentDeviceID   string        `json:"parentDeviceId,omitempty"`
}

// Messages describes available messages
type Messages []struct {
	Type        string `json:"@type"`
	ID          string `json:"id"`
	MessageCode struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	} `json:"messageCode"`
	SourceType string   `json:"sourceType"`
	SourceID   string   `json:"sourceId"`
	Timestamp  int64    `json:"timestamp"`
	Flags      []string `json:"flags"`
	Arguments  struct {
		Version string `json:"version"`
	} `json:"arguments"`
}

// DoorWindow describes a single door or window within OpenDoorsWindows
type DoorWindow struct {
	Name     string `json:"name"`
	RoomName string `json:"roomName"`
}

// OpenDoorsWindows describes all doors and windows and those that are open
type OpenDoorsWindows struct {
	AllDoors       []DoorWindow `json:"allDoors"`
	OpenDoors      []DoorWindow `json:"openDoors"`
	UnknownDoors   []DoorWindow `json:"unknownDoors"`
	AllWindows     []DoorWindow `json:"allWindows"`
	OpenWindows    []DoorWindow `json:"openWindows"`
	UnknownWindows []DoorWindow `json:"unknownWindows"`
	AllOthers      []DoorWindow `json:"allOthers"`
	OpenOthers     []DoorWindow `json:"openOthers"`
	UnknownOthers  []DoorWindow `json:"unknownOthers"`
}

const uriRooms = ":8444/smarthome/rooms"
const uriDevices = ":8444/smarthome/devices"
const uriMessages = ":8444/smarthome/messages"

// Rooms returns the rooms configured by the user. (Requires client authentication)
func (b *boschShcAPI) Rooms() (r Rooms, e error) {
	e = b.get(uriRooms, &r)

	return r, e
}
