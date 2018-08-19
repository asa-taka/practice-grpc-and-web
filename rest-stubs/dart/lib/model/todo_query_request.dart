part of swagger.api;

@Entity()
class TodoQueryRequest {
  
  @Property(name: 'start')
  DateTime start = null;
  

  @Property(name: 'end')
  DateTime end = null;
  
  TodoQueryRequest();

  @override
  String toString()  {
    return 'TodoQueryRequest[start=$start, end=$end, ]';
  }
}

