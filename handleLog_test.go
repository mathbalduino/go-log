package loxeLog

import (
	"reflect"
	"testing"
)

func TestLog_Field(t *testing.T) {
	t.Run("Search adHocFields first", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ccc"}},
			syncFields:  LogFields{"a": "bbb"},
			adHocFields: []LogFields{{"a": "aaa"}},
		}
		v := l.Field("a")
		if v.(string) != "aaa" {
			t.Fatalf("Expected to get the value from the log adHocFields")
		}
	})
	t.Run("Search syncFields after not found at adHocFields", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ccc"}},
			syncFields:  LogFields{"a": "bbb"},
			adHocFields: []LogFields{{"b": "aaa"}},
		}
		v := l.Field("a")
		if v.(string) != "bbb" {
			t.Fatalf("Expected to get the value from the log syncFields")
		}
	})
	t.Run("Search logger fields after not found at syncFields", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ccc"}},
			syncFields:  LogFields{"b": "bbb"},
			adHocFields: []LogFields{{"b": "aaa"}},
		}
		v := l.Field("a")
		if v.(string) != "ccc" {
			t.Fatalf("Expected to get the value from the logger fields")
		}
	})
}

// Since it will call "handleLog", it needs to lock the read mutex
func TestHandleLog(t *testing.T) {
	t.Run("Should clone the logger fields into another map, adding the 'lvl' and 'msg' keys", raceFreeTest(func(t *testing.T) {
		c := DefaultConfig()
		expectedFields := LogFields{"a": "aaa", "b": "bbb", "c": "ccc", "lvl": uint64(0), "msg": "my msg"}
		calls := 0
		l := Log{lvl: 0, msg: "my msg", logger: &Logger{
			configuration: &c,
			outputs: []Output{func(lvl uint64, msg string, receivedFields LogFields) {
				calls += 1
				if reflect.ValueOf(receivedFields).Pointer() == reflect.ValueOf(expectedFields).Pointer() {
					t.Fatalf("Expected to create another map")
				}
				if !reflect.DeepEqual(receivedFields, expectedFields) {
					t.Fatalf("Expected an equivalent fields")
				}
			}},
			fields: LogFields{"a": "aaa", "b": "bbb", "c": "ccc"},
		}}
		handleLog(l)
		if calls != 1 {
			t.Fatalf("Expected to call the output")
		}
	}, rHandleLog))
}
