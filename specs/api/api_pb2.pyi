# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from events_pb2 import (
    Event as events_pb2___Event,
)

from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
)

from google.protobuf.internal.containers import (
    RepeatedCompositeFieldContainer as google___protobuf___internal___containers___RepeatedCompositeFieldContainer,
)

from google.protobuf.message import (
    Message as google___protobuf___message___Message,
)

from typing import (
    Iterable as typing___Iterable,
    Optional as typing___Optional,
    Text as typing___Text,
)

from typing_extensions import (
    Literal as typing_extensions___Literal,
)


class CreateProjectRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    ProjectId = ... # type: typing___Text
    ProjectName = ... # type: typing___Text

    def __init__(self,
        *,
        ProjectId : typing___Optional[typing___Text] = None,
        ProjectName : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> CreateProjectRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"ProjectId",u"ProjectName"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"ProjectId",b"ProjectId",u"ProjectName",b"ProjectName"]) -> None: ...

class CreateProjectResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    def __init__(self,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> CreateProjectResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...

class OkResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    def __init__(self,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> OkResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...

class ScenarioRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    Name = ... # type: typing___Text
    timestamp = ... # type: int

    @property
    def Events(self) -> google___protobuf___internal___containers___RepeatedCompositeFieldContainer[events_pb2___Event]: ...

    def __init__(self,
        *,
        Name : typing___Optional[typing___Text] = None,
        Events : typing___Optional[typing___Iterable[events_pb2___Event]] = None,
        timestamp : typing___Optional[int] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> ScenarioRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"Events",u"Name",u"timestamp"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"Events",b"Events",u"Name",b"Name",u"timestamp",b"timestamp"]) -> None: ...

class Error(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    Message = ... # type: typing___Text

    def __init__(self,
        *,
        Message : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> Error: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"Message"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"Message",b"Message"]) -> None: ...

class ScenarioResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    @property
    def Error(self) -> Error: ...

    def __init__(self,
        *,
        Error : typing___Optional[Error] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> ScenarioResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"Error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"Error"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"Error",b"Error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"Error",b"Error"]) -> None: ...

class KillRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    def __init__(self,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> KillRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...

class PingRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    def __init__(self,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> PingRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
