---
title: /session
position_number: 2.1
type: post
desc: Create session

auths:
  - type: API Key
    key: GoChat-Token
    value: 433578ab-84c2-4e02-4656-de55a8097c9f
    desc:

l_code_blocks:
  - code: |-
      {
        "name": "root",
        "max_role": 4,
        "password": "12345678"
      }
    title: Body
    language: json
  - code: |-
      curl --location --request POST 'localhost:1213/v1/user' \
      --header 'GoChat-Token: 433578ab-84c2-4e02-4656-de55a8097c9f' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "name": "root",
          "max_role": 4,
          "password": "12345678"
      }'
    title: cURL
    language: bash
  - code: |-
      var settings = {
        "url": "localhost:1213/v1/user",
        "method": "POST",
        "timeout": 0,
        "headers": {
          "GoChat-Token": "433578ab-84c2-4e02-4656-de55a8097c9f",
          "Content-Type": "application/json"
        },
        "data": JSON.stringify({
          "name": "root",
          "max_role": 4,
          "password": "12345678"
        }),
      };

      $.ajax(settings).done(function (response) {
        console.log(response);
      });
    title: jQuery
    language: javascript

r_code_blocks:
  - code: |-
      {
        "id": 3,
        "uuid": "...",
        "name": "root",
        "password": "...",
        "max_role": 4,
        "created_at": "..."
      }
    title: Response
    language: json
  - code: |-
      {
        "error": "..."
      }
    title: Error
    language: json
---


