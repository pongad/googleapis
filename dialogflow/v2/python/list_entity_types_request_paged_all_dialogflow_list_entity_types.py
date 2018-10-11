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

# DO NOT EDIT! This is a generated sample ("RequestPagedAll",  "dialogflow_list_entity_types")

# To install the latest published package dependency, execute the following:
#   pip install google-cloud-dialogflow

import sys

# [START dialogflow_list_entity_types]

from google.cloud import dialogflow_v2


def sample_list_entity_types():
    """List Entity Types"""

    # [START dialogflow_list_entity_types_core]

    client = dialogflow_v2.EntityTypesClient()

    parent = client.project_agent_path('[PROJECT]')

    # Iterate over all results
    for response_item in client.list_entity_types(parent):
        print('Entity type name: {}'.format(response_item.name))
        print('Entity type display name: {}'.format(
            response_item.display_name))
        print('Entities:')
        for entity in response_item.entities:
            print('Entity value:'.format(entity.value))

    # [END dialogflow_list_entity_types_core]


# [END dialogflow_list_entity_types]


def main():
    import argparse

    parser = argparse.ArgumentParser()
    args = parser.parse_args()

    sample_list_entity_types()


if __name__ == '__main__':
    main()
