part of swagger.api;

@Entity()
class TodoCreateRequest {
  
  @Property(name: 'record')
  TodoRecordInput record = null;
  
  TodoCreateRequest();

  @override
  String toString()  {
    return 'TodoCreateRequest[record=$record, ]';
  }
}

