# Skimdex

A minimal, fast text based search engine built in Go and React for searching through your local Markdown files.

## What is this?
Skimdex is a tool that crawls a directory on your computer, reads your Markdown files, and builds a "search index." This allows you to search for keywords across hundreds of files instantly, rather than waiting for a slow system search.

## How it works
The project is split into two main parts:

### 1. Core (The Backend - Go)
- **Crawler:** Walks through your folders to find `.md` files.
- **Parser:** Reads the files and extracts the title and keywords.
- **Indexer:** Creates an "Inverted Index" (a map of words to file IDs) so lookups are lightning fast.
- **Storage:** Saves the index to your disk using Go's `gob` format so it persists between runs.

### 2. Web (The Frontend - React + Vite)
- A clean UI to type your search queries and see the results.
- Communicates with the Go server to fetch matches.

## Current Project Structure
```text
core/
├── cmd/
│   ├── index/    # Command to crawl and build the index
│   └── server/   # Command to start the search API
├── internal/
│   ├── crawler/  # Logic for finding files
│   ├── parser/   # Logic for reading Markdown
│   ├── indexer/  # The "brain" that maps words to files
│   └── storage/  # Logic for saving/loading data to disk
web/              # React frontend (TypeScript + Vite)
```

## Setup (WIP)
Currently, you can build the indexer by running:
```bash
cd core
go run cmd/index/main.go
```

## Why Go?
I used Go because it's great at handling file systems and is very fast for data processing tasks like indexing. It also makes persistence easy with the packages form the standard library like `gob`, `regexp` and many others. It's super simple to read yet very powerful and effecient.
