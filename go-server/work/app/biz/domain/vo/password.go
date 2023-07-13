package vo

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofrs/uuid"
	"go-service/pkg/errno"
	"regexp"
	"strings"
)

type Password struct {
	plainPwd   string
	salt       string
	credential string
}

func NewPassword(plainPwd string) Password {
	return Password{
		plainPwd: strings.TrimSpace(plainPwd),
	}
}

func NextRandomPassword() Password {
	return Password{
		plainPwd: NewRandGenerator(RandPasswordGenerator).Next(),
	}
}

func NewEncryptedPassword(salt, credential string) Password {
	return Password{
		salt:       salt,
		credential: credential,
	}
}

func (p Password) Encrypt() Password {
	return p.encrypt()
}

func (p Password) CheckFormat() (bool, error) {
	if p.salt != "" && p.credential != "" {
		return true, nil
	}
	if p.plainPwd == "" {
		return false, errno.NewSimpleBizError(errno.ErrPasswordEmpty, nil)
	}
	if !isPwdValid(p.plainPwd) {
		return false, errno.NewSimpleBizError(errno.ErrInvalidPhoneFormat, nil)
	}
	return true, nil
}

func (p Password) GetPlainPassword() string {
	return p.plainPwd
}

func (p Password) GetSalt() string {
	return p.salt
}

func (p Password) GetCredential() string {
	return p.credential
}

func (p Password) encrypt() Password {
	if p.plainPwd == "" {
		return p
	}
	if p.salt != "" && p.credential != "" {
		return p
	}
	if u, err := uuid.NewV4(); err == nil {
		p.salt = u.String()
	}
	p.credential = encryptPwd(p.salt, p.plainPwd)
	return p
}

func encryptPwd(salt string, pwd string) (result string) {
	return sha256Encrypt(sha256Encrypt(pwd) + salt)
}

func sha256Encrypt(str string) (result string) {
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return s
}

func isPwdValid(pwd string) bool {
	if len(pwd) > 20 && len(pwd) < 8 {
		return false
	}
	var (
		numberReg = `[0-9]{1}`
		letterReg = `[a-zA-Z]{1}`
		symbolReg = `[!"#$%&'()*+,-./:;<=>?@[\]^_{|}\~]{1}`
		pattern   = `[A-Za-z\d!"#$%&'()*+,-./:;<=>?@[\]^_{|}\~]{8,20}$`
	)
	if !regexp.MustCompile(pattern).MatchString(pwd) {
		return false
	}
	sum := 0
	if ok, _ := regexp.MatchString(numberReg, pwd); ok {
		sum++
	}
	if ok, _ := regexp.MatchString(letterReg, pwd); ok {
		sum++
	}
	if ok, _ := regexp.MatchString(symbolReg, pwd); ok {
		sum++
	}
	return sum >= 2
}
