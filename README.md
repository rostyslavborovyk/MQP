# MQP - Message Queue Populator

Populate queues with messages by a specified in
config.json way

## Usage
1. Download MQP binary from the last release
2. Create mqp-config.json
3. Run ./MQP

## Config example

```json
{
  "services": [{
    "type": "rabbitmq",
    "url": "amqp://guest:guest@localhost:5672/",
    "queues": [
      {
        "name": "non-random-queue",
        "message": {
          "frequency": 0.5,
          "bodyVariations": {
            "type": "text/plain",
            "variations": [
              "Non random text 1",
              "Non random text 2"
            ]
          },
          "includeTimestamp": false,
          "includeRandom": false
        }
      },
      {
        "name": "random-queue",
        "message": {
          "frequency": 0.8,
          "bodyVariations": {
            "type": "applications/json",
            "variations": [
              {
                "event": "OrderCreated",
                "orderId": 1
              },
              {
                "event": "OrderCreated",
                "orderId": 13
              },
              {
                "event": "OrderCreated",
                "orderId": 14
              }
            ]
          },
          "includeTimestamp": false,
          "includeRandom": true,
          "randomConfig": {
            "erlangOrder": 20
          }
        }
      }
    ]
  }]
}
```
`services` - specifies a list of services in which MQP will 
create queues

`services[i].type` - type of service, possible values: `rabbitmq`, redis and kafka tbd

`services[i].url` - connection string for a service

`services[i].queues` - list of queues that will be created for a 
specific service

`services[i].queues[i].name` - name of a queue that will be populated

`services[i].queues[i].message` - message config for a specific queue

`services[i].queues[i].message.frequency` - messages per second

`services[i].queues[i].message.bodyVariations` - variations of a message

`services[i].queues[i].message.bodyVariations.type` - type of message, possible values `text/plain`, `application/json`

`services[i].queues[i].message.bodyVariations.variations` - array of values that are randomly chosen for each sent message

`services[i].queues[i].message.includeTimestamp` - adds timestamp to a message body if body variation type equals to `text/plain`

`services[i].queues[i].message.includeRandom` - adds random Poisson distribution features

`services[i].queues[i].message.randomConfig.erlangOrder` - [erlang order](https://en.wikipedia.org/wiki/Erlang_distribution)

## Development status

Providers
- [X] RabbitMQ
- [ ] Kafka
- [ ] Redis

