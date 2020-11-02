# mimage

[![CI](https://github.com/dnovikoff/mimage/workflows/CI/badge.svg?branch=master&event=push)](https://github.com/dnovikoff/mimage/actions?query=workflow%3ACI)
[![Coverage Status](https://img.shields.io/codecov/c/github/dnovikoff/mimage.svg)](https://codecov.io/gh/dnovikoff/mimage)
[![Go Report Card](https://goreportcard.com/badge/github.com/dnovikoff/mimage)](https://goreportcard.com/report/github.com/dnovikoff/mimage)

This is a library for creating images with mahjong tiles.
There is a server to show fuctionalty.

Run the server from docker

`docker run -p 8080:8080 -it --rm tempai/mimage`

Or start with go

`go run ./cmd/mimage`

## Example urls
- Simple string http://localhost:8080/1379m2568p23456s_6z.png
- Rotated tiles http://localhost:8080/123445679m_-123s_4-444s.png
- Kan http://localhost:8080/I50Im_5-5-05s.png

## Examples of resulting images
- ![Simple](https://raw.githubusercontent.com/dnovikoff/mimage/master/pkg/image/test_data/1379m2568p23456s_6z.png)
- ![Rotated](https://raw.githubusercontent.com/dnovikoff/mimage/master/pkg/image/test_data/123445679m_-123s_4-444s.png)
- ![Kan](https://raw.githubusercontent.com/dnovikoff/mimage/master/pkg/image/test_data/I50Im_5-5-05s.png)