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

  All **response** and **error feedback** using the right JSON format:
r_code_blocks:
  - code: |-
      {
        "info": "..."
      }
    title: Response
    language: json
  - code: |-
      {
        "error": "error message here"
      }
    title: Error
    language: json
---