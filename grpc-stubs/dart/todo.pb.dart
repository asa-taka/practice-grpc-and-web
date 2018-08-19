///
//  Generated code. Do not modify.
///
// ignore_for_file: non_constant_identifier_names,library_prefixes

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:protobuf/protobuf.dart';

import 'google/protobuf/timestamp.pb.dart' as $google$protobuf;

import 'todo.pbenum.dart';

export 'todo.pbenum.dart';

class Todo_Record extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_Record')
    ..a<int>(1, 'id', PbFieldType.O3)
    ..a<$google$protobuf.Timestamp>(2, 'createdAt', PbFieldType.OM, $google$protobuf.Timestamp.getDefault, $google$protobuf.Timestamp.create)
    ..aOS(3, 'title')
    ..aOS(4, 'detail')
    ..a<$google$protobuf.Timestamp>(5, 'deadline', PbFieldType.OM, $google$protobuf.Timestamp.getDefault, $google$protobuf.Timestamp.create)
    ..e<Todo_Record_Status>(6, 'status', PbFieldType.OE, Todo_Record_Status.STATUS_TODO, Todo_Record_Status.valueOf, Todo_Record_Status.values)
    ..hasRequiredFields = false
  ;

  Todo_Record() : super();
  Todo_Record.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_Record.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_Record clone() => new Todo_Record()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_Record create() => new Todo_Record();
  static PbList<Todo_Record> createRepeated() => new PbList<Todo_Record>();
  static Todo_Record getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_Record();
    return _defaultInstance;
  }
  static Todo_Record _defaultInstance;
  static void $checkItem(Todo_Record v) {
    if (v is! Todo_Record) checkItemFailed(v, 'Todo_Record');
  }

  int get id => $_get(0, 0);
  set id(int v) { $_setSignedInt32(0, v); }
  bool hasId() => $_has(0);
  void clearId() => clearField(1);

  $google$protobuf.Timestamp get createdAt => $_getN(1);
  set createdAt($google$protobuf.Timestamp v) { setField(2, v); }
  bool hasCreatedAt() => $_has(1);
  void clearCreatedAt() => clearField(2);

  String get title => $_getS(2, '');
  set title(String v) { $_setString(2, v); }
  bool hasTitle() => $_has(2);
  void clearTitle() => clearField(3);

  String get detail => $_getS(3, '');
  set detail(String v) { $_setString(3, v); }
  bool hasDetail() => $_has(3);
  void clearDetail() => clearField(4);

  $google$protobuf.Timestamp get deadline => $_getN(4);
  set deadline($google$protobuf.Timestamp v) { setField(5, v); }
  bool hasDeadline() => $_has(4);
  void clearDeadline() => clearField(5);

  Todo_Record_Status get status => $_getN(5);
  set status(Todo_Record_Status v) { setField(6, v); }
  bool hasStatus() => $_has(5);
  void clearStatus() => clearField(6);
}

class _ReadonlyTodo_Record extends Todo_Record with ReadonlyMessageMixin {}

class Todo_RecordInput extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_RecordInput')
    ..a<int>(1, 'id', PbFieldType.O3)
    ..aOS(2, 'title')
    ..aOS(3, 'detail')
    ..a<$google$protobuf.Timestamp>(4, 'deadline', PbFieldType.OM, $google$protobuf.Timestamp.getDefault, $google$protobuf.Timestamp.create)
    ..hasRequiredFields = false
  ;

  Todo_RecordInput() : super();
  Todo_RecordInput.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_RecordInput.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_RecordInput clone() => new Todo_RecordInput()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_RecordInput create() => new Todo_RecordInput();
  static PbList<Todo_RecordInput> createRepeated() => new PbList<Todo_RecordInput>();
  static Todo_RecordInput getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_RecordInput();
    return _defaultInstance;
  }
  static Todo_RecordInput _defaultInstance;
  static void $checkItem(Todo_RecordInput v) {
    if (v is! Todo_RecordInput) checkItemFailed(v, 'Todo_RecordInput');
  }

  int get id => $_get(0, 0);
  set id(int v) { $_setSignedInt32(0, v); }
  bool hasId() => $_has(0);
  void clearId() => clearField(1);

  String get title => $_getS(1, '');
  set title(String v) { $_setString(1, v); }
  bool hasTitle() => $_has(1);
  void clearTitle() => clearField(2);

  String get detail => $_getS(2, '');
  set detail(String v) { $_setString(2, v); }
  bool hasDetail() => $_has(2);
  void clearDetail() => clearField(3);

  $google$protobuf.Timestamp get deadline => $_getN(3);
  set deadline($google$protobuf.Timestamp v) { setField(4, v); }
  bool hasDeadline() => $_has(3);
  void clearDeadline() => clearField(4);
}

class _ReadonlyTodo_RecordInput extends Todo_RecordInput with ReadonlyMessageMixin {}

class Todo_GetRequest extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_GetRequest')
    ..a<int>(1, 'id', PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  Todo_GetRequest() : super();
  Todo_GetRequest.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_GetRequest.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_GetRequest clone() => new Todo_GetRequest()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_GetRequest create() => new Todo_GetRequest();
  static PbList<Todo_GetRequest> createRepeated() => new PbList<Todo_GetRequest>();
  static Todo_GetRequest getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_GetRequest();
    return _defaultInstance;
  }
  static Todo_GetRequest _defaultInstance;
  static void $checkItem(Todo_GetRequest v) {
    if (v is! Todo_GetRequest) checkItemFailed(v, 'Todo_GetRequest');
  }

  int get id => $_get(0, 0);
  set id(int v) { $_setSignedInt32(0, v); }
  bool hasId() => $_has(0);
  void clearId() => clearField(1);
}

class _ReadonlyTodo_GetRequest extends Todo_GetRequest with ReadonlyMessageMixin {}

class Todo_GetResponse extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_GetResponse')
    ..a<Todo_Record>(1, 'record', PbFieldType.OM, Todo_Record.getDefault, Todo_Record.create)
    ..hasRequiredFields = false
  ;

  Todo_GetResponse() : super();
  Todo_GetResponse.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_GetResponse.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_GetResponse clone() => new Todo_GetResponse()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_GetResponse create() => new Todo_GetResponse();
  static PbList<Todo_GetResponse> createRepeated() => new PbList<Todo_GetResponse>();
  static Todo_GetResponse getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_GetResponse();
    return _defaultInstance;
  }
  static Todo_GetResponse _defaultInstance;
  static void $checkItem(Todo_GetResponse v) {
    if (v is! Todo_GetResponse) checkItemFailed(v, 'Todo_GetResponse');
  }

  Todo_Record get record => $_getN(0);
  set record(Todo_Record v) { setField(1, v); }
  bool hasRecord() => $_has(0);
  void clearRecord() => clearField(1);
}

class _ReadonlyTodo_GetResponse extends Todo_GetResponse with ReadonlyMessageMixin {}

class Todo_QueryRequest extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_QueryRequest')
    ..a<$google$protobuf.Timestamp>(1, 'start', PbFieldType.OM, $google$protobuf.Timestamp.getDefault, $google$protobuf.Timestamp.create)
    ..a<$google$protobuf.Timestamp>(2, 'end', PbFieldType.OM, $google$protobuf.Timestamp.getDefault, $google$protobuf.Timestamp.create)
    ..hasRequiredFields = false
  ;

  Todo_QueryRequest() : super();
  Todo_QueryRequest.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_QueryRequest.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_QueryRequest clone() => new Todo_QueryRequest()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_QueryRequest create() => new Todo_QueryRequest();
  static PbList<Todo_QueryRequest> createRepeated() => new PbList<Todo_QueryRequest>();
  static Todo_QueryRequest getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_QueryRequest();
    return _defaultInstance;
  }
  static Todo_QueryRequest _defaultInstance;
  static void $checkItem(Todo_QueryRequest v) {
    if (v is! Todo_QueryRequest) checkItemFailed(v, 'Todo_QueryRequest');
  }

  $google$protobuf.Timestamp get start => $_getN(0);
  set start($google$protobuf.Timestamp v) { setField(1, v); }
  bool hasStart() => $_has(0);
  void clearStart() => clearField(1);

  $google$protobuf.Timestamp get end => $_getN(1);
  set end($google$protobuf.Timestamp v) { setField(2, v); }
  bool hasEnd() => $_has(1);
  void clearEnd() => clearField(2);
}

class _ReadonlyTodo_QueryRequest extends Todo_QueryRequest with ReadonlyMessageMixin {}

class Todo_QueryResponse extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_QueryResponse')
    ..pp<Todo_Record>(1, 'records', PbFieldType.PM, Todo_Record.$checkItem, Todo_Record.create)
    ..hasRequiredFields = false
  ;

  Todo_QueryResponse() : super();
  Todo_QueryResponse.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_QueryResponse.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_QueryResponse clone() => new Todo_QueryResponse()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_QueryResponse create() => new Todo_QueryResponse();
  static PbList<Todo_QueryResponse> createRepeated() => new PbList<Todo_QueryResponse>();
  static Todo_QueryResponse getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_QueryResponse();
    return _defaultInstance;
  }
  static Todo_QueryResponse _defaultInstance;
  static void $checkItem(Todo_QueryResponse v) {
    if (v is! Todo_QueryResponse) checkItemFailed(v, 'Todo_QueryResponse');
  }

  List<Todo_Record> get records => $_getList(0);
}

class _ReadonlyTodo_QueryResponse extends Todo_QueryResponse with ReadonlyMessageMixin {}

class Todo_CreateRequest extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_CreateRequest')
    ..a<Todo_RecordInput>(1, 'record', PbFieldType.OM, Todo_RecordInput.getDefault, Todo_RecordInput.create)
    ..hasRequiredFields = false
  ;

  Todo_CreateRequest() : super();
  Todo_CreateRequest.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_CreateRequest.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_CreateRequest clone() => new Todo_CreateRequest()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_CreateRequest create() => new Todo_CreateRequest();
  static PbList<Todo_CreateRequest> createRepeated() => new PbList<Todo_CreateRequest>();
  static Todo_CreateRequest getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_CreateRequest();
    return _defaultInstance;
  }
  static Todo_CreateRequest _defaultInstance;
  static void $checkItem(Todo_CreateRequest v) {
    if (v is! Todo_CreateRequest) checkItemFailed(v, 'Todo_CreateRequest');
  }

  Todo_RecordInput get record => $_getN(0);
  set record(Todo_RecordInput v) { setField(1, v); }
  bool hasRecord() => $_has(0);
  void clearRecord() => clearField(1);
}

class _ReadonlyTodo_CreateRequest extends Todo_CreateRequest with ReadonlyMessageMixin {}

class Todo_UpdateRequest extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_UpdateRequest')
    ..a<Todo_RecordInput>(1, 'record', PbFieldType.OM, Todo_RecordInput.getDefault, Todo_RecordInput.create)
    ..a<int>(2, 'id', PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  Todo_UpdateRequest() : super();
  Todo_UpdateRequest.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_UpdateRequest.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_UpdateRequest clone() => new Todo_UpdateRequest()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_UpdateRequest create() => new Todo_UpdateRequest();
  static PbList<Todo_UpdateRequest> createRepeated() => new PbList<Todo_UpdateRequest>();
  static Todo_UpdateRequest getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_UpdateRequest();
    return _defaultInstance;
  }
  static Todo_UpdateRequest _defaultInstance;
  static void $checkItem(Todo_UpdateRequest v) {
    if (v is! Todo_UpdateRequest) checkItemFailed(v, 'Todo_UpdateRequest');
  }

  Todo_RecordInput get record => $_getN(0);
  set record(Todo_RecordInput v) { setField(1, v); }
  bool hasRecord() => $_has(0);
  void clearRecord() => clearField(1);

  int get id => $_get(1, 0);
  set id(int v) { $_setSignedInt32(1, v); }
  bool hasId() => $_has(1);
  void clearId() => clearField(2);
}

class _ReadonlyTodo_UpdateRequest extends Todo_UpdateRequest with ReadonlyMessageMixin {}

class Todo_DeleteRequest extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_DeleteRequest')
    ..a<int>(1, 'id', PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  Todo_DeleteRequest() : super();
  Todo_DeleteRequest.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_DeleteRequest.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_DeleteRequest clone() => new Todo_DeleteRequest()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_DeleteRequest create() => new Todo_DeleteRequest();
  static PbList<Todo_DeleteRequest> createRepeated() => new PbList<Todo_DeleteRequest>();
  static Todo_DeleteRequest getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_DeleteRequest();
    return _defaultInstance;
  }
  static Todo_DeleteRequest _defaultInstance;
  static void $checkItem(Todo_DeleteRequest v) {
    if (v is! Todo_DeleteRequest) checkItemFailed(v, 'Todo_DeleteRequest');
  }

  int get id => $_get(0, 0);
  set id(int v) { $_setSignedInt32(0, v); }
  bool hasId() => $_has(0);
  void clearId() => clearField(1);
}

class _ReadonlyTodo_DeleteRequest extends Todo_DeleteRequest with ReadonlyMessageMixin {}

class Todo_MutateResponse extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo_MutateResponse')
    ..a<Todo_Record>(1, 'record', PbFieldType.OM, Todo_Record.getDefault, Todo_Record.create)
    ..hasRequiredFields = false
  ;

  Todo_MutateResponse() : super();
  Todo_MutateResponse.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo_MutateResponse.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo_MutateResponse clone() => new Todo_MutateResponse()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo_MutateResponse create() => new Todo_MutateResponse();
  static PbList<Todo_MutateResponse> createRepeated() => new PbList<Todo_MutateResponse>();
  static Todo_MutateResponse getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo_MutateResponse();
    return _defaultInstance;
  }
  static Todo_MutateResponse _defaultInstance;
  static void $checkItem(Todo_MutateResponse v) {
    if (v is! Todo_MutateResponse) checkItemFailed(v, 'Todo_MutateResponse');
  }

  Todo_Record get record => $_getN(0);
  set record(Todo_Record v) { setField(1, v); }
  bool hasRecord() => $_has(0);
  void clearRecord() => clearField(1);
}

class _ReadonlyTodo_MutateResponse extends Todo_MutateResponse with ReadonlyMessageMixin {}

class Todo extends GeneratedMessage {
  static final BuilderInfo _i = new BuilderInfo('Todo')
    ..hasRequiredFields = false
  ;

  Todo() : super();
  Todo.fromBuffer(List<int> i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Todo.fromJson(String i, [ExtensionRegistry r = ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Todo clone() => new Todo()..mergeFromMessage(this);
  BuilderInfo get info_ => _i;
  static Todo create() => new Todo();
  static PbList<Todo> createRepeated() => new PbList<Todo>();
  static Todo getDefault() {
    if (_defaultInstance == null) _defaultInstance = new _ReadonlyTodo();
    return _defaultInstance;
  }
  static Todo _defaultInstance;
  static void $checkItem(Todo v) {
    if (v is! Todo) checkItemFailed(v, 'Todo');
  }
}

class _ReadonlyTodo extends Todo with ReadonlyMessageMixin {}

