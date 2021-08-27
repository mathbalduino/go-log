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
		expectedFields := LogFields{"a": "aaa", "b": "bbb", "c": "ccc", c.LvlFieldName: uint64(0), c.MsgFieldName: "my msg"}
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
	t.Run("AdHocFields should override asyncHooks functions, that needs to be called to override SyncFields, that override base fields", raceFreeTest(func(t *testing.T) {
		c := DefaultConfig()
		expectedFields := LogFields{
			"field":         "value",
			"overrideField": "newValue",
			"sync":          "value",
			"overrideSync":  "newValue",
			"async":         "value",
			"overrideAsync": "newValue",
			"adHoc":         "value",
			c.LvlFieldName:  uint64(0),
			c.MsgFieldName:  "my msg",
		}
		calls := 0
		l := Log{lvl: 0, msg: "my msg",
			logger: &Logger{
				fields: LogFields{"overrideField": "shouldOverrideThis", "field": "value"},
				asyncHooks: Hooks{
					"overrideSync":  func(log Log) interface{} { return "newValue" },
					"overrideAsync": func(log Log) interface{} { return "shouldOverrideThis" },
					"async":         func(log Log) interface{} { return "value" },
				},
				configuration: &c,
				outputs: []Output{func(lvl uint64, msg string, receivedFields LogFields) {
					calls += 1
					if !reflect.DeepEqual(receivedFields, expectedFields) {
						t.Fatalf("Expected an equivalent fields")
					}
				}},
			},
			syncFields:  LogFields{"overrideField": "newValue", "overrideSync": "shouldOverrideThis", "sync": "value"},
			adHocFields: []LogFields{{"overrideAsync": "newValue", "adHoc": "value"}},
		}
		handleLog(l)
		if calls != 1 {
			t.Fatalf("Expected to call the output")
		}
	}, rHandleLog))
	t.Run("Should call every registered output, with the correct log", raceFreeTest(func(t *testing.T) {
		c := DefaultConfig()
		callsA, callsB, callsC := 0, 0, 0
		expectedFields := LogFields{c.LvlFieldName: uint64(0), c.MsgFieldName: "my msg"}
		l := Log{lvl: 0, msg: "my msg",
			logger: &Logger{
				configuration: &c,
				outputs: []Output{
					func(lvl uint64, msg string, receivedFields LogFields) {
						callsA += 1
						if !reflect.DeepEqual(receivedFields, expectedFields) {
							t.Fatalf("Expected an equivalent fields")
						}
					},
					func(lvl uint64, msg string, receivedFields LogFields) {
						callsB += 1
						if !reflect.DeepEqual(receivedFields, expectedFields) {
							t.Fatalf("Expected an equivalent fields")
						}
					},
					func(lvl uint64, msg string, receivedFields LogFields) {
						callsC += 1
						if !reflect.DeepEqual(receivedFields, expectedFields) {
							t.Fatalf("Expected an equivalent fields")
						}
					},
				},
			},
		}
		handleLog(l)
		if callsA != 1 || callsB != 1 || callsC != 1 {
			t.Fatalf("Expected to call every output")
		}
	}, rHandleLog))
	t.Run("Should pass the correct lvl and msg as separate arguments, and inside the map", raceFreeTest(func(t *testing.T) {
		c := DefaultConfig()
		calls := 0
		l := Log{lvl: 0, msg: "my msg",
			logger: &Logger{
				configuration: &c,
				outputs: []Output{func(lvl uint64, msg string, receivedFields LogFields) {
					calls += 1
					if lvl != 0 || receivedFields[c.LvlFieldName] != uint64(0) {
						t.Fatalf("Wrong lvl")
					}
					if msg != "my msg" || receivedFields[c.MsgFieldName] != "my msg" {
						t.Fatalf("Wrong msg")
					}
				}},
			},
		}
		handleLog(l)
		if calls != 1 {
			t.Fatalf("Expected to call every output")
		}
	}, rHandleLog))
}
