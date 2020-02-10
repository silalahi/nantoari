# nantoari

> Currently, nantoari is in highly development process. Some API may breaks. Do not use for production, yet.

:zap: Simple CSV to API ready, instantly!

[![asciicast](https://asciinema.org/a/226316.png)](https://asciinema.org/a/226316)

##### And here are some ideas:
- Configurable (either flag or file base).
- Simple middleware to store total hit or request in Redis


GET /:uuid -> Retrieve JSON
POST / -> (Post CSV file and convert to JSON) -> Return UUID
GET /:uuid/stats