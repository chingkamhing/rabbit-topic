# RabbitMQ topic queue

## Setup
* setup a rabbitmq service
    + change .env accordingly to set the rabbitmq's username and password
    + invoke "docker-compose up"
* setup a consumer
    + e.g. ./rabbit-topic receive --username rabbitmq_user --password rabbitmq_password --exchange worldleader wildcow
* setup a producer
    + e.g. ./rabbit-topic send --username rabbitmq_user --password rabbitmq_password --exchange worldleader --route wildcow "Hello WildCow"

## References:
* [Topics](https://www.rabbitmq.com/tutorials/tutorial-five-go.html)
* [RabbitMQ container with Docker Compose](https://zgadzaj.com/development/docker/docker-compose/containers/rabbitmq)
