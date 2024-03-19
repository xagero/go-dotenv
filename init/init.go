package init

import (
	"github.com/xagero/go-dotenv"
	"os"
)

func init() {

	filename := ".env"
	if _, err := os.Stat(filename); err == nil {
		dotenv.ReadFromFile(filename)
	}
}
