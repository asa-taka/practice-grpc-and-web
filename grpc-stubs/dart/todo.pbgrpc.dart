///
//  Generated code. Do not modify.
///
// ignore_for_file: non_constant_identifier_names,library_prefixes

import 'dart:async';

import 'package:grpc/grpc.dart';

import 'todo.pb.dart';
export 'todo.pb.dart';

class TodoServiceClient extends Client {
  static final _$queryTodo =
      new ClientMethod<Todo_QueryResponse, Todo_QueryRequest>(
          '/todo.TodoService/QueryTodo',
          (Todo_QueryResponse value) => value.writeToBuffer(),
          (List<int> value) => new Todo_QueryRequest.fromBuffer(value));
  static final _$getTodo = new ClientMethod<Todo_GetRequest, Todo_GetResponse>(
      '/todo.TodoService/GetTodo',
      (Todo_GetRequest value) => value.writeToBuffer(),
      (List<int> value) => new Todo_GetResponse.fromBuffer(value));
  static final _$createTodo =
      new ClientMethod<Todo_CreateRequest, Todo_MutateResponse>(
          '/todo.TodoService/CreateTodo',
          (Todo_CreateRequest value) => value.writeToBuffer(),
          (List<int> value) => new Todo_MutateResponse.fromBuffer(value));
  static final _$updateTodo =
      new ClientMethod<Todo_UpdateRequest, Todo_MutateResponse>(
          '/todo.TodoService/UpdateTodo',
          (Todo_UpdateRequest value) => value.writeToBuffer(),
          (List<int> value) => new Todo_MutateResponse.fromBuffer(value));
  static final _$deleteTodo =
      new ClientMethod<Todo_DeleteRequest, Todo_MutateResponse>(
          '/todo.TodoService/DeleteTodo',
          (Todo_DeleteRequest value) => value.writeToBuffer(),
          (List<int> value) => new Todo_MutateResponse.fromBuffer(value));

  TodoServiceClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<Todo_QueryRequest> queryTodo(Todo_QueryResponse request,
      {CallOptions options}) {
    final call = $createCall(_$queryTodo, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }

  ResponseFuture<Todo_GetResponse> getTodo(Todo_GetRequest request,
      {CallOptions options}) {
    final call = $createCall(_$getTodo, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }

  ResponseFuture<Todo_MutateResponse> createTodo(Todo_CreateRequest request,
      {CallOptions options}) {
    final call = $createCall(_$createTodo, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }

  ResponseFuture<Todo_MutateResponse> updateTodo(Todo_UpdateRequest request,
      {CallOptions options}) {
    final call = $createCall(_$updateTodo, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }

  ResponseFuture<Todo_MutateResponse> deleteTodo(Todo_DeleteRequest request,
      {CallOptions options}) {
    final call = $createCall(_$deleteTodo, new Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class TodoServiceBase extends Service {
  String get $name => 'todo.TodoService';

  TodoServiceBase() {
    $addMethod(new ServiceMethod<Todo_QueryResponse, Todo_QueryRequest>(
        'QueryTodo',
        queryTodo_Pre,
        false,
        false,
        (List<int> value) => new Todo_QueryResponse.fromBuffer(value),
        (Todo_QueryRequest value) => value.writeToBuffer()));
    $addMethod(new ServiceMethod<Todo_GetRequest, Todo_GetResponse>(
        'GetTodo',
        getTodo_Pre,
        false,
        false,
        (List<int> value) => new Todo_GetRequest.fromBuffer(value),
        (Todo_GetResponse value) => value.writeToBuffer()));
    $addMethod(new ServiceMethod<Todo_CreateRequest, Todo_MutateResponse>(
        'CreateTodo',
        createTodo_Pre,
        false,
        false,
        (List<int> value) => new Todo_CreateRequest.fromBuffer(value),
        (Todo_MutateResponse value) => value.writeToBuffer()));
    $addMethod(new ServiceMethod<Todo_UpdateRequest, Todo_MutateResponse>(
        'UpdateTodo',
        updateTodo_Pre,
        false,
        false,
        (List<int> value) => new Todo_UpdateRequest.fromBuffer(value),
        (Todo_MutateResponse value) => value.writeToBuffer()));
    $addMethod(new ServiceMethod<Todo_DeleteRequest, Todo_MutateResponse>(
        'DeleteTodo',
        deleteTodo_Pre,
        false,
        false,
        (List<int> value) => new Todo_DeleteRequest.fromBuffer(value),
        (Todo_MutateResponse value) => value.writeToBuffer()));
  }

  Future<Todo_QueryRequest> queryTodo_Pre(
      ServiceCall call, Future request) async {
    return queryTodo(call, await request);
  }

  Future<Todo_GetResponse> getTodo_Pre(ServiceCall call, Future request) async {
    return getTodo(call, await request);
  }

  Future<Todo_MutateResponse> createTodo_Pre(
      ServiceCall call, Future request) async {
    return createTodo(call, await request);
  }

  Future<Todo_MutateResponse> updateTodo_Pre(
      ServiceCall call, Future request) async {
    return updateTodo(call, await request);
  }

  Future<Todo_MutateResponse> deleteTodo_Pre(
      ServiceCall call, Future request) async {
    return deleteTodo(call, await request);
  }

  Future<Todo_QueryRequest> queryTodo(
      ServiceCall call, Todo_QueryResponse request);
  Future<Todo_GetResponse> getTodo(ServiceCall call, Todo_GetRequest request);
  Future<Todo_MutateResponse> createTodo(
      ServiceCall call, Todo_CreateRequest request);
  Future<Todo_MutateResponse> updateTodo(
      ServiceCall call, Todo_UpdateRequest request);
  Future<Todo_MutateResponse> deleteTodo(
      ServiceCall call, Todo_DeleteRequest request);
}
