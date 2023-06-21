package test

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"
)

func TestTime(t *testing.T) {
	for _, d := range []string{
		"2022-07-18",
		"2022-07-19",
		"2022-07-20",
		"2022-07-21",
		"2022-07-22",
		"2022-07-23",
		"2022-07-24",
	} {

		t1, _ := time.Parse("2006-01-02", d)
		t.Logf(t1.Format("Monday"))
	}
}

func f1(s string) int {
	return bytes.Count([]byte(s), nil) - 1
}

func f2(s string) int {
	return strings.Count(s, "") - 1
}

func f3(s string) int {
	return len([]rune(s))
}

func f4(s string) int {
	return utf8.RuneCountInString(s)
}

func TestLen(t *testing.T) {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println([]byte(s))
	fmt.Println(f1(s))
	fmt.Println(f2(s))
	fmt.Println(f3(s))
	fmt.Println(f4(s))
}

func TestLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		j := 0
	loop:
		for {
			for {
				j++
				if j == 1 {
					break loop
				}
				t.Logf("can not print me")
			}
			t.Logf("can not print me")
		}
		t.Logf("%d", i)
	}
}

func Test_1(t *testing.T) {
	names := []string{"lilei", "xiaoming", "hanmeimei"}
	for _, v := range names {
		fmt.Println("pointer:", &v, "name:", v)
	}
}

type History struct {
	Id int64
}

func Test_2(t *testing.T) {
	l := make([]*History, 1)
	history := &History{
		Id: 10,
	}
	l[0] = history
	history = &History{
		Id: 20,
	}
	t.Logf("l: %#v", l[0])
}

func Test_3(t *testing.T) {
	offset := 0
	if true {
		t.Logf("in offset: %d", offset)
		offset := 200
		t.Logf("in offset: %d", offset)
	}
	t.Logf("out offset: %d", offset)
}

func Test_4(t *testing.T) {
	jobs := make(chan func(), 10)
	for i := 0; i < 3; i++ {
		for _, char := range []string{"a", "b", "c"} {

			j := i
			charLocal := char

			f := func() {
				t.Logf("i: %d char: %s", j, charLocal)
			}
			jobs <- f
		}
	}
	for i := 0; i < 9; i++ {
		f := <-jobs
		f()
	}
}

func Test_5(t *testing.T) {
	//l := make([]int, 0)
	//l = append(l, 1)
	//t.Logf("%v", l[:1])
	produceFuncList := make([]func(), 0)
	var wgLocal sync.WaitGroup
	var count int64
	for i := 0; i < 1024; i++ {
		for j := 0; j < 3; j++ {
			produceFuncList = append(produceFuncList, func() {
				defer func() {
					wgLocal.Done()
				}()
				atomic.AddInt64(&count, 1)
				//t.Logf("index: %d", index)
			})
		}
	}
	for part := 0; part < len(produceFuncList)/50+1; part++ {
		start, end := part*50, (part+1)*50
		if start >= len(produceFuncList) {
			break
		}
		if end > len(produceFuncList) {
			end = len(produceFuncList)
		}
		for _, f := range produceFuncList[start:end] {
			wgLocal.Add(1)
			go f()
		}
		wgLocal.Wait()
		t.Logf("run part: %d count: %d", part, atomic.LoadInt64(&count))
	}
}

func Test_MapRead(t *testing.T) {
	var m map[int64]bool
	t.Logf("%t", m[1])
}

func TestTypeTransfer(t *testing.T) {
	type User struct {
		Name string
		Age  int64
	}
	f := func(i interface{}) {
		user := i.(*User)
		t.Logf("%#v", user)
	}
	f(&User{Name: "slk", Age: 30})
	f(1)
}

func TestPaginate(t *testing.T) {
	items := make([]int64, 0)
	for i := int64(0); i < 30; i++ {
		items = append(items, i)
	}
	offset, count := 0, 5
	for {
		start, end := offset, offset+count
		offset = end
		if start > len(items) {
			break
		}
		if end > len(items) {
			end = len(items)
		}
		fmt.Printf("%#v offset: %d\n", items[start:end], offset)
	}
}

// slice可以不用初始化，但是map一定要初始化
func TestSliceDeclare(t *testing.T) {
	var userIdList []int64
	userIdMap := make(map[int64]struct{})
	for i := int64(0); i < 20; i++ {
		userIdList = append(userIdList, i)
		userIdMap[i] = struct{}{}
	}
	t.Logf("userIdList: %#v userIdMap: %#v", userIdList, userIdMap)
}

type Argument struct {
	A int64
}
type RoamingMemoryPaginate struct {
	Param struct {
		Args struct {
			A int64
		}
	}
	Variable struct {
		UserId       int64
		FromCache    bool
		DeviceType   string
		HaveNextPage bool
		NextOffset   int64
	}
	MaxQueryCount int64
}

func TestStructLoop(t *testing.T) {
	mp := new(RoamingMemoryPaginate)
	t.Logf("%d A: %d", mp.Variable.UserId, mp.Param.Args.A)
}

var _gConf *GlobalConfig

type GlobalConfig struct {
	Id int64
}

func InitConfig() {
	_gConf = &GlobalConfig{Id: 1}
}

func GetConfig() GlobalConfig {
	return *_gConf
}

func EditConfig(cfg GlobalConfig) {
	cfg.Id = 2
}

func EditSlice(s []int64) {
	s[0] = 2
}

func TestStructVal(t *testing.T) {
	InitConfig()
	cfg := GetConfig()
	cfg.Id = 2
	t.Logf("%#v", GetConfig())
	cfgV2 := GlobalConfig{Id: 1}
	EditConfig(cfgV2)
	t.Logf("%#v", cfgV2)
	s := []int64{1}
	EditSlice(s)
	t.Logf("%#v", s)
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("%d", rand.Intn(100000))
	}
}

func TestMap(t *testing.T) {
	var m map[int64]string
	m = nil
	fmt.Println(m[10])
}

func TestListNil(t *testing.T) {
	l := make([]int64, 0, 10)
	t.Logf("l==nil is %t len(l)=%d", l == nil, len(l))
}

type Jobs struct {
	TableNum int64
	FileId   int64
}

func TestChannelMemoryCost(t *testing.T) {
	c := make(chan *Jobs, 15000000)
	t.Logf("len(c)=%d", len(c))
	time.Sleep(1 * time.Hour)
}

func TestInt64Slice(t *testing.T) {
	v := make([]int64, 10)
	t.Logf("v=%v", v)
	v2 := make([]string, 10)
	t.Logf("v=%v", v2)
	v3 := make([]rune, 10)
	t.Logf("v=%v", v3)
}
