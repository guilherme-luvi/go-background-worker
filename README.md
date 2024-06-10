# go-queue-worker

# Kafka Worker em Go

Este projeto implementa um worker em Go que consome mensagens de um tópico Kafka, processa as mensagens e atualiza um cache Redis com os dados processados.

## Arquivos Principais

- **main.go**: Ponto de entrada da aplicação.
- **config.json**: Arquivo de configuração contendo detalhes do Kafka e Redis.
- **pkg/config/config.go**: Carrega a configuração do arquivo config.json.
- **internal/kafka/consumer.go**: Implementa o consumidor Kafka.
- **internal/database/cache.go**: Atualiza o cache Redis com os dados processados.

## Explicando execução principal:

### Arquivo Consumer.go

A função Consume é responsável por consumir mensagens de um tópico Kafka:

- A função Consume recebe um contexto ctx e as configurações da aplicação como parâmetros.
- Ela cria um canal chamado messageChannel que é usado para passar mensagens do consumidor Kafka para os workers.
- Inicia um pool de workers. O número de workers é determinado pelo valor NumWorkers na configuração. 
- Cada worker é uma goroutine que executa a função worker.
- A função entra em um loop infinito onde ela lê continuamente mensagens do consumidor Kafka. Se o contexto for cancelado, ela fecha o canal de mensagens e espera que todos os workers terminem antes de retornar.
- Se uma mensagem é lida com sucesso do consumidor Kafka, ela é enviada para o canal de mensagens. 
- Os workers estão lendo continuamente deste canal e processam qualquer mensagem que recebem.

A função worker é responsável por processar as mensagens:
- Ela lê continuamente do canal de mensagens até que o canal seja fechado ou o contexto seja cancelado. 
- Para cada mensagem, ela chama a função handleMessage para processar a mensagem. 
- Se ocorrer um erro ao processar a mensagem, ele é registrado.