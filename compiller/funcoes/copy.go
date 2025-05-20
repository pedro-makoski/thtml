package funcoes

import (
	"path/filepath"
	"thtml/file"
)

func Copy(line string, data map[string]any) error {
	args, err := GetQuery(map[string]ValueOptionsNumber{
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
	}, line)
	if err != nil {
		return err
	}

	args, err = ConcatAll(args, data)
	if err != nil {
		return err
	}

	origem := args[0]
	pastaOrigem := args[1]
	origem = filepath.Join(origem, pastaOrigem)
	nomeDoNovoArquivo := args[2]
	if nomeDoNovoArquivo == "" {
		nomeDoNovoArquivo = args[0]
	}
	pastaDestino := args[3]
	destino := filepath.Join(pastaDestino, nomeDoNovoArquivo)

	err = file.InteligentCopy(origem, destino)

	return nil
}
