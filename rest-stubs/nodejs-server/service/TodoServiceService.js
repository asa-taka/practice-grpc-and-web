'use strict';


/**
 *
 * body TodoCreateRequest 
 * returns TodoMutateResponse
 **/
exports.createTodo = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "record" : {
    "created_at" : "2000-01-23T04:56:07.000+00:00",
    "id" : 0,
    "detail" : "detail",
    "title" : "title",
    "deadline" : "2000-01-23T04:56:07.000+00:00",
    "status" : { }
  }
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * id Integer 
 * returns TodoMutateResponse
 **/
exports.deleteTodo = function(id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "record" : {
    "created_at" : "2000-01-23T04:56:07.000+00:00",
    "id" : 0,
    "detail" : "detail",
    "title" : "title",
    "deadline" : "2000-01-23T04:56:07.000+00:00",
    "status" : { }
  }
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * id Integer 
 * returns TodoGetResponse
 **/
exports.getTodo = function(id) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "record" : {
    "created_at" : "2000-01-23T04:56:07.000+00:00",
    "id" : 0,
    "detail" : "detail",
    "title" : "title",
    "deadline" : "2000-01-23T04:56:07.000+00:00",
    "status" : { }
  }
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * returns TodoQueryRequest
 **/
exports.queryTodo = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "start" : "2000-01-23T04:56:07.000+00:00",
  "end" : "2000-01-23T04:56:07.000+00:00"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * id Integer 
 * body TodoUpdateRequest 
 * returns TodoMutateResponse
 **/
exports.updateTodo = function(id,body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "record" : {
    "created_at" : "2000-01-23T04:56:07.000+00:00",
    "id" : 0,
    "detail" : "detail",
    "title" : "title",
    "deadline" : "2000-01-23T04:56:07.000+00:00",
    "status" : { }
  }
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

