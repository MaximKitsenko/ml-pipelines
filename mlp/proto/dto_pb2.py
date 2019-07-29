# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/dto.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto import events_pb2 as proto_dot_events__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/dto.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x0fproto/dto.proto\x1a\x12proto/events.proto\"w\n\x0bProjectData\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x15\n\rdataset_count\x18\x03 \x01(\x05\x12\x11\n\traw_bytes\x18\x04 \x01(\x03\x12\x11\n\tzip_bytes\x18\x05 \x01(\x03\x12\x11\n\tjob_count\x18\x06 \x01(\x05\"0\n\x06\x45xpert\x12\x11\n\texpert_id\x18\x01 \x01(\t\x12\x13\n\x0b\x65xpert_name\x18\x02 \x01(\t\"\x81\x05\n\x0b\x44\x61tasetData\x12\x12\n\ndataset_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nproject_id\x18\x03 \x01(\t\x12\x14\n\x0crecord_count\x18\x04 \x01(\x03\x12\x18\n\x10record_count_set\x18\x05 \x01(\x08\x12\x12\n\nfile_count\x18\x06 \x01(\x03\x12\x16\n\x0e\x66ile_count_set\x18\x07 \x01(\x08\x12\x11\n\traw_bytes\x18\x08 \x01(\x03\x12\x15\n\rraw_bytes_set\x18\t \x01(\x08\x12\x11\n\tzip_bytes\x18\n \x01(\x03\x12\x15\n\rzip_bytes_set\x18\x0b \x01(\x08\x12\x13\n\x0bsample_body\x18\x0c \x01(\x0c\x12!\n\x0bsample_kind\x18\r \x01(\x0e\x32\x0c.SAMPLE_KIND\x12\x12\n\nsample_set\x18\x0e \x01(\x08\x12\x18\n\x10update_timestamp\x18\x0f \x01(\x03\x12\x1c\n\x14update_timestamp_set\x18\x10 \x01(\x08\x12\x13\n\x0b\x64\x61ta_format\x18\x11 \x01(\t\x12\x17\n\x0f\x64\x61ta_format_set\x18\x12 \x01(\x08\x12\x13\n\x0b\x64\x65scription\x18\x13 \x01(\t\x12\x17\n\x0f\x64\x65scription_set\x18\x14 \x01(\x08\x12\x15\n\rupstream_jobs\x18\x15 \x03(\t\x12\x17\n\x0f\x64ownstream_jobs\x18\x16 \x03(\t\x12\x18\n\x07\x65xperts\x18\x17 \x03(\x0b\x32\x07.Expert\x12\x12\n\nstorage_id\x18\x18 \x01(\t\x12\x16\n\x0estorage_id_set\x18\x19 \x01(\x08\x12\x18\n\x10storage_location\x18\x1a \x01(\t\x12\x1c\n\x14storage_location_set\x18\x1b \x01(\x08\"v\n\x03Job\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12\x10\n\x08job_name\x18\x02 \x01(\t\x12\x12\n\nproject_id\x18\x03 \x01(\t\x12\x0e\n\x06inputs\x18\x04 \x03(\t\x12\x0f\n\x07outputs\x18\x05 \x03(\t\x12\x18\n\x07\x65xperts\x18\x06 \x03(\x0b\x32\x07.Expert\"\x8f\x01\n\x0bTenantStats\x12\x15\n\rproject_count\x18\x01 \x01(\x05\x12\x15\n\rdataset_count\x18\x02 \x01(\x05\x12\x11\n\tjob_count\x18\x03 \x01(\x05\x12\x14\n\x0creport_count\x18\x04 \x01(\x05\x12\x14\n\x0c\x65xpert_count\x18\x05 \x01(\x05\x12\x13\n\x0bmodel_count\x18\x06 \x01(\x05\"*\n\nAssetCache\x12\x0e\n\x06\x64igest\x18\x01 \x01(\t\x12\x0c\n\x04\x62ody\x18\x02 \x01(\x0c\"K\n\tChunkData\x12\x0c\n\x04hash\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08raw_size\x18\x03 \x01(\x03\x12\x10\n\x08zip_size\x18\x04 \x01(\x03\"M\n\nSourceLink\x12\x17\n\x04kind\x18\x01 \x01(\x0e\x32\t.LinkType\x12\x13\n\x0bsource_name\x18\x02 \x01(\t\x12\x11\n\tuser_name\x18\x03 \x01(\t\"\xb8\x01\n\x0bVersionData\x12\x0f\n\x07version\x18\x01 \x01(\x05\x12\x1a\n\x06\x63hunks\x18\x02 \x03(\x0b\x32\n.ChunkData\x12\x1c\n\x07sources\x18\x03 \x03(\x0b\x32\x0b.SourceLink\x12\x10\n\x08raw_size\x18\x04 \x01(\x03\x12\x10\n\x08zip_size\x18\x05 \x01(\x03\x12\x12\n\ndataset_id\x18\x06 \x01(\t\x12&\n\ndescriptor\x18\x07 \x01(\x0b\x32\x12.DatasetDescriptor\")\n\x0b\x46ieldSchema\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\"l\n\x11\x44\x61tasetDescriptor\x12\x0c\n\x04type\x18\x01 \x01(\t\x12\x14\n\x0crecord_count\x18\x02 \x01(\x03\x12\x1c\n\x06schema\x18\x03 \x03(\x0b\x32\x0c.FieldSchema\x12\x15\n\rsample_string\x18\x04 \x01(\t*%\n\x08LinkType\x12\r\n\tUNIVERSAL\x10\x00\x12\n\n\x06UPLOAD\x10\x01\x62\x06proto3')
  ,
  dependencies=[proto_dot_events__pb2.DESCRIPTOR,])

_LINKTYPE = _descriptor.EnumDescriptor(
  name='LinkType',
  full_name='LinkType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNIVERSAL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UPLOAD', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1660,
  serialized_end=1697,
)
_sym_db.RegisterEnumDescriptor(_LINKTYPE)

LinkType = enum_type_wrapper.EnumTypeWrapper(_LINKTYPE)
UNIVERSAL = 0
UPLOAD = 1



_PROJECTDATA = _descriptor.Descriptor(
  name='ProjectData',
  full_name='ProjectData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='ProjectData.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='id', full_name='ProjectData.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_count', full_name='ProjectData.dataset_count', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_bytes', full_name='ProjectData.raw_bytes', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zip_bytes', full_name='ProjectData.zip_bytes', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='job_count', full_name='ProjectData.job_count', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=39,
  serialized_end=158,
)


_EXPERT = _descriptor.Descriptor(
  name='Expert',
  full_name='Expert',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='expert_id', full_name='Expert.expert_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='expert_name', full_name='Expert.expert_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=160,
  serialized_end=208,
)


_DATASETDATA = _descriptor.Descriptor(
  name='DatasetData',
  full_name='DatasetData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='DatasetData.dataset_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='DatasetData.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetData.project_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record_count', full_name='DatasetData.record_count', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record_count_set', full_name='DatasetData.record_count_set', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count', full_name='DatasetData.file_count', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count_set', full_name='DatasetData.file_count_set', index=6,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_bytes', full_name='DatasetData.raw_bytes', index=7,
      number=8, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_bytes_set', full_name='DatasetData.raw_bytes_set', index=8,
      number=9, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zip_bytes', full_name='DatasetData.zip_bytes', index=9,
      number=10, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zip_bytes_set', full_name='DatasetData.zip_bytes_set', index=10,
      number=11, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_body', full_name='DatasetData.sample_body', index=11,
      number=12, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_kind', full_name='DatasetData.sample_kind', index=12,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_set', full_name='DatasetData.sample_set', index=13,
      number=14, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp', full_name='DatasetData.update_timestamp', index=14,
      number=15, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp_set', full_name='DatasetData.update_timestamp_set', index=15,
      number=16, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format', full_name='DatasetData.data_format', index=16,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format_set', full_name='DatasetData.data_format_set', index=17,
      number=18, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='DatasetData.description', index=18,
      number=19, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description_set', full_name='DatasetData.description_set', index=19,
      number=20, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='upstream_jobs', full_name='DatasetData.upstream_jobs', index=20,
      number=21, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='downstream_jobs', full_name='DatasetData.downstream_jobs', index=21,
      number=22, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts', full_name='DatasetData.experts', index=22,
      number=23, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_id', full_name='DatasetData.storage_id', index=23,
      number=24, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_id_set', full_name='DatasetData.storage_id_set', index=24,
      number=25, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_location', full_name='DatasetData.storage_location', index=25,
      number=26, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_location_set', full_name='DatasetData.storage_location_set', index=26,
      number=27, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=211,
  serialized_end=852,
)


_JOB = _descriptor.Descriptor(
  name='Job',
  full_name='Job',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='job_id', full_name='Job.job_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='job_name', full_name='Job.job_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='Job.project_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs', full_name='Job.inputs', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='Job.outputs', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts', full_name='Job.experts', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=854,
  serialized_end=972,
)


_TENANTSTATS = _descriptor.Descriptor(
  name='TenantStats',
  full_name='TenantStats',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_count', full_name='TenantStats.project_count', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_count', full_name='TenantStats.dataset_count', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='job_count', full_name='TenantStats.job_count', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='report_count', full_name='TenantStats.report_count', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='expert_count', full_name='TenantStats.expert_count', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='model_count', full_name='TenantStats.model_count', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=975,
  serialized_end=1118,
)


_ASSETCACHE = _descriptor.Descriptor(
  name='AssetCache',
  full_name='AssetCache',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='digest', full_name='AssetCache.digest', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body', full_name='AssetCache.body', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1120,
  serialized_end=1162,
)


_CHUNKDATA = _descriptor.Descriptor(
  name='ChunkData',
  full_name='ChunkData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hash', full_name='ChunkData.hash', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='ChunkData.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_size', full_name='ChunkData.raw_size', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zip_size', full_name='ChunkData.zip_size', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1164,
  serialized_end=1239,
)


_SOURCELINK = _descriptor.Descriptor(
  name='SourceLink',
  full_name='SourceLink',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='kind', full_name='SourceLink.kind', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source_name', full_name='SourceLink.source_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user_name', full_name='SourceLink.user_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1241,
  serialized_end=1318,
)


_VERSIONDATA = _descriptor.Descriptor(
  name='VersionData',
  full_name='VersionData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version', full_name='VersionData.version', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='chunks', full_name='VersionData.chunks', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sources', full_name='VersionData.sources', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_size', full_name='VersionData.raw_size', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zip_size', full_name='VersionData.zip_size', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='VersionData.dataset_id', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='descriptor', full_name='VersionData.descriptor', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1321,
  serialized_end=1505,
)


_FIELDSCHEMA = _descriptor.Descriptor(
  name='FieldSchema',
  full_name='FieldSchema',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='FieldSchema.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='FieldSchema.type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1507,
  serialized_end=1548,
)


_DATASETDESCRIPTOR = _descriptor.Descriptor(
  name='DatasetDescriptor',
  full_name='DatasetDescriptor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='DatasetDescriptor.type', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record_count', full_name='DatasetDescriptor.record_count', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='schema', full_name='DatasetDescriptor.schema', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_string', full_name='DatasetDescriptor.sample_string', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1550,
  serialized_end=1658,
)

_DATASETDATA.fields_by_name['sample_kind'].enum_type = proto_dot_events__pb2._SAMPLE_KIND
_DATASETDATA.fields_by_name['experts'].message_type = _EXPERT
_JOB.fields_by_name['experts'].message_type = _EXPERT
_SOURCELINK.fields_by_name['kind'].enum_type = _LINKTYPE
_VERSIONDATA.fields_by_name['chunks'].message_type = _CHUNKDATA
_VERSIONDATA.fields_by_name['sources'].message_type = _SOURCELINK
_VERSIONDATA.fields_by_name['descriptor'].message_type = _DATASETDESCRIPTOR
_DATASETDESCRIPTOR.fields_by_name['schema'].message_type = _FIELDSCHEMA
DESCRIPTOR.message_types_by_name['ProjectData'] = _PROJECTDATA
DESCRIPTOR.message_types_by_name['Expert'] = _EXPERT
DESCRIPTOR.message_types_by_name['DatasetData'] = _DATASETDATA
DESCRIPTOR.message_types_by_name['Job'] = _JOB
DESCRIPTOR.message_types_by_name['TenantStats'] = _TENANTSTATS
DESCRIPTOR.message_types_by_name['AssetCache'] = _ASSETCACHE
DESCRIPTOR.message_types_by_name['ChunkData'] = _CHUNKDATA
DESCRIPTOR.message_types_by_name['SourceLink'] = _SOURCELINK
DESCRIPTOR.message_types_by_name['VersionData'] = _VERSIONDATA
DESCRIPTOR.message_types_by_name['FieldSchema'] = _FIELDSCHEMA
DESCRIPTOR.message_types_by_name['DatasetDescriptor'] = _DATASETDESCRIPTOR
DESCRIPTOR.enum_types_by_name['LinkType'] = _LINKTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProjectData = _reflection.GeneratedProtocolMessageType('ProjectData', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTDATA,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:ProjectData)
  ))
_sym_db.RegisterMessage(ProjectData)

Expert = _reflection.GeneratedProtocolMessageType('Expert', (_message.Message,), dict(
  DESCRIPTOR = _EXPERT,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:Expert)
  ))
_sym_db.RegisterMessage(Expert)

DatasetData = _reflection.GeneratedProtocolMessageType('DatasetData', (_message.Message,), dict(
  DESCRIPTOR = _DATASETDATA,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:DatasetData)
  ))
_sym_db.RegisterMessage(DatasetData)

Job = _reflection.GeneratedProtocolMessageType('Job', (_message.Message,), dict(
  DESCRIPTOR = _JOB,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:Job)
  ))
_sym_db.RegisterMessage(Job)

TenantStats = _reflection.GeneratedProtocolMessageType('TenantStats', (_message.Message,), dict(
  DESCRIPTOR = _TENANTSTATS,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:TenantStats)
  ))
_sym_db.RegisterMessage(TenantStats)

AssetCache = _reflection.GeneratedProtocolMessageType('AssetCache', (_message.Message,), dict(
  DESCRIPTOR = _ASSETCACHE,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:AssetCache)
  ))
_sym_db.RegisterMessage(AssetCache)

ChunkData = _reflection.GeneratedProtocolMessageType('ChunkData', (_message.Message,), dict(
  DESCRIPTOR = _CHUNKDATA,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:ChunkData)
  ))
_sym_db.RegisterMessage(ChunkData)

SourceLink = _reflection.GeneratedProtocolMessageType('SourceLink', (_message.Message,), dict(
  DESCRIPTOR = _SOURCELINK,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:SourceLink)
  ))
_sym_db.RegisterMessage(SourceLink)

VersionData = _reflection.GeneratedProtocolMessageType('VersionData', (_message.Message,), dict(
  DESCRIPTOR = _VERSIONDATA,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:VersionData)
  ))
_sym_db.RegisterMessage(VersionData)

FieldSchema = _reflection.GeneratedProtocolMessageType('FieldSchema', (_message.Message,), dict(
  DESCRIPTOR = _FIELDSCHEMA,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:FieldSchema)
  ))
_sym_db.RegisterMessage(FieldSchema)

DatasetDescriptor = _reflection.GeneratedProtocolMessageType('DatasetDescriptor', (_message.Message,), dict(
  DESCRIPTOR = _DATASETDESCRIPTOR,
  __module__ = 'proto.dto_pb2'
  # @@protoc_insertion_point(class_scope:DatasetDescriptor)
  ))
_sym_db.RegisterMessage(DatasetDescriptor)


# @@protoc_insertion_point(module_scope)
