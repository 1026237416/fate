package fate_test

import (
	"github.com/godcong/fate"
	"testing"
	"time"
)

func TestFate_FirstRunInit(t *testing.T) {
	f := fate.NewFate("张", time.Now())
	eng := fate.InitMysql("localhost:3306", "root", "111111")

	f.SetDB(eng)
	e := f.MakeName()
	if e != nil {
		t.Fatal(e)
	}
}
