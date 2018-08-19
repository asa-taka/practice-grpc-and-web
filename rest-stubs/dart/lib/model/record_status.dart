part of swagger.api;

@Entity()
class RecordStatus {
  /// The underlying value of this enum member.
  final String value;

  const RecordStatus._internal(this.value);

  static const RecordStatus tODO_ = const RecordStatus._internal(&quot;STATUS_TODO&quot;);
  static const RecordStatus iNPROGRESS_ = const RecordStatus._internal(&quot;STATUS_INPROGRESS&quot;);
  static const RecordStatus dONE_ = const RecordStatus._internal(&quot;STATUS_DONE&quot;);
}

class RecordStatusTypeTransformer extends TypeTransformer<RecordStatus> {

  @override
  dynamic encode(RecordStatus data) {
    return data.value;
  }

  @override
  RecordStatus decode(dynamic data) {
    switch (data) {
      case &quot;STATUS_TODO&quot;: return RecordStatus.tODO_;
      case &quot;STATUS_INPROGRESS&quot;: return RecordStatus.iNPROGRESS_;
      case &quot;STATUS_DONE&quot;: return RecordStatus.dONE_;
      default: throw('Unknown enum value to decode: $data');
    }
  }
}

