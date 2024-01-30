# <p align=center>Adriano Sela Aviles</p>

###### <p align=center> [ adriano.selaviles@gmail.com ] [ 778 681 3106 ] </p>

##### Motivated Senior Software and Security Engineer with a passion for learning. Experienced with ownership of globally distributed full stack services including the network, DNS, monitoring and logging, the platform, API & client code bases, documentation, and the infrastructure pieces that live in cloud providers.

## Skills

**Scripting:** Bash (+ awk/sed/grep), Makefile, Python

**Programming:** Go, Java, C/C++, Python, Lua

**Systems & Infrastructure:** Docker/docker-compose, Kubernetes, Terraform, Vagrant, AWS EC2 and Google Compute Engine

**Networking:** TCP/IP, DNS, TLS, BGP, in-depth understanding of the OSI structure, Linux implementation of TCP/IP suite, Berkeley Packet Filter (BPF), and experience writing reliable data transport protocols (i.e. a self-written TCP)

**Cloud Computing & Services:** Experience with all of AWS, GCP, and Azure, microservice architecture design, serverless (i.e. AWS Lambda, GCP Cloud Functions)

**Version Control & CI/CD:** Git, Jenkins, Concourse, AWS CodeBuild + CodeDeploy, GitHub Actions

**Identity & Access Management:** IAM flows (i.e. OAuth2.0, SAML), experience writing custom identity providers

**Security:** Deep understanding of symmetric and asymmetric key encryption, block ciphers, hashing functions, key sharing (e.g. elliptic curve Diffie-Hellman), experience writing public key infrastructures (PKIs). Attended internal Cisco R00tcamp - an intensive 7-day bootcamp on Kali Linux tools (e.g. Metasploit framework)

## Experience

#### **Senior Software Engineer / Security, Border0 Inc.** <sub><sup>(December 2022 - present, Vancouver BC)</sup></sub>

- Drove several security features from conception to completion including the introduction of robust credential stuffing mitigations and more secure authentication and authorization flows
- Directly responsible for triaging and addressing 1st and 3rd party software vulnerabilites
- Lead internal penetration testing and coordinated external penetration testing
- Wrote software to integrate with identity providers through protocols such as OAuth, OIDC, SAML, specific provider APIs (e.g. Okta, Google Workspace), and directory synchronization protocols e.g. SCIM, LDAP
- Full stack development including Go and Python back-end applications, JavaScript front-end applications, reverse proxies programmed in Lua, and the infrastructure pieces that live in various cloud providers (managed via infrastructure-as-code)

> Technologies: Go, Bash, Lua, AWS, GCP, Azure, Terraform, CloudFormation, Linux

#### **Senior Application Security Engineer, Thumbtack** <sub><sup>(May 2022 - November 2022, Vancouver BC)</sup></sub>

- Lead security incident response efforts and coordinated external penetration testing
- Primary recipient and triager of company-wide "responsible disclosure" mailbox
- Lead the introduction of tools for static and dynamic analysis of software through a company-wide week-long event to instrument the relevant codebases with the appropriate tooling
- Performed on-demand security design and architecture reviews

> Technologies: Go, Bash, Lua, AWS, Linux

#### **Senior Application Security Engineer, Cisco Systems Inc.** <sub><sup>(August 2021 - May 2022, Vancouver BC)</sup></sub>

- Driving the design, architecture, and development of large-scale, globally distributed security software applications and services
- Researching, designing, building, and operating "highly-available" security software services at scale, using Golang, Bash, and other languages/tools as fits the situation's needs
- Participating in and driving threat-modelling and risk-management of Cisco's infrastructure and products
- Supporting internal Cisco customers by extending existing services and creating new ones
- Developing solutions for computer networking and orchestration problems with availability, scale, security, and performance in mind

> Technologies: Go, Bash, Python, AWS, Jenkins, Linux

#### **Software Engineer, Amazon Web Services** <sub><sup>(July 2020 - July 2021, Vancouver BC)</sup></sub>

- Built and maintained several threat detection microservices at AWS scale
- Continuously designed and implemented features for AWS GuardDuty
- Hardened systems against increasingly sophisticated security threats
- Took part in launching the GuardDuty service to multiple AWS regions
- Provided 24/7 on-call support

> Technologies: Java, JavaScript, AWS, Linux, Networking primitives

#### **Security Operations Intern, Mozilla** <sub><sup>(May 2019 - Sept. 2019, San Francisco CA)</sup></sub>

- Apache Beam pipeline and cloud architecture design for processing AWS GuardDuty and GCP Event Threat Detection findings as part of the [FoxSec pipeline](https://github.com/mozilla-services/foxsec-pipeline), a collection of data processing Beam pipelines which alert on relevant threats to Firefox Infrastructure.

> Technologies: Java 9, Apache Beam SDK, Terraform, AWS {GuardDuty, CloudWatch, Kinesis}, GCP {Event Threat Detection, StackDriver, Pub/Sub}, Google Cloud Dataflow.

- Client library and CLI for [IPrepd](https://github.com/mozilla-services/iprepd), an Open-Source IP reputation daemon. Small additions to [SOPS](https://github.com/mozilla/sops), a secrets management solution.

> Technologies: Go

#### **Cloud Infra. Engineering Intern, Cisco Systems Inc.** <sub><sup>(May 2017 - Sept. 2018, Vancouver BC)</sup></sub>

- Collaborately built Vinz, a PKI (and client libraries) for provisioning Cisco servers with SSL certificates via programmatic API calls. The API is capable of delivering certificates and metadata within seconds of a request.

> Technologies: Go, DNS in depth, internal platforms and tooling

- Implemented an OpenID Connect identity provider in Golang and integrated it with AWS IAM as well as several other internal services which consumed our JSON web tokens (i.e. a compute platform, a metrics & logging service, ...)

> Technologies: Go, OpenID, AWS {IAM, DynamoDB, Lambda}, Terraform, Docker + Kubernetes, internal platforms and tooling

- Designed and implemented a high availability (HA) solution for multiplexing DNS among multiple redundant servers in accordance to frequent service healthchecks.

> Technologies: Go, AWS {Route53, CloudWatch, Lambda}

- Migrated builds and deployments from a legacy Jenkins server to Concourse

> Technologies: Jenkins, Councourse, AWS, internal platforms and tooling

## Projects

- [**sslmgr**](https://github.com/adrianosela/sslmgr) - layer of abstraction the around acme/autocert certificate manager [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#security) [![Documentation](https://godoc.org/github.com/adrianosela/sslmgr?status.svg)](https://godoc.org/github.com/adrianosela/sslmgr)
- [**rdtp**](https://github.com/adrianosela/rdtp) - cross OS transport layer service implemented as a daemon process [![Documentation](https://godoc.org/github.com/adrianosela/rdtp?status.svg)](https://godoc.org/github.com/adrianosela/rdtp)
- [**padl**](https://github.com/adrianosela/padl) - secrets management as-a-service API [![Documentation](https://godoc.org/github.com/adrianosela/padl?status.svg)](https://godoc.org/github.com/adrianosela/padl)
- [**multikey**](https://github.com/adrianosela/multikey) - library for implementing complex encryption schemes leveraging Shamir's Secret Sharing algorithm [![Documentation](https://godoc.org/github.com/adrianosela/multikey?status.svg)](https://godoc.org/github.com/adrianosela/multikey)
- [**spoof**](https://github.com/adrianosela/spoof) - command line utility for injecting spoofed frames into a network [![Documentation](https://godoc.org/github.com/adrianosela/spoof?status.svg)](https://godoc.org/github.com/adrianosela/spoof)
- [**certcache**](https://github.com/adrianosela/certcache) - a collection of useful implementations of the acme/certcache interface for common datastores (i.e. AWS DynamoDB, MongoDB, S3, and Google Firestore) [![Documentation](https://godoc.org/github.com/adrianosela/sslmgr?status.svg)](https://godoc.org/github.com/adrianosela/sslmgr)
- [**auth**](https://github.com/adrianosela/auth) - experimenting with RSA Keys, JWT Tokens, and OpenID Connect [![Documentation](https://godoc.org/github.com/adrianosela/auth?status.svg)](https://godoc.org/github.com/adrianosela/auth)
- [**keystore**](https://github.com/adrianosela/keystore) - filesystem Keystore with a REST API [![Documentation](https://godoc.org/github.com/adrianosela/keystore?status.svg)](https://godoc.org/github.com/adrianosela/keystore)
- [**cliprepd**](https://github.com/adrianosela/cliprepd) - command line client for Mozilla's IPrepd API [![Documentation](https://godoc.org/github.com/adrianosela/cliprepd?status.svg)](https://godoc.org/github.com/adrianosela/cliprepd)
- [**iprepd-firewall**](https://github.com/adrianosela/iprepd-firewall) - seamless IP reputation based application-layer firewall for services written in Go using a (Mozilla) IPrepd server [![Documentation](https://godoc.org/github.com/adrianosela/iprepd-firewall?status.svg)](https://godoc.org/github.com/adrianosela/iprepd-firewall/fwmw)
- [**ecdh**](https://github.com/adrianosela/ecdh) - elliptic curve diffie-hellman key exchange library in Go [![Documentation](https://godoc.org/github.com/adrianosela/ecdh?status.svg)](https://godoc.org/github.com/adrianosela/ecdh)
- [**botnet**](https://github.com/adrianosela/botnet) - command line managed bot master with a websockets based Command and Control (C&C) server [![Documentation](https://godoc.org/github.com/adrianosela/botnet?status.svg)](https://godoc.org/github.com/adrianosela/botnet)
- [**GoFaceTrainer**](https://github.com/adrianosela/GoFaceTrainer) - train a facial recognition model in Go using OpenCV and Machine Box [![Documentation](https://godoc.org/github.com/adrianosela/GoFaceTrainer?status.svg)](https://godoc.org/github.com/adrianosela/GoFaceTrainer)
- [**GoAway**](https://github.com/adrianosela/GoAway) - An OpenCV based motion detector and utilities in Go [![Documentation](https://godoc.org/github.com/adrianosela/GoAway?status.svg)](https://godoc.org/github.com/adrianosela/GoAway)
- [**streamer**](https://github.com/adrianosela/streamer) - python3 video streaming client and server implementing RSTP and RTP
- [**traceroute**](https://github.com/adrianosela/traceroute) - homemade implementation of a UDP traceroute program in python 
- [**ping**](https://github.com/adrianosela/traceroute) - homemade implementation of a an ICMP pinger program in python
- [**neuron**](https://github.com/adrianosela/neuron) - example predictive neural network in python
- [**mapp**](https://github.com/adrianosela/mapp) - events lookup mobile app written in Dart and Node JS back-end
- [**ovf**](https://github.com/adrianosela/ovf) - example stack overflow exploit in C

## Education

**BASc in Computer Engineering, University of British Columbia University** (2015-2020, Vancouver BC)

- Dean's Honour List all 4 academic years
- Coop Status: completed 4/4 (mandatory) work terms + 1 (additional) work term
