# noverde

### Executar
> Sem compilar
```bash
$ go run main.go ./contas.csv ./transacoes.csv
```
> Compilando
```bash
$ go build
$ ./noverde ./contas.csv ./transacoes.csv
```

### Testar
> Para realizar os testes é preciso instalar o `goblin` usando o gerenciador de dependências `glide`
```bash
$ glide install
$ go test ./account
```
