package vo

import "crypto/rand"

type UserUid struct {
	uid string
}

func NewUserUid(uid string) UserUid {
	return UserUid{uid: uid}
}

func NextUserUid() UserUid {
	return UserUid{uid: generateNextUid()}
}

func (u UserUid) GetUid() string {
	return u.uid
}

func generateNextUid() string {
	return "1000" + randomUid()
}

func randomUid() string {
	prime, err := rand.Prime(rand.Reader, 24)
	for ; err != nil; prime, err = rand.Prime(rand.Reader, 24) {
	}
	random := prime.String()
	if len(random) < 8 {
		for i := len(random); i < 8; i++ {
			random = "0" + random
		}
	} else if len(random) > 8 {
		random = string([]rune(random)[:8])
	}
	return random
}
