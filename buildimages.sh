#!/bin/bash

docker build -t policy/demoservice demo-service/
docker build -t policy/userservice user-service/
docker build -t policy/productservice product-service/