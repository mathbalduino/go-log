package logger

import (
	"reflect"
	"testing"
	"time"
)

func TestLog_Field(t *testing.T) {
	t.Run("Search postFields first", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ddd"}},
			preFields:   LogFields{"a": "ccc"},
			adHocFields: []LogFields{{"a": "bbb"}},
			postFields:  LogFields{"a": "aaa"},
		}
		v := l.Field("a")
		if v.(string) != "aaa" {
			t.Fatalf("Expected to get the value from the log postFields")
		}
	})
	t.Run("Search adHocFields after postFields", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ddd"}},
			preFields:   LogFields{"a": "ccc"},
			adHocFields: []LogFields{{"a": "bbb", "z": "xxx"}, {"z": "zzz"}},
			postFields:  LogFields{"b": "aaa"},
		}
		v := l.Field("a")
		if v.(string) != "bbb" {
			t.Fatalf("Expected to get the value from the log adHocFields")
		}
		v = l.Field("z")
		if v.(string) != "zzz" {
			t.Fatalf("Expected to get the value from the log adHocFields (the latter one)")
		}
	})
	t.Run("Search preFields after adHocFields", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ddd"}},
			preFields:   LogFields{"a": "ccc"},
			adHocFields: []LogFields{{"b": "bbb"}},
			postFields:  LogFields{"b": "aaa"},
		}
		v := l.Field("a")
		if v.(string) != "ccc" {
			t.Fatalf("Expected to get the value from the log preFields")
		}
	})
	t.Run("Search fields after preFields", func(t *testing.T) {
		l := Log{
			logger:      &Logger{fields: LogFields{"a": "ddd"}},
			preFields:   LogFields{"b": "ccc"},
			adHocFields: []LogFields{{"b": "bbb"}},
			postFields:  LogFields{"b": "aaa"},
		}
		v := l.Field("a")
		if v.(string) != "ddd" {
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
	t.Run("PostFields should override adHocFields, that override PreFields, that override fields", raceFreeTest(func(t *testing.T) {
		c := DefaultConfig()
		expectedFields := LogFields{
			"field":               "value",
			"overriddenField":     "newValue",
			"preField":            "value",
			"overriddenPreField":  "newValue",
			"adHoc":               "value",
			"overriddenAdHoc_1":     "newValue",
			"overriddenAdHoc_2":     "newValue",
			"postField":           "value",
			c.LvlFieldName:        uint64(0),
			c.MsgFieldName:        "my msg",
		}
		calls := 0
		l := Log{lvl: 0, msg: "my msg",
			logger: &Logger{
				fields: LogFields{"overriddenField": "shouldOverrideThis", "field": "value"},
				postHooks: Hooks{
					"overriddenAdHoc_2":  func(log Log) interface{} { return "newValue" },
					"postField":         func(log Log) interface{} { return "value" },
				},
				configuration: &c,
				outputs: []Output{func(lvl uint64, msg string, receivedFields LogFields) {
					calls += 1
					if !reflect.DeepEqual(receivedFields, expectedFields) {
						t.Fatalf("Expected an equivalent fields")
					}
				}},
			},
			preFields:   LogFields{"overriddenField": "newValue", "overriddenPreField": "shouldOverrideThis", "preField": "value"},
			adHocFields: []LogFields{
				{"overriddenAdHoc_1": "shouldOverrideThis", "overriddenAdHoc_2": "shouldOverrideThis", "adHoc": "value", "overriddenPreField": "newValue"},
				{"overriddenAdHoc_1": "newValue"}},
		}
		handleLog(l)
		if calls != 1 {
			t.Fatalf("Expected to call the output")
		}
	}, rHandleLog))
	t.Run("Should call every registered output, with the correct log, in the correct order", raceFreeTest(func(t *testing.T) {
		c := DefaultConfig()
		callsA, callsB, callsC := 0, 0, 0
		timeA, timeB, timeC := time.Time{}, time.Time{}, time.Time{}
		expectedFields := LogFields{c.LvlFieldName: uint64(0), c.MsgFieldName: "my msg"}
		l := Log{lvl: 0, msg: "my msg",
			logger: &Logger{
				configuration: &c,
				outputs: []Output{
					func(lvl uint64, msg string, receivedFields LogFields) {
						callsA += 1
						timeA = time.Now()
						if !reflect.DeepEqual(receivedFields, expectedFields) {
							t.Fatalf("Expected an equivalent fields")
						}
					},
					func(lvl uint64, msg string, receivedFields LogFields) {
						callsB += 1
						timeB = time.Now()
						if !reflect.DeepEqual(receivedFields, expectedFields) {
							t.Fatalf("Expected an equivalent fields")
						}
					},
					func(lvl uint64, msg string, receivedFields LogFields) {
						callsC += 1
						timeC = time.Now()
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
		if !timeA.Before(timeB) || !timeB.Before(timeC) {
			t.Fatal("Expected the outputs to be called in order")
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
