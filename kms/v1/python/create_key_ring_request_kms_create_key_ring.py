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

# DO NOT EDIT! This is a generated sample ("Request",  "kms_create_key_ring")

# To install the latest published package dependency, execute the following:
#   pip install google-cloud-kms

import sys

# [START kms_create_key_ring]

from google.cloud import kms_v1
import six


def sample_create_key_ring(key_ring_id):
    """Create a KeyRing"""

    # [START kms_create_key_ring_core]

    client = kms_v1.KeyManagementServiceClient()

    # # TODO: Initialize ``key_ring_id``:
    # key_ring_id = ''

    if isinstance(key_ring_id, six.binary_type):
        key_ring_id = key_ring_id.decode('utf-8')
    parent = client.location_path('[PROJECT]', '[LOCATION]')

    # TODO: Initialize ``key_ring``:
    key_ring = {}

    response = client.create_key_ring(parent, key_ring_id, key_ring)
    print('Name: {}'.format(response.name))

    # [END kms_create_key_ring_core]


# [END kms_create_key_ring]


def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument('--key_ring_id', type=str, default='')
    args = parser.parse_args()

    sample_create_key_ring(args.key_ring_id)


if __name__ == '__main__':
    main()
