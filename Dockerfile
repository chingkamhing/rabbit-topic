FROM rabbitmq:3.7-management-alpine

RUN rabbitmq-plugins enable --offline rabbitmq_mqtt rabbitmq_web_mqtt
RUN rabbitmq-plugins enable --offline rabbitmq_auth_backend_http
RUN rabbitmq-plugins enable --offline rabbitmq_auth_backend_cache

# define environment variables
ENV RABBITMQ_DEFAULT_USER=${RABBITMQ_DEFAULT_USER}
ENV RABBITMQ_DEFAULT_PASS=${RABBITMQ_DEFAULT_PASS}
ENV RABBITMQ_PID_FILE /var/lib/rabbitmq/mnesia/rabbitmq

ADD init.sh /init.sh
RUN chmod +x /init.sh

# define default command
CMD ["/init.sh"]

# AMQP
EXPOSE 5672
# rabbitMQ management
EXPOSE 15672
# Web MQTT
EXPOSE 15675
# Prometheus ??
EXPOSE 15692
# MQTT TCP
EXPOSE 1883
# MQTT SSL
EXPOSE 8883
