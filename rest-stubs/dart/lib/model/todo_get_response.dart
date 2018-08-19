part of swagger.api;

@Entity()
class TodoGetResponse {
  
  @Property(name: 'record')
  TodoRecord record = null;
  
  TodoGetResponse();

  @override
  String toString()  {
    return 'TodoGetResponse[record=$record, ]';
  }
}

