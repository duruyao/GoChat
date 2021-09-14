---
title: Errors
position_number: 3
content_markdown: |-
  | CODE | STATUS | DESCRIPTION |
  | :-: | :-: | :-: |
  | 200 | OK | Success |
  | 201 | Created | Created Successful |
  | 400 | Bad Request | We could not process that action |
  | 403 | Forbidden | We couldn't authenticate you |
  | 500 | Internal Server Error | Error occurred in the backend service |

  All **response** and **error feedback** using the right JSON format **------------>**

r_code_blocks:
  - code: |-
      {
          "id": 1,
          "uuid": "bb16ab21-6968-4182-4902-da0e04416619",
          "name": "root",
          "password": "fa585d89c851dd338a70dcf535aa",
          "max_role": 4,
          "created_at": "2021-09-12T17:36:30.4119796+08:00"
      }
    title: Response
    language: json
  - code: |-
      {
          "error": "sql: no rows in result set"
      }
    title: Error
    language: json
---