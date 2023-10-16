# APRENDENDO GO

- Lista de Comandos
    - go run 
    - go build
    - go install
    - go module init "nome do module" 
    - go get "link do pacote" Exemplo : go get github.com/badoux/checkmail
    - go mod tidy //Exclui todas as dependencias


- Criação de Files
    - Todos os arquivos de projetos deve conter o o main.go com a função func main () {}
    - Para criar modulos internos: criar pasta , usar o go module init e inserir a linha package "nome do module" no arquivo .go

- Tipos de Dados
    - Declaração explicita:
        - string -> variable string = "value
        - int -> variable int = 10
            - int - Pega a quantidade de bytes da máquina onde está sendo executado
            - int16 - 16 bytes
            - int8 -  8 bytes
            - int32  - 32 bytes
            - int64 - 64 bytes
        - float -> variable float = 2.333
            - float - Pega quantidade de bytes da maquina
            - float32,float64
    - Declaração implicita
        - variable := "value" --> Automaticamente string
        - variable := 10   --> int (valor default)
        - variable := 3.222 --> float
    