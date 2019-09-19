
if you want to update the security key,
you can use redis client to set the key(pub_key) value

set pub_key ```"{\"id\": \"9g1\", \"key\": \"hg92g02fewh7qk11\", \"iv\": \"546d4t546tw32e1f\"}"```

- id。code salt. once you change it, you should change client's id size and application server's is size
- key。aes encrypt key
- iv。aes encrypt iv


