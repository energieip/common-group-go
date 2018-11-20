package groupmodel

import (
	"encoding/json"

	"github.com/energieip/common-led-go/pkg/driverled"
	"github.com/energieip/common-sensor-go/pkg/driversensor"
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

//GroupBase
type GroupBase struct {
	Group              int     `gorethink:"group" json:"group"` //groupID
	SensorRule         *string `gorethink:"sensor_rule" json:"sensorRule"`
	Auto               *bool   `gorethink:"auto" json:"auto"`
	SlopeStart         *int    `gorethink:"slope_start" json:"slopeStart"`
	SlopeStop          *int    `gorethink:"slope_stop" json:"slopeStop"`
	Watchdog           *int    `gorethink:"watchdog" json:"watchdog"`
	CorrectionInterval *int    `gorethink:"correction_interval" json:"correctionInterval"`
	GroupRules         *Rule   `gorethink:"group_rules" json:"groupRules"`
}

//GroupConfig representation
type GroupConfig struct {
	GroupBase
	Leds    []string `gorethink:"leds" json:"leds"` //Mac address list
	Sensors []string `gorethink:"sensors" json:"sensors"`
}

// Rule when the group is in automatic mode
type Rule struct {
	Brightness *int `gorethink:"brightness" json:"brightness"`
	Presence   *int `gorethink:"presence" json:"presence"`
}

type Setpoint struct {
	SpLeds *int `gorethink:"sp_leds" json:"spLeds"`
}

//GroupRuntime runtime execution
type GroupRuntime struct {
	GroupBase
	Setpoints *Setpoint             `gorethink:"setpoints" json:"setpoints"`
	Leds      []driverled.Led       `gorethink:"leds" json:"leds"`
	Sensors   []driversensor.Sensor `gorethink:"sensors" json:"sensors"`
}

//GroupStatus status dump to the server
type GroupStatus struct {
	ID                 string   `gorethink:"id,omitempty" json:"ID"` //database if
	Group              int      `gorethink:"group" json:"group"`     //groupID
	SensorRule         string   `gorethink:"sensor_rule" json:"sensorRule"`
	Auto               bool     `gorethink:"auto" json:"auto"`
	SlopeStart         int      `gorethink:"slope_start" json:"slopeStart"`
	SlopeStop          int      `gorethink:"slope_stop" json:"slopeStop"`
	Watchdog           int      `gorethink:"watchdog" json:"watchdog"`
	CorrectionInterval int      `gorethink:"correction_interval" json:"correctionInterval"`
	GroupRules         Rule     `gorethink:"group_rules" json:"groupRules"`
	Error              int      `gorethink:"error" json:"error"`
	TimeToAuto         int      `gorethink:"time_to_auto" json:"timeToAuto"`
	SetpointLeds       int      `gorethink:"setpoint_leds" json:"setpoint_leds"`
	Presence           bool     `gorethink:"presence" json:"presence"`
	TimeToLeave        int      `gorethink:"time_to_leave" json:"timeToLeave"`
	Leds               []string `gorethink:"leds" json:"leds"` //Mac address list
	Sensors            []string `gorethink:"sensors" json:"sensors"`
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
func ToGroupConfig(val interface{}) (*GroupStatus, error) {
	var group GroupStatus
	inrec, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inrec, &group)
	return &group, err
}

// ToMapInterface convert group struct in Map[string] interface{}
func (group GroupRuntime) ToMapInterface() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(group)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

// ToJSON dump group struct
func (group GroupRuntime) ToJSON() (string, error) {
	inrec, err := json.Marshal(group)
	if err != nil {
		return "", err
	}
	return string(inrec[:]), err
}

//ToGroupRuntime convert interface to group runtime object
func ToGroupRuntime(val interface{}) (*GroupRuntime, error) {
	var group GroupRuntime
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
