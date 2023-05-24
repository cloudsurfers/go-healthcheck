# go-healthcheck
A simple replacement for curl calls in docker-compose.yml. The primary use case are mircroservices based on scratch image. 

## Usage
'''healthcechk -h http:///xxxx/your-healthcheck-url'''. The executable exits either wie exit code 0 when Response Http Code is 200 or 1 otherweise.

##Downloads:
### Linux binaries
[Version 1.0.0](https://github.com/cloudsurfers/go-healthcheck/raw/main/builds/hc1.0.0)
