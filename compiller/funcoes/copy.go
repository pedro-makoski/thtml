package funcoes

import (
	"path/filepath"
	"thtml/compiller/querycomand"
	"thtml/file"
)

func Copy(line string, data map[string]any) error {
	args, err := querycomand.GetQuery(map[string]querycomand.ValueOptionsNumber{
		"--": {
			Required: true,
			Number:   0,
		},
		"from": {
			Required: true,
			Number:   1,
		},
		"as": {
			Required: false,
			Number:   2,
		},
		"on": {
			Required: true,
			Number:   3,
		},
	}, line, true, data, []string{"--", "from", "as", "on"})
	if err != nil {
		return err
	}

	origem := args[0]
	pastaOrigem := args[1]
	origem = filepath.Join(pastaOrigem, origem)
	nomeDoNovoArquivo := args[2]
	if nomeDoNovoArquivo == "" {
		nomeDoNovoArquivo = filepath.Base(args[0])
	}
	pastaDestino := args[3]
	destino := filepath.Join(pastaDestino, nomeDoNovoArquivo)

	err = file.InteligentCopy(origem, destino)

	return err
}
