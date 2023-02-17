package conf

type Config struct {
	Port    string
	UseCors bool
	Domain  string
	UseTLS  bool
	TLSPem  string
	TLSKey  string
	SSLHost string
}

var (
	Conf = Config{
		Port:    ":8880",                 // 网站访问端口
		UseCors: true,                    // 是否允许跨域访问
		Domain:  ".liruinian.top",        // setcookie 时使用的域名
		UseTLS:  true,                    // 是否使用TLS加密（https）*使用加密需要填写以下字段
		TLSPem:  "api.liruinian.top.pem", // pem路径
		TLSKey:  "api.liruinian.top.key", // key路径
		SSLHost: "api.liruinian.top",     // ssl证书生效的域名
	}
)
