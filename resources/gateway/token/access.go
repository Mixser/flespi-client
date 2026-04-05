package flespi_token

import "encoding/json"

// AccessType represents the token access mode.
type AccessType int

const (
	AccessTypeStandard AccessType = 0
	AccessTypeMaster   AccessType = 1
	AccessTypeACL      AccessType = 2
)

// TokenAccess represents the token access field.
// Type 0 = Standard, Type 1 = Master, Type 2 = ACL (requires ACL field).
type TokenAccess struct {
	Type AccessType `json:"type"`
	ACL  []ACE      `json:"acl,omitempty"`
}

func MasterAccess() TokenAccess   { return TokenAccess{Type: AccessTypeMaster} }
func StandardAccess() TokenAccess { return TokenAccess{Type: AccessTypeStandard} }
func ACLAccess(acl []ACE) TokenAccess {
	return TokenAccess{Type: AccessTypeACL, ACL: acl}
}

// ACEIDs represents the "ids" field of an ACE: "all" | "in-groups" | []int64.
type ACEIDs struct {
	all      bool
	inGroups bool
	list     []int64
}

var (
	ACEIDsAll      = ACEIDs{all: true}
	ACEIDsInGroups = ACEIDs{inGroups: true}
)

func ACEIDsList(ids ...int64) ACEIDs { return ACEIDs{list: ids} }

func (a ACEIDs) MarshalJSON() ([]byte, error) {
	switch {
	case a.all:
		return json.Marshal("all")
	case a.inGroups:
		return json.Marshal("in-groups")
	default:
		return json.Marshal(a.list)
	}
}

func (a *ACEIDs) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		switch s {
		case "all":
			a.all = true
		case "in-groups":
			a.inGroups = true
		}
		return nil
	}
	return json.Unmarshal(data, &a.list)
}

// ACE is a single access control entry, discriminated by URI.
type ACE struct {
	URI        string      `json:"uri"`
	Methods    []string    `json:"methods,omitempty"`
	IDs        *ACEIDs     `json:"ids,omitempty"`
	Submodules []Submodule `json:"submodules,omitempty"`
	// mqtt-specific fields
	Topic   string   `json:"topic,omitempty"`
	Actions []string `json:"actions,omitempty"`
}

// Submodule is a named submodule entry within an ACE.
type Submodule struct {
	Name    string   `json:"name"`
	Methods []string `json:"methods,omitempty"`
}

// ACE URI constants
const (
	ACEURIGwChannels        = "gw/channels"
	ACEURIGwDevices         = "gw/devices"
	ACEURIGwGroups          = "gw/groups"
	ACEURIGwStreams         = "gw/streams"
	ACEURIGwModems          = "gw/modems"
	ACEURIGwCalcs           = "gw/calcs"
	ACEURIGwPlugins         = "gw/plugins"
	ACEURIGwGeofences       = "gw/geofences"
	ACEURIGwAssets          = "gw/assets"
	ACEURIStorageContainers = "storage/containers"
	ACEURIStorageCDNs       = "storage/cdns"
	ACEURIMqtt              = "mqtt"
	ACEURIAI                = "ai"
)

// Submodule name constants
const (
	SubmoduleLogs           = "logs"
	SubmoduleMessages       = "messages"
	SubmoduleTelemetry      = "telemetry"
	SubmoduleSettings       = "settings"
	SubmoduleConnections    = "connections"
	SubmoduleDevices        = "devices"
	SubmoduleChannels       = "channels"
	SubmoduleGroups         = "groups"
	SubmoduleGeofences      = "geofences"
	SubmoduleCalculate      = "calculate"
	SubmoduleMedia          = "media"
	SubmodulePackets        = "packets"
	SubmoduleCommands       = "commands"
	SubmoduleCommandsQueue  = "commands-queue"
	SubmoduleCommandsResult = "commands-result"
	SubmoduleSMS            = "sms"
	SubmoduleFiles          = "files"
	SubmoduleIntervals      = "intervals"
	SubmoduleIdents         = "idents"
)

// NewACE creates a basic ACE for the given URI and methods.
func NewACE(uri string, methods []string) ACE {
	return ACE{URI: uri, Methods: methods}
}

// NewACEWithIDs creates an ACE with an IDs restriction.
func NewACEWithIDs(uri string, methods []string, ids ACEIDs) ACE {
	return ACE{URI: uri, Methods: methods, IDs: &ids}
}

// NewACEWithSubmodules creates an ACE with IDs and submodule restrictions.
func NewACEWithSubmodules(uri string, methods []string, ids ACEIDs, submodules []Submodule) ACE {
	return ACE{URI: uri, Methods: methods, IDs: &ids, Submodules: submodules}
}

// NewMQTTACE creates an ACE for the mqtt URI.
func NewMQTTACE(topic string, actions []string, methods []string) ACE {
	return ACE{URI: ACEURIMqtt, Topic: topic, Actions: actions, Methods: methods}
}
