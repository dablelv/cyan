# Synopsis
Go common and huge utility functions help you to create your go program quickly and easily.

# encode_util
Some useful functions can be used to handle conversion of different character encoding, such as gbk to utf8.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

gbkStr := []byte{0xC4, 0xE3, 0xBA, 0xC3} // 你好 in gbk
utf8Str, _ := huge.GbkToUtf8(gbkStr)     // 你好 in utf8
fmt.Println(string(utf8Str))             // 你好

gbkStrRes, _ := huge.Utf8ToGbk(utf8Str)
fmt.Println(gbkStrRes) // [196 227 186 195]
```

# net_util
Some useful functions can be used to handle network. For example you can use `IPv4StrToU32()` transform ipv4 string to uint32 value.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

huge.IsReservedIP("0.0.0.256")    // -1 invalid ip
huge.IsReservedIP("39.156.69.79") // 0 public ip
huge.IsReservedIP("127.0.0.1")    // 1 reserved ip

huge.GetNativeEndian()              // LittleEndian
huge.IPv4StrToU32("127.0.0.1")      // 2130706433
huge.U32ToIPv4Str(2130706433)       // "127.0.0.1"
```
# slice_util
Some useful functions can be used to handle slice.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

huge.UniqueIntSlice([]int{1, 2, 2, 3})              // [1 2 3]
huge.UniqueUintSlice([]uint{1, 2, 2, 3})            // [1 2 3]
huge.UniqueStrSlice([]string{"a", "b", "b", "c"})   // [a b c]

huge.ReverseIntSlice([]int{1, 2, 3})                // [3 2 1]
huge.ReverseUintSlice([]uint{1, 2, 3})              // [3 2 1]
huge.ReverseStrSlice([]string{"a", "b", "c"})       // [c b a]

huge.SumSlice([]int{1,2,3})                         // 6
huge.SumSlice([]uint{1,2,3})                        // 6
huge.SumSlice([]float32{1.1, 2.2, 3.3})             // 6.6
huge.SumSlice([]float64{1.1, 2.2, 3.3})             // 6.6

huge.JoinSliceWithSep([]int{1, 2, 3}, ",")              // 1,2,3
huge.JoinSliceWithSep([]uint{1, 2, 3}, ",")             // 1,2,3
huge.JoinSliceWithSep([]float64{1.1, 2.2, 3.3}, ",")    // 1.1,2.2,3.3
huge.JoinSliceWithSep([]string{"a", "b", "c"}, ",")     // a,b,c

// CRUD(Create Read Update Delete) on slice by index
fib := []int{1, 1, 2, 3, 5, 8}
r, _ := huge.InsertSliceE(fib, 5, 13)			// [1 1 2 3 5 8 13]
r, _ := huge.DeleteSliceE(fib, 0)				// [1 2 3 5 8]
r, _ := huge.UpdateSliceE(fib, 5, 88)			// [1 1 2 3 5 88]
r, _ := huge.GetEleIndexesSliceE(fib, 1)		// [0 1]

// or
r := huge.InsertIntSlice(fib, 5, 13)		// [1 1 2 3 5 8 13]
r := huge.DeleteIntSliceE(fib, 0)			// [1 2 3 5 8
r := huge.UpdateIntSliceE(fib, 5, 88)		// [1 1 2 3 5 88]
r := huge.GetEleIndexesSlice(fib, 1)		// [0 1]
```

# sql_util
Some useful functions can be used to handle sql statement.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

var sql = `select * from t where field1="{name}"`
huge.FormatSql(sql, map[string]string{"name": "dablelv"}, false)    // select * from t where field1="dablelv"
```

# string_util
Some useful functions can be used to handle string.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

huge.StrToMapTrue("a,b,c", ",")     // map[a:true b:true c:true]
huge.StrToMapFalse("a,b,c", ",")    // map[a:false b:false c:false]
huge.StrToIntSlice("1,2,3", ",")    // [1 2 3]
huge.StrToUintSlice("1,2,3", ",")   // [1 2 3]

huge.Split("a,b,c", ",")            // [a b c]

huge.JoinStrSkipEmpty(",", "a", "", "b")    // a,b
huge.JoinStrNoSkipEmpty(",", "a", "", "b")  // a,,b

huge.ReverseStr("abc")                      // cba
```

# struct_util
Some useful functions can be used to handle string.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

var st = struct {
    I int
    S string
}{I: 1, S: "a"}

huge.Struct2Map(st)         // map[I:1 S:a]
huge.Struct2MapString(st)   // map[I:1 S:a]
```

# time_util
Some useful functions can be used to handle date and time.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

tc := huge.NewTimeCounter()
// ...                          // do your statements
tc.GetMs()                      // get the time cost in millisecond

huge.GetNowDate()               // just a example, the same below. 2020-05-16
huge.GetNowTime()               // 00:15:42
huge.GetNowDateTime()           // 2020-05-16 00:15:42
huge.GetNowDateTimeZ()          // 2020-05-16 00:15:42 +08:00

huge.GetNowS()                  // 1589559342
huge.GetNowMs()                 // 1589559342963
huge.GetNowUs()                 // 1589559342963062
huge.GetNowNs()                 // 1589559342964063200

huge.GetDayBeginMoment(time.Now())  // 2020-05-16 00:00:00 +0800 CST
huge.GetDayEndMoment(time.Now())    // 2020-05-16 23:59:59.999999999 +0800 CST
```

# typeconv_util
Some useful functions can be used to type conversion. The most functions are from github.com/spf13/cast and some new functions is added by myself. 

Example ToString:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

huge.ToString("mayonegg")         // "mayonegg"
huge.ToString(8)                  // "8"
huge.ToString(8.31)               // "8.31"
huge.ToString([]byte("one time")) // "one time"
huge.ToString(nil)                // ""

var foo interface{} = "one more time"
huge.ToString(foo)                // "one more time"

// support type definition
type Int32Alias int32
var i32Alias IntAlias = 1
huge.ToString(i32Alias)         // "1"

type Uint32Alias uint32
var u32Alias Uint32Alias = 1
huge.ToString(u32Alias)         // "1"

type StringAlias string
var sAlias StringAlias = "a"
huge.ToString(sAlias)         // "a"
```

Example ToInt:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

huge.ToInt(8)                  // 8
huge.ToInt(8.31)               // 8
huge.ToInt("8")                // 8
huge.ToInt(true)               // 1
huge.ToInt(false)              // 0

var eight interface{} = 8
huge.ToInt(eight)              // 8
huge.ToInt(nil)                // 0
```

# url_util
Some useful functions can be used to handle url.

Example:
```
import (
    huge "github.com/dablelv/go-huge-util"
)

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

# crypto_util
Some useful functions can be used to create Hash, HMAC and crypt data.
```
import (
    huge "github.com/dablelv/go-huge-util"
)

// Hash functions
huge.Hash_md5("")   // d41d8cd98f00b204e9800998ecf8427e
huge.Hash_MD5("")   // D41D8CD98F00B204E9800998ECF8427E

huge.Hash_sha1("")  // da39a3ee5e6b4b0d3255bfef95601890afd80709
huge.Hash_SHA1("")  // DA39A3EE5E6B4B0D3255BFEF95601890AFD80709

huge.Hash_sha224("")    // d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f
huge.Hash_SHA224("")    // D14A028C2A3A2BC9476102BB288234C415A2B01F828EA62AC5B3E42F

huge.Hash_sha256("")    // e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
huge.Hash_SHA256("")    // E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855

huge.Hash_sha384("")    // 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b
huge.Hash_SHA384("")    // 38B060A751AC96384CD9327EB1B1E36A21FDB71114BE07434C0CC7BF63F6E1DA274EDEBFE76F65FBD51AD2F14898B95B

huge.Hash_sha512("")    // cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e
huge.Hash_SHA512("")    // CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E

// HMAC functions
huge.HMAC_md5("", "")   // 74e6f7298a9c2d168935f58c001bad88
huge.HMAC_MD5("", "")   // 74E6F7298A9C2D168935F58C001BAD88

huge.HMAC_sha1("", "")  // fbdb1d1b18aa6c08324b7d64b71fb76370690e1d
huge.HMAC_SHA1("", "")  // FBDB1D1B18AA6C08324B7D64B71FB76370690E1D

huge.HMAC_sha224("", "")    // 5ce14f72894662213e2748d2a6ba234b74263910cedde2f5a9271524
huge.HMAC_SHA224("", "")    // 5CE14F72894662213E2748D2A6BA234B74263910CEDDE2F5A9271524

huge.HMAC_sha256("", "")    // b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad
huge.HMAC_SHA256("", "")    // B613679A0814D9EC772F95D778C35FC5FF1697C493715653C6C712144292C5AD

huge.HMAC_sha384("", "")    // 6c1f2ee938fad2e24bd91298474382ca218c75db3d83e114b3d4367776d14d3551289e75e8209cd4b792302840234adc
huge.HMAC_SHA384("", "")    // 6C1F2EE938FAD2E24BD91298474382CA218C75DB3D83E114B3D4367776D14D3551289E75E8209CD4B792302840234ADC

huge.HMAC_sha512("", "")    // b936cee86c9f87aa5d3c6f2e84cb5a4239a5fe50480a6ec66b70ab5b1f4ac6730c6c515421b327ec1d69402e53dfb49ad7381eb067b338fd7b0cb22247225d47
huge.HMAC_SHA512("", "")    // B936CEE86C9F87AA5D3C6F2E84CB5A4239A5FE50480A6EC66B70AB5B1F4AC6730C6C515421B327EC1D69402E53DFB49AD7381EB067B338FD7B0CB22247225D47

// Encryption functions
p := []byte("plaintext")
key16 := []byte("12345678abcdefgh")
c, _ := Base64AESCBCEncrypt(p, key16) // A67NhD3RBiNaMgG6HTm8LQ==
p, _ = Base64AESCBCDecrypt(c, key16)  // plaintext

key8 := []byte("12345678")
c, _ := Base64DESCBCEncrypt(p, key8) // UZS/y4By6ksePYMBbvZdig==
p, _ := Base64DESCBCDecrypt(c, key8) // plaintext

key24 := []byte("12345678abcdefgh12345678")
c, _ := Base64TripleDESCBCEncrypt(p, key24) // dau0DzmDGQbHasZaOvxxwg==
p, _ := Base64TripleDESCBCDecrypt(c, key24) // plaintext
```

# rand_util
Some functions to create a real non-negative random int number, specified length random string, and so on.
```
huge.GetRandInt()               // 2040723487295132865
huge.GetRandIntn(100)           // 49
huge.GetRandIntRange(0, 100)    // 44
huge.GetRandByteSlice(3)        // [241 16 101]
huge.GetRandStr(3)              // dAt
huge.GetRandLowerStr(3)         // lts
huge.GetRandUpperStr(3)         // YUT
```

# other_util
Some useful functions now unclassified. Of course, it may be classified and moved to a new go file in the future.
```
huge.IsLittleEndian()   // true

student := struct {
    Hobby   string
    Age     int32
}{
    "hobby": "pingpopng",
    "age":   28,
}
huge.ToFormattedJSON(&student)
/*
output:
{
    "Hobby": "pingpopng",
    "Age": 28
}
*/
```

# Summary
The above example is just the tip of the iceberg. For more usage, please read the source code.

Due to the limited personal ability, you are welcome to criticize and correct. Of course, welcome to join in the construction of this library.