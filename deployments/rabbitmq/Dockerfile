FROM rabbitmq:3.9.23-management

COPY ./rabbitmq_delayed_message_exchange-3.9.0.ez /opt/rabbitmq/plugins/

RUN rabbitmq-plugins enable rabbitmq_delayed_message_exchange
