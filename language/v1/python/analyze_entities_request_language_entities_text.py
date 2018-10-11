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

# DO NOT EDIT! This is a generated sample ("Request",  "language_entities_text")

# To install the latest published package dependency, execute the following:
#   pip install google-cloud-language

import sys

# [START language_entities_text]

from google.cloud import language_v1
from google.cloud.language_v1 import enums
import six


def sample_analyze_entities(content):
    """Analyze entities in text"""

    # [START language_entities_text_core]

    client = language_v1.LanguageServiceClient()

    # content = 'Your text to analyze, e.g. Hello, world!'

    if isinstance(content, six.binary_type):
        content = content.decode('utf-8')
    type_ = enums.Document.Type.PLAIN_TEXT
    document = {'type': type_, 'content': content}

    response = client.analyze_entities(document)
    entities = response.entities
    for entity in entities:
        print('Entity name: {}'.format(entity.name))
        print('Entity type: {}'.format(entity.type))

    # [END language_entities_text_core]


# [END language_entities_text]


def main():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--content',
        type=str,
        default='Your text to analyze, e.g. Hello, world!')
    args = parser.parse_args()

    sample_analyze_entities(args.content)


if __name__ == '__main__':
    main()
