# go-queue-worker

# Kafka Worker em Go

Este projeto implementa um worker em Go que consome mensagens de um tópico Kafka, processa as mensagens e atualiza um cache Redis com os dados processados.

## Arquivos Principais

- **main.go**: Ponto de entrada da aplicação.
- **config.json**: Arquivo de configuração contendo detalhes do Kafka e Redis.
- **internal/config/config.go**: Carrega a configuração do arquivo config.json.
- **internal/kafka/consumer.go**: Implementa o consumidor Kafka.
- **internal/database/cache.go**: Atualiza o cache Redis com os dados processados.
