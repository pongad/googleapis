🐄 ` ` **G**.a.**P**.i.**C** samples` `

----

## Samples! 🦑

 - [ ] `<==`
   - [x] means that the sample is generated in this repo _*may not actually run yet_

### Dialogflow
   
  - **Detect Intents**
    - [ ] `dialogflow_detect_intent_text`
    - [ ] `dialogflow_detect_intent_audio`
    - [ ] `dialogflow_detect_intent_knowledge`

 - **Intent Management**
    - [ ] `dialogflow_create_intent`
    - [ ] `dialogflow_list_intents`
    - [ ] `dialogflow_delete_intent`

 - **Context Management**
    - [ ] `dialogflow_create_context`
    - [ ] `dialogflow_list_context`
    - [ ] `dialogflow_delete_context`

  - **Entity Management**
    - [ ] `dialogflow_create_entity`
    - [ ] `dialogflow_list_entities`
    - [ ] `dialogflow_delete_entity`

 - **Entity Type Management**
    - [x] **`dialogflow_create_entity_type`**
    - [x] **`dialogflow_list_entity_types`**
    - [x] **`dialogflow_delete_entity_type`**

  - **Session Entity Type Management**
    - [ ] `dialogflow_create_session_entity_type`
    - [ ] `dialogflow_list_session_entity_types`
    - [ ] `dialogflow_delete_session_entity_type`

  - **Document Management**
    - [ ] `dialogflow_create_document`
    - [ ] `dialogflow_list_document`
    - [ ] `dialogflow_delete_document`
    - [ ] `dialogflow_get_document`

  - **Knowledgebase Management**
    - [ ] `dialogflow_create_knowledge_base`
    - [ ] `dialogflow_list_knowledge_base`
    - [ ] `dialogflow_delete_knowledge_base`
    - [ ] `dialogflow_get_knowledge_base`

----

### Data Loss Prevention

 - [ ] `dlp_categorical_stats`
 - [ ] `dlp_create_template`
 - [ ] `dlp_create_trigger`
 - [ ] `dlp_deidentify_date_shift`
 - [ ] `dlp_deidentify_fpe`
 - [ ] `dlp_deidentify_masking`
 - [ ] `dlp_delete_job`
 - [ ] `dlp_delete_template`
 - [ ] `dlp_delete_trigger`
 - [ ] `dlp_inspect_bigquery`
 - [ ] `dlp_inspect_datastore`
 - [ ] `dlp_inspect_file`
 - [ ] `dlp_inspect_gcs`
 - [x] **`dlp_inspect_string`**
 - [ ] `dlp_k_anonymity`
 - [ ] `dlp_k_map`
 - [ ] `dlp_l_diversity`
 - [ ] `dlp_list_info_types`
 - [ ] `dlp_list_jobs`
 - [ ] `dlp_list_templates`
 - [ ] `dlp_list_triggers`
 - [ ] `dlp_numerical_stats`
 - [ ] `dlp_redact_image`
 - [ ] `dlp_reidentify_fpe`

----

### Job Discovery

 - [ ] sample
 - [ ] sample

----

### Key Management Service

 - **Encrypt / Decrypt**
   - [ ] `kms_encrypt`
   - [ ] `kms_decrypt`

 - **Encrypt / Decrypt (RSA)**
   - [ ] `kms_encrypt_rsa`
   - [ ] `kms_decrypt_rsa`

 - **Signing & Verifying**
   - [ ] `kms_sign_asymmetric`
   - [ ] `kms_verify_signature_rsa`
   - [ ] `kms_verify_signature_ec`

 - **Managing Keyrings**
   - [x] **`kms_create_keyring`**
   - [ ] `kms_get_keyring_policy`
   - [ ] `kms_add_member_to_keyring_policy`

 - **Managing Cryptokeys**
   - [ ] `kms_create_cryptokey`
   - [ ] `kms_create_cryptokey_version`
   - [ ] `kms_enable_cryptokey_version`
   - [ ] `kms_disable_cryptokey_version`
   - [ ] `kms_destroy_cryptokey_version`
   - [ ] `kms_destroy_cryptokey_version`
   - [ ] `kms_set_cryptokey_primary_version`
   - [ ] `kms_create_cryptokey_with_rotation_schedule`
   - [ ] `kms_add_member_to_cryptokey_policy`
   - [ ] `kms_remove_member_from_cryptokey_policy`

----

### Natural Language

  - **Analyzing Sentiment**
    - [x] **`language_sentiment_text`**
    - [x] **`language_sentiment_gcs`**

  - **Analyzing Entities**
    - [x] **`language_entities_text`**
    - [x] **language_entities_gcs`**

  - **Analyzing Syntax**
    - [x] **`language_syntax_text`**
    - [x] **`language_syntax_gcs`**

  - **Analyzing Entity Sentiment**
    - [ ] `language_entity_sentiment_text`
    - [ ] `language_entity_sentiment_gcs`

  - **Classifying Content**
    - [ ] `language_classify_text`
    - [ ] `language_classify_gcs`

----

### Natural Language – AutoML

  - **Prediction**
    - [ ] `automl_language_predict`

  - **Import/Export Data**
    - [ ] `automl_language_import_data`
    - [ ] `automl_language_export_data`

  - **Managing Datasets** 
    - [ ] `automl_language_create_dataset`
    - [ ] `automl_language_list_datasets`
    - [ ] `automl_language_delete_dataset`
    - [ ] `automl_language_get_dataset`

  - **Managing Models**
    - [ ] `automl_language_create_model`
    - [ ] `automl_language_list_models`
    - [ ] `automl_language_delete_model`
    - [ ] `automl_language_get_model`

    - **Model Evaluations**
      - [ ] `automl_language_list_model_evaluations`
      - [ ] `automl_language_list_model_evaluations`
      - [ ] `automl_language_display_evaluation`

----

### Speech-to-Text

 - **Perfoming Speech Recognition**
   - [ ] `speech_transcribe_sync`
   - [x] **`speech_transcribe_sync_gcs`**
   - [ ] `speech_transcribe_async`
   - [ ] `speech_transcribe_async_gcs`

 - **Configuring Recognition Requests**
   - _Enabling Word-Level Confidence_
       - [ ] `speech_transcribe_word_level_confidence_beta`
   - _Separating Different Speakers_
     - [ ] `speech_transcribe_diarization_beta`
     - [ ] `speech_transcribe_diarization_gcs_beta`
   - _Detecting Language Spoken Automatically_
     - [ ] `speech_transcribe_multilanguage_beta`
     - [ ] `speech_transcribe_multilanguage_gcs_beta`
   - _Transcribing Audio with Multiple Channels_
     - [ ] `speech_transcribe_multilanguage_beta`
     - [ ] `speech_transcribe_multilanguage_gcs_beta`
   - _Getting Punctuation_
     - [ ] `speech_transcribe_auto_punctuation`
   - _Getting Word Timestamps_
     - [ ] `speech_transcribe_async_word_time_offsets_gcs`
   - _Adding Recognition Metadata_
     - [ ] `speech_transcribe_recognition_metadata_beta`
   - _Using Enhanced Models_
     - [ ] `speech_transcribe_model_selection`
     - [ ] `speech_transcribe_model_selection_beta`
     - [ ] `speech_transcribe_model_selection_gcs_beta`
     - [ ] `speech_transcribe_enhanced_model_beta`

----

### Text-to-Speech

 - [ ] `tts_list_voices`
 - [ ] `tts_synthesize_text`
 - [ ] `tts_synthesize_text_file`
 - [ ] `tts_synthesize_ssml`
 - [ ] `tts_synthesize_ssml_file`

----

### Translate – AutoML

  - **Prediction**
    - [ ] `automl_translate_predict`

  - **Import Data**
    - [ ] `automl_translate_import_data`

  - **Managing Datasets** 
    - [ ] `automl_translate_create_dataset`
    - [ ] `automl_translate_list_datasets`
    - [ ] `automl_translate_delete_dataset`
    - [ ] `automl_translate_get_dataset`

  - **Managing Models**
    - [ ] `automl_translate_create_model`
    - [ ] `automl_translate_list_models`
    - [ ] `automl_translate_delete_model`
    - [ ] `automl_translate_get_model`

    - **Model Evaluations**
      - [ ] `automl_translate_list_model_evaluations`
      - [ ] `automl_translate_list_model_evaluations`

----

### Video Intelligence

 - [ ] `video_analyze_labels`
 - [ ] `video_analyze_labels_gcs`
 - [ ] `video_analyze_shots`
 - [ ] `video_analyze_explicit_content`
 - [ ] `video_speech_transcription_gcs`

----

### Vision

 - [ ] sample
 - [ ] sample

----

### Vision – AutoML

  - **Prediction**
    - [ ] `automl_vision_predict`

  - **Import/Export Data**
    - [ ] `automl_vision_import_data`
    - [ ] `automl_vision_export_data`

  - **Managing Datasets** 
    - [ ] `automl_vision_create_dataset`
    - [ ] `automl_vision_list_datasets`
    - [ ] `automl_vision_delete_dataset`
    - [ ] `automl_vision_get_dataset`

  - **Managing Models**
    - [ ] `automl_vision_create_model`
    - [ ] `automl_vision_list_models`
    - [ ] `automl_vision_delete_model`
    - [ ] `automl_vision_get_model`

    - **Model Evaluations**
      - [ ] `automl_vision_list_model_evaluations`
      - [ ] `automl_vision_list_model_evaluations`
      - [ ] `automl_vision_display_evaluation`

----

### Vision – Product Search

 - **Managing Products**
   - [ ] `create_product`
   - [ ] `get_product`
   - [ ] `list_products`
   - [ ] `delete_product`
   - [ ] `update_product_labels`

 - **Managing Product Sets**
   - [ ] `create_product_set`
   - [ ] `get_product_set`
   - [ ] `list_product_sets`
   - [ ] `delete_product_set`

 - **Managing Products in Product Sets**
   - [ ] `add_product_to_product_set`
   - [ ] `list_products_in_product_set`
   - [ ] `remove_product_from_product_set`

 - **Managing Reference Images**
   - [ ] `create_reference_image`
   - [ ] `get_reference_image`
   - [ ] `list_reference_images`
   - [ ] `delete_reference_image`

 - **Similar Products**
   - [ ] `get_similar_products_file`
   - [ ] `get_similar_products_uri`

----

#### `$ generate-samples`

```php
Usage:

  generate-samples ALL
  generate-samples ALL python
  generate-samples ALL python csharp
  generate-samples speech
  generate-samples speech dlp.v2
  generate-samples dialogflow
  generate-samples dialogflow.v1
  generate-samples dialogflow.v1 python
  generate-samples dialogflow.v1 python csharp
  generate-samples --help      # prints this message :)
  generate-samples --parallel  # parallel generation [TODO] 🦇
```
