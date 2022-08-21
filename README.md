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
          "body": "Non random message",
          "includeTimestamp": true,
          "includeRandom": false
        }
      },
      {
        "name": "random-queue",
        "message": {
          "frequency": 0.8,
          "body": "Random message",
          "includeTimestamp": true,
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

`services[i].queues` - list of queues that will be created for a 
specific service

`services[i].queues[i].name` - name of a queue that will be populated

`services[i].queues[i].message` - message config for a specific queue

`services[i].queues[i].message.frequency` - messages per second

`services[i].queues[i].message.body` - content of a message

`services[i].queues[i].message.includeTimestamp` - adds timestamp to a message body if true

`services[i].queues[i].message.includeRandom` - adds random Poisson distribution features

`services[i].queues[i].message.randomConfig.erlangOrder` - [erlang order](https://en.wikipedia.org/wiki/Erlang_distribution)