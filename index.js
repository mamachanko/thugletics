'use strict';

var yaml = require('js-yaml');
var fs = require('fs');
var express = require('express');
var _ = require('lodash');

var app = express();

function readExercises() {
    return fs.readFileSync('./exercises.yaml', 'utf8');
}

function parseExercises() {
    return yaml.safeLoad(readExercises());
}

function getWorkout() {
    var workout = {'1': '10 min cardio'};
    var exercises = parseExercises();
    var i = 2;
    var muscleGroups = _.shuffle(Object.keys(exercises));
    muscleGroups.forEach(function (muscleGroup) {
        workout[i.toString()] = _.sample(exercises[muscleGroup]);
        i++;
    })
    return workout;
}

function workoutHandler(req, res) {
    res.status(201);
    res.set('Content-Type', 'application/json');
    res.send({'data': getWorkout()});
};

app.get('/', workoutHandler);
app.post('/workout', workoutHandler);

if (module === require.main) {
    var server = app.listen(process.env.PORT || 8080, function () {
        var host = server.address().address;
        var port = server.address().port;

        console.log('App listening at http://%s:%s', host, port);
    });
}

module.exports = app;
