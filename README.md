# mimage

[![Build Status](https://travis-ci.org/dnovikoff/mimage.svg?branch=master)](https://travis-ci.org/dnovikoff/mimage)
[![Coverage Status](https://img.shields.io/codecov/c/github/dnovikoff/mimage.svg)](https://codecov.io/gh/dnovikoff/mimage)
[![Go Report Card](https://goreportcard.com/badge/github.com/dnovikoff/mimage)](https://goreportcard.com/report/github.com/dnovikoff/mimage)

This is a library for creating images with mahjong tiles.
There is a server to show fuctionalty.

Run the server from docker

`docker run -p 8080:8080 -it --rm tempai/mimage`

Or start with go

`go run ./cmd/mimage`

## Example urls
- Simple string http://localhost:8080/123s123p123z.png
- Rotated tiles http://localhost:8080/123445679m_-123s.png
- Kan http://localhost:8080/I50Im.png
- Upgraded http://localhost:8080/5-5-05s.png