package types

type Postgres struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     int64
}

type GraylogConfig struct {
	Host string
	Port int
}

type WebServer struct {
	Enable bool
	Port   int64
	Url    string
}
type EmailConfig struct {
	Sender     string // email отправителя
	Password   string
	Host       string
	Port       int64
	SenderName string //название отправителя
	SenderLogo string
	IsSendWithEmptySender bool // признак что не прописывать отправителя
}







