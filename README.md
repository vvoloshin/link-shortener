# golang-link-shortener (REST)
Link-shortener with saving data in the SQLite database.

###### Request to create short link example:

curl --request POST 'https://shortlink.com/processing' \  
--header 'x-api-key: 777' \  
--header 'Content-Type: text/plain' \  
--data-raw 'https://www.google.ru/search?q=golang'

###### Response example:

`https://short.com/tfptjz7A`

###### Get the original link from short link (with redirecting to it):

GET (from browser) https://shortlink.com/tfptjz7A
