package tools

import (
	"Homework_Refactor/conf"
	"github.com/gookit/color"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	DB   *gorm.DB
	Conf conf.Config
)

func init() {
	DB = GetDB()
	Conf = conf.Conf
}

func GetPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}

func VerifyUserIfAdmin(username string, lToken string, aToken string) bool {
	if VerifyUser(username, lToken) && VerifyToken(aToken, "admin_token", username+" is admin") {
		return true
	}
	return false
}
func VerifyUser(username string, lToken string) bool {
	if VerifyToken(lToken, "user_token", username+time.Now().Format("2023-02-01")) {
		return true
	}
	return false
}
func CreateToken(tokenName string, tokenCtx string) string {
	token, err := JwtEncoder(tokenName, tokenCtx, int64(3*time.Hour))
	if err != nil {
		log.Println(color.FgRed.Render("Create Token Failed! name:" + tokenName + " ctx:" + tokenCtx))
		log.Println(color.FgRed.Render(err.Error()))
		return ""
	}
	return token
}

func VerifyToken(token string, tokenName string, tokenCtx string) bool {
	name, ctx, err := JwtDecoder(token)
	if err != nil {
		log.Println(color.FgRed.Render("Verify Token Failed! Decoder Error!"))
		log.Println(color.FgRed.Render(err.Error()))
		return false
	}
	if name != tokenName {
		log.Println(color.FgRed.Render("Verify Token Failed! name Not Valid!"))
		log.Println(color.FgRed.Render(err.Error()))
	} else {
		if ctx == tokenCtx {
			return true
		}
	}

	return false
}
