package E

import "flag"

func E设置命令行(name string, defaultvalue string, help string, value *string) {
	flag.StringVar(value, name, defaultvalue, help)
}

func E命令行解析() {
	flag.Parse()
}
