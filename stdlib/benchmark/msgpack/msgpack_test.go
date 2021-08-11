package msgpack

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

var s = Status{
	Age:   26,
	Name:  "bohler",
	Money: 0.1,
	Areas: []string{"SZ", "HN"},
	QinFu: map[string]int{"dml": 25, "lzy": 25},
}

func BenchmarkMarshalByMsgPack(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(s)
	}
}

func BenchmarkUnMarshalByMsgPack(b *testing.B) {
	bytes, _ := msgpack.Marshal(s)
	ss := Status{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(bytes, &ss)
	}
}

func BenchmarkMarshalByJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(s)
	}
}

func BenchmarkUnMarshalByJson(b *testing.B) {
	bytes, _ := json.Marshal(s)
	ss := Status{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bytes, &ss)
	}
}

func BenchmarkMarshalByMsgp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.MarshalMsg(nil)
	}
}

func BenchmarkUnMarshalByMsgp(b *testing.B) {
	bytes, _ := msgpack.Marshal(s)
	ss := &Status{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss.UnmarshalMsg(bytes)
	}
}

/*
func BenchmarkMarshalByGjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gjson.
	}
}
*/

func BenchmarkMarshalByJsonIter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(s)
	}
}

func BenchmarkUnMarshalByJsonIter(b *testing.B) {
	bytes, _ := jsoniter.Marshal(s)
	ss := Status{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(bytes, &ss)
	}
}
