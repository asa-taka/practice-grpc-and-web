///
//  Generated code. Do not modify.
///
// ignore_for_file: non_constant_identifier_names,library_prefixes

const Todo$json = const {
  '1': 'Todo',
  '3': const [Todo_Record$json, Todo_RecordInput$json, Todo_GetRequest$json, Todo_GetResponse$json, Todo_QueryRequest$json, Todo_QueryResponse$json, Todo_CreateRequest$json, Todo_UpdateRequest$json, Todo_DeleteRequest$json, Todo_MutateResponse$json],
};

const Todo_Record$json = const {
  '1': 'Record',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    const {'1': 'created_at', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    const {'1': 'title', '3': 3, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'detail', '3': 4, '4': 1, '5': 9, '10': 'detail'},
    const {'1': 'deadline', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'deadline'},
    const {'1': 'status', '3': 6, '4': 1, '5': 14, '6': '.todo.Todo.Record.Status', '10': 'status'},
  ],
  '4': const [Todo_Record_Status$json],
  '7': const {},
};

const Todo_Record_Status$json = const {
  '1': 'Status',
  '2': const [
    const {'1': 'STATUS_TODO', '2': 0},
    const {'1': 'STATUS_INPROGRESS', '2': 1},
    const {'1': 'STATUS_DONE', '2': 2},
  ],
};

const Todo_RecordInput$json = const {
  '1': 'RecordInput',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    const {'1': 'title', '3': 2, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'detail', '3': 3, '4': 1, '5': 9, '10': 'detail'},
    const {'1': 'deadline', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'deadline'},
  ],
};

const Todo_GetRequest$json = const {
  '1': 'GetRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
  ],
  '7': const {},
};

const Todo_GetResponse$json = const {
  '1': 'GetResponse',
  '2': const [
    const {'1': 'record', '3': 1, '4': 1, '5': 11, '6': '.todo.Todo.Record', '10': 'record'},
  ],
};

const Todo_QueryRequest$json = const {
  '1': 'QueryRequest',
  '2': const [
    const {'1': 'start', '3': 1, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'start'},
    const {'1': 'end', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'end'},
  ],
};

const Todo_QueryResponse$json = const {
  '1': 'QueryResponse',
  '2': const [
    const {'1': 'records', '3': 1, '4': 3, '5': 11, '6': '.todo.Todo.Record', '10': 'records'},
  ],
};

const Todo_CreateRequest$json = const {
  '1': 'CreateRequest',
  '2': const [
    const {'1': 'record', '3': 1, '4': 1, '5': 11, '6': '.todo.Todo.RecordInput', '10': 'record'},
  ],
};

const Todo_UpdateRequest$json = const {
  '1': 'UpdateRequest',
  '2': const [
    const {'1': 'record', '3': 1, '4': 1, '5': 11, '6': '.todo.Todo.RecordInput', '10': 'record'},
    const {'1': 'id', '3': 2, '4': 1, '5': 5, '10': 'id'},
  ],
};

const Todo_DeleteRequest$json = const {
  '1': 'DeleteRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
  ],
};

const Todo_MutateResponse$json = const {
  '1': 'MutateResponse',
  '2': const [
    const {'1': 'record', '3': 1, '4': 1, '5': 11, '6': '.todo.Todo.Record', '10': 'record'},
  ],
};

