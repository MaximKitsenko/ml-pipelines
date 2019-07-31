
from django import urls

from explore.specs import when, preset
from test import *
from utils import pretty

from proto import events_pb2 as evt

def given_a_dataset_with_storage_location(t: test.Env):
    """show storage location on all related screens"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.metadata.storage_location = "url://aws"
    ds.metadata.set_fields.append(evt.FIELD_STORAGE_LOCATION)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .storage-loc', "url://aws"),
    )


def given_a_dataset_without_storage_location(t: test.Env):
    """show N/A instead"""
    prj = preset.project_created(t)
    ds = preset.dataset_created(t, prj)

    ds.metadata.del_fields.append(evt.FIELD_STORAGE_LOCATION)

    t.given_events(prj, ds)

    t.scenario(
        when.view_dataset(ds.dataset_id),
        then.text(f'main #ds-{ds.dataset_id} .storage-loc', "N/A"),
    )