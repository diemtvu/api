# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: security/v1beta1/mtls.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from type.v1beta1 import selector_pb2 as type_dot_v1beta1_dot_selector__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='security/v1beta1/mtls.proto',
  package='istio.security.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z\035istio.io/api/security/v1beta1'),
  serialized_pb=_b('\n\x1bsecurity/v1beta1/mtls.proto\x12\x16istio.security.v1beta1\x1a\x1btype/v1beta1/selector.proto\"q\n\nMTLSPolicy\x12\x35\n\x04mode\x18\x01 \x01(\x0e\x32\'.istio.security.v1beta1.MTLSPolicy.Mode\",\n\x04Mode\x12\x08\n\x04NONE\x10\x00\x12\x0e\n\nPERMISSIVE\x10\x01\x12\n\n\x06STRICT\x10\x02\x42\x1fZ\x1distio.io/api/security/v1beta1b\x06proto3')
  ,
  dependencies=[type_dot_v1beta1_dot_selector__pb2.DESCRIPTOR,])



_MTLSPOLICY_MODE = _descriptor.EnumDescriptor(
  name='Mode',
  full_name='istio.security.v1beta1.MTLSPolicy.Mode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PERMISSIVE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STRICT', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=153,
  serialized_end=197,
)
_sym_db.RegisterEnumDescriptor(_MTLSPOLICY_MODE)


_MTLSPOLICY = _descriptor.Descriptor(
  name='MTLSPolicy',
  full_name='istio.security.v1beta1.MTLSPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.security.v1beta1.MTLSPolicy.mode', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MTLSPOLICY_MODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=84,
  serialized_end=197,
)

_MTLSPOLICY.fields_by_name['mode'].enum_type = _MTLSPOLICY_MODE
_MTLSPOLICY_MODE.containing_type = _MTLSPOLICY
DESCRIPTOR.message_types_by_name['MTLSPolicy'] = _MTLSPOLICY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MTLSPolicy = _reflection.GeneratedProtocolMessageType('MTLSPolicy', (_message.Message,), {
  'DESCRIPTOR' : _MTLSPOLICY,
  '__module__' : 'security.v1beta1.mtls_pb2'
  # @@protoc_insertion_point(class_scope:istio.security.v1beta1.MTLSPolicy)
  })
_sym_db.RegisterMessage(MTLSPolicy)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
