# APACHE KAFKA

 __"O Apache Kafka é uma plataforma distribuída de streaming de eventos open-source que é utilizada por milhares de empresas para uma alta performance em pipeline de dados, stream, de analytics, integração de dados e aplicações de missão crítica."__
 https://kafka.apache.org

zukeeper

##  Repositório do Projeto
    https://github.com/codeedu/fc2-kafka.git

__Host confluent local__
> http://localhost:9021/

__Criar topicos__
> kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3 

__Listar Topics__
> kafka-topics --list --bootstrap-server=localhost:9092

__Describe Topics__
> kafka-topics -topic=teste --bootstrap-server=localhost:9092

__Consumir topicos__
> kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning  

__Produzir mensagem para um topic__
> kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste 

__Criar topico em group Consumer Groups__
> kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning  --group=x

__Listar consumer grups__
> kafka-consumer-groups --bootstrap-server=localhost:9092 --group=x --describe

__Git Lib Kafka__
> https://github.com/edenhill/librdkafka/

__extra_hosts - "host.docker.internal:172.17.0.1"__

    Permite comunicação de qualquer conrtainer na rede docker
    é necessário que o /etc/hosts tenha uma linha configurado da seguinte forma
    127.0.0.1 docker.host.internal

__Acessar o container do go__
> docker exec -it gokafka bash

__Iniciando modulo Go dentro do container go__
> go mod init github.com/ricardosantossantana/fc2-gokafka

__Lembrete 01__
	O parametro "bootstrap.servers": deve ser preenchido com o nome do container que estiver rodando o kafka
    
    configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka_kafka_1:9092",
	}