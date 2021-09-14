---
title: Authentication
position_number: 2

content_markdown: |-
  You need to be authenticated for all API requests. You can generate an API key in your developer dashboard.

  Add the API key to all the **Request Header**.

  Nothing will work unless you include this API key
  {: .error}

l_code_blocks:
  - code: |-
      curl --location --request GET 'localhost:1213/v1/api' \
      --header 'GoChat-Token: 433578ab-84c2-4e02-4656-de55a8097c9f'
    title: cURL
    language: bash
  - code: |-
      var settings = {
        "url": "localhost:1213/v1/api",
        "method": "GET",
        "timeout": 0,
        "headers": {
          "GoChat-Token": "433578ab-84c2-4e02-4656-de55a8097c9f"
        },
      };
      
      $.ajax(settings).done(function (response) {
        console.log(response);
      });
    title: JQuery
    language: javascript
---