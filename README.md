# marathon-schema

DC/OS integration utility providing schema based validation of marathon.json
files.

Validator constrains `marathon.json` files to the correct Marathon App
Definition JSON format. The App Definition schema is already provided by the
Marathon open source project itself so this tool only needs to pull down that
schema file, include some schema checking logic in a single binary, and validate
a test `marathon.json` file.

## Background

Simple answer, I needed an extremely simple and easily deployable tool for
continuously integrating `marathon.json` files edited by developers.

The origin of this project was to help establish a standards based build
pipeline that pushed hundreds of developers into the same
practices. `marathon.json` helped us decompose services into single repos and
drive automation influenced directly by engineering teams and not some random
operations team two continents away.

## Schemas

Marathon provides the following JSON schema files.

- [AppDefinition.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/AppDefinition.json)
- [Group.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/Group.json)
