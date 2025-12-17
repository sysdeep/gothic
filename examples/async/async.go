package main

import (
	"fmt"
	"time"

	"github.com/sysdeep/gothic/gothic"
)

func main() {
	ir := gothic.NewInterpreter(`
		pack [ttk::progressbar .bar1] -padx 20 -pady 20
		pack [ttk::progressbar .bar2] -padx 20 -pady 20
		pack [tk::label .label -text lbl] -padx 20 -pady 20

		pack [ttk::button .exit_btn -text Exit -command exit ] -padx 20 -pady 20

		bind . <Control-q> { exit }
	`)

	go func() {
		i := 0
		inc := -1
		for {
			if i > 99 || i < 1 {
				inc = -inc
			}
			i += inc
			time.Sleep(5e7)
			ir.Eval(`.bar1 configure -value %{}`, i)
		}
	}()

	go func() {
		i := 0
		inc := -1

		for {
			if i > 99 || i < 1 {
				inc = -inc
			}
			i += inc
			time.Sleep(1e8)
			ir.Eval(`.bar2 configure -value %{}`, i)
		}
	}()

	go func() {
		i := 0
		inc := -1

		for {
			if i > 99 || i < 1 {
				inc = -inc
			}
			i += inc
			time.Sleep(time.Second * 1)
			ir.Eval(fmt.Sprintf(".label configure -text %d", i))
		}
	}()

	<-ir.Done
}
