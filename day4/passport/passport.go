package passport

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	byrMin   = 1920
	byrMax   = 2002
	iyrMin   = 2010
	iyrMax   = 2020
	eyrMin   = 2020
	eyrMax   = 2030
	hgtCMMin = 150
	hgtCMMax = 193
	hgtInMin = 59
	hgtInMax = 76
)

var (
	hclExp = regexp.MustCompile(`^#[a-f0-9]{6}$`)
	eclExp = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	pidExp = regexp.MustCompile(`^\d{9}$`)
)

// Passport stores fields that can be validated based on the challenge rules.
type Passport struct {
	byr *YearField   // Birth Year
	iyr *YearField   // Issue Year
	eyr *YearField   // Expiration Year
	hgt *HeightField // Height
	hcl *TextField   // Hair Color
	ecl *TextField   // Eye Color
	pid *TextField   // Passport ID
}

// YearField stores a year value that must exist within a provided date range.
type YearField struct {
	value int
	min   int
	max   int
}

// NewYearField creates a YearField with the specified min and max year values.
func NewYearField(year string, min, max int) (*YearField, error) {
	i, err := strconv.Atoi(year)
	if err != nil {
		return nil, err
	}
	return &YearField{
		value: i,
		min:   min,
		max:   max,
	}, nil
}

func (y YearField) isValid() bool {
	return y.value <= y.max && y.value >= y.min
}

// HeightField stores a height value that must be between the min and max values based on the measurement unit.
type HeightField struct {
	value int
	unit  string
	min   int
	max   int
}

// NewHeightField creates a HeightField with the min and max select by measurement unit.
func NewHeightField(height, unit string) (*HeightField, error) {
	var h HeightField
	i, err := strconv.Atoi(height)
	if err != nil {
		return nil, err
	}
	h.value = i

	switch unit {
	case "cm":
		h.min = hgtCMMin
		h.max = hgtCMMax
	case "in":
		h.min = hgtInMin
		h.max = hgtInMax
	default:
		return nil, fmt.Errorf("cannot use \"%s\" as a unit for height", unit)
	}

	return &h, nil
}

func (h HeightField) isValid() bool {
	return h.value <= h.max && h.value >= h.min
}

// TextField stores a text value that must match a provided regexp.
type TextField struct {
	value string
	exp   *regexp.Regexp
}

// NewTextField creates a TextField that must match the provided regexp to be valid.
func NewTextField(value string, exp *regexp.Regexp) (*TextField, error) {
	if exp == nil {
		return nil, errors.New("regexp must be defined for TextField")
	}
	return &TextField{
		value: value,
		exp:   exp,
	}, nil
}

func (t TextField) isValid() bool {
	return t.exp.MatchString(t.value)
}

// ValidateStrict checks for matching fields and validates correct formatting for each match.
func (p *Passport) ValidateStrict(input string) {
	byr := regexp.MustCompile(`byr:(\d*)`)
	if byr.MatchString(input) {
		matches := byr.FindStringSubmatch(input)
		f, err := NewYearField(matches[1], byrMin, byrMax)
		if err == nil {
			p.byr = f
		}
	}

	iyr := regexp.MustCompile(`iyr:(\d*)`)
	if iyr.MatchString(input) {
		matches := iyr.FindStringSubmatch(input)
		f, err := NewYearField(matches[1], iyrMin, iyrMax)
		if err == nil {
			p.iyr = f
		}
	}

	eyr := regexp.MustCompile(`eyr:(\d*)`)
	if eyr.MatchString(input) {
		matches := eyr.FindStringSubmatch(input)
		f, err := NewYearField(matches[1], eyrMin, eyrMax)
		if err == nil {
			p.eyr = f
		}
	}

	hgt := regexp.MustCompile(`hgt:(\d*)(cm|in)?`)
	if hgt.MatchString(input) {
		matches := hgt.FindStringSubmatch(input)
		f, err := NewHeightField(matches[1], matches[2])
		if err == nil {
			p.hgt = f
		}
	}

	hcl := regexp.MustCompile(`hcl:(#[a-f0-9]*)`)
	if hcl.MatchString(input) {
		matches := hcl.FindStringSubmatch(input)
		f, err := NewTextField(matches[1], hclExp)
		if err == nil {
			p.hcl = f
		}
	}

	ecl := regexp.MustCompile(`ecl:(\w*)`)
	if ecl.MatchString(input) {
		matches := ecl.FindStringSubmatch(input)
		f, err := NewTextField(matches[1], eclExp)
		if err == nil {
			p.ecl = f
		}
	}

	pid := regexp.MustCompile(`pid:(\d*)`)
	if pid.MatchString(input) {
		matches := pid.FindStringSubmatch(input)
		f, err := NewTextField(matches[1], pidExp)
		if err == nil {
			p.pid = f
		}
	}
}

// ValidateSimple checks for matching fields in the given input string and updates the passport for each match.
func (p *Passport) ValidateSimple(input string) {
	if strings.Contains(input, "byr:") {
		p.byr = &YearField{}
	}
	if strings.Contains(input, "iyr:") {
		p.iyr = &YearField{}
	}
	if strings.Contains(input, "eyr:") {
		p.eyr = &YearField{}
	}
	if strings.Contains(input, "hgt:") {
		p.hgt = &HeightField{}
	}
	if strings.Contains(input, "hcl:") {
		p.hcl = &TextField{}
	}
	if strings.Contains(input, "ecl:") {
		p.ecl = &TextField{}
	}
	if strings.Contains(input, "pid:") {
		p.pid = &TextField{}
	}
}

// HasAllRequiredFields checks if the passport has all of the required fields
func (p *Passport) HasAllRequiredFields() bool {
	return p.byr != nil &&
		p.iyr != nil &&
		p.eyr != nil &&
		p.hgt != nil &&
		p.hcl != nil &&
		p.ecl != nil &&
		p.pid != nil
}

// HasAllValidFields checks if the passport has all of the required fields
func (p *Passport) HasAllValidFields() bool {
	return p.byr != nil && p.byr.isValid() &&
		p.iyr != nil && p.iyr.isValid() &&
		p.eyr != nil && p.eyr.isValid() &&
		p.hgt != nil && p.hgt.isValid() &&
		p.hcl != nil && p.hcl.isValid() &&
		p.ecl != nil && p.ecl.isValid() &&
		p.pid != nil && p.pid.isValid()
}

// Rules from the challenge:
// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
