# MQTT client sample in Go

## Limitations
Many parameters are hard-coded:
- Port 1883 only can be used.
- QoS level is 1 only.
- ...

## Installation
```
$ go get github.com/eclipse/paho.mqtt.golang
```

## Example
A MQTT Server's IP address is 192.168.56.64.

Subscriber:
```
$ cd sub
$ go run sub.go -s 192.168.56.64 -i sub_a -t 'topic'
```

Subscriber (shared subscription):

```
$ cd sub
$ go run sub.go -s 192.168.56.64 -i sub_a -t '$share/group1/topic'
```

Publisher:

```
$ cd pub
$ go run pub.go -s 192.168.56.64 -i pub_a -t 'topic'
```
