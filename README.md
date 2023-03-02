<div align=center>

![Go version](https://img.shields.io/github/go-mod/go-version/dablelv/go-huge-util)
[![GitHub latest tag](https://img.shields.io/github/tag/dablelv/go-huge-util)](https://github.com/dablelv/go-huge-util)
[![Go Report Card](https://goreportcard.com/badge/github.com/dablelv/go-huge-util)](https://goreportcard.com/report/github.com/dablelv/go-huge-util)
[![Unit Test](https://github.com/dablelv/go-huge-util/workflows/go-test/badge.svg)](https://github.com/dablelv/go-huge-util/actions)
[![Coverage Status](https://coveralls.io/repos/github/dablelv/go-huge-util/badge.svg?branch=master)](https://coveralls.io/github/dablelv/go-huge-util?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/dablelv/go-huge-util.svg)](https://pkg.go.dev/github.com/dablelv/go-huge-util)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/dablelv/go-huge-util/blob/main/LICENSE)
[![Stars](https://img.shields.io/github/stars/dablelv/go-huge-util?style=social)](https://img.shields.io/github/stars/dablelv/go-huge-util?style=social)

</div>

# Synopsis
Go common used utility functions help you to create go program quickly and easily.

# Encode
Some useful functions can be used to handle conversion of different character encoding, such as gbk to utf8.
```go
gbkStr := []byte{0xC4, 0xE3, 0xBA, 0xC3}    // 你好 in gbk
utf8Str, _ := GbkToUtf8(gbkStr)             // 你好 in utf8
gbkStrRes, _ := Utf8ToGbk(utf8Str)          // [196 227 186 195]
```
# Net
Some useful functions can be used to handle network. For example you can use `IPv4StrToU32()` transform ipv4 string to uint32 value.

```go
IsReservedIP("0.0.0.256")    // -1 invalid ip
IsReservedIP("39.156.69.79") // 0 public ip
IsReservedIP("127.0.0.1")    // 1 reserved ip

IPv4StrToU32("127.0.0.1")      // 2130706433
U32ToIPv4Str(2130706433)       // "127.0.0.1"
GetNativeEndian()              // LittleEndian
IsLittleEndian()               // true
```
# Slice
Some useful functions can be used to handle slice.

```go
UniqueIntSlice([]int{1, 2, 2, 3})              // [1 2 3]
UniqueUintSlice([]uint{1, 2, 2, 3})            // [1 2 3]
UniqueStrSlice([]string{"a", "b", "b", "c"})   // [a b c]

ReverseIntSlice([]int{1, 2, 3})                // [3 2 1]
ReverseUintSlice([]uint{1, 2, 3})              // [3 2 1]
ReverseStrSlice([]string{"a", "b", "c"})       // [c b a]

SumSlice([]int{1,2,3})                         // 6
SumSlice([]uint{1,2,3})                        // 6
SumSlice([]float32{1.1, 2.2, 3.3})             // 6.6
SumSlice([]float64{1.1, 2.2, 3.3})             // 6.6

JoinSliceWithSep([]int{1, 2, 3}, ",")              // 1,2,3
JoinSliceWithSep([]uint{1, 2, 3}, ",")             // 1,2,3
JoinSliceWithSep([]float64{1.1, 2.2, 3.3}, ",")    // 1.1,2.2,3.3
JoinSliceWithSep([]string{"a", "b", "c"}, ",")     // a,b,c

// CRUD(Create Read Update Delete) on slice by index
fib := []int{1, 1, 2, 3, 5, 8}
r, _ := InsertSliceE(fib, 6, 13)    // [1 1 2 3 5 8 13]
r, _ := DeleteSliceE(fib, 0)        // [1 2 3 5 8]
r, _ := UpdateSliceE(fib, 5, 88)    // [1 1 2 3 5 88]
r, _ := GetEleIndexesSliceE(fib, 1) // [0 1]

// or
r := InsertIntSlice(fib, 5, 13)		// [1 1 2 3 5 8 13]
r := DeleteIntSliceE(fib, 0)		// [1 2 3 5 8]
r := UpdateIntSliceE(fib, 5, 88)	// [1 1 2 3 5 88]
r := GetElemIndexesSlice(fib, 1)	// [0 1]
```

# SQL
Some useful functions can be used to handle sql statement.

```go
var sql = `select * from t where field1="{name}"`
huge.FormatSql(sql, map[string]string{"name": "dablelv"}, false)    // select * from t where field1="dablelv"
```

# String
Some useful functions can be used to handle string.

```go
Split("a,b,c", ",")                    // []string{"a", "b", "c"}

SplitSeps("a,b|c", ",", "|")           // []string{"a", "b", "c"}
SplitSeps("a,bSEPc", ",", "SEP")       // []string{"a", "b", "c"}

JoinStr(",", "a", "", "b")             // "a,,b"
JoinStrSkipEmpty(",", "a", "", "b")    // "a,b"

ReverseStr("abc")                      // "cba"

GetAlphanumericNumByASCII("108条梁山man")     // 6
GetAlphanumericNumByASCIIV2("108条梁山man")   // 6
GetAlphanumericNumByRegExp("108条梁山man")    // 6
```

# Time
Some useful functions can be used to handle date and time.

```go
tc := huge.NewTimeCounter()
// do your statements
tc.GetMs()                  // get the time cost in millisecond

GetNowDate()                // just a example, the same below. 2020-05-16
GetNowTime()                // 00:15:42
GetNowDateTime()            // 2020-05-16 00:15:42
GetNowDateTimeZ()           // 2020-05-16 00:15:42 +08:00

GetNowS()                   // 1589559342
GetNowMs()                  // 1589559342963
GetNowUs()                  // 1589559342963062
GetNowNs()                  // 1589559342964063200

GetDayBeginMoment(time.Now())  // 2020-05-16 00:00:00 +0800 CST
GetDayEndMoment(time.Now())    // 2020-05-16 23:59:59.999999999 +0800 CST
```

# Type Conversion
Some useful functions can be used to convert one type to another types, such as to map or slice.

## to another type
```go
// to string
s, err := ToAny[string]("foo") // "foo"
s, err := ToAny[string](8)     // "8"
s, err := ToAny[string](8.31)  // "8.31"
s, err := ToAny[string]([]byte("one time")) // "one time"
s, err := ToAny[string](nil)                // ""

var foo any = "one more time"
s, err := ToAny[string](foo)                // "one more time"

// to int
i, err :=  ToAny[int](8)                  // 8
i, err :=  ToAny[int](8.31)               // 8
i, err :=  ToAny[int]("8")                // 8
i, err :=  ToAny[int](true)               // 1
i, err :=  ToAny[int](false)              // 0
i, err :=  ToAny[int](nil)                // 0

var eight any = 8
i, err :=  ToAny[int](eight)              // 8

// to bool
b, err := ToAny[bool]("true")           // true
b, err := ToAny[bool]("false")          // false
b, err := ToAny[bool]("True")           // true
b, err := ToAny[bool]("False")          // false
b, err := ToAny[bool](1)                // true
b, err := ToAny[bool](0)                // false
b, err := ToAny[bool](nil)              // false

var one any = 1
b, err := ToAny[bool](one)           // true
```
## to set
```go
// Convert bool slice or array to set.
bools := []bool{true, false, true}
set := conv.ToBoolSet(bools)
set, _ := conv.ToBoolSetE(bools)
set := conv.ToSet[bool](bools)
set, _ := conv.ToSetE[bool](bools)

// Convert int slice or array to set.
ints := []int{1, 2, 3}
set := conv.ToIntSet(ints)
set, _ := conv.ToIntSetE(ints)
set := conv.ToSetG[int](ints)
set, _ := conv.ToSetE[int](ints)

// Convert string slice or array to set.
strs := []string{"foo", "bar", "baz"}
set := conv.ToStrSet(strs)
set, _ := conv.ToStrSetE(strs)
set := conv.ToSet[string](strs)
set, _ := conv.ToSetE[string](strs)

// Split string to set.
conv.SplitStrToSet("a,b,c", ",")  // map[a:{}, b:{}, c:{}]
```

## to slice
```go
// Convert string separated by white space character to string slice.
sl := conv.ToStrSlice("a b c") // []string{"a","b","c"}
sl, _ := conv.ToStrSliceE("a b c") // []string{"a","b","c"}

// Convert int slice or array to string slice.
sl := conv.ToStrSlice([]int{1, 2, 3}) // []string{"1","2","3"}
sl, _ := conv.ToStrSliceE([]int{1, 2, 3}) // []string{"1","2","3"}

// Convert map to slice in random order.
ks, vs := conv.Map2Slice(map[int]int{1:1, 2:2, 3:3})
ks, vs, _ := conv.Map2SliceE(map[int]int{1:1, 2:2, 3:3})

slK, _ : = ks.([]int)
slV, _ : = vs.([]int)

ks, vs := conv.Map2Slice(map[string]int{"foo":1, "bar":2, "baz":3})
ks, vs, _ := conv.Map2SliceE(map[string]int{"foo":1, "bar":2, "baz":3})

slK, _ : = ks.([]string)
slV, _ : = vs.([]int)

// Split string to slice.
// int[1,2,3]
ints := conv.SplitStrToSlice[int]("1,2,3", ",")
ints, _ := conv.SplitStrToSliceE[int]("1,2,3", ",")
// uint[1,2,3]      
uints := conv.SplitStrToSlice[uint]("1,2,3", ",")
uints, _ := conv.SplitStrToSliceE[uint]("1,2,3", ",")
// float64[1.1,2.2,3.3]
f64s := conv.SplitStrToSlice[float64]("1.1,2.2,3.3", ",")
f64s, _ := conv.SplitStrToSliceE[float64]("1.1,2.2,3.3", ",")
// bool[true,false,true,false]
bs := conv.SplitStrToSlice[bool]("1,0,true,false", ",")
bs, _ := conv.SplitStrToSliceE[bool]("1,0,true,false", ",")
```

## to map

```go
var st = struct {
    I int
    S string
}{I: 1, S: "a"}

// to map[string]any
Struct2Map(st)         // map["I":1 "S":"a"]
// to map[string]string
Struct2MapString(st)   // map["I":"1" "S":"a"]

// any type tp map[string]string
m := ToMapStrStr(`{"foo":"foo","bar":"bar","baz":"baz"}`)       // map["foo":"foo" "bar":"bar" "baz":"baz"]
m, err := ToMapStrStrE(`{"foo":"foo","bar":"bar","baz":"baz"}`) // map["foo":"foo" "bar":"bar" "baz":"baz"], nil
```

# URL

Some useful functions can be used to handle url.

```go
var rawUrl=`http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name`
huge.RawUrlGetDomain(rawUrl)    // "www.aspxfans.com"
huge.RawUrlGetPort(rawUrl)      // "8080"

huge.RawURLGetParam(rawUrl, "page")         // 1 <nil>
huge.RawURLGetParams(rawUrl, "page")        // [1 2] <nil>
huge.RawURLGetAllParams(rawUrl)             // map[boardID:[520] page:[1 2]] <nil>

huge.RawURLAddParam(rawUrl, "keyword", "dog")   // http://www.aspxfans.com:8080/news/index.asp?boardID=520&keyword=dog&page=1&page=2#name
huge.RawURLDelParam(rawUrl, "page")             // http://www.aspxfans.com:8080/news/index.asp?boardID=520#name
huge.RawURLSetParam(rawUrl, "boardID", "521")   // http://www.aspxfans.com:8080/news/index.asp?boardID=521&page=1&page=2#name
```

# Crypto
Some useful functions can be used to create Hash, HMAC and crypt data.

```go
// Hash functions.
MD5L("")   // d41d8cd98f00b204e9800998ecf8427e
MD5U("")   // D41D8CD98F00B204E9800998ECF8427E

SHA1L("")  // da39a3ee5e6b4b0d3255bfef95601890afd80709
SHA1U("")  // DA39A3EE5E6B4B0D3255BFEF95601890AFD80709

SHA224L("")    // d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f
SHA224U("")    // D14A028C2A3A2BC9476102BB288234C415A2B01F828EA62AC5B3E42F

SHA256L("")    // e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
SHA256U("")    // E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855

SHA384L("")    // 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b
SHA384U("")    // 38B060A751AC96384CD9327EB1B1E36A21FDB71114BE07434C0CC7BF63F6E1DA274EDEBFE76F65FBD51AD2F14898B95B

SHA512L("")    // cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e
SHA512U("")    // CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E

// HMAC functions.
HMACMD5L("", "")   // 74e6f7298a9c2d168935f58c001bad88
HMACMD5U("", "")   // 74E6F7298A9C2D168935F58C001BAD88

HMACSHA1L("", "")  // fbdb1d1b18aa6c08324b7d64b71fb76370690e1d
HMACSHA1U("", "")  // FBDB1D1B18AA6C08324B7D64B71FB76370690E1D

HMACSHA224L("", "")    // 5ce14f72894662213e2748d2a6ba234b74263910cedde2f5a9271524
HMACSHA224U("", "")    // 5CE14F72894662213E2748D2A6BA234B74263910CEDDE2F5A9271524

HMACSHA256L("", "")    // b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad
HMACSHA256U("", "")    // B613679A0814D9EC772F95D778C35FC5FF1697C493715653C6C712144292C5AD

HMACSHA384L("", "")    // 6c1f2ee938fad2e24bd91298474382ca218c75db3d83e114b3d4367776d14d3551289e75e8209cd4b792302840234adc
HMACSHA384U("", "")    // 6C1F2EE938FAD2E24BD91298474382CA218C75DB3D83E114B3D4367776D14D3551289E75E8209CD4B792302840234ADC

HMACSHA512L("", "")    // b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47
HMACSHA512U("", "")    // B936CEE86C9F87AA5D3C6F2E84CB5A4239A5FE50480A6EC66B70AB5B1F4AC6730C6C515421B327EC1D69402E53DFB49AD7381EB067B338FD7B0CB22247225D47

// Encryption functions.
p := []byte("plaintext")
key16 := []byte("12345678abcdefgh")
c, _ := Base64AESCBCEncrypt(p, key16) // A67NhD3RBiNaMgG6HTm8LQ==
p, _ = Base64AESCBCDecrypt(c, key16)  // plaintext

key8 := []byte("12345678")
c, _ := Base64DESCBCEncrypt(p, key8) // UZS/y4By6ksePYMBbvZdig==
p, _ := Base64DESCBCDecrypt(c, key8) // plaintext

key24 := []byte("12345678abcdefgh12345678")
c, _ := Base64TriDESCBCEncrypt(p, key24) // dau0DzmDGQbHasZaOvxxwg==
p, _ := Base64TriDESCBCDecrypt(c, key24) // plaintext
```

# Rand

Some functions to create a real non-negative random int number, specified length random string, and so on.

```go
math.GetRandInt()               // 2040723487295132865
math.GetRandIntn(100)           // 49
math.GetRandIntRange(0, 100)    // 44
math.GetRandByteSlice(3)        // [241 16 101]
math.GetRandStr(3)              // dAt
math.GetRandLowerStr(3)         // lts
math.GetRandUpperStr(3)         // YUT
```

# JSON

Some converting function to json.

```go
student := struct {
    Hobby   string
    Age     int32
}{
    Hobby: "pingpopng",
    Age:   28,
}
encoding.ToIndentJSON(&student)
/*
output:
{
    "Hobby": "pingpopng",
    "Age": 28
}
*/
```

# Comparison

Some useful functions to compare.

```go
// Compare two any type value.
cmp.Cmp(888, 889)       // LT
cmp.Cmp(888, 888)       // EQ
cmp.Cmp(889, 888)       // GT
cmp.Cmp(88.8, 88.9)     // LT
cmp.Cmp(88.8, 88.8)     // EQ
cmp.Cmp(88.9, 88.8)     // GT
cmp.Cmp("abc", "b")     // LT
cmp.Cmp("abc", "abc")   // EQ
cmp.Cmp("b", "abc")     // GT

// Compare semantic version. 
b, _ := cmp.VerGTVer("1.0.5", "1.0.4")   // true
b, _ := cmp.VerLTVer("1.0.5", "2.0.4")   // true
b, _ := cmp.VerGEVer("2.0.4", "2.0.4")   // true
b, _ := cmp.VerLEVer("1.0.5", "1.0.5")   // true
```

# Others

Some useful functions now unclassified. Of course, it may be classified and moved to a new subdirectory in the future.

```go
// NO NOW.
```

# Star History

[![Star History Chart](https://api.star-history.com/svg?repos=dablelv/go-huge-util&type=Date)](https://star-history.com/#dablelv/go-huge-util&Date)

# How to Contribute

We really appreciate any code commits which make this lib powerful. Please follow the rules below to create your pull request.

1. Fork the repository.
2. Add and Commit your changes.
3. Push to your forked repository.
4. Create new pull request.

# Summary

The above examples are just the tip of the iceberg. For more usage, please read the source code.

Due to the limited personal ability, you are welcome to criticize and correct. Of course, welcome to join in the construction of this library.