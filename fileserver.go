package main

import (
	"io"
	"net/http"
	//"fmt"
	"strings"
	//"log"
	. "strconv"
	"syscall"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if !strings.HasSuffix(path, "/") && strings.Count(path, "/") == 1 {
			http.Redirect(w, req, path+"/", http.StatusFound)
			return
		}
		if path != "/" {
			//arr := strings.Split(path,"/")
			//fmt.Fprintf(w,"%s\n",path)
			//fmt.Fprintf(w,"%v",arr)
			h := path[1:2]
			t := path[2:]
			name := h + ":" + t
			//fmt.Fprintf(w,"\n%s,%s-->%s\n",h,t,name)
			//
			http.ServeFile(w, req, name)
			return
		}
		io.WriteString(w, "<p>共享硬盘：</p>")
		for _, s := range GetLogicalDrives() {
			s = s[0:1]
			io.WriteString(w, "<a href='"+s+"/'>"+s+"</a>\n")
		}
	})
	http.ListenAndServe(":8080", nil)
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("c:/"))))
}

func GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := FormatInt(int64(n), 2)

	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}
	temp := drives_all[0:len(s)]

	var d []string
	for i, v := range s {

		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}

	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives

}
