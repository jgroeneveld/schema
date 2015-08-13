package schema

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"testing"
)

func dataFromJSON(t *testing.T, jsonStr string) interface{} {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("can not get runtime caller")
	}
	location := fmt.Sprintf("%s:%d", path.Base(file), line)

	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Fatalf("Fail to unmarshal JSON: %s\nsource: %s", err, location)
	}
	return data
}
