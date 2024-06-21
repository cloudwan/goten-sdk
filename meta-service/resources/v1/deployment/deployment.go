package deployment

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DbDataVersion []int

func ParseAsDbDataVersion(dataVersion string) (DbDataVersion, error) {
	var verRegexp = regexp.MustCompile("^v[0-9.]{1,64}$")
	if !verRegexp.MatchString(dataVersion) {
		return nil, fmt.Errorf("value %s does not match regexp v[0-9.]{1,64}", dataVersion)
	}
	versionNumbers := strings.Split(dataVersion[1:], ".")
	verNumbers := make([]int, 0, len(versionNumbers))
	for _, numStr := range versionNumbers {
		if numStr == "" {
			verNumbers = append(verNumbers, 0)
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(fmt.Errorf("error parsing %s as integer", numStr))
		}
		verNumbers = append(verNumbers, num)
	}
	return verNumbers, nil
}

func CompareDbDataVersions(left, right DbDataVersion) int {
	for e := 0; e < len(left) && e < len(right); e++ {
		if left[e] < right[e] {
			return -1
		} else if left[e] > right[e] {
			return 1
		}
	}
	if len(left) < len(right) {
		return -1
	} else if len(left) > len(right) {
		return 1
	}
	return 0
}

func IsDbDataVersionSmaller(left, right DbDataVersion) bool {
	return CompareDbDataVersions(left, right) < 0
}
