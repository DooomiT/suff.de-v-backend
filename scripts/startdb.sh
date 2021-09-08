#!/bin/bash
docker build internal/db/ -t suffde-db
docker run  --name dev-postgres -p 5432:5432 suffde-db