TINYINT               127                            占1byte

UNSIGNED TINYINT       255                   占1byte

SMALLINT               32768                         占2byte

UNSIGNED SMALLINT    65535                占2byte

INT                  2147483647                         占4byte

UNSIGNED INT    4294967295                占4byte

BIGINT      9223372036854775807             占8byte

UNSIGNED BIGINT  18446744073709551615  占8byte

DATETIME    1000-01-01 00:00:00/9999-12-31 23:59:59    占8byte

TIMESTAMP   1970-01-01 00:00:00/2038-1-19 11:14:07     占4byte

1. mysql 开发中有些会定义成 int(8), 其实是没个卵用，因为 int 类型永远都是占
4个字节，int(8)中的8表示的是显示宽度，不足用0 补充，超过的无视，但要完整显示，
还是需要设置unsigned zerofill 才生效
2. 定义datetime 和timestamp 时，因为占用大小不同，一个8一个4，所能表示的时间
范围也不同， datetime 能显示0000到9999年， timestamp 能显示到2038年，还有一个
区别，datetime类型插入为null时，该列就是null, 但timestamp默认不为空，插入为null
时，取当前默认值填充