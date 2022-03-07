package Other

import (
	"flag"
	"fmt"
	"os"
)

var (
	h          bool
	TargetFile string
	Thread     int
)

/*开始把这一块写main函数了，还又写了初始化函数用来初始化这个值，其实大可不必，放其他包里不就好了，本来也是引用类型的，也不怕啥，哪里调用都行*/
func init() {
	fmt.Println(`__   ___   _ _____ _   _      ____  _____ ____ 
\ \ / / | | |_   _| | | |    / ___|| ____/ ___|
 \ V /| | | | | | | | | |____\___ \|  _|| |    
  | | | |_| | | | | |_| |_____|__) | |__| |___ 
  |_|  \___/  |_|  \___/     |____/|_____\____|

`)
	flag.StringVar(&TargetFile, "TF", "", "目标文件")
	flag.IntVar(&Thread, "t", 500, "并发数量")
	flag.BoolVar(&h, "h", false, "Help")

	// 修改提示信息

	flag.Usage = usage
	flag.Parse()
	if h || (TargetFile == "") {
		flag.Usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `YUTU-SEC
Usage: SpEL.exe -TF url.txt -t 100
Options:
`)
	flag.PrintDefaults()

}
