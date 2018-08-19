// package: todo
// file: todo.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_api_annotations_pb from "./google/api/annotations_pb";
import * as protoc_gen_swagger_options_annotations_pb from "./protoc-gen-swagger/options/annotations_pb";

export class Todo extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todo.AsObject;
  static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todo;
  static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
  export type AsObject = {
  }

  export class Record extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    hasCreatedAt(): boolean;
    clearCreatedAt(): void;
    getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): void;

    getTitle(): string;
    setTitle(value: string): void;

    getDetail(): string;
    setDetail(value: string): void;

    hasDeadline(): boolean;
    clearDeadline(): void;
    getDeadline(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setDeadline(value?: google_protobuf_timestamp_pb.Timestamp): void;

    getStatus(): Todo.Record.Status;
    setStatus(value: Todo.Record.Status): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Record.AsObject;
    static toObject(includeInstance: boolean, msg: Record): Record.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Record, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Record;
    static deserializeBinaryFromReader(message: Record, reader: jspb.BinaryReader): Record;
  }

  export namespace Record {
    export type AsObject = {
      id: number,
      createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      title: string,
      detail: string,
      deadline?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      status: Todo.Record.Status,
    }

    export enum Status {
      STATUS_TODO = 0,
      STATUS_INPROGRESS = 1,
      STATUS_DONE = 2,
    }
  }

  export class RecordInput extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    getTitle(): string;
    setTitle(value: string): void;

    getDetail(): string;
    setDetail(value: string): void;

    hasDeadline(): boolean;
    clearDeadline(): void;
    getDeadline(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setDeadline(value?: google_protobuf_timestamp_pb.Timestamp): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RecordInput.AsObject;
    static toObject(includeInstance: boolean, msg: RecordInput): RecordInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RecordInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RecordInput;
    static deserializeBinaryFromReader(message: RecordInput, reader: jspb.BinaryReader): RecordInput;
  }

  export namespace RecordInput {
    export type AsObject = {
      id: number,
      title: string,
      detail: string,
      deadline?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
  }

  export class GetRequest extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetRequest;
    static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;
  }

  export namespace GetRequest {
    export type AsObject = {
      id: number,
    }
  }

  export class GetResponse extends jspb.Message {
    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): Todo.Record | undefined;
    setRecord(value?: Todo.Record): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetResponse;
    static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;
  }

  export namespace GetResponse {
    export type AsObject = {
      record?: Todo.Record.AsObject,
    }
  }

  export class QueryRequest extends jspb.Message {
    hasStart(): boolean;
    clearStart(): void;
    getStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setStart(value?: google_protobuf_timestamp_pb.Timestamp): void;

    hasEnd(): boolean;
    clearEnd(): void;
    getEnd(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setEnd(value?: google_protobuf_timestamp_pb.Timestamp): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryRequest.AsObject;
    static toObject(includeInstance: boolean, msg: QueryRequest): QueryRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryRequest;
    static deserializeBinaryFromReader(message: QueryRequest, reader: jspb.BinaryReader): QueryRequest;
  }

  export namespace QueryRequest {
    export type AsObject = {
      start?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      end?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
  }

  export class QueryResponse extends jspb.Message {
    clearRecordsList(): void;
    getRecordsList(): Array<Todo.Record>;
    setRecordsList(value: Array<Todo.Record>): void;
    addRecords(value?: Todo.Record, index?: number): Todo.Record;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryResponse.AsObject;
    static toObject(includeInstance: boolean, msg: QueryResponse): QueryResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryResponse;
    static deserializeBinaryFromReader(message: QueryResponse, reader: jspb.BinaryReader): QueryResponse;
  }

  export namespace QueryResponse {
    export type AsObject = {
      recordsList: Array<Todo.Record.AsObject>,
    }
  }

  export class CreateRequest extends jspb.Message {
    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): Todo.RecordInput | undefined;
    setRecord(value?: Todo.RecordInput): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateRequest): CreateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateRequest;
    static deserializeBinaryFromReader(message: CreateRequest, reader: jspb.BinaryReader): CreateRequest;
  }

  export namespace CreateRequest {
    export type AsObject = {
      record?: Todo.RecordInput.AsObject,
    }
  }

  export class UpdateRequest extends jspb.Message {
    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): Todo.RecordInput | undefined;
    setRecord(value?: Todo.RecordInput): void;

    getId(): number;
    setId(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRequest;
    static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
  }

  export namespace UpdateRequest {
    export type AsObject = {
      record?: Todo.RecordInput.AsObject,
      id: number,
    }
  }

  export class DeleteRequest extends jspb.Message {
    getId(): number;
    setId(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRequest;
    static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
  }

  export namespace DeleteRequest {
    export type AsObject = {
      id: number,
    }
  }

  export class MutateResponse extends jspb.Message {
    hasRecord(): boolean;
    clearRecord(): void;
    getRecord(): Todo.Record | undefined;
    setRecord(value?: Todo.Record): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MutateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: MutateResponse): MutateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MutateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MutateResponse;
    static deserializeBinaryFromReader(message: MutateResponse, reader: jspb.BinaryReader): MutateResponse;
  }

  export namespace MutateResponse {
    export type AsObject = {
      record?: Todo.Record.AsObject,
    }
  }
}

