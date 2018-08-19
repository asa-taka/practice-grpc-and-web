part of swagger.api;

@Entity()
class TodoRecordInput {
  
  @Property(name: 'id')
  int id = null;
  

  @Property(name: 'title')
  String title = null;
  

  @Property(name: 'detail')
  String detail = null;
  

  @Property(name: 'deadline')
  DateTime deadline = null;
  
  TodoRecordInput();

  @override
  String toString()  {
    return 'TodoRecordInput[id=$id, title=$title, detail=$detail, deadline=$deadline, ]';
  }
}

