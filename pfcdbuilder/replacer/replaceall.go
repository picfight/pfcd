package replacer

import "github.com/jfixby/coinknife"

func ReplaceAll(data string) string {
	data = coinknife.Replace(data, "decred/dcrd", "picfight/pfcd")
	return data
}
