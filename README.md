# Golang Exercises

## Basic Exercises

|NO.|Sample|Description|
|----|----|----|
|001|[hello-world](./001-hello-world/hello_world.go)|first sample: hello,world!|
|002|[if-statement](./002-if-statement/if_statement.go)|if statement sample|
|003|[switch-statement](./003-switch-statement/switch_statement.go)|switch statement sample|
|004|[for-statement](./004-for-statement/for_statement.go)|for statement sample|
|005|[slice-type](./005-slice-type/slice_type.go)|slice data type sample|
|006|[map-type](./006-map-type/map_type.go)|map data type sample|
|007|[interface-type](./007-interface-type/interface_type.go)|interface type sample|
|008|[channel-type](./008-channel-type/channel_type.go)|channel type sample|
|009|[class-inheric](./009-class-inherit/class_inherit.go)|class inheritance sample|
|010|[generic-function](./010-generic-function/generic_function.go)|generic function sample|
|011|[dynamic-function](./011-dynamic-function/dynamic_function.go)|dynamic function sample|
|012|[type-size](./012-type-size/type_size.go)|type size evaluation sample|
|013|[type-assert](./013-type-assert/type_assert.go)|type assertion sample|
|014|[interface-trap](./014-interface-trap/interface_trap.go)|an interface trap sample|
|015|[time-ticker](./015-time-ticker/time_ticker.go)|time ticker sample|
|016|[channel-close](./016-channel-close/channel_close.go)|the behavior of closed channel sample|
|017|[panic-recover](./017-panic-recover/panic_recover.go)|panic recover sample|
|018|[net-context](./018-net-context/with_cancel.go)|net context with cancel / deadline / timeout samples|
|019|[map-assert](./019-map-assert/map_assert.go)|map assertion sample|
|020|[goroutine-pool](./020-goroutine-pool/goroutine_pool.go)|goroutine pool sample|
|021|[unnamed-struct](./021-unnamed-struct/unnamed_struct.go)|unnamed struct type sample|

## PKG Exercises
|PKG|Sample|Description|
|----|----|----|
|[amqp](https://github.com/streadway/amqp)|[amqp-simple-mode](./pkgs/amqp/01-simple-mode/)|amqp simple mode sample|
|[amqp](https://github.com/streadway/amqp)|[amqp-work-queue-mode](./pkgs/amqp/02-work-queues/)|amqp work queue mode sample|
|[amqp](https://github.com/streadway/amqp)|[amqp-fanout-mode](./pkgs/amqp/03-fanout/)|amqp fanout mode sample|
|[amqp](https://github.com/streadway/amqp)|[amqp-direct-route-mode](./pkgs/amqp/04-direct-route/)|amqp direct route mode sample|
|[amqp](https://github.com/streadway/amqp)|[amqp-topic-route-mode](./pkgs/amqp/05-topic-route/)|amqp topic route mode sample|
|[amqp](https://github.com/streadway/amqp)|[amqp-rpc-mode](./pkgs/amqp/06-rpc-demo/)|amqp rpc mode sample|
|encoding/base64|[base64](./pkgs/base64/)|base64 encoding/decoding sample|
|regexp|[regexp](./pkgs/regexp/)|regular expression sample|
|encoding/xml|[xml](./pkgs/xml/)|xml decoding sample|
|[sego](https://github.com/huichen/sego)|[sego](./pkgs/sego/)|sego word segment sample|
|fabcar|[fabcar](./pkgs/fabcar/)|Hyperledger fabric v2.2 fabcar chaincode invoking sample|
|[bccsp](https://github.com/hyperledger/fabric/bccsp)|[bccsp](./pkgs/bccsp/)|Hyperledger fabric bccsp package sample|
|[bcmsp](https://github.com/hyperledger/fabric-sdk-go/pkg/client/msp)|[msp](./pkgs/bcmsp/)|Hyperledger fabric-sdk-go msp package sample|
|[rwriter](https://github.com/ewangplay/rwriter)|[rwriter](./pkgs/rwriter/)|Rotate writer rwriter package sample|
|[cryptolib](https://github.com/ewangplay/cryptolib)|[cryptolib](./pkgs/cryptolib/)|common cryptographical library cryptolib package sample|
|[did wallet](https://github.com/ewangplay/did-wallet)|[did-wallet](./pkgs/did-wallet/)|did wallet package sample|
|[elastic](https://github.com/olivere/elastic/v7)|[elastic](./pkgs/elastic/)|elastic search sdk v7 package sample|
|[gfa-eps-sdk-go](https://github.com/gfacloud/gfa-eps-sdk-go)|[epssdk](./pkgs/epssdk/)|gfa eps service sdk package sample|
|[gfa-eps-sdk-go](https://github.com/gfacloud/gfa-eps-sdk-go)|[epssdk](./pkgs/epssdk/)|gfa eps service sdk package sample|
|[pkcs11](https://github.com/miekg/pkcs11)|[pkcs11](./pkgs/pkcs11/)|pkcs11 package testing with SoftHSM2 sample|
|[pkcs12](https://golang.org/x/crypto/pkcs12)|[pkcs12](./pkgs/pkcs12/)|pkcs12 package sample|
|tls|[tls](./pkgs/tls/)|TLS certificate and connection related sample|


## Learning Go

`Learning Go` is a primer ebook for studying golang, please refer to it.

|Chapter|Topic|Description|
|----|----|----|
|Q2|[for-loop](./learning-go/Q2/for_loop.go)|For-loop exercise|
|Q3|[FizzBuzz](./learning-go/Q3/FizzBuzz.go)|FizzBuzz exercise|
|Q4|[strings](./learning-go/Q4/strings.go)|Strings exercise|

# GOPL

`GOPL` refers to "The Go Programming Language" book, written by `Alan A.A.Donovan` and `Brian W.Kernighan`. 

|Chaper|Topic|Description|
|----|----|----|
|ch6|[intlist](./gopl/ch6/intlist/intlist.go)|A linked list of integers sample|
|ch6|[intset](./gopl/ch6/intset/intset.go)|A non-negativ integer set sample|
|ch7|[counter](./gopl/ch7/counter/counter.go)|Byte Counter, Word Counter, Line Counter implementation satisfied to Writer Interface|
