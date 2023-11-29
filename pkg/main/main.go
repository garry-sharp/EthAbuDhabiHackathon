package main

import (
	"bufio"
	"os"
	"xxx/pkg/config"
	"xxx/pkg/contract"
)

func main() {
	config.ParseConfig()
	conf := config.GetConfig()

	if conf.Deploy {
		contract.Deploy()
	}

	if conf.Batch {
		_tkns := []*config.Token{}
		for i := 0; i < len(conf.Tokens); i++ {
			_tkns = append(_tkns, &(conf.Tokens[i]))
		}
		for i := 0; i < len(conf.Dexes); i++ {
			contract.UniswapV2DexQuery(&conf.Dexes[i], _tkns)
		}
		contract.UniswapV2PricePusher(_tkns)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

}
