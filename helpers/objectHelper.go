package helper

import "encoding/base64"

func ImgToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
