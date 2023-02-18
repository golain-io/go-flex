# DEPRECATED 
Replaced by: https://github.com/golain-io/go-fluxbuilder

# Go-Flex
Standardised Query Builder for FluxQL (InfluxDB 2.x)  
Heavily inspired by github.com/Masterminds/squirrel  

#### Why Builder Pattern?
From Calhoun.io :
> The builder pattern is a technique where a developer uses a "builder" to construct an object. The end result could be anything - a string, a struct instance, or even a closure - but it is built using a builder. More often than not a builder is used because the object being created is complex and needs to be constructed in multiple steps, so the builder helps isolate each step and prevent bugs.

#### Example Code
```
import fluxq "github.com/golain-io/go-flex"

query := fluxq.From("sample-bucket").Range("-1h", "-30s")
```
