'use strict';

var yaml = require('js-yaml');
var fs = require('fs');
var express = require('express');

var app = express();

function readExercises() {
    return fs.readFileSync('./exercises.yaml', 'utf8');
}

function parseExercises() {
    return yaml.safeLoad(readExercises());
}

app.post('/workout', function(req, res) {
    res.set('Content-Type', 'application/json');
    res.status(201);
    res.send(parseExercises());
});

if (module === require.main) {
    var server = app.listen(process.env.PORT || 8080, function () {
        var host = server.address().address;
        var port = server.address().port;

        console.log('App listening at http://%s:%s', host, port);
    });
}

module.exports = app;
