const express = require('express');
const caller = require('grpc-caller');
const PROTO_PATH = __dirname + '/../article/article.proto';
const client = caller('localhost:50051', PROTO_PATH, 'ArticleService');
const app = express();
const PORT = 3000;

app.get('/', (req, res) => {
    client.first({}, (err, response) => {
        console.log(response);
        res.send(response.content)
    });
});

app.listen(PORT, () => console.log(`Listening on port ${PORT}`));
