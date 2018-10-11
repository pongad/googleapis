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

# DO NOT EDIT! This is a generated sample ("Request",  "automl_vision_create_dataset")

# To install the latest published package dependency, execute the following:
#   pip install google-cloud-automl

import sys

# [START automl_vision_create_dataset]

from google.cloud import automl_v1beta1
import six


def sample_create_dataset(display_name):
    """Transcribe audio file"""

    # [START automl_vision_create_dataset_core]

    client = automl_v1beta1.AutoMlClient()

    # display_name = 'Your Dataset display name'

    if isinstance(display_name, six.binary_type):
        display_name = display_name.decode('utf-8')
    parent = client.location_path('[PROJECT]', '[LOCATION]')
    dataset = {'display_name': display_name}

    response = client.create_dataset(parent, dataset)
    dataset = response
    print('Dataset name: {}'.format(dataset.name))
    print('Dataset display name: {}'.format(dataset.display_name))
    print('Image classification dataset metadata:')
    print('	{}'.format(dataset.image_classification_dataset_metadata))
    print('Dataset example count: {}'.format(dataset.example_count))
    print('Dataset create time:')
    print('	seconds: {}'.format(dataset.create_time.seconds))
    print('	nanos: {}'.format(dataset.create_time.nanos))

    # [END automl_vision_create_dataset_core]


# [END automl_vision_create_dataset]


def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--display_name', type=str, default='Your Dataset display name')
    args = parser.parse_args()

    sample_create_dataset(args.display_name)


if __name__ == '__main__':
    main()
