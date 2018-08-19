'use strict';

var utils = require('../utils/writer.js');
var TodoService = require('../service/TodoServiceService');

module.exports.createTodo = function createTodo (req, res, next) {
  var body = req.swagger.params['body'].value;
  TodoService.createTodo(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.deleteTodo = function deleteTodo (req, res, next) {
  var id = req.swagger.params['id'].value;
  TodoService.deleteTodo(id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getTodo = function getTodo (req, res, next) {
  var id = req.swagger.params['id'].value;
  TodoService.getTodo(id)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.queryTodo = function queryTodo (req, res, next) {
  TodoService.queryTodo()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.updateTodo = function updateTodo (req, res, next) {
  var id = req.swagger.params['id'].value;
  var body = req.swagger.params['body'].value;
  TodoService.updateTodo(id,body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
