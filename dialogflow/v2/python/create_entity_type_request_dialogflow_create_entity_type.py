# -*- coding: utf-8 -*-
#
# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# DO NOT EDIT! This is a generated sample ("Request",  "dialogflow_create_entity_type")

# To install the latest published package dependency, execute the following:
#   pip install google-cloud-dialogflow

import sys

# [START dialogflow_create_entity_type]

from google.cloud import dialogflow_v2
from google.cloud.dialogflow_v2 import enums


def sample_create_entity_type():
    """Create an EntityType"""

    # [START dialogflow_create_entity_type_core]

    client = dialogflow_v2.EntityTypesClient()

    parent = client.project_agent_path('[PROJECT]')
    kind = enums.EntityType.Kind.KIND_LIST
    entity_type = {'kind': kind}

    response = client.create_entity_type(parent, entity_type)
    print('Entity type created with display name: {}'.format(
        response.display_name))

    # [END dialogflow_create_entity_type_core]


# [END dialogflow_create_entity_type]


def main():
    import argparse

    parser = argparse.ArgumentParser()
    args = parser.parse_args()

    sample_create_entity_type()


if __name__ == '__main__':
    main()
