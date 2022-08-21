# MQP - Message Queue Populator

Populate queues with messages by a specified in
config.json way

## Config example

```json
{
  "services": [{
    "type": "rabbitmq",
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
          "includeRandom": true
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