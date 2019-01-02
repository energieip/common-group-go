package groupmodel

import (
	"encoding/json"
)

const (
	// Sensor rules when there are several sensors in the same group
	SensorAverage = "average"
	SensorMin     = "min"
	SensorMax     = "max"

	DbStatusName    = "status"
	TableStatusName = "groups"

	GroupConfigUrl = "update/settings"
)

//GroupConfig representation
type GroupConfig struct {
	Group              int      `json:"group"` //groupID
	SensorRule         *string  `json:"sensorRule,omitempty"`
	Auto               *bool    `json:"auto,omitempty"`
	SlopeStart         *int     `json:"slopeStart,omitempty"`
	SlopeStop          *int     `json:"slopeStop,omitempty"`
	Watchdog           *int     `json:"watchdog,omitempty"`
	CorrectionInterval *int     `json:"correctionInterval,omitempty"`
	RuleBrightness     *int     `json:"ruleBrightness,omitempty"`
	RulePresence       *int     `json:"rulePresence,omitempty"`
	SetpointLeds       *int     `json:"setpointLeds,omitempty"`
	FriendlyName       *string  `json:"friendlyName,omitempty"`
	Leds               []string `json:"leds"` //Mac address list
	Sensors            []string `json:"sensors"`
}

//GroupStatus status dump to the server
type GroupStatus struct {
	ID                 string   `json:"ID,omitempty"` //database id
	Group              int      `json:"group"`        //groupID
	SensorRule         string   `json:"sensorRule"`
	Auto               bool     `json:"auto"`
	SlopeStart         int      `json:"slopeStart"`
	SlopeStop          int      `json:"slopeStop"`
	Watchdog           int      `json:"watchdog"`
	CorrectionInterval int      `json:"correctionInterval"`
	RuleBrightness     *int     `json:"ruleBrightness,omitempty"`
	RulePresence       *int     `json:"rulePresence,omitempty"`
	Error              int      `json:"error"`
	TimeToAuto         int      `json:"timeToAuto"`
	SetpointLeds       int      `json:"setpointLeds"`
	Presence           bool     `json:"presence"`
	TimeToLeave        int      `json:"timeToLeave"`
	Leds               []string `json:"leds"` //Mac address list
	Sensors            []string `json:"sensors"`
	FriendlyName       string   `json:"friendlyName"`
}

// ToMapInterface convert group struct in Map[string] interface{}
func (group GroupConfig) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(group)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

// ToJSON dump group struct
func (group GroupConfig) ToJSON() (string, error) {
	inrec, err := json.Marshal(group)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

//ToGroupConfig convert interface to group config object
func ToGroupConfig(val interface{}) (*GroupConfig, error) {
	var group GroupConfig
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &group)
	return &group, err
}

// ToMapInterface convert GroupStatus struct in Map[string] interface{}
func (group GroupStatus) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(group)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

// ToJSON dump group GroupStatus
func (group GroupStatus) ToJSON() (string, error) {
	inrec, err := json.Marshal(group)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

//ToGroupStatus convert interface to status object
func ToGroupStatus(val interface{}) (*GroupStatus, error) {
	var group GroupStatus
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &group)
	return &group, err
}
