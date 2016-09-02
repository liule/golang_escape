package mysql_escape

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
)

var phpCode string = `// use z to seperate, so no z in data
$src = array(
    "normal",
    "aaa\ro", // \r
    "iam\0ooo", // \0
    "\"shou\x08ldbe", // " and \b
    "+123",
    "'youyou'", // '
    "\n gogogo",
    "\026haha",
    "tab \t iam tab",
    "yoyo\\ yoyo",
    "heros never die",
    "heros%",
    "heros_",
    "heros$(&&^*&$##%^(*(*&$##&^^&(*)",
    "??>><<>><?<:||)(&___*((*&%@%@$@#GJGJGFDOIUO*&*JKJKL137813"
);
$dest ="";

foreach ($src as $s){
    $dest = $dest."z".$s."z".mysql_escape_string($s);
}
file_put_contents("sql_util_test.dat", $dest);
`

// the test data is generated by php, using mysql_escape_string. // data are sperated by 'z'
// phpCode are pasted above.

func TestMysqlEscapeString(t *testing.T) {
	Convey("testMysqlEscape", t, func() {
		data, err := ioutil.ReadFile("test.dat")
		So(err, ShouldBeNil)
		pairs := bytes.Split(data, []byte("z"))
		pairs = pairs[1:]
		for i := 0; i < len(pairs); i += 2 {
			originalData := pairs[i]
			destData := pairs[i+1]
			escaped := MysqlEscapeString(string(originalData))
			byteEqual(t, []byte(escaped), destData)
		}
	})
}

func byteEqual(t *testing.T, src, dest []byte) {
	t.Log(src)
	t.Log(dest)
	if len(src) != len(dest) {
		t.Error("len check fail")
	}
	for i, v := range src {
		if dest[i] != v {
			t.Error("fail in index %v", i)
		}
	}
}
