# marathon-schema

Validator for constraining `marathon.json` files to the Marathon App Definition
JSON format. The app definition schema is already provided by the Marathon open
source project itself so it only needs to be pulled in by this binary and
validated against a file.

[AppDefinition.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/AppDefinition.json)
[Group.json](https://github.com/mesosphere/marathon/blob/master/docs/docs/rest-api/public/api/v2/schema/Group.json)
