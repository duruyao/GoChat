#!/bin/bash

echo "make sure you have installed commitizen by using \`npm install -g commitizen\`"

commitizen init cz-conventional-changelog --save --save-exact

echo "using \`git cz\` instead of \`git commit\`"