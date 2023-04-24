#!/usr/bin/env -S python3 -u

"""
Run version compatibility tests
"""
import os
from get_latest_release_versions import update_helm_repo, get_latest_release_versions
from compatibility_test import make_compatibility_test_runner
from clusters import GKECluster

# set required test parameters
os.environ["ORCHESTRATOR_FLAVOR"] = "k8s"
os.environ["ROX_POSTGRES_DATASTORE"] = "true"

update_helm_repo()
chart_versions=get_latest_release_versions(4)

gkecluster=GKECluster("compat-test")

failing_sensor_versions = []
for version in chart_versions:
    os.environ["SENSOR_CHART_VERSION"] = version
    try:
        make_compatibility_test_runner(cluster=gkecluster).run()
    except Exception:
        print(f"Exception \"{Exception}\" raised in compatibility test for sensor version {version} and central version {chart_versions[0]}")
        failing_sensor_versions += version

if len(chart_versions) > 1:
    os.environ["SENSOR_CHART_VERSION"] = chart_versions[0]
    os.environ["CENTRAL_CHART_VERSION"] = chart_versions[1]
    try:
        make_compatibility_test_runner(cluster=gkecluster).run()
    except Exception:
        print(f"Exception \"{Exception}\" raised in compatibility test for sensor version {chart_versions[0]} and central version {chart_versions[1]}")
        failing_sensor_versions += version


if len(failing_sensor_versions) > 0:
    raise SensorVersionsFailure(f"Compatibility tests failed for Sensor versions " + ', '.join(failing_sensor_versions) + " and Central versions " + ', '.join(failing_sensor_versions[:2]))

class SensorVersionsFailure(Exception):
    pass
