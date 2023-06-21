package source_code

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestTimeIsZero(t *testing.T) {
	var now time.Time
	t.Logf("now.IsZero() = %v", now.IsZero())
}

func TestZeroTime(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2022-11-06")
	t2, _ := time.Parse("2006-01-02", "2022-11-05")
	t.Logf("t1.Sub(t2)=%f", t1.Sub(t2).Hours())
}

func TestDailyLoop(t *testing.T) {
	closeChan := make(chan struct{}, 0)
	_sigs := make(chan os.Signal, 1)
	signal.Notify(_sigs, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
	alreadyRunDate := ""
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//defer func() {
	//	wg.Done()
	//}()
	go func() {
		msg := <-_sigs
		t.Logf("msg: %#v", msg)
		close(closeChan)
		time.Sleep(1 * time.Second)
	}()
ForLoop:
	for {
		select {
		case <-closeChan:
			break ForLoop
		default:
		}

		now := time.Now()
		//now = time.Unix(1670256005, 0)
		t.Logf("now=%s now.Hour()=%d now.Minite()=%d alreadyRunDate: %s", now, now.Hour(), now.Minute(), alreadyRunDate)
		if now.Hour() == 14 && now.Minute() <= 15 && now.Format("2006-01-02") != alreadyRunDate {
			alreadyRunDate = now.Format("2006-01-02")
			t.Logf("run")
		} else {
			select {
			case <-closeChan:
				t.Logf("close chan")
				break ForLoop
			case <-time.After(10 * time.Second):
				t.Logf("sleep over")
			}
		}
	}
	//}()

	//wg.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			t.Logf("tick")
		}
	}
}
