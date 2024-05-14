package bridge

import (
	"github.com/westarver/constantly/generate"
	"github.com/westarver/constantly/app"
)

func LoadFromFile() {
	generate.LoadFromFile()
}

func SaveToFile(bool) {
	generate.SaveToFile(bool)
}

func PreviewString() {
	generate.PreviewString()
}

func WriteConstants() {
	generate.WriteConstants
}
