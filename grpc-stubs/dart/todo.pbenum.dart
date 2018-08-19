///
//  Generated code. Do not modify.
///
// ignore_for_file: non_constant_identifier_names,library_prefixes

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' show int, dynamic, String, List, Map;
import 'package:protobuf/protobuf.dart';

class Todo_Record_Status extends ProtobufEnum {
  static const Todo_Record_Status STATUS_TODO = const Todo_Record_Status._(0, 'STATUS_TODO');
  static const Todo_Record_Status STATUS_INPROGRESS = const Todo_Record_Status._(1, 'STATUS_INPROGRESS');
  static const Todo_Record_Status STATUS_DONE = const Todo_Record_Status._(2, 'STATUS_DONE');

  static const List<Todo_Record_Status> values = const <Todo_Record_Status> [
    STATUS_TODO,
    STATUS_INPROGRESS,
    STATUS_DONE,
  ];

  static final Map<int, dynamic> _byValue = ProtobufEnum.initByValue(values);
  static Todo_Record_Status valueOf(int value) => _byValue[value] as Todo_Record_Status;
  static void $checkItem(Todo_Record_Status v) {
    if (v is! Todo_Record_Status) checkItemFailed(v, 'Todo_Record_Status');
  }

  const Todo_Record_Status._(int v, String n) : super(v, n);
}

