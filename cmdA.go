package cmd
import (
	"os"
	"os/exec"
	"log"
)


func ExecuteBAT(){
	// fmt.Println(utf8.RuneCountInString(person.text))	//字节长度	unicode/utf8
	err := os.Chdir("C:\\Program Files\\apache-tomcat-7.0.85")
	// currentDir, _ := os.Getwd()
	if err != nil{
		log.Fatal("change DIR failed: ",err)
	}
	cmd := exec.Command("start ping","127.0.0.1")
	runResult1 := cmd.Run()
	if runResult1 != nil{
		log.Fatal(err)
	}
}
