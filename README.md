# APRENDENDO GO

- Lista de Comandos
    . go run 
    . go build
    . go install
    . go module init "nome do module" 
    . go get "link do pacote" Exemplo : go get github.com/badoux/checkmail
    . go mod tidy //Exclui todas as dependencias


- Criação de Files
    . Todos os arquivos de projetos deve conter o o main.go com a função func main () {}
    . Para criar modulos internos: criar pasta , usar o go module init e inserir a linha package "nome do module" no arquivo .go