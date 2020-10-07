package parser

import (
	"app/common"
	"app/orchestration/trigger"
	"errors"
)

// GenerateTriggerFromDict creates a trigger object from the dict generated from yml file
func GenerateTriggerFromDict(dict common.Dict) (*trigger.Trigger, error) {
	if err := triggerDictIsNotValid(dict); err != nil {
		return nil, err
	}

	triggerDict := dict["trigger"].(common.Dict)
	trigger := trigger.Trigger{
		Type: triggerDict["type"].(string),
		Info: triggerDict["info"].(string),
	}

	return &trigger, nil
}

func triggerDictIsNotValid(dict common.Dict) error {
	trigger, exist := dict["trigger"]

	if !exist {
		return errors.New("symphony doesn't have a trigger definition")
	}

	triggerDict := trigger.(common.Dict)

	// TODO: validate if type is valid and all info are strings

	if _, exist := triggerDict["type"]; !exist {
		return errors.New("type field not found")
	}

	if _, exist := triggerDict["info"]; !exist {
		return errors.New("info field not found")
	}

	return nil
}
