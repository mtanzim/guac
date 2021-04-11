make && 
echo '{"start": "2021-01-01", "end":"2021-12-31" }' | sam local invoke --event - "WakaDynamoFetch" --env-vars env.json