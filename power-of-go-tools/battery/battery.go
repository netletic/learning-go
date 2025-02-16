package battery

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

type Battery struct {
	Name             string
	ID               int64
	ChargePercent    int
	TimeToFullCharge string
	Present          bool
}

func (batt *Battery) ToJSON() string {
	data, err := json.MarshalIndent(batt, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func GetStatus() (Status, error) {
	text, err := GetPmsetOutput()
	if err != nil {
		return Status{}, nil
	}
	return ParsePmsetOutput(text)
}

func GetPmsetOutput() (string, error) {
	data, err := exec.Command("pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var pmsetOutput = regexp.MustCompile(`(\d{1,3})%`)

func ParsePmsetOutput(text string) (Status, error) {
	matches := pmsetOutput.FindStringSubmatch(text)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", text)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{
		ChargePercent: charge,
	}, nil
}
