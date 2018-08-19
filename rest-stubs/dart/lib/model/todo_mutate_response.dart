part of swagger.api;

@Entity()
class TodoMutateResponse {
  
  @Property(name: 'record')
  TodoRecord record = null;
  
  TodoMutateResponse();

  @override
  String toString()  {
    return 'TodoMutateResponse[record=$record, ]';
  }
}

