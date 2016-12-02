# marathon-schema

DC/OS integration utility providing schema based validation of marathon.json
files.

Validator constrains `marathon.json` files to the correct Marathon App
Definition JSON format. The App Definition schema is already provided by the
Marathon open source project itself so this tool only needs to pull down that
schema file, include some schema checking logic in a single binary, and validate
a test `marathon.json` file.

**WARNING** I'm still learning Go and this is far from feature complete. Dose
  heavily with salt...

## Background

Simple answer, I needed an extremely simple and easily deployable tool for
continuously integrating `marathon.json` files edited by developers.

The origin of this project was to help establish a standards based build
pipeline that pushed hundreds of developers into the same
practices. `marathon.json` helped us decompose services into single repos and
drive automation influenced directly by engineering teams and not some random
operations team two continents away.

## Usage (Future)

Validate `marathon.json`, as an app def (`default`), against remote master
schema (`default`).

```sh
$ marathon-schema marathon.json
```

Validate multiple files (`marathon1.json`, `marathon2.json`), as app defs
(`default`), against remote master schema (`default`).

```sh
$ marathon-schema marathon1.json marathon2.json
```

Validate `marathon.json`, as an app def (`-a`), against binary cached version of
`1.3.3` (`-v`) app def schema.

```sh
$ marathon-schema -m 1.3.3 -a marathon.json
```

Validate `group.json`, as a group (`-g`), against binary cached version of
`1.3.3` (`-v`) group schema.

```sh
$ marathon-schema -m 1.3.3 -g group.json
```

Validate `file.json`, as the type of schema we've loaded (`-f`).

```sh
$ marathon-schema -f schema.json file.json
```

## Schemas

Marathon provides the following JSON schema files.

- [AppDefinition.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/AppDefinition.json)
- [Group.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/Group.json)

## Links

- [Marathon Application Definitions](https://docs.mesosphere.com/1.8/usage/marathon/application-basics/)
