package live

import (
	"fmt"
	"os/exec"
)

type Siri struct {
	statements chan string
}

var siri *Siri

func init() {
	siri = &Siri{
		statements: make(chan string),
	}
	go func() {
		for str := range siri.statements {
			cmd := exec.Command("say", "--voice=Mei-Jia", str)

			//cmd := exec.Command("say", "--voice=Sin-ji", str)
			_, err := cmd.CombinedOutput()
			// 检查命令是否执行成功
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
}

func Speak(str string) {
	siri.statements <- str
}
