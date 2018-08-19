part of swagger.api;

@Entity()
class TodoUpdateRequest {
  
  @Property(name: 'record')
  TodoRecordInput record = null;
  

  @Property(name: 'id')
  int id = null;
  
  TodoUpdateRequest();

  @override
  String toString()  {
    return 'TodoUpdateRequest[record=$record, id=$id, ]';
  }
}

