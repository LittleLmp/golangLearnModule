package main
import (
	"os"
	"os/exec"
	"log"
	"time"
	"net/smtp"
	"strings"
	"fmt"
)

func main(){
	// auto restart web APP 
	time.Sleep(1*time.Second)
	result := Fn_skoda()
	time.Sleep(2*time.Second)
	if result != "" {
		Send(result)
		log.Fatal(result)
	}
	time.Sleep(5*time.Second)
	// 高效拼接字符串（'+'也可以的）
	// 利用bytes.Buffer缓冲
	// WriteString()将字节写入到buffer中
	// String()将buffer中的字符串拼接起来
	// http.HandleFunc("/",handler)
	// log.Fatal(http.ListenAndServe(":8080",nil))
	// fmt.Println(utf8.RuneCountInString(person.text))	//字节长度	unicode/utf8
}

func Fn_skoda() string{
	err := ""
	err1 := os.Chdir("D:\\Program Files\\tomcat\\apache-tomcat-7.0.85\\bin")
	// currentDir, _ := os.Getwd()
	// fmt.Println(currentDir)
	if err1 != nil{
		log.Fatal("change DIR failed: ",err1)
		err = err+"change DIR failed - "
	}

	err2 := exec.Command(".\\shutdown.bat").Start()
	if err2 != nil{
		log.Fatal("run shell-script 'shutdown.bat' failed: ",err2)
		err = err+"run shell-script 'shutdown.bat' failed - "
	}
	time.Sleep(3*time.Second)
	err3 := exec.Command("taskkill","/f","/t","/im","java.exe").Start()
	if err3 != nil{
		log.Fatal("kill process 'java.exe' failed: ",err3)
		err = err+"kill process 'java.exe' failed - "
	}
	time.Sleep(2*time.Second)
	err4 := exec.Command(".\\startup.bat").Start()
	if err4 != nil{
		log.Fatal("run shell-script 'startup.bat' failed: ",err4)
		err = err+"run shell-script 'startup.bat' failed "
	}
	time.Sleep(5*time.Second)
	return err
}

// send E-mail Controller!
func Send(result string){
	user := "m15836786217@aliyun.com"
	password := "52q,salyx112110"
	host := "smtp.aliyun.com:25"
	// smtp.aliyun.com	smtp.163.com
	to := "m15836786217@163.com"
	subject := "Skoda web APP Auto restart"
	body := `
		<html>
		<body>
		<h3>
		` +result+ `
		</h3>
		</body>
		</html>
		`
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}
// send E-mail main body!
func SendMail(user, password, host, to, subject, body, mailtype string) error{
	hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
    } else {
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }

    msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}
