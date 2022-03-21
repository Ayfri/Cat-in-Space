package main

import (
	"sort"
	"strings"
)

func Is500Error(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "500") {
		return true
	}
	return false
}

func IndexOf(array []string, value string) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func SortStreamersByLivingThenList(streamers []UserData, nameList []string) []UserData {
	var livingStreamers []UserData
	var notLivingStreamers []UserData

	for _, streamer := range streamers {
		if streamer.IsLive {
			livingStreamers = append(livingStreamers, streamer)
		} else {
			notLivingStreamers = append(notLivingStreamers, streamer)
		}
	}

	sort.Slice(livingStreamers, func(i, j int) bool {
		return IndexOf(nameList, livingStreamers[i].Login) < IndexOf(nameList, livingStreamers[j].Login)
	})

	sort.Slice(notLivingStreamers, func(i, j int) bool {
		return IndexOf(nameList, notLivingStreamers[i].Login) < IndexOf(nameList, notLivingStreamers[j].Login)
	})

	return append(livingStreamers, notLivingStreamers...)
}
