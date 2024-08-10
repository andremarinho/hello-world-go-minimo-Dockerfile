# Hello World Go Mínimo com Docker

Este repositório contém um exemplo básico de um aplicativo "Hello World" em Go, com um Dockerfile mínimo para empacotamento e execução.

## Estrutura do Projeto

main.go: Arquivo principal do aplicativo Go.
Dockerfile: Arquivo de configuração para criação da imagem Docker.
Pré-requisitos
Go instalado
Docker instalado
Como Executar
Localmente
Compile o aplicativo Go:




## Criando imagem, rodando container e enviando para do docker hub:

```python

#Gere a imagem
docker build -t hello-world-go .



#tagueie a imagem
docker tag hello-world-go seu-usuario/hello-world-go

#Faça login no docker
docker login

#Subindo a imagem para o docker hub
docker push seu-usuario/hello-world-go

#Realizar o pull da imagem
docker pull andrebezerramarinho/fullcycle:latest

#Execute o container
docker run --rm andrebezerramarinho/fullcycle:latest

#A saída esperada é:
Full Cycle Rocks!!

```

## Contributing

Contribuição
Sinta-se à vontade para abrir issues ou enviar pull requests para melhorias e correções.

## License

[MIT](https://choosealicense.com/licenses/mit/)
