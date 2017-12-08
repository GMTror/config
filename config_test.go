package config

import (
	"bytes"
	"os"
	"testing"
	"time"
)

type Time struct {
	time.Duration
}

func (t *Time) UnmarshalENV(data []byte) error {
	pt, err := time.ParseDuration(bytes.NewBuffer(data).String())
	if err != nil {
		return err
	}
	*t = Time{pt}
	return nil
}

type ConfigENV struct {
	Byte         byte               `env:"BYTE"default:"100"`
	Int          int                `env:"INT"default:"2147483647"`
	Int8         int8               `env:"INT8"default:"127"`
	Int16        int16              `env:"INT16"default:"32767"`
	Int32        int32              `env:"INT32"default:"2147483647"`
	Int64        int64              `env:"INT64"default:"9223372036854775807"`
	Uint         uint               `env:"UINT"default:"4294967295"`
	Uint8        uint8              `env:"UINT8"default:"255"`
	Uint16       uint16             `env:"UINT16"default:"65535"`
	Uint32       uint32             `env:"UINT32"default:"4294967295"`
	Uint64       uint64             `env:"UINT64"default:"18446744073709551615"`
	Float32      float32            `env:"FLOAT32"default:"3.40282346638528859811704183484516925440e+38"`
	Float64      float64            `env:"FLOAT64"default:"1.797693134862315708145274237317043567981e+308"`
	String       string             `env:"STRING"default:"test_string"description:"this is description"`
	Bool         bool               `env:"BOOL"default:"true"`
	ArrayByte    []byte             `env:"ARRAY_BYTE"default:"10,100"`
	ArrayInt     []int              `env:"ARRAY_INT"default:"-2147483648,2147483647"`
	ArrayInt8    []int8             `env:"ARRAY_INT8"default:"-128,127"`
	ArrayInt16   []int16            `env:"ARRAY_INT16"default:"-32768,32767"`
	ArrayInt32   []int32            `env:"ARRAY_INT32"default:"-2147483648,2147483647"`
	ArrayInt64   []int64            `env:"ARRAY_INT64"default:"-9223372036854775808,9223372036854775807"`
	ArrayUint    []uint             `env:"ARRAY_UINT"default:"0,4294967295"`
	ArrayUint8   []uint8            `env:"ARRAY_UINT8"default:"0,255"`
	ArrayUint16  []uint16           `env:"ARRAY_UINT16"default:"0,65535"`
	ArrayUint32  []uint32           `env:"ARRAY_UINT32"default:"0,4294967295"`
	ArrayUint64  []uint64           `env:"ARRAY_UINT64"default:"0,18446744073709551615"`
	ArrayFloat32 []float32          `env:"ARRAY_FLOAT32"default:"1.401298464324817070923729583289916131280e-45,3.40282346638528859811704183484516925440e+38"`
	ArrayFloat64 []float64          `env:"ARRAY_FLOAT64"default:"4.940656458412465441765687928682213723651e-324,1.797693134862315708145274237317043567981e+308"`
	ArrayString  []string           `env:"ARRAY_STRING"default:"test_string,test_string2"`
	ArrayBool    []bool             `env:"ARRAY_BOOL"default:"false,true"`
	HashByte     map[string]byte    `env:"HASH_BYTE"default:"a:1,b:2,c:3"`
	HashInt      map[string]int     `env:"HASH_INT"default:"a:-2147483648,b:0,c:2147483647"`
	HashInt8     map[string]int8    `env:"HASH_INT8"default:"a:-128,b:0,c:127"`
	HashInt16    map[string]int16   `env:"HASH_INT16"default:"a:-32768,b:0,c:32767"`
	HashInt32    map[string]int32   `env:"HASH_INT32"default:"a:-2147483648,b:0,c:2147483647"`
	HashInt64    map[string]int64   `env:"HASH_INT64"default:"a:-9223372036854775808,b:0,c:9223372036854775807"`
	HashUint     map[string]uint    `env:"HASH_UINT"default:"a:0,b:4294967295"`
	HashUint8    map[string]uint8   `env:"HASH_UINT8"default:"a:0,b:255"`
	HashUint16   map[string]uint16  `env:"HASH_UINT16"default:"a:0,b:65535"`
	HashUint32   map[string]uint32  `env:"HASH_UINT32"default:"a:0,b:4294967295"`
	HashUint64   map[string]uint64  `env:"HASH_UINT64"default:"a:0,b:18446744073709551615"`
	HashFloat32  map[string]float32 `env:"HASH_FLOAT32"default:"a:1.401298464324817070923729583289916131280e-45,b:0,c:3.40282346638528859811704183484516925440e+38"`
	HashFloat64  map[string]float64 `env:"HASH_FLOAT64"default:"a:4.940656458412465441765687928682213723651e-324,b:0,c:1.797693134862315708145274237317043567981e+308"`
	HashString   map[string]string  `env:"HASH_STRING"default:"a:,b:test_string"`
	HashBool     map[string]bool    `env:"HASH_BOOL"default:"a:false,b:true"`
	Struct       struct {
		Struct struct {
			String string `env:"STRING"default:"test_string"description:"this is description"`
		}
		PasStruct struct {
			String string `env:"PAS_STRING"default:"test_string"description:"this is description"`
		} `env:"-"description:"this is description"`
	} `env:"STRUCT"description:"this is description"`
}

var envs = map[string]string{
	"BYTE":              "10",
	"INT":               "-2147483648",
	"INT8":              "-128",
	"INT16":             "-32768",
	"INT32":             "-2147483648",
	"INT64":             "-9223372036854775808",
	"UINT":              "0",
	"UINT8":             "0",
	"UINT16":            "0",
	"UINT32":            "0",
	"UINT64":            "0",
	"FLOAT32":           "1.401298464324817070923729583289916131280e-45",
	"FLOAT64":           "4.940656458412465441765687928682213723651e-324",
	"STRING":            "test_string_2",
	"BOOL":              "false",
	"ARRAY_BYTE":        "1,2,3",
	"ARRAY_INT":         "-1,0,1",
	"ARRAY_INT8":        "-1,0,1",
	"ARRAY_INT16":       "-1,0,1",
	"ARRAY_INT32":       "-1,0,1",
	"ARRAY_INT64":       "-1,0,1",
	"ARRAY_UINT":        "0,1,2",
	"ARRAY_UINT8":       "0,1,2",
	"ARRAY_UINT16":      "0,1,2",
	"ARRAY_UINT32":      "0,1,2",
	"ARRAY_UINT64":      "0,1,2",
	"ARRAY_FLOAT32":     "-1.1,0,1.1",
	"ARRAY_FLOAT64":     "-1.1,0,1.1",
	"ARRAY_STRING":      "a,b,aa,bb",
	"ARRAY_BOOL":        "1,t,T,TRUE,true,True,0,f,F,FALSE,false,False",
	"HASH_BYTE":         "a:4,b:5",
	"HASH_INT":          "a:0,b:1",
	"HASH_INT8":         "a:0,b:1",
	"HASH_INT16":        "a:0,b:1",
	"HASH_INT32":        "a:0,b:1",
	"HASH_INT64":        "a:0,b:1",
	"HASH_UINT":         "a:1,b:0",
	"HASH_UINT8":        "a:1,b:0",
	"HASH_UINT16":       "a:1,b:0",
	"HASH_UINT32":       "a:1,b:0",
	"HASH_UINT64":       "a:1,b:0",
	"HASH_FLOAT32":      "a:0",
	"HASH_FLOAT64":      "a:0",
	"HASH_STRING":       "a:test_string,b:",
	"HASH_BOOL":         "a:true,b:false",
	"STRUCT_STRING":     "test_string_2",
	"STRUCT_PAS_STRING": "test_string",
}

func TestReadEnv(t *testing.T) {
	for k, _ := range envs {
		if err := os.Setenv(k, ""); err != nil {
			t.Error(err)
		}
	}

	c := &ConfigENV{}

	if err := ReadENV(c); err != nil {
		t.Error(err)
	}

	if c.Byte != 100 {
		t.Error("default env BYTE")
	}

	if c.Int != 2147483647 {
		t.Error("default env INT")
	}

	if c.Int8 != 127 {
		t.Error("default env INT8")
	}

	if c.Int16 != 32767 {
		t.Error("default env INT16")
	}

	if c.Int32 != 2147483647 {
		t.Error("default env INT32")
	}

	if c.Int64 != 9223372036854775807 {
		t.Error("default env INT64")
	}

	if c.Uint != 4294967295 {
		t.Error("default env UINT")
	}

	if c.Uint8 != 255 {
		t.Error("default env UINT8")
	}

	if c.Uint16 != 65535 {
		t.Error("default env UINT16")
	}

	if c.Uint32 != 4294967295 {
		t.Error("default env UINT32")
	}

	if c.Uint64 != 18446744073709551615 {
		t.Error("default env UINT64")
	}

	if c.Float32 != 3.40282346638528859811704183484516925440e+38 {
		t.Error("default env FLOAT32")
	}

	if c.Float64 != 1.797693134862315708145274237317043567981e+308 {
		t.Error("default env FLOAT64")
	}

	if c.String != "test_string" {
		t.Error("default env STRING")
	}

	if !c.Bool {
		t.Error("default env BOOL")
	}

	if c.ArrayByte[0] != 10 || c.ArrayByte[1] != 100 {
		t.Error("default env ARRAY_BYTE")
	}

	if c.ArrayInt[0] != -2147483648 || c.ArrayInt[1] != 2147483647 {
		t.Error("default env ARRAY_INT")
	}

	if c.ArrayInt8[0] != -128 || c.ArrayInt8[1] != 127 {
		t.Error("default env ARRAY_INT8")
	}

	if c.ArrayInt16[0] != -32768 || c.ArrayInt16[1] != 32767 {
		t.Error("default env ARRAY_INT16")
	}

	if c.ArrayInt32[0] != -2147483648 || c.ArrayInt32[1] != 2147483647 {
		t.Error("default env ARRAY_INT32")
	}

	if c.ArrayInt64[0] != -9223372036854775808 || c.ArrayInt64[1] != 9223372036854775807 {
		t.Error("default env ARRAY_INT64")
	}

	if c.ArrayUint[0] != 0 || c.ArrayUint[1] != 4294967295 {
		t.Error("default env ARRAY_UINT")
	}

	if c.ArrayUint8[0] != 0 || c.ArrayUint8[1] != 255 {
		t.Error("default env ARRAY_UINT8")
	}

	if c.ArrayUint16[0] != 0 || c.ArrayUint16[1] != 65535 {
		t.Error("default env ARRAY_UINT16")
	}

	if c.ArrayUint32[0] != 0 || c.ArrayUint32[1] != 4294967295 {
		t.Error("default env ARRAY_UINT32")
	}

	if c.ArrayUint64[0] != 0 || c.ArrayUint64[1] != 18446744073709551615 {
		t.Error("default env ARRAY_UINT64")
	}

	if c.ArrayFloat32[0] != 1.401298464324817070923729583289916131280e-45 || c.ArrayFloat32[1] != 3.40282346638528859811704183484516925440e+38 {
		t.Error("default env ARRAY_FLOAT32")
	}

	if c.ArrayFloat64[0] != 4.940656458412465441765687928682213723651e-324 || c.ArrayFloat64[1] != 1.797693134862315708145274237317043567981e+308 {
		t.Error("default env ARRAY_FLOAT64")
	}

	if c.ArrayString[0] != "test_string" || c.ArrayString[1] != "test_string2" {
		t.Error("default env ARRAY_STRING")
	}

	if c.ArrayBool[0] || !c.ArrayBool[1] {
		t.Error("default env ARRAY_BOOL")
	}

	if c.HashByte["a"] != 1 || c.HashByte["b"] != 2 || c.HashByte["c"] != 3 {
		t.Error("default env HASH_BYTE")
	}

	if c.HashInt["a"] != -2147483648 || c.HashInt["b"] != 0 || c.HashInt["c"] != 2147483647 {
		t.Error("default env HASH_INT")
	}

	if c.HashInt8["a"] != -128 || c.HashInt8["b"] != 0 || c.HashInt8["c"] != 127 {
		t.Error("default env HASH_INT8")
	}

	if c.HashInt16["a"] != -32768 || c.HashInt16["b"] != 0 || c.HashInt16["c"] != 32767 {
		t.Error("default env HASH_INT16")
	}

	if c.HashInt32["a"] != -2147483648 || c.HashInt32["b"] != 0 || c.HashInt32["c"] != 2147483647 {
		t.Error("default env HASH_INT32")
	}

	if c.HashInt64["a"] != -9223372036854775808 || c.HashInt64["b"] != 0 || c.HashInt64["c"] != 9223372036854775807 {
		t.Error("default env HASH_INT64")
	}

	if c.HashUint["a"] != 0 || c.HashUint["b"] != 4294967295 {
		t.Error("default env HASH_UINT")
	}

	if c.HashUint8["a"] != 0 || c.HashUint8["b"] != 255 {
		t.Error("default env HASH_UINT8")
	}

	if c.HashUint16["a"] != 0 || c.HashUint16["b"] != 65535 {
		t.Error("default env HASH_UINT16")
	}

	if c.HashUint32["a"] != 0 || c.HashUint32["b"] != 4294967295 {
		t.Error("default env HASH_UINT32")
	}

	if c.HashUint64["a"] != 0 || c.HashUint64["b"] != 18446744073709551615 {
		t.Error("default env HASH_UINT64")
	}

	if c.HashFloat32["a"] != 1.401298464324817070923729583289916131280e-45 || c.HashFloat32["b"] != 0 || c.HashFloat32["c"] != 3.40282346638528859811704183484516925440e+38 {
		t.Error("default env HASH_FLOAT32")
	}

	if c.HashFloat64["a"] != 4.940656458412465441765687928682213723651e-324 || c.HashFloat64["b"] != 0 || c.HashFloat64["c"] != 1.797693134862315708145274237317043567981e+308 {
		t.Error("default env HASH_FLOAT64")
	}

	if c.HashString["a"] != "" || c.HashString["b"] != "test_string" {
		t.Error("default env HASH_STRING")
	}

	if c.HashBool["a"] || !c.HashBool["b"] {
		t.Error("default env HASH_BOOL")
	}

	if c.Struct.Struct.String != "test_string" {
		t.Error("read env STRUCT_STRING")
	}

	if c.Struct.PasStruct.String != "" {
		t.Error("read env STRUCT_PAS_STRING")
	}

	for k, v := range envs {
		if err := os.Setenv(k, v); err != nil {
			t.Error(err)
		}
	}

	c = &ConfigENV{}

	if err := ReadENV(c); err != nil {
		t.Error(err)
	}

	if c.Byte != 10 {
		t.Error("default env BYTE")
	}

	if c.Int != -2147483648 {
		t.Error("default env INT")
	}

	if c.Int8 != -128 {
		t.Error("default env INT8")
	}

	if c.Int16 != -32768 {
		t.Error("default env INT16")
	}

	if c.Int32 != -2147483648 {
		t.Error("default env INT32")
	}

	if c.Int64 != -9223372036854775808 {
		t.Error("default env INT64")
	}

	if c.Uint != 0 {
		t.Error("default env UINT")
	}

	if c.Uint8 != 0 {
		t.Error("default env UINT8")
	}

	if c.Uint16 != 0 {
		t.Error("default env UINT16")
	}

	if c.Uint32 != 0 {
		t.Error("default env UINT32")
	}

	if c.Uint64 != 0 {
		t.Error("default env UINT64")
	}

	if c.Float32 != 1.401298464324817070923729583289916131280e-45 {
		t.Error("default env FLOAT32")
	}

	if c.Float64 != 4.940656458412465441765687928682213723651e-324 {
		t.Error("default env FLOAT64")
	}

	if c.String != "test_string_2" {
		t.Error("default env STRING")
	}

	if c.Bool {
		t.Error("default env BOOL")
	}

	if c.ArrayByte[0] != 1 || c.ArrayByte[1] != 2 || c.ArrayByte[2] != 3 {
		t.Error("read env value ARRAY_BYTE")
	}

	if c.ArrayInt[0] != -1 || c.ArrayInt[1] != 0 || c.ArrayInt[2] != 1 {
		t.Error("read env value ARRAY_INT")
	}

	if c.ArrayInt8[0] != -1 || c.ArrayInt8[1] != 0 || c.ArrayInt8[2] != 1 {
		t.Error("read env value ARRAY_INT8")
	}

	if c.ArrayInt16[0] != -1 || c.ArrayInt16[1] != 0 || c.ArrayInt16[2] != 1 {
		t.Error("read env value ARRAY_INT16")
	}

	if c.ArrayInt32[0] != -1 || c.ArrayInt32[1] != 0 || c.ArrayInt32[2] != 1 {
		t.Error("read env value ARRAY_INT32")
	}

	if c.ArrayInt64[0] != -1 || c.ArrayInt64[1] != 0 || c.ArrayInt64[2] != 1 {
		t.Error("read env value ARRAY_INT64")
	}

	if c.ArrayUint[0] != 0 || c.ArrayUint[1] != 1 || c.ArrayUint[2] != 2 {
		t.Error("read env value ARRAY_UINT")
	}

	if c.ArrayUint8[0] != 0 || c.ArrayUint8[1] != 1 || c.ArrayUint8[2] != 2 {
		t.Error("read env value ARRAY_UINT8")
	}

	if c.ArrayUint16[0] != 0 || c.ArrayUint16[1] != 1 || c.ArrayUint16[2] != 2 {
		t.Error("read env value ARRAY_UINT16")
	}

	if c.ArrayUint32[0] != 0 || c.ArrayUint32[1] != 1 || c.ArrayUint32[2] != 2 {
		t.Error("read env value ARRAY_UINT32")
	}

	if c.ArrayUint64[0] != 0 || c.ArrayUint64[1] != 1 || c.ArrayUint64[2] != 2 {
		t.Error("read env value ARRAY_UINT64")
	}

	if c.ArrayFloat32[0] != -1.1 || c.ArrayFloat32[1] != 0 || c.ArrayFloat32[2] != 1.1 {
		t.Error("read env value ARRAY_FLOAT32")
	}

	if c.ArrayFloat64[0] != -1.1 || c.ArrayFloat64[1] != 0 || c.ArrayFloat64[2] != 1.1 {
		t.Error("read env value ARRAY_FLOAT64")
	}

	if c.ArrayString[0] != "a" || c.ArrayString[1] != "b" || c.ArrayString[2] != "aa" || c.ArrayString[3] != "bb" {
		t.Error("read env value ARRAY_STRING")
	}

	if !c.ArrayBool[0] || !c.ArrayBool[1] || !c.ArrayBool[2] || !c.ArrayBool[3] || !c.ArrayBool[4] || !c.ArrayBool[5] || c.ArrayBool[6] || c.ArrayBool[7] || c.ArrayBool[8] || c.ArrayBool[9] || c.ArrayBool[10] || c.ArrayBool[11] {
		t.Error("read env value ARRAY_BOOL")
	}

	if c.HashByte["a"] != 4 || c.HashByte["b"] != 5 {
		t.Error("read env HASH_BYTE")
	}

	if c.HashInt["a"] != 0 || c.HashInt["b"] != 1 {
		t.Error("read env HASH_INT")
	}

	if c.HashInt8["a"] != 0 || c.HashInt8["b"] != 1 {
		t.Error("read env HASH_INT8")
	}

	if c.HashInt16["a"] != 0 || c.HashInt16["b"] != 1 {
		t.Error("read env HASH_INT16")
	}

	if c.HashInt32["a"] != 0 || c.HashInt32["b"] != 1 {
		t.Error("read env HASH_INT32")
	}

	if c.HashInt64["a"] != 0 || c.HashInt64["b"] != 1 {
		t.Error("read env HASH_INT64")
	}

	if c.HashUint["a"] != 1 || c.HashUint["b"] != 0 {
		t.Error("read env HASH_UINT")
	}

	if c.HashUint8["a"] != 1 || c.HashUint8["b"] != 0 {
		t.Error("read env HASH_UINT8")
	}

	if c.HashUint16["a"] != 1 || c.HashUint16["b"] != 0 {
		t.Error("read env HASH_UINT16")
	}

	if c.HashUint32["a"] != 1 || c.HashUint32["b"] != 0 {
		t.Error("read env HASH_UINT32")
	}

	if c.HashUint64["a"] != 1 || c.HashUint64["b"] != 0 {
		t.Error("read env HASH_UINT64")
	}

	if c.HashFloat32["a"] != 0 {
		t.Error("read env HASH_FLOAT32")
	}

	if c.HashFloat64["a"] != 0 {
		t.Error("read env HASH_FLOAT64")
	}

	if c.HashString["a"] != "test_string" || c.HashString["b"] != "" {
		t.Error("read env HASH_STRING")
	}

	if !c.HashBool["a"] || c.HashBool["b"] {
		t.Error("read env HASH_BOOL")
	}

	if c.Struct.Struct.String != "test_string_2" {
		t.Error("read env STRUCT_STRING")
	}

	if c.Struct.PasStruct.String != "" {
		t.Error("read env STRUCT_PAS_STRING")
	}
}
