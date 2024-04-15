#!/bin/bash
cd webserver/src
go build -o ../app/main  main.go data.go account.go content.go account_data.go add_book.go add_readinglist.go update_book.go update_readinglist.go authenicate_book.go