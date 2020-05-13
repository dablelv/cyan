[toc]

# Synopsis
Go common and huge utility functions help you to create your go program quickly and easily.

# encode_util
Some useful functions can be used to handle conversion of different character encoding, such as gbk to utf8.

# net_util
Some useful functions can be used to handle network. For example you can use `IPv4StrToU32()` transform ipv4 string to uint32 value.

# slice_util
Some useful functions can be used to handle slice.

# sql_util
Some useful functions can be used to handle sql statement.

# string_util
Some useful functions can be used to handle string.

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