package mockserver

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

func msDecode(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil, Result: output, TagName: "json",
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func addToStringArray(arr []string, val string) []string {
	for _, v := range arr {
		if v == val {
			return arr
		}
	}
	return append(arr, val)
}

func removeFromStringArray(arr []string, val string) []string {
	ret := []string{}
	for _, v := range arr {
		if v != val {
			ret = append(ret, v)
		}
	}
	return ret
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func validateRequestBody(
	b io.Reader, requiredFields, allowedFields, acceptedFields []string,
) (
	map[string]interface{}, int, error,
) {
	dec := json.NewDecoder(b)
	var a interface{}
	if err := dec.Decode(&a); err != nil {
		return nil, 400, fmt.Errorf(
			"failed to parse the request body as JSON: %s", err)
	}
	body, ok := a.(map[string]interface{})
	if !ok {
		return nil, 400, fmt.Errorf(
			"failed to parse the request body as a JSON object: %s", a)
	}
	rqf := makeHash(requiredFields)
	for k, _ := range rqf {
		if _, ok := body[k]; !ok {
			return body, 400, fmt.Errorf(
				`in the request body the field "%s" is required`, k)
		}
	}
	alf := makeHash(allowedFields)
	if len(alf) != 0 {
		for k, _ := range rqf {
			alf[k] = nil
		}
		arr := make([]string, len(alf))
		i := 0
		for k, _ := range alf {
			arr[i] = k
			i++
		}
		for k, _ := range body {
			if _, ok := alf[k]; !ok {
				return body, 400, fmt.Errorf(
					`in the request body an invalid field is found: "%s"\nThe allowed fields: %s`,
					k, strings.Join(arr, ", "))
			}
		}
	}
	acf := makeHash(acceptedFields)
	if len(acf) != 0 {
		for k, _ := range rqf {
			acf[k] = nil
		}
		for k, _ := range alf {
			acf[k] = nil
		}
		for k, _ := range body {
			if _, ok := acf[k]; !ok {
				delete(body, k)
			}
		}
	}
	return body, 200, nil
}

func makeHash(arr []string) map[string]interface{} {
	if arr == nil {
		return map[string]interface{}{}
	}
	h := map[string]interface{}{}
	for _, k := range arr {
		h[k] = nil
	}
	return h
}
