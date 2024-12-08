# Barticle

A text summarization and question answering tool

## Requirements

go
docker
python

## Building

Docker can be used to build the api
'''bash
docker build -t barticle ./api
'''
To build the client using go do the following
'''bash
go build -o bin/Barticle
'''
To run the docker image in a container
'''bash
docker run -p 8000:8000 barticle
'''

## Running the application

using the cli
'''bash
./bin/Barticle summary textfile1.txt textfile2.txt
./bin/Barticle summary --url <https://en.wikipedia.org/wiki/The_Jungle_Book>
'''
to run the TUI
'''bash
./bin/Barticle
'''
