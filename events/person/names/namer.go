package names

import (
	"math/rand"
	"time"
)

//New returns a new set of names
func New(length int) []string {

	rand.Seed(time.Now().UnixNano())

	maleNamesLength := len(maleFirstNames)
	femaleNamesLength := len(femaleFirstNames)
	surnameLength := len(surnames)

	list := make([]string, length)

	for i := 0; i < length; i += 2 {
		list[i] = maleFirstNames[rand.Intn(maleNamesLength)] + " " + surnames[rand.Intn(surnameLength)]
		list[i+1] = femaleFirstNames[rand.Intn(femaleNamesLength)] + " " + surnames[rand.Intn(surnameLength)]
	}
	return list

}
