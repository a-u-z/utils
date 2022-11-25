package main

import (
	"log"
	"sort"
	"strings"
)

// 國小/班級/座號
// 忠孝國小/A/25
// 忠孝國小/A/30
// 民權國小/C/10
// 民權國小/C/11
// 民權國小/F/1
// 民權國小/F/2

func stringSliceToLadderMap(s []string) (specialStruct map[string]map[string][]string) {
	sort.Strings(s)
	for i, singleS := range s {
		empIDSlice := strings.Split(singleS, "/")
		if len(empIDSlice) != 3 {
			log.Printf("here is 長度不對")
			continue
		}

		eventID := empIDSlice[0]  // "32986793-e
		marketID := empIDSlice[1] // 1_11500075
		productID := empIDSlice[2]

		m, exist := specialStruct[eventID]
		if exist {
			ps, exist := m[marketID]
			if exist {
				temp := append(ps, productID)
				specialStruct[eventID][marketID] = temp
			} else {
				specialStruct[eventID][marketID] = []string{productID}
			}
		} else {
			if i == 0 {
				specialStruct = map[string]map[string][]string{
					eventID: {
						marketID: {productID},
					},
				}
			}
			specialStruct[eventID] = map[string][]string{
				marketID: {productID},
			}
		}

	}
	return
}
