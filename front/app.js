const express = require('express');
const caller = require('grpc-caller');
const PROTO_PATH = __dirname + '/../article/article.proto';
const SERVER_HOST = process.env.SERVER_HOST || 'localhost';
const client = caller(`${SERVER_HOST}:50051`, PROTO_PATH, 'ArticleService');
const app = express();
const PORT = 3000;
const bodyParser = require('body-parser');

app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.json());

app.set('views', __dirname + '/views');
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
    client.first({}, (_err, response) => {
        res.render('index', response);
    });
});

app.get('/new', (req, res) => {
    res.render('new');
});

app.post('/new', (req, res) => {
    let data = {
        title: req.body.title,
        content: req.body.content,
        status: parseInt(req.body.status, 10)
    };
    console.log(data);
    client.post(data, (err, _response) => {
        if (err == null) {
            res.redirect('/');
        } else {
            res.redirect('/new');
        }
    });
});

app.listen(PORT, () => console.log(`Listening on port ${PORT}`));
