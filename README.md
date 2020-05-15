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

huge.UniqueIntSlice([]int{1,2,3,3,4})   // [1,2,3,4]
huge.UniqueUintSlice([]int{1,2,3,3,4})  // [1,2,3,4]

huge.UniqueUintSlice([]string{"a", "b", "b", "c"})  // [a b c]
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

```

# struct_util
Some useful functions can be used to handle string.

# time_util
Some useful functions can be used to handle date and time.

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
```

# other_util
Some useful functions now unclassified.

# Summary
Due to the limited personal ability, you are welcome to criticize and correct. Of course, welcome to join in the construction of this library.