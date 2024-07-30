#!/usr/bin/env bash

rm -r public/assets
rm public/index.html
rm public/favicon.ico
cd client-v2
bun install
bun run build
cd ..
cp -r client-v2/dist/* public