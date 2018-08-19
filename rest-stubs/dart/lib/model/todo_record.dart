part of swagger.api;

@Entity()
class TodoRecord {
  
  @Property(name: 'id')
  int id = null;
  

  @Property(name: 'created_at')
  DateTime createdAt = null;
  

  @Property(name: 'title')
  String title = null;
  

  @Property(name: 'detail')
  String detail = null;
  

  @Property(name: 'deadline')
  DateTime deadline = null;
  

  @Property(name: 'status')
  RecordStatus status = null;
  
  TodoRecord();

  @override
  String toString()  {
    return 'TodoRecord[id=$id, createdAt=$createdAt, title=$title, detail=$detail, deadline=$deadline, status=$status, ]';
  }
}

