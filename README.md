## English | [简体中文](assets/readme-cn.md)

<p align="center">
<img width="500px" src="https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/logo.png">
</p>

<div align=center>

[![Go Report](https://goreportcard.com/badge/github.com/zhufuyi/sponge)](https://goreportcard.com/report/github.com/zhufuyi/sponge)
[![codecov](https://codecov.io/gh/zhufuyi/sponge/branch/main/graph/badge.svg)](https://codecov.io/gh/zhufuyi/sponge)
[![Go Reference](https://pkg.go.dev/badge/github.com/zhufuyi/sponge.svg)](https://pkg.go.dev/github.com/zhufuyi/sponge)
[![Go](https://github.com/zhufuyi/sponge/workflows/Go/badge.svg?branch=main)](https://github.com/zhufuyi/sponge/actions)
[![Awesome Go](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go)
[![License: MIT](https://img.shields.io/github/license/zhufuyi/sponge)](https://img.shields.io/github/license/zhufuyi/sponge)

</div>

**Sponge** is a basic development framework that integrates `code auto generation`, `Gin and GRPC`, sponge has a rich set of code generation commands, generating different functional code can be combined into a complete service (similar to the way that artificially broken sponge cells can automatically recombine into a new sponge). The code is decoupled and modularly designed, it is easy to build a complete project from development to deployment, just fill in the business logic code on the generated template code, greatly improved development efficiency and reduced development difficulty, the use of Go can also be "low-code development".

<br>

If you are develop RESTful web service or microservice with a simple CRUD API interface, you don't need to write any Go code can be compiled and deployed to the linux server, docker, k8s, just need to connect to the database (mysql、mongodb、postgresql、tidb、sqlite) to generate a complete service code.

If you develop generic web or microservices, just focus on the three core parts of `define database tables`, `define api interfaces in proto files`, `fill in business logic code in the generated template files`, and the rest of the go code is automatically generated by sponge.

<br>

### Sponge Generates the Code Framework

sponge is mainly based on `SQL` and `Protobuf` two ways to generate code, each way has to generate code for different functions. `SQL` supports databases **mysql**, **mongodb**, **postgresql**, **tidb**, **sqlite**.

#### Generate Code Framework

<p align="center">
<img width="1500px" src="https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/sponge-framework.png">
</p>

<br>

#### Generate Code Framework Corresponding UI Interface

<p align="center">
<img width="1200px" src="https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/en_sponge-ui.png">
</p>

<br>

#### Generate Service Code for Egg Model

The sponge code generation process strips away the business logic and non-business logic of the two major parts of the code. Sponge's code generation function as a hen, the generated service code is the egg, take the generated web service backend code as an example, egg model profiling diagram:

<p align="center">
<img width="1200px" src="https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/en_web-http-pb-anatomy.png">
</p>

In addition to the egg model of web service backend code, there are egg models of grpc service code and grpc gateway service code, [click here to view](https://go-sponge.com/learn-about-sponge?id=%f0%9f%8f%b7egg-model-for-generate-service-code).

<br>

### Services framework

Sponge is essentially a microservice framework that includes code generation capabilities. The microservice framework is shown in the following figure, which is a typical microservice hierarchical structure, with high performance, high scalability, contains commonly used service governance features, you can easily replace or add their own service governance features.

<p align="center">
<img width="1000px" src="https://raw.githubusercontent.com/zhufuyi/sponge/main/assets/microservices-framework.png">
</p>

<br>

### Key Features

- Web framework [gin](https://github.com/gin-gonic/gin)
- RPC framework [grpc](https://github.com/grpc/grpc-go)
- Configuration parsing [viper](https://github.com/spf13/viper)
- Configuration center [nacos](https://github.com/alibaba/nacos)
- Logging component [zap](https://github.com/uber-go/zap)
- Database ORM component [gorm](https://github.com/go-gorm/gorm)
- Cache component [go-redis](https://github.com/go-redis/redis), [ristretto](https://github.com/dgraph-io/ristretto)
- Automated API documentation [swagger](https://github.com/swaggo/swag), [protoc-gen-openapiv2](https://github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2)
- Authentication [jwt](https://github.com/golang-jwt/jwt)
- Message Queue [rabbitmq](https://github.com/rabbitmq/amqp091-go)
- Distributed Transaction Manager [dtm](https://github.com/dtm-labs/dtm)
- Parameter validation [validator](https://github.com/go-playground/validator)
- Adaptive rate limiting [ratelimit](https://github.com/zhufuyi/sponge/tree/main/pkg/shield/ratelimit)
- Adaptive circuit breaking [circuitbreaker](https://github.com/zhufuyi/sponge/tree/main/pkg/shield/circuitbreaker)
- Distributed Tracing [opentelemetry](https://github.com/open-telemetry/opentelemetry-go)
- Metrics monitoring [prometheus](https://github.com/prometheus/client_golang/prometheus), [grafana](https://github.com/grafana/grafana)
- Service registration and discovery [etcd](https://github.com/etcd-io/etcd), [consul](https://github.com/hashicorp/consul), [nacos](https://github.com/alibaba/nacos)
- Adaptive collecting [profile](https://go.dev/blog/pprof)
- Resource statistics [gopsutil](https://github.com/shirou/gopsutil)
- Code quality checking [golangci-lint](https://github.com/golangci/golangci-lint)
- Continuous integration and deployment [jenkins](https://github.com/jenkinsci/jenkins), [docker](https://www.docker.com/), [kubernetes](https://github.com/kubernetes/kubernetes)

<br>

### Project Code Directory Structure

The project code directory structure created by sponge follows the [project-layout](https://github.com/golang-standards/project-layout) convention and is structured as follows:

```bash
.
├── api            # Directory for exposing external API interfaces, typically containing proto files and generated *.pb.go files. The directory structure is typically in the form `api/xxx/v1`, where v1 indicates the version.
├── assets         # Store various static resources, such as images, markdown files, etc.
├── cmd            # Program entry directory
│    └── serviceName
│         ├── initial     # Program initialization, consisting of three files: initApp initializes configurations, registerServers registers services (HTTP or grpc), and registerClose registers resource cleanup.
│         └── main.go     # Program entry file
├── configs        # Directory for configuration files
├── deployments    # Directory for deployment scripts, supporting binary, Docker and Kubernetes deployments.
├─ docs            # Directory for API interface Swagger documentation.
├── internal       # Directory for code of private applications and libraries.
│    ├── cache        # Cache directory wrapped around business logic.
│    ├── config       # Directory for Go structure configuration files.
│    ├── dao          # Data access directory.
│    ├── ecode        # Directory for system error codes and custom business error codes.
│    ├── handler      # Directory for implementing HTTP business functionality (specific to web services).
│    ├── model        # Database model directory.
│    ├── routers      # HTTP routing directory.
│    ├── rpcclient    # Directory for client-side code that connects to grpc services.
│    ├── server       # Directory for creating services, including HTTP and grpc.
│    ├── service      # Directory for implementing grpc business functionality (specific to grpc services).
│    └── types        # Directory for defining request and response parameter structures for HTTP.
├── pkg            # Directory for shared libraries.
├── scripts        # Directory for scripts, including compilation, execution, code generation, and deployment scripts.
├── test           # Directory for scripts required for testing services  and test SQL.
└── third_party    # Directory for external helper programs, forked code, and other third-party tools.
```

<br>

### Quick start

#### Installation sponge

sponge can be installed on Windows, macOS, Linux and Docker environments. Click here for [instructions on installing sponge](https://github.com/zhufuyi/sponge/blob/main/assets/install-en.md).

#### Starting UI service

After installing the sponge, start the UI service:

```bash
sponge run
```

Access `http://localhost:24631` in a local browser and manipulate the generated code on the UI page.

> If you want to access it on a cross-host browser, you need to specify the host ip or domain name when starting the UI, example `sponge run -a http://your_host_ip:24631`. It is also possible to start the UI service on docker to support cross-host access, Click for instructions on [starting the sponge UI service in docker](https://github.com/zhufuyi/sponge/blob/main/assets/install-en.md#docker-environment).

<br>

### Sponge Development Documentation

Detailed instructions for operating, configuring, and deploying a project using sponge development, Click here to view the [sponge development documentation](https://go-sponge.com/)

<br>

### Examples of use

#### Simple examples

No specific business logic code is included.

- [1_web-gin-CRUD](https://github.com/zhufuyi/sponge_examples/tree/main/1_web-gin-CRUD)
- [2_micro-grpc-CRUD](https://github.com/zhufuyi/sponge_examples/tree/main/2_micro-grpc-CRUD)
- [3_web-gin-protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/3_web-gin-protobuf)
- [4_micro-grpc-protobuf](https://github.com/zhufuyi/sponge_examples/tree/main/4_micro-grpc-protobuf)
- [5_micro-gin-rpc-gateway](https://github.com/zhufuyi/sponge_examples/tree/main/5_micro-gin-rpc-gateway)
- [6_micro-cluster-demo](https://github.com/zhufuyi/sponge_examples/tree/main/6_micro-cluster)

#### Complete project examples

- [7_community-single](https://github.com/zhufuyi/sponge_examples/tree/main/7_community-single)
- [8_community-cluster](https://github.com/zhufuyi/sponge_examples/tree/main/8_community-cluster)

#### Distributed transaction examples

- [9_order-system](https://github.com/zhufuyi/sponge_examples/tree/main/9_order-grpc-distributed-transaction)

<br>

### License

See the [LICENSE](LICENSE) file for licensing information.

<br>

### How to contribute

You are more than welcome to join us, raise an Issue or Pull Request.

Pull Request instructions.

1. Fork the code
2. Create your own branch: `git checkout -b feat/xxxx`
3. Commit your changes: `git commit -am 'feat: add xxxxx'`
4. Push your branch: `git push origin feat/xxxx`
5. Commit your pull request

<br>

**If it's help to you, give it a star ⭐.**

<br>
