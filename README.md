# marathon-schema

[![Build Status](https://travis-ci.org/cheapRoc/marathon-schema.svg?branch=master)](https://travis-ci.org/cheapRoc/marathon-schema)

DC/OS integration utility providing schema based validation of marathon.json
files.

Validator constrains `marathon.json` files to the correct Marathon App
Definition JSON format. The App Definition schema is already provided by the
Marathon open source project itself so this tool only needs to pull down that
schema file, include some schema checking logic in a single binary, and validate
a test `marathon.json` file.

**WARNING** This is current feature incomplete. Dose heavily...

![salt guy](http://i.giphy.com/l4Jz3a8jO92crUlWM.gif)

## Background

Simple answer, I needed an extremely simple and easily deployable tool for
continuously integrating `marathon.json` files edited by developers.

The origin of this project was to help establish a standards based build
pipeline that pushed hundreds of developers into the same
practices. `marathon.json` helped us decompose services into single repos and
drive automation influenced directly by engineering teams and not some random
operations team two continents away.

## Usage (Future)

Validate `marathon.json`, as an app def (`default`), against binary stored
schema (`default`).

```sh
$ marathon-schema ./marathon.json
$ cat ./marathon.json | marathon-schema
```

Validate multiple files (`marathon1.json`, `marathon2.json`), as app defs
(`default`), against binary stored schema (`default`).

```sh
$ marathon-schema ./marathon1.json ./marathon2.json
```

Validate `marathon.json`, as an app def (`-a`), against remote Git tag `1.3.3`
(`-t`) of the app def schema.

```sh
$ marathon-schema -a -t 1.3.3 ./marathon.json
$ cat ./marathon.json | marathon-schema -a -t 1.3.3
```

Validate `group.json`, as a group (`-g`), against remote Git tag `1.3.3` (`-t`)
of the app def schema.

```sh
$ marathon-schema -t 1.3.3 -g ./group.json
$ cat ./group.json | marathon-schema -g -t 1.3.3
```

## Schemas

Marathon provides the following remote JSON schema files...

- [AppDefinition.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/AppDefinition.json)
- [Group.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/Group.json)

We replace with the `-t` flag like so...

- [https://raw.githubusercontent.com/mesosphere/marathon/v${TAG}/docs/docs/rest-api/public/api/v2/schema/AppDefinition.json](https://github.com/mesosphere/marathon/blob/v1.3.3/docs/docs/rest-api/public/api/v2/schema/AppDefinition.json)

## Links

- [Marathon Application Definitions](https://docs.mesosphere.com/1.8/usage/marathon/application-basics/)
