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

huge.JoinSkipEmpty(",", "a", "", "b")   // a,b
huge.JoinNoSkipEmpty(",", "a", "", "b") // a,,b

huge.ReverseStr("abc")                       // cba
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

# other_util
Some useful functions now unclassified.

# Summary
The above example is just the tip of the iceberg. For more usage, please read the source code.

Due to the limited personal ability, you are welcome to criticize and correct. Of course, welcome to join in the construction of this library.