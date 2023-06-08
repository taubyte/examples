#!/bin/bash

yarn install
yarn audit fix
yarn build && mv dist/* /out