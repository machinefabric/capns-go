# CapDag-Go Test Catalog

**Total Tests:** 654

**Numbered Tests:** 586

**Unnumbered Tests:** 68

**Numbered Tests Missing Descriptions:** 0

**Numbering Mismatches:** 0

All numbered test numbers are unique.

This catalog lists all tests in the CapDag-Go codebase.

| Test # | Function Name | Description | File |
|--------|---------------|-------------|------|
| test001 | `Test001_cap_urn_creation` | TEST001: Test that cap URN is created with tags parsed correctly and direction specs accessible | urn/cap_urn_test.go:33 |
| test002 | `Test002_direction_specs_default_to_wildcard` | TEST002: Test that missing 'in' or 'out' defaults to media: wildcard | urn/cap_urn_test.go:57 |
| test003 | `Test003_direction_matching` | TEST003: Test that direction specs must match exactly, different in/out types don't match, wildcard matches any | urn/cap_urn_test.go:76 |
| test004 | `Test004_unquoted_values_lowercased` | TEST004: Test that unquoted keys and values are normalized to lowercase | urn/cap_urn_test.go:98 |
| test005 | `Test005_quoted_values_preserve_case` | TEST005: Test that quoted values preserve case while unquoted are lowercased | urn/cap_urn_test.go:120 |
| test006 | `Test006_quoted_value_special_chars` | TEST006: Test that quoted values can contain special characters (semicolons, equals, spaces) | urn/cap_urn_test.go:146 |
| test007 | `Test007_quoted_value_escape_sequences` | TEST007: Test that escape sequences in quoted values (\" and \\) are parsed correctly | urn/cap_urn_test.go:167 |
| test008 | `Test008_mixed_quoted_unquoted` | TEST008: Test that mixed quoted and unquoted values in same URN parse correctly | urn/cap_urn_test.go:188 |
| test009 | `Test009_unterminated_quote_error` | TEST009: Test that unterminated quote produces UnterminatedQuote error | urn/cap_urn_test.go:202 |
| test010 | `Test010_invalid_escape_sequence_error` | TEST010: Test that invalid escape sequences (like \n, \x) produce InvalidEscapeSequence error | urn/cap_urn_test.go:212 |
| test011 | `Test011_serialization_smart_quoting` | TEST011: Test that serialization uses smart quoting (no quotes for simple lowercase, quotes for special chars/uppercase) | urn/cap_urn_test.go:229 |
| test012 | `Test012_round_trip_simple` | TEST012: Test that simple cap URN round-trips (parse -> serialize -> parse equals original) | urn/cap_urn_test.go:260 |
| test013 | `Test013_round_trip_quoted` | TEST013: Test that quoted values round-trip preserving case and spaces | urn/cap_urn_test.go:271 |
| test014 | `Test014_round_trip_escapes` | TEST014: Test that escape sequences round-trip correctly | urn/cap_urn_test.go:285 |
| test015 | `Test015_cap_prefix_required` | TEST015: Test that cap: prefix is required and case-insensitive | urn/cap_urn_test.go:299 |
| test016 | `Test016_trailing_semicolon_equivalence` | TEST016: Test that trailing semicolon is equivalent (same hash, same string, matches) | urn/cap_urn_test.go:320 |
| test017 | `Test017_tag_matching` | TEST017: Test tag matching: exact match, subset match, wildcard match, value mismatch | urn/cap_urn_test.go:334 |
| test018 | `Test018_matching_case_sensitive_values` | TEST018: Test that quoted values with different case do NOT match (case-sensitive) | urn/cap_urn_test.go:364 |
| test019 | `Test019_missing_tag_handling` | TEST019: Missing tag in instance causes rejection — pattern's tags are constraints | urn/cap_urn_test.go:378 |
| test020 | `Test020_specificity` | TEST020: Test specificity calculation (direction specs use MediaUrn tag count, wildcards don't count) | urn/cap_urn_test.go:407 |
| test021 | `Test021_builder` | TEST021: Test builder creates cap URN with correct tags and direction specs | urn/cap_urn_test.go:431 |
| test022 | `Test022_builder_requires_direction` | TEST022: Test builder requires both in_spec and out_spec | urn/cap_urn_test.go:450 |
| test023 | `Test023_builder_preserves_case` | TEST023: Test builder lowercases keys but preserves value case | urn/cap_urn_test.go:471 |
| test024 | `Test024_directional_accepts` | TEST024: Directional accepts — pattern's tags are constraints, instance must satisfy | urn/cap_urn_test.go:485 |
| test025 | `Test025_best_match` | TEST025: Test find_best_match returns most specific matching cap | urn/cap_urn_test.go:515 |
| test026 | `Test026_merge_and_subset` | TEST026: Test merge combines tags from both caps, subset keeps only specified tags | urn/cap_urn_test.go:544 |
| test027 | `Test027_wildcard_tag` | TEST027: Test with_wildcard_tag sets tag to wildcard, including in/out | urn/cap_urn_test.go:577 |
| test028 | `Test028_empty_cap_urn_defaults_to_wildcard` | TEST028: Test empty cap URN defaults to media: wildcard | urn/cap_urn_test.go:594 |
| test029 | `Test029_minimal_cap_urn` | TEST029: Test minimal valid cap URN has just in and out, empty tags | urn/cap_urn_test.go:612 |
| test030 | `Test030_extended_character_support` | TEST030: Test extended characters (forward slashes, colons) in tag values | urn/cap_urn_test.go:621 |
| test031 | `Test031_wildcard_restrictions` | TEST031: Test wildcard rejected in keys but accepted in values | urn/cap_urn_test.go:636 |
| test032 | `Test032_duplicate_key_rejection` | TEST032: Test duplicate keys are rejected with DuplicateKey error | urn/cap_urn_test.go:654 |
| test033 | `Test033_numeric_key_restriction` | TEST033: Test pure numeric keys rejected, mixed alphanumeric allowed, numeric values allowed | urn/cap_urn_test.go:664 |
| test034 | `Test034_empty_value_error` | TEST034: Test empty values are rejected | urn/cap_urn_test.go:690 |
| test035 | `Test035_has_tag_case_sensitive` | TEST035: Test has_tag is case-sensitive for values, case-insensitive for keys, works for in/out | urn/cap_urn_test.go:701 |
| test036 | `Test036_with_tag_preserves_value` | TEST036: Test with_tag preserves value case | urn/cap_urn_test.go:715 |
| test037 | `Test037_with_tag_rejects_empty_value` | TEST037: Test with_tag rejects empty value | urn/cap_urn_test.go:725 |
| test038 | `Test038_semantic_equivalence` | TEST038: Test semantic equivalence of unquoted and quoted simple lowercase values | urn/cap_urn_test.go:733 |
| test039 | `Test039_get_tag_returns_direction_specs` | TEST039: Test get_tag returns direction specs (in/out) with case-insensitive lookup | urn/cap_urn_test.go:742 |
| test040 | `Test040_matching_semantics_exact_match` | TEST040: Matching semantics - exact match succeeds | urn/cap_urn_test.go:768 |
| test041 | `Test041_matching_semantics_cap_missing_tag` | TEST041: Matching semantics - cap missing tag matches (implicit wildcard) | urn/cap_urn_test.go:779 |
| test042 | `Test042_matching_semantics_cap_has_extra_tag` | TEST042: Pattern rejects instance missing required tags | urn/cap_urn_test.go:796 |
| test043 | `Test043_matching_semantics_request_has_wildcard` | TEST043: Matching semantics - request wildcard matches specific cap value | urn/cap_urn_test.go:813 |
| test044 | `Test044_matching_semantics_cap_has_wildcard` | TEST044: Matching semantics - cap wildcard matches specific request value | urn/cap_urn_test.go:824 |
| test045 | `Test045_matching_semantics_value_mismatch` | TEST045: Matching semantics - value mismatch does not match | urn/cap_urn_test.go:835 |
| test046 | `Test046_matching_semantics_fallback_pattern` | TEST046: Matching semantics - fallback pattern (cap missing tag = implicit wildcard) | urn/cap_urn_test.go:846 |
| test047 | `Test047_matching_semantics_thumbnail_void_input` | TEST047: Matching semantics - thumbnail fallback with void input | urn/cap_urn_test.go:862 |
| test048 | `Test048_matching_semantics_wildcard_direction_matches_anything` | TEST048: Matching semantics - wildcard direction matches anything | urn/cap_urn_test.go:879 |
| test049 | `Test049_matching_semantics_cross_dimension_independence` | TEST049: Non-overlapping tags — neither direction accepts | urn/cap_urn_test.go:895 |
| test050 | `Test050_matching_semantics_direction_mismatch` | TEST050: Matching semantics - direction mismatch prevents matching | urn/cap_urn_test.go:912 |
| test051 | `Test051_input_validation_success` | TEST051: Test input validation succeeds with valid positional argument | cap/validation_test.go:50 |
| test052 | `Test052_input_validation_missing_required` | TEST052: Test input validation fails with MissingRequiredArgument when required arg missing | cap/validation_test.go:67 |
| test053 | `Test053_input_validation_wrong_type` | TEST053: Test input validation fails with InvalidArgumentType when wrong type provided | cap/validation_test.go:85 |
| test054 | `Test054_xv5_inline_spec_redefinition_detected` | TEST054: XV5 - Test inline media spec redefinition of existing registry spec is detected and rejected | cap/schema_validation_test.go:702 |
| test055 | `Test055_xv5_new_inline_spec_allowed` | TEST055: XV5 - Test new inline media spec (not in registry) is allowed | cap/schema_validation_test.go:724 |
| test056 | `Test056_xv5_empty_media_specs_allowed` | TEST056: XV5 - Test empty media_specs (no inline specs) passes XV5 validation | cap/schema_validation_test.go:745 |
| test060 | `Test060_wrong_prefix_fails` | TEST060: Test wrong prefix fails with InvalidPrefix error showing expected and actual prefix | urn/media_urn_test.go:39 |
| test061 | `Test061_is_binary` | TEST061: Test is_binary returns true when textable tag is absent (binary = not textable) | urn/media_urn_test.go:50 |
| test062 | `Test062_is_record` | TEST062: Test is_record returns true when record marker tag is present indicating key-value structure | urn/media_urn_test.go:87 |
| test063 | `Test063_is_scalar` | TEST063: Test is_scalar returns true when list marker tag is absent (scalar is default) | urn/media_urn_test.go:116 |
| test064 | `Test064_is_list` | TEST064: Test is_list returns true when list marker tag is present indicating ordered collection | urn/media_urn_test.go:137 |
| test065 | `Test065_is_opaque` | TEST065: Test is_opaque returns true when record marker is absent (opaque is default) | urn/media_urn_test.go:152 |
| test066 | `Test066_is_json` | TEST066: Test is_json returns true only when json marker tag is present for JSON representation | urn/media_urn_test.go:183 |
| test067 | `Test067_is_text` | TEST067: Test is_text returns true only when textable marker tag is present | urn/media_urn_test.go:198 |
| test068 | `Test068_is_void` | TEST068: Test is_void returns true when void flag or type=void tag is present | urn/media_urn_test.go:213 |
| test071 | `Test071_to_string_roundtrip` | TEST071: Test to_string roundtrip ensures serialization and deserialization preserve URN structure | urn/media_urn_test.go:241 |
| test072 | `Test072_constants_parse` | TEST072: Test all media URN constants parse successfully as valid media URNs | urn/media_urn_test.go:254 |
| test073 | `Test073_extension_helpers` | TEST073: Test extension helper functions create media URNs with ext tag and correct format | urn/media_urn_test.go:296 |
| test074 | `Test074_media_urn_matching` | TEST074: Test media URN conforms_to using tagged URN semantics with specific and generic requirements | urn/media_urn_test.go:305 |
| test075 | `Test075_matching` | TEST075: Test accepts with implicit wildcards where handlers with fewer tags can handle more requests | urn/media_urn_test.go:326 |
| test076 | `Test076_specificity` | TEST076: Test specificity increases with more tags for ranking conformance | urn/media_urn_test.go:338 |
| test077 | `Test077_serde_roundtrip` | TEST077: Test serde roundtrip serializes to JSON string and deserializes back correctly | urn/media_urn_test.go:350 |
| test078 | `Test078_object_does_not_conform_to_string` | TEST078: conforms_to behavior between MEDIA_OBJECT and MEDIA_STRING | urn/media_urn_test.go:367 |
| test088 | `Test088_resolve_from_registry_str` | TEST088: Test resolving string media URN from registry returns correct media type and profile | media/spec_test.go:26 |
| test089 | `Test089_resolve_from_registry_obj` | TEST089: Test resolving JSON media URN from registry returns JSON media type | media/spec_test.go:35 |
| test090 | `Test090_resolve_from_registry_binary` | TEST090: Test resolving binary media URN returns octet-stream and is_binary true | media/spec_test.go:43 |
| test091 | `Test091_resolve_custom_media_spec` | TEST091: Test resolving custom media URN from local media_specs takes precedence over registry | media/spec_test.go:52 |
| test092 | `Test092_resolve_custom_with_schema` | TEST092: Test resolving custom record media spec with schema from local media_specs | media/spec_test.go:78 |
| test093 | `Test093_resolve_unresolvable_fails_hard` | TEST093: Test resolving unknown media URN fails with UnresolvableMediaUrn error | media/spec_test.go:109 |
| test094 | `Test094_local_overrides_registry` | TEST094: Test local media_specs definition overrides registry definition for same URN | media/spec_test.go:119 |
| test095 | `Test095_media_spec_def_serialize` | TEST095: Test MediaSpecDef serializes with required fields and skips None fields | media/spec_test.go:149 |
| test096 | `Test096_media_spec_def_deserialize` | TEST096: Test deserializing MediaSpecDef from JSON object | media/spec_test.go:175 |
| test097 | `Test097_validate_no_duplicate_urns_catches_duplicates` | TEST097: Test duplicate URN validation catches duplicates | media/spec_test.go:191 |
| test098 | `Test098_validate_no_duplicate_urns_passes_for_unique` | TEST098: Test duplicate URN validation passes for unique URNs | media/spec_test.go:203 |
| test099 | `Test099_resolved_is_binary` | TEST099: Test ResolvedMediaSpec is_binary returns true when textable tag is absent | media/spec_test.go:217 |
| test100 | `Test100_resolved_is_map` | TEST100: Test ResolvedMediaSpec is_record returns true when record marker is present | media/spec_test.go:235 |
| test101 | `Test101_resolved_is_scalar` | TEST101: Test ResolvedMediaSpec is_scalar returns true when list marker is absent | media/spec_test.go:255 |
| test102 | `Test102_resolved_is_list` | TEST102: Test ResolvedMediaSpec is_list returns true when list marker is present | media/spec_test.go:273 |
| test103 | `Test103_resolved_is_json` | TEST103: Test ResolvedMediaSpec is_json returns true when json tag is present | media/spec_test.go:291 |
| test104 | `Test104_resolved_is_text` | TEST104: Test ResolvedMediaSpec is_text returns true when textable tag is present | media/spec_test.go:309 |
| test105 | `Test105_metadata_propagation` | TEST105: Test metadata propagates from media spec def to resolved media spec | media/spec_test.go:331 |
| test106 | `Test106_metadata_with_validation` | TEST106: Test metadata and validation can coexist in media spec definition | media/spec_test.go:358 |
| test107 | `Test107_extensions_propagation` | TEST107: Test extensions field propagates from media spec def to resolved | media/spec_test.go:400 |
| test108 | `Test108_cap_creation` | TEST108: Test creating new cap with URN, title, and command verifies correct initialization | cap/definition_test.go:23 |
| test109 | `Test109_cap_with_metadata` | TEST109: Test creating cap with metadata initializes and retrieves metadata correctly | cap/definition_test.go:42 |
| test110 | `Test110_cap_matching` | TEST110: Test cap matching with subset semantics for request fulfillment | cap/definition_test.go:68 |
| test111 | `Test111_cap_title` | TEST111: Test getting and setting cap title updates correctly | cap/definition_test.go:82 |
| test112 | `Test112_cap_definition_equality` | TEST112: Test cap equality based on URN and title matching | cap/definition_test.go:97 |
| test113 | `Test113_cap_stdin` | TEST113: Test cap stdin support via args with stdin source and serialization roundtrip | cap/definition_test.go:113 |
| test114 | `Test114_arg_source_types` | TEST114: Test ArgSource type variants stdin, position, and cli_flag with their accessors | cap/definition_test.go:150 |
| test115 | `Test115_cap_arg_serialization` | TEST115: Test CapArg serialization and deserialization with multiple sources | cap/definition_test.go:180 |
| test116 | `Test116_cap_arg_constructors` | TEST116: Test CapArg constructor methods basic and with_description create args correctly | cap/definition_test.go:206 |
| test117 | `Test117_register_and_find_cap_set` | TEST117: Test registering cap set and finding by exact and subset matching | cap_matrix_test.go:32 |
| test118 | `Test118_best_cap_set_selection` | TEST118: Test selecting best cap set based on specificity ranking With is_dispatchable semantics: - Provider must satisfy ALL request constraints - General request matches specific provider (provider refines request) - Specific request does NOT match general provider (provider lacks constraints) | cap_matrix_test.go:92 |
| test119 | `Test119_invalid_urn_handling` | TEST119: Test invalid URN returns InvalidUrn error | cap_matrix_test.go:149 |
| test120 | `Test120_accepts_request` | TEST120: Test accepts_request checks if registry can handle a capability request | cap_matrix_test.go:166 |
| test121 | `Test121_cap_block_more_specific_wins` | TEST121: Test CapBlock selects more specific cap over less specific regardless of registry order | cap_matrix_test.go:225 |
| test122 | `Test122_cap_block_tie_goes_to_first` | TEST122: Test CapBlock breaks specificity ties by first registered registry | cap_matrix_test.go:275 |
| test123 | `Test123_cap_block_polls_all` | TEST123: Test CapBlock polls all registries to find most specific match | cap_matrix_test.go:309 |
| test124 | `Test124_cap_block_no_match` | TEST124: Test CapBlock returns error when no registries match the request | cap_matrix_test.go:348 |
| test125 | `Test125_cap_block_fallback_scenario` | TEST125: Test CapBlock prefers specific cartridge over generic provider fallback | cap_matrix_test.go:368 |
| test126 | `Test126_cap_block_can_method` | TEST126: Test composite can method returns CapCaller for capability execution | cap_matrix_test.go:436 |
| test127 | `Test127_cap_graph_basic_construction` | TEST127: Test CapGraph adds nodes and edges from capability definitions | cap_matrix_test.go:526 |
| test128 | `Test128_cap_graph_outgoing_incoming` | TEST128: Test CapGraph tracks outgoing and incoming edges for spec conversions | cap_matrix_test.go:566 |
| test129 | `Test129_cap_graph_can_convert` | TEST129: Test CapGraph detects direct and indirect conversion paths between specs | cap_matrix_test.go:602 |
| test130 | `Test130_cap_graph_find_path` | TEST130: Test CapGraph finds shortest path for spec conversion chain | cap_matrix_test.go:646 |
| test131 | `Test131_cap_graph_find_all_paths` | TEST131: Test CapGraph finds all conversion paths sorted by length | cap_matrix_test.go:703 |
| test132 | `Test132_cap_graph_get_direct_edges` | TEST132: Test CapGraph returns direct edges sorted by specificity | cap_matrix_test.go:739 |
| test133 | `Test133_cap_graph_with_cap_block` | TEST133: Test CapBlock graph integration with multiple registries and conversion paths | cap_matrix_test.go:828 |
| test134 | `Test134_cap_graph_stats` | TEST134: Test CapGraph stats provides counts of nodes and edges | cap_matrix_test.go:787 |
| test135 | `Test135_registry_creation` | TEST135: Test registry creation with temporary cache directory succeeds | cap/registry_test.go:25 |
| test136 | `Test136_cache_key_generation` | TEST136: Test cache key generation produces consistent hashes for same URN | cap/registry_test.go:32 |
| test137 | `Test137_parse_registry_json` | TEST137: Test parsing registry JSON without stdin args verifies cap structure | cap/registry_test.go:86 |
| test138 | `Test138_parse_registry_json_with_stdin` | TEST138: Test parsing registry JSON with stdin args verifies stdin media URN extraction | cap/registry_test.go:102 |
| test139 | `Test139_url_keeps_cap_prefix_literal` | TEST139: Test URL construction keeps cap prefix literal and only encodes tags part | cap/registry_test.go:142 |
| test140 | `Test140_url_encodes_media_urns` | TEST140: Test URL encodes media URNs with proper percent encoding for special characters | cap/registry_test.go:153 |
| test141 | `Test141_url_format_is_valid` | TEST141: Test exact URL format contains properly encoded media URN components | cap/registry_test.go:165 |
| test142 | `Test142_normalize_handles_different_tag_orders` | TEST142: Test normalize handles different tag orders producing same canonical form | cap/registry_test.go:182 |
| test143 | `Test143_default_config` | TEST143: Test default config uses capdag.com or environment variable values | cap/registry_test.go:193 |
| test144 | `Test144_custom_registry_url` | TEST144: Test custom registry URL updates both registry and schema base URLs | cap/registry_test.go:206 |
| test145 | `Test145_custom_registry_and_schema_url` | TEST145: Test custom registry and schema URLs set independently | cap/registry_test.go:214 |
| test146 | `Test146_schema_url_not_overwritten_when_explicit` | TEST146: Test schema URL not overwritten when set explicitly before registry URL | cap/registry_test.go:223 |
| test147 | `Test147_registry_for_test_with_config` | TEST147: Test registry for test with custom config creates registry with specified URLs | cap/registry_test.go:233 |
| test148 | `Test148_cap_manifest_creation` | TEST148: Test creating cap manifest with name, version, description, and caps | bifaci/manifest_test.go:23 |
| test149 | `Test149_cap_manifest_with_author` | TEST149: Test cap manifest with author field sets author correctly | bifaci/manifest_test.go:44 |
| test150 | `Test150_cap_manifest_json_serialization` | TEST150: Test cap manifest JSON serialization and deserialization roundtrip | bifaci/manifest_test.go:85 |
| test151 | `Test151_cap_manifest_required_fields` | TEST151: Test cap manifest deserialization fails when required fields are missing | bifaci/manifest_test.go:129 |
| test152 | `Test152_cap_manifest_with_multiple_caps` | TEST152: Test cap manifest with multiple caps stores and retrieves all capabilities | bifaci/manifest_test.go:150 |
| test153 | `Test153_cap_manifest_empty_caps` | TEST153: Test cap manifest with empty caps list serializes and deserializes correctly | bifaci/manifest_test.go:177 |
| test154 | `Test154_cap_manifest_optional_fields` | TEST154: Test cap manifest optional author field skipped in serialization when None | bifaci/manifest_test.go:198 |
| test155 | `Test155_component_metadata_interface` | TEST155: Test ComponentMetadata trait provides manifest and caps accessor methods | bifaci/manifest_test.go:247 |
| test156 | `Test156_stdin_source_data_creation` | TEST156: Test creating StdinSource Data variant with byte vector | cap/caller_test.go:231 |
| test157 | `Test157_stdin_source_file_reference_creation` | TEST157: Test creating StdinSource FileReference variant with all required fields | cap/caller_test.go:242 |
| test158 | `Test158_stdin_source_empty_data` | TEST158: Test StdinSource Data with empty vector stores and retrieves correctly | cap/caller_test.go:265 |
| test159 | `Test159_stdin_source_binary_content` | TEST159: Test StdinSource Data with binary content like PNG header bytes | cap/caller_test.go:274 |
| test160 | `Test160_stdin_source_data_clone` | TEST160: Test StdinSource Data clone creates independent copy with same data | cap/caller_test.go:287 |
| test161 | `Test161_stdin_source_file_reference_clone` | TEST161: Test StdinSource FileReference clone creates independent copy with same fields | cap/caller_test.go:305 |
| test162 | `Test162_stdin_source_debug` | TEST162: Test StdinSource Debug format displays variant type and relevant fields | cap/caller_test.go:324 |
| test163 | `Test163_schema_validator_validate_argument_with_schema_success` | TEST163: Test argument schema validation succeeds with valid JSON matching schema | cap/schema_validation_test.go:32 |
| test164 | `Test164_schema_validator_validate_argument_with_schema_failure` | TEST164: Test argument schema validation fails with JSON missing required fields | cap/schema_validation_test.go:71 |
| test165 | `Test165_schema_validator_validate_output_with_schema_success` | TEST165: Test output schema validation succeeds with valid JSON matching schema | cap/schema_validation_test.go:134 |
| test166 | `Test166_schema_validator_skip_validation_without_schema` | TEST166: Test validation skipped when resolved media spec has no schema | cap/schema_validation_test.go:765 |
| test167 | `Test167_schema_validator_unresolvable_media_urn_fails_hard` | TEST167: Test validation fails hard when media URN cannot be resolved from any source | cap/schema_validation_test.go:790 |
| test168 | `Test168_response_wrapper_from_json` | TEST168: Test ResponseWrapper from JSON deserializes to correct structured type | cap/response_test.go:22 |
| test169 | `Test169_response_wrapper_as_int` | TEST169: Test ResponseWrapper converts to primitive types integer, float, boolean, string | cap/response_test.go:75 |
| test170 | `Test170_response_wrapper_from_binary` | TEST170: Test ResponseWrapper from binary stores and retrieves raw bytes correctly | cap/response_test.go:58 |
| test171 | `Test171_frame_type_roundtrip` | TEST171: Test all FrameType discriminants roundtrip through u8 conversion preserving identity | bifaci/frame_test.go:13 |
| test172 | `Test172_frame_type_valid_range` | TEST172: Test FrameType::from_u8 returns None for values outside the valid discriminant range | bifaci/frame_test.go:38 |
| test173 | `Test173_frame_type_wire_protocol_values` | TEST173: Test FrameType discriminant values match the wire protocol specification exactly | bifaci/frame_test.go:71 |
| test174 | `Test174_message_id_new_uuid_roundtrip` | TEST174: Test MessageId::new_uuid generates valid UUID that roundtrips through string conversion | bifaci/frame_test.go:106 |
| test175 | `Test175_message_id_uuid_distinct` | TEST175: Test two MessageId::new_uuid calls produce distinct IDs (no collisions) | bifaci/frame_test.go:127 |
| test176 | `Test176_message_id_uint_no_uuid_string` | TEST176: Test MessageId::Uint does not produce a UUID string, to_uuid_string returns None | bifaci/frame_test.go:137 |
| test177 | `Test177_message_id_from_invalid_uuid_str` | TEST177: Test MessageId::from_uuid_str rejects invalid UUID strings | bifaci/frame_test.go:150 |
| test178 | `Test178_message_id_as_bytes` | TEST178: Test MessageId::as_bytes produces correct byte representations for Uuid and Uint variants | bifaci/frame_test.go:160 |
| test179 | `Test179_message_id_default` | TEST179: Test MessageId::default creates a UUID variant (not Uint) | bifaci/frame_test.go:181 |
| test180 | `Test180_frame_hello_without_manifest` | TEST180: Test Frame::hello without manifest produces correct HELLO frame for host side | bifaci/frame_test.go:192 |
| test181 | `Test181_frame_hello_with_manifest` | TEST181: Test Frame::hello_with_manifest produces HELLO with manifest bytes for cartridge side | bifaci/frame_test.go:207 |
| test182 | `Test182_frame_req` | TEST182: Test Frame::req stores cap URN, payload, and content_type correctly | bifaci/frame_test.go:223 |
| test184 | `Test184_frame_chunk` | TEST184: Test Frame::chunk stores seq and payload for streaming (with stream_id) | bifaci/frame_test.go:248 |
| test185 | `Test185_frame_err` | TEST185: Test Frame::err stores error code and message in metadata | bifaci/frame_test.go:273 |
| test186 | `Test186_frame_log` | TEST186: Test Frame::log stores level and message in metadata | bifaci/frame_test.go:292 |
| test187 | `Test187_frame_end_with_payload` | TEST187: Test Frame::end with payload sets eof and optional final payload | bifaci/frame_test.go:311 |
| test188 | `Test188_frame_end_without_payload` | TEST188: Test Frame::end without payload still sets eof marker | bifaci/frame_test.go:329 |
| test189 | `Test189_frame_chunk_with_offset` | TEST189: Test chunk_with_offset sets offset on all chunks but len only on seq=0 (with stream_id) | bifaci/frame_test.go:345 |
| test190 | `Test190_frame_heartbeat` | TEST190: Test Frame::heartbeat creates minimal frame with no payload or metadata | bifaci/frame_test.go:395 |
| test191 | `Test191_error_accessors_on_non_err_frame` | TEST191: Test error_code and error_message return None for non-Err frame types | bifaci/frame_test.go:411 |
| test192 | `Test192_log_accessors_on_non_log_frame` | TEST192: Test log_level and log_message return None for non-Log frame types | bifaci/frame_test.go:427 |
| test193 | `Test193_hello_accessors_on_non_hello_frame` | TEST193: Test hello_max_frame and hello_max_chunk return None for non-Hello frame types | bifaci/frame_test.go:438 |
| test194 | `Test194_frame_new_defaults` | TEST194: Test Frame::new sets version and defaults correctly, optional fields are None | bifaci/frame_test.go:449 |
| test195 | `Test195_frame_default_type` | TEST195: Test Frame::default creates a Req frame (the documented default) | bifaci/frame_test.go:489 |
| test196 | `Test196_is_eof_when_none` | TEST196: Test is_eof returns false when eof field is None (unset) | bifaci/frame_test.go:500 |
| test197 | `Test197_is_eof_when_false` | TEST197: Test is_eof returns false when eof field is explicitly Some(false) | bifaci/frame_test.go:508 |
| test198 | `Test198_limits_default` | TEST198: Test Limits::default provides the documented default values | bifaci/frame_test.go:518 |
| test199 | `Test199_protocol_version_constant` | TEST199: Test PROTOCOL_VERSION is 2 | bifaci/frame_test.go:536 |
| test200 | `Test200_key_constants` | TEST200: Test integer key constants match the protocol specification | bifaci/frame_test.go:543 |
| test201 | `Test201_hello_manifest_binary_data` | TEST201: Test hello_with_manifest preserves binary manifest data (not just JSON text) | bifaci/frame_test.go:580 |
| test202 | `Test202_message_id_equality_and_hash` | TEST202: Test MessageId Eq/Hash semantics: equal UUIDs are equal, different ones are not | bifaci/frame_test.go:602 |
| test203 | `Test203_message_id_cross_variant_inequality` | TEST203: Test Uuid and Uint variants of MessageId are never equal even for coincidental byte values | bifaci/frame_test.go:630 |
| test204 | `Test204_req_frame_empty_payload` | TEST204: Test Frame::req with empty payload stores Some(empty vec) not None | bifaci/frame_test.go:641 |
| test205 | `Test205_req_frame_roundtrip` | TEST205: Test REQ frame encode/decode roundtrip preserves all fields | bifaci/io_test.go:27 |
| test206 | `Test206_hello_frame_roundtrip` | TEST206: Test HELLO frame encode/decode roundtrip preserves max_frame, max_chunk, max_reorder_buffer | bifaci/io_test.go:59 |
| test207 | `Test207_err_frame_roundtrip` | TEST207: Test ERR frame encode/decode roundtrip preserves error code and message | bifaci/io_test.go:84 |
| test208 | `Test208_log_frame_roundtrip` | TEST208: Test LOG frame encode/decode roundtrip preserves level and message | bifaci/io_test.go:109 |
| test210 | `Test210_end_frame_roundtrip` | TEST210: Test END frame encode/decode roundtrip preserves eof marker and optional payload | bifaci/io_test.go:136 |
| test211 | `Test211_hello_with_manifest_roundtrip` | TEST211: Test HELLO with manifest encode/decode roundtrip preserves manifest bytes and limits | bifaci/io_test.go:163 |
| test212 | `Test212_chunk_with_offset_roundtrip` | TEST212: Test chunk_with_offset encode/decode roundtrip preserves offset, len, eof (with stream_id) | bifaci/io_test.go:189 |
| test213 | `Test213_heartbeat_roundtrip` | TEST213: Test heartbeat frame encode/decode roundtrip preserves ID with no extra fields | bifaci/io_test.go:233 |
| test214 | `Test214_frame_io_roundtrip` | TEST214: Test write_frame/read_frame IO roundtrip through length-prefixed wire format | bifaci/io_test.go:256 |
| test215 | `Test215_read_multiple_frames` | TEST215: Test reading multiple sequential frames from a single buffer | bifaci/io_test.go:281 |
| test216 | `Test216_write_frame_rejects_oversized` | TEST216: Test write_frame rejects frames exceeding max_frame limit | bifaci/io_test.go:315 |
| test217 | `Test217_read_frame_rejects_oversized` | TEST217: Test read_frame rejects incoming frames exceeding the negotiated max_frame limit | bifaci/io_test.go:334 |
| test218 | `Test218_write_chunked` | TEST218: Test write_chunked splits data into chunks respecting max_chunk and reconstructs correctly Chunks from write_chunked have seq=0. SeqAssigner at the output stage assigns final seq. Chunk ordering within a stream is tracked by chunk_index (chunk_index field). | bifaci/io_test.go:355 |
| test219 | `Test219_write_chunked_empty` | TEST219: Test write_chunked with empty data produces a single EOF chunk | bifaci/io_test.go:413 |
| test220 | `Test220_write_chunked_exact_chunk_size` | TEST220: Test write_chunked with data exactly equal to max_chunk produces exactly one chunk | bifaci/io_test.go:456 |
| test221 | `Test221_read_frame_eof` | TEST221: Test read_frame returns Ok(None) on clean EOF (empty stream) | bifaci/io_test.go:493 |
| test222 | `Test222_read_frame_truncated_length_prefix` | TEST222: Test read_frame handles truncated length prefix (fewer than 4 bytes available) | bifaci/io_test.go:504 |
| test223 | `Test223_read_frame_truncated_body` | TEST223: Test read_frame returns error on truncated frame body (length prefix says more bytes than available) | bifaci/io_test.go:515 |
| test224 | `Test224_message_id_uint_roundtrip` | TEST224: Test MessageId::Uint roundtrips through encode/decode | bifaci/io_test.go:531 |
| test225 | `Test225_decode_non_map_value` | TEST225: Test decode_frame rejects non-map CBOR values (e.g., array, integer, string) | bifaci/io_test.go:551 |
| test226 | `Test226_decode_missing_version` | TEST226: Test decode_frame rejects CBOR map missing required version field | bifaci/io_test.go:562 |
| test227 | `Test227_decode_invalid_frame_type_value` | TEST227: Test decode_frame rejects CBOR map with invalid frame_type value | bifaci/io_test.go:577 |
| test228 | `Test228_decode_missing_id` | TEST228: Test decode_frame rejects CBOR map missing required id field | bifaci/io_test.go:591 |
| test229 | `Test229_frame_reader_writer_set_limits` | TEST229: Test FrameReader/FrameWriter set_limits updates the negotiated limits | bifaci/io_test.go:605 |
| test230 | `Test230_sync_handshake` | TEST230: Test async handshake exchanges HELLO frames and negotiates minimum limits | bifaci/io_test.go:629 |
| test231 | `Test231_handshake_rejects_non_hello` | TEST231: Test handshake fails when peer sends non-HELLO frame | bifaci/io_test.go:727 |
| test232 | `Test232_handshake_rejects_missing_manifest` | TEST232: Test handshake fails when cartridge HELLO is missing required manifest | bifaci/io_test.go:763 |
| test233 | `Test233_binary_payload_all_byte_values` | TEST233: Test binary payload with all 256 byte values roundtrips through encode/decode | bifaci/io_test.go:797 |
| test234 | `Test234_decode_garbage_bytes` | TEST234: Test decode_frame handles garbage CBOR bytes gracefully with an error | bifaci/io_test.go:822 |
| test235 | `Test235_response_chunk_fields` | TEST235: Test ResponseChunk stores payload, seq, offset, len, and eof fields correctly | bifaci/host_test.go:10 |
| test236 | `Test236_response_chunk_all_fields_populated` | TEST236: Test ResponseChunk with all fields populated preserves offset, len, and eof | bifaci/host_test.go:34 |
| test237 | `Test237_cartridge_response_single_final_payload` | TEST237: Test CartridgeResponse::Single final_payload returns the single payload slice | bifaci/host_test.go:56 |
| test238 | `Test238_cartridge_response_single_empty_payload` | TEST238: Test CartridgeResponse::Single with empty payload returns empty slice and empty vec | bifaci/host_test.go:68 |
| test239 | `Test239_cartridge_response_streaming_concatenated` | TEST239: Test CartridgeResponse::Streaming concatenated joins all chunk payloads in order | bifaci/host_test.go:79 |
| test240 | `Test240_cartridge_response_streaming_final_payload` | TEST240: Test CartridgeResponse::Streaming final_payload returns the last chunk's payload | bifaci/host_test.go:96 |
| test241 | `Test241_cartridge_response_streaming_empty_chunks` | TEST241: Test CartridgeResponse::Streaming with empty chunks vec returns empty concatenation | bifaci/host_test.go:113 |
| test242 | `Test242_cartridge_response_streaming_preallocation` | TEST242: Test CartridgeResponse::Streaming concatenated capacity is pre-allocated correctly for large payloads | bifaci/host_test.go:127 |
| test243 | `Test243_host_error_variants` | TEST243: Test AsyncHostError variants display correct error messages | bifaci/host_test.go:146 |
| test244 | `Test244_host_error_conversion` | TEST244: Test AsyncHostError::from converts CborError to Cbor variant | bifaci/host_test.go:195 |
| test245 | `Test245_host_error_io_variant` | TEST245: Test AsyncHostError::from converts io::Error to Io variant | bifaci/host_test.go:206 |
| test246 | `Test246_response_chunk_copy` | TEST246: Test AsyncHostError Clone implementation produces equal values | bifaci/host_test.go:217 |
| test247 | `Test247_response_chunk_clone` | TEST247: Test ResponseChunk Clone produces independent copy with same data | bifaci/host_test.go:241 |
| test248 | `Test248_register_and_find_handler` | TEST248: Test register_op and find_handler by exact cap URN | bifaci/cartridge_runtime_test.go:99 |
| test249 | `Test249_raw_handler` | TEST249: Test register_op handler echoes bytes directly | bifaci/cartridge_runtime_test.go:117 |
| test250 | `Test250_typed_handler_deserialization` | TEST250: Test Op handler collects input and processes it | bifaci/cartridge_runtime_test.go:155 |
| test251 | `Test251_typed_handler_rejects_invalid_json` | TEST251: Test Op handler propagates errors through RuntimeError::Handler | bifaci/cartridge_runtime_test.go:195 |
| test252 | `Test252_find_handler_unknown_cap` | TEST252: Test find_handler returns None for unregistered cap URNs | bifaci/cartridge_runtime_test.go:224 |
| test253 | `Test253_handler_is_send_sync` | TEST253: Test OpFactory can be cloned via Arc and sent across tasks (Send + Sync) | bifaci/cartridge_runtime_test.go:237 |
| test254 | `Test254_no_peer_invoker` | TEST254: Test NoPeerInvoker always returns PeerRequest error | bifaci/cartridge_runtime_test.go:275 |
| test255 | `Test255_no_peer_invoker_with_arguments` | TEST255: Test NoPeerInvoker call_with_bytes also returns error | bifaci/cartridge_runtime_test.go:287 |
| test256 | `Test256_new_cartridge_runtime_with_valid_json` | TEST256: Test CartridgeRuntime::with_manifest_json stores manifest data and parses when valid | bifaci/cartridge_runtime_test.go:299 |
| test257 | `Test257_new_cartridge_runtime_with_invalid_json` | TEST257: Test CartridgeRuntime::new with invalid JSON still creates runtime (manifest is None) | bifaci/cartridge_runtime_test.go:314 |
| test258 | `Test258_new_cartridge_runtime_with_manifest_struct` | TEST258: Test CartridgeRuntime::with_manifest creates runtime with valid manifest data | bifaci/cartridge_runtime_test.go:329 |
| test259 | `Test259_extract_effective_payload_non_cbor` | TEST259: Test extract_effective_payload with non-CBOR content_type returns raw payload unchanged | bifaci/cartridge_runtime_test.go:349 |
| test260 | `Test260_extract_effective_payload_no_content_type` | TEST260: Test extract_effective_payload with None content_type returns raw payload unchanged | bifaci/cartridge_runtime_test.go:361 |
| test261 | `Test261_extract_effective_payload_cbor_match` | TEST261: Test extract_effective_payload with CBOR content extracts matching argument value | bifaci/cartridge_runtime_test.go:373 |
| test262 | `Test262_extract_effective_payload_cbor_no_match` | TEST262: Test extract_effective_payload with CBOR content fails when no argument matches expected input | bifaci/cartridge_runtime_test.go:388 |
| test263 | `Test263_extract_effective_payload_invalid_cbor` | TEST263: Test extract_effective_payload with invalid CBOR bytes returns deserialization error | bifaci/cartridge_runtime_test.go:400 |
| test264 | `Test264_extract_effective_payload_cbor_not_array` | TEST264: Test extract_effective_payload with CBOR non-array (e.g. map) returns error | bifaci/cartridge_runtime_test.go:411 |
| test270 | `Test270_multiple_handlers` | TEST270: Test registering multiple Op handlers for different caps and finding each independently | bifaci/cartridge_runtime_test.go:431 |
| test271 | `Test271_handler_replacement` | TEST271: Test Op handler replacing an existing registration for the same cap URN | bifaci/cartridge_runtime_test.go:481 |
| test272 | `Test272_extract_effective_payload_multiple_args` | TEST272: Test extract_effective_payload CBOR with multiple arguments selects the correct one | bifaci/cartridge_runtime_test.go:508 |
| test273 | `Test273_ExtractEffectivePayloadBinaryValue` | TEST273: Test extract_effective_payload with binary data in CBOR value (not just text) | bifaci/cartridge_runtime_test.go:522 |
| test274 | `Test274_cap_argument_value_new` | TEST274: Test CapArgumentValue::new stores media_urn and raw byte value | cap/caller_test.go:349 |
| test275 | `Test275_cap_argument_value_from_str` | TEST275: Test CapArgumentValue::from_str converts string to UTF-8 bytes | cap/caller_test.go:356 |
| test276 | `Test276_cap_argument_value_as_str_valid` | TEST276: Test CapArgumentValue::value_as_str succeeds for UTF-8 data | cap/caller_test.go:363 |
| test277 | `Test277_cap_argument_value_as_str_invalid_utf8` | TEST277: Test CapArgumentValue::value_as_str fails for non-UTF-8 binary data | cap/caller_test.go:371 |
| test278 | `Test278_cap_argument_value_empty` | TEST278: Test CapArgumentValue::new with empty value stores empty vec | cap/caller_test.go:378 |
| test279 | `Test279_cap_argument_value_clone` | TEST279: Test CapArgumentValue Clone produces independent copy with same data | cap/caller_test.go:387 |
| test280 | `Test280_cap_argument_value_debug` | TEST280: Test CapArgumentValue Debug format includes media_urn and value | cap/caller_test.go:404 |
| test281 | `Test281_cap_argument_value_string_types` | TEST281: Test CapArgumentValue::new accepts Into<String> for media_urn (String and &str) | cap/caller_test.go:413 |
| test282 | `Test282_cap_argument_value_unicode` | TEST282: Test CapArgumentValue::from_str with Unicode string preserves all characters | cap/caller_test.go:423 |
| test283 | `Test283_cap_argument_value_large_binary` | TEST283: Test CapArgumentValue with large binary payload preserves all bytes | cap/caller_test.go:431 |
| test284 | `Test284_HandshakeHostCartridge` | TEST284: Handshake exchanges HELLO frames, negotiates limits | bifaci/integration_test.go:467 |
| test285 | `Test285_RequestResponseSimple` | TEST285: Simple request-response flow (REQ → END with payload) | bifaci/integration_test.go:509 |
| test286 | `Test286_StreamingChunks` | TEST286: Streaming response with multiple CHUNK frames | bifaci/integration_test.go:571 |
| test287 | `Test287_HeartbeatFromHost` | TEST287: Host-initiated heartbeat | bifaci/integration_test.go:649 |
| test290 | `Test290_LimitsNegotiation` | TEST290: Limit negotiation picks minimum | bifaci/integration_test.go:840 |
| test291 | `Test291_BinaryPayloadRoundtrip` | TEST291: Binary payload roundtrip (all 256 byte values) | bifaci/integration_test.go:880 |
| test292 | `Test292_MessageIdUniqueness` | TEST292: Sequential requests get distinct MessageIds | bifaci/integration_test.go:954 |
| test293 | `Test293_CartridgeRuntimeHandlerRegistration` | TEST293: Test CartridgeRuntime Op registration and lookup by exact and non-existent cap URN | bifaci/integration_test.go:1025 |
| test299 | `Test299_EmptyPayloadRoundtrip` | TEST299: Empty payload request/response roundtrip | bifaci/integration_test.go:1361 |
| test304 | `Test304_media_availability_output_constant` | TEST304: Test MEDIA_AVAILABILITY_OUTPUT constant parses as valid media URN with correct tags | urn/media_urn_test.go:379 |
| test305 | `Test305_media_path_output_constant` | TEST305: Test MEDIA_PATH_OUTPUT constant parses as valid media URN with correct tags | urn/media_urn_test.go:388 |
| test306 | `Test306_availability_and_path_output_distinct` | TEST306: Test MEDIA_AVAILABILITY_OUTPUT and MEDIA_PATH_OUTPUT are distinct URNs | urn/media_urn_test.go:397 |
| test307 | `Test307_model_availability_urn` | TEST307: Test model_availability_urn builds valid cap URN with correct op and media specs | standard/caps_test.go:11 |
| test308 | `Test308_model_path_urn` | TEST308: Test model_path_urn builds valid cap URN with correct op and media specs | standard/caps_test.go:19 |
| test309 | `Test309_model_availability_and_path_are_distinct` | TEST309: Test model_availability_urn and model_path_urn produce distinct URNs | standard/caps_test.go:27 |
| test310 | `Test310_llm_generate_text_urn_shape` | TEST310: llm_generate_text_urn() produces a valid cap URN with textable in/out specs | standard/caps_test.go:35 |
| test312 | `Test312_all_urn_builders_produce_valid_urns` | TEST312: Test all URN builders produce parseable cap URNs | standard/caps_test.go:46 |
| test320 | `Test320_cartridge_info_construction` | TEST320-335: CartridgeRepoServer and CartridgeRepoClient tests | bifaci/cartridge_repo_test.go:40 |
| test321 | `Test321_cartridge_info_is_signed` | TEST321: CartridgeInfo.is_signed() returns true when signature is present | bifaci/cartridge_repo_test.go:71 |
| test322 | `Test322_cartridge_info_build_for_platform` | TEST322: CartridgeInfo.build_for_platform() returns the build matching the current platform | bifaci/cartridge_repo_test.go:98 |
| test323 | `Test323_cartridge_repo_server_validate_registry` | TEST323: CartridgeRepoServer validates registry JSON schema version | bifaci/cartridge_repo_test.go:152 |
| test324 | `Test324_cartridge_repo_server_transform_to_array` | TEST324: CartridgeRepoServer transforms v3 registry JSON into flat cartridge array | bifaci/cartridge_repo_test.go:183 |
| test325 | `Test325_cartridge_repo_server_get_cartridges` | TEST325: CartridgeRepoServer.get_cartridges() returns all parsed cartridges | bifaci/cartridge_repo_test.go:232 |
| test326 | `Test326_cartridge_repo_server_get_cartridge_by_id` | TEST326: CartridgeRepoServer.get_cartridge() returns cartridge matching the given ID | bifaci/cartridge_repo_test.go:262 |
| test327 | `Test327_cartridge_repo_server_search_cartridges` | TEST327: CartridgeRepoServer.search_cartridges() filters by text query against name and description | bifaci/cartridge_repo_test.go:300 |
| test328 | `Test328_cartridge_repo_server_get_by_category` | TEST328: CartridgeRepoServer.get_by_category() filters cartridges by category tag | bifaci/cartridge_repo_test.go:339 |
| test329 | `Test329_cartridge_repo_server_get_by_cap` | TEST329: CartridgeRepoServer.get_suggestions_for_cap() finds cartridges providing a given cap URN | bifaci/cartridge_repo_test.go:378 |
| test330 | `Test330_cartridge_repo_client_update_cache` | TEST330: CartridgeRepoClient updates its local cache from server response | bifaci/cartridge_repo_test.go:419 |
| test331 | `Test331_cartridge_repo_client_get_suggestions` | TEST331: CartridgeRepoClient.get_suggestions_for_cap() returns cartridge suggestions for a cap URN | bifaci/cartridge_repo_test.go:448 |
| test332 | `Test332_cartridge_repo_client_get_cartridge` | TEST332: CartridgeRepoClient.get_cartridge() retrieves a specific cartridge by ID from cache | bifaci/cartridge_repo_test.go:483 |
| test333 | `Test333_cartridge_repo_client_get_all_caps` | TEST333: CartridgeRepoClient.get_all_caps() returns aggregate cap URNs from all cached cartridges | bifaci/cartridge_repo_test.go:515 |
| test334 | `Test334_cartridge_repo_client_needs_sync` | TEST334: CartridgeRepoClient.needs_sync() returns true when cache TTL has expired | bifaci/cartridge_repo_test.go:565 |
| test335 | `Test335_cartridge_repo_server_client_integration` | TEST335: Server creates registry response and client consumes it end-to-end | bifaci/cartridge_repo_test.go:582 |
| test336 | `Test336_FilePathReadsFilePassesBytes` | TEST336: Single file-path arg with stdin source reads file and passes bytes to handler | bifaci/cartridge_runtime_test.go:599 |
| test337 | `Test337_FilePathWithoutStdinPassesString` | TEST337: file-path arg without stdin source passes path as string (no conversion) | bifaci/cartridge_runtime_test.go:674 |
| test338 | `Test338_FilePathViaCliFlag` | TEST338: file-path arg reads file via --file CLI flag | bifaci/cartridge_runtime_test.go:715 |
| test339 | `Test339_FilePathArrayGlobExpansion` | TEST339: file-path-array reads multiple files with glob pattern | bifaci/cartridge_runtime_test.go:755 |
| test340 | `Test340_FileNotFoundClearError` | TEST340: File not found error provides clear message | bifaci/cartridge_runtime_test.go:823 |
| test341 | `Test341_StdinPrecedenceOverFilePath` | TEST341: stdin takes precedence over file-path in source order | bifaci/cartridge_runtime_test.go:862 |
| test342 | `Test342_FilePathPositionZeroReadsFirstArg` | TEST342: file-path with position 0 reads first positional arg as file | bifaci/cartridge_runtime_test.go:905 |
| test343 | `Test343_NonFilePathArgsUnaffected` | TEST343: Non-file-path args are not affected by file reading | bifaci/cartridge_runtime_test.go:945 |
| test344 | `Test344_FilePathArrayInvalidJSONFails` | TEST344: file-path-array with nonexistent path fails clearly | bifaci/cartridge_runtime_test.go:981 |
| test345 | `Test345_FilePathArrayOneFileMissingFailsHard` | TEST345: file-path-array with literal nonexistent path fails hard | bifaci/cartridge_runtime_test.go:1020 |
| test346 | `Test346_LargeFileReadsSuccessfully` | TEST346: Large file (1MB) reads successfully | bifaci/cartridge_runtime_test.go:1068 |
| test347 | `Test347_EmptyFileReadsAsEmptyBytes` | TEST347: Empty file reads as empty bytes | bifaci/cartridge_runtime_test.go:1112 |
| test348 | `Test348_FilePathConversionRespectsSourceOrder` | TEST348: file-path conversion respects source order | bifaci/cartridge_runtime_test.go:1152 |
| test349 | `Test349_FilePathMultipleSourcesFallback` | TEST349: file-path arg with multiple sources tries all in order | bifaci/cartridge_runtime_test.go:1195 |
| test350 | `Test350_FullCLIModeWithFilePathIntegration` | TEST350: Integration test - full CLI mode invocation with file-path | bifaci/cartridge_runtime_test.go:1237 |
| test351 | `Test351_FilePathArrayEmptyArray` | TEST351: file-path array with empty CBOR array returns empty (CBOR mode) | bifaci/cartridge_runtime_test.go:1313 |
| test352 | `Test352_FilePermissionDeniedClearError` | TEST352: file permission denied error is clear (Unix-specific) | bifaci/cartridge_runtime_test.go:1354 |
| test353 | `Test353_CBORPayloadFormatConsistency` | TEST353: CBOR payload format matches between CLI and CBOR mode | bifaci/cartridge_runtime_test.go:1404 |
| test354 | `Test354_GlobPatternNoMatchesEmptyArray` | TEST354: Glob pattern with no matches fails hard (NO FALLBACK) | bifaci/cartridge_runtime_test.go:1462 |
| test355 | `Test355_GlobPatternSkipsDirectories` | TEST355: Glob pattern skips directories | bifaci/cartridge_runtime_test.go:1509 |
| test356 | `Test356_MultipleGlobPatternsCombined` | TEST356: Multiple glob patterns combined | bifaci/cartridge_runtime_test.go:1574 |
| test357 | `Test357_SymlinksFollowed` | TEST357: Symlinks are followed when reading files | bifaci/cartridge_runtime_test.go:1644 |
| test358 | `Test358_BinaryFileNonUTF8` | TEST358: Binary file with non-UTF8 data reads correctly | bifaci/cartridge_runtime_test.go:1697 |
| test359 | `Test359_InvalidGlobPatternFails` | TEST359: Invalid glob pattern fails with clear error | bifaci/cartridge_runtime_test.go:1743 |
| test360 | `Test360_ExtractEffectivePayloadWithFileData` | TEST360: Extract effective payload handles file-path data correctly | bifaci/cartridge_runtime_test.go:1781 |
| test361 | `Test361_CLIModeFilePath` | TEST361: CLI mode with file path - pass file path as command-line argument | bifaci/cartridge_runtime_test.go:1831 |
| test362 | `Test362_CLIModePipedBinary` | TEST362: CLI mode with binary piped in - pipe binary data via stdin This test simulates real-world conditions: - Pure binary data piped to stdin (NOT CBOR) - CLI mode detected (command arg present) - Cap accepts stdin source - Binary is chunked on-the-fly and accumulated - Handler receives complete CBOR payload | bifaci/cartridge_runtime_test.go:1880 |
| test363 | `Test363_CBORModeChunkedContent` | TEST363: CBOR mode with chunked content - send file content streaming as chunks | bifaci/cartridge_runtime_test.go:1967 |
| test364 | `Test364_CBORModeFilePath` | TEST364: CBOR mode with file path - send file path in CBOR arguments (auto-conversion) | bifaci/cartridge_runtime_test.go:2110 |
| test365 | `Test365_stream_start_frame` | TEST365: Frame::stream_start stores request_id, stream_id, and media_urn | bifaci/frame_test.go:652 |
| test366 | `Test366_stream_end_frame` | TEST366: Frame::stream_end stores request_id and stream_id | bifaci/frame_test.go:674 |
| test367 | `Test367_stream_start_with_empty_stream_id` | TEST367: StreamStart frame with empty stream_id still constructs (validation happens elsewhere) | bifaci/frame_test.go:695 |
| test368 | `Test368_stream_start_with_empty_media_urn` | TEST368: StreamStart frame with empty media_urn still constructs (validation happens elsewhere) | bifaci/frame_test.go:714 |
| test389 | `Test389_stream_start_roundtrip` | TEST389: StreamStart encode/decode roundtrip preserves stream_id and media_urn | bifaci/io_test.go:831 |
| test390 | `Test390_stream_end_roundtrip` | TEST390: StreamEnd encode/decode roundtrip preserves stream_id, no media_urn | bifaci/io_test.go:859 |
| test395 | `Test395_BuildPayloadSmall` | TEST395: Small payload (< max_chunk) produces correct CBOR arguments | bifaci/cartridge_runtime_test.go:2165 |
| test396 | `Test396_BuildPayloadLarge` | TEST396: Large payload (> max_chunk) accumulates across chunks correctly | bifaci/cartridge_runtime_test.go:2214 |
| test397 | `Test397_BuildPayloadEmpty` | TEST397: Empty reader produces valid empty CBOR arguments | bifaci/cartridge_runtime_test.go:2258 |
| test398 | `Test398_BuildPayloadIOError` | TEST398: IO error from reader propagates as RuntimeError::Io | bifaci/cartridge_runtime_test.go:2301 |
| test399 | `Test399_relay_notify_discriminant_roundtrip` | TEST399: Verify RelayNotify frame type discriminant roundtrips through u8 (value 10) | bifaci/frame_test.go:733 |
| test400 | `Test400_relay_state_discriminant_roundtrip` | TEST400: Verify RelayState frame type discriminant roundtrips through u8 (value 11) | bifaci/frame_test.go:746 |
| test401 | `Test401_relay_notify_factory_and_accessors` | TEST401: Verify relay_notify factory stores manifest and limits, and accessors extract them | bifaci/frame_test.go:759 |
| test402 | `Test402_relay_state_factory_and_payload` | TEST402: Verify relay_state factory stores resource payload in frame payload field | bifaci/frame_test.go:802 |
| test403 | `Test403_frame_type_one_past_cancel` | TEST403: Verify from_u8 returns None for values past the last valid frame type | bifaci/frame_test.go:816 |
| test404 | `Test404_slave_sends_relay_notify_on_connect` | TEST404: Slave sends RelayNotify on connect (initial_notify parameter) | bifaci/relay_test.go:15 |
| test405 | `Test405_master_reads_relay_notify` | TEST405: Master reads RelayNotify and extracts manifest + limits | bifaci/relay_test.go:70 |
| test406 | `Test406_slave_stores_relay_state` | TEST406: Slave stores RelayState from master | bifaci/relay_test.go:112 |
| test407 | `Test407_protocol_frames_pass_through` | TEST407: Protocol frames pass through slave transparently (both directions) | bifaci/relay_test.go:166 |
| test408 | `Test408_relay_frames_not_forwarded` | TEST408: RelayNotify/RelayState are NOT forwarded through relay | bifaci/relay_test.go:293 |
| test409 | `Test409_slave_injects_relay_notify_midstream` | TEST409: Slave can inject RelayNotify mid-stream (cap change) | bifaci/relay_test.go:374 |
| test410 | `Test410_master_receives_updated_relay_notify` | TEST410: Master receives updated RelayNotify (cap change callback via read_frame) | bifaci/relay_test.go:449 |
| test411 | `Test411_socket_close_detection` | TEST411: Socket close detection (both directions) | bifaci/relay_test.go:543 |
| test412 | `Test412_bidirectional_concurrent_flow` | TEST412: Bidirectional concurrent frame flow through relay | bifaci/relay_test.go:589 |
| test413 | `Test413_register_cartridge_adds_cap_table` | TEST413: Register cartridge adds entries to cap_table | bifaci/host_multi_test.go:34 |
| test414 | `Test414_capabilities_empty_initially` | TEST414: capabilities() returns empty JSON initially (no running cartridges) | bifaci/host_multi_test.go:52 |
| test415 | `Test415_req_triggers_spawn` | TEST415: REQ for known cap triggers spawn attempt (verified by expected spawn error for non-existent binary) | bifaci/host_multi_test.go:63 |
| test416 | `Test416_attach_cartridge_handshake` | TEST416: Attach cartridge performs HELLO handshake, extracts manifest, updates capabilities | bifaci/host_multi_test.go:100 |
| test417 | `Test417_route_req_by_cap_urn` | TEST417: Route REQ to correct cartridge by cap_urn (with two attached cartridges) | bifaci/host_multi_test.go:138 |
| test418 | `Test418_route_continuation_by_req_id` | TEST418: Route STREAM_START/CHUNK/STREAM_END/END by req_id (not cap_urn) Verifies that after the initial REQ→cartridge routing, all subsequent continuation frames with the same req_id are routed to the same cartridge — even though no cap_urn is present on those frames. | bifaci/host_multi_test.go:240 |
| test419 | `Test419_heartbeat_local_handling` | TEST419: Cartridge HEARTBEAT handled locally (not forwarded to relay) | bifaci/host_multi_test.go:329 |
| test420 | `Test420_cartridge_frames_forwarded_to_relay` | TEST420: Cartridge non-HELLO/non-HB frames forwarded to relay (pass-through) | bifaci/host_multi_test.go:410 |
| test421 | `Test421_cartridge_death_updates_caps` | TEST421: Cartridge death updates capability list (caps removed) | bifaci/host_multi_test.go:500 |
| test422 | `Test422_cartridge_death_sends_err` | TEST422: Cartridge death sends ERR for all pending requests via relay | bifaci/host_multi_test.go:555 |
| test423 | `Test423_multi_cartridge_distinct_caps` | TEST423: Multiple cartridges registered with distinct caps route independently | bifaci/host_multi_test.go:624 |
| test424 | `Test424_concurrent_requests_same_cartridge` | TEST424: Concurrent requests to the same cartridge are handled independently | bifaci/host_multi_test.go:745 |
| test425 | `Test425_find_cartridge_for_cap_unknown` | TEST425: find_cartridge_for_cap returns None for unregistered cap | bifaci/host_multi_test.go:850 |
| test426 | `Test426_relay_switch_single_master_req_response` | TEST426: Single master REQ/response routing | bifaci/relay_switch_test.go:10 |
| test427 | `Test427_relay_switch_multi_master_cap_routing` | TEST427: Multi-master cap routing | bifaci/relay_switch_test.go:76 |
| test428 | `Test428_relay_switch_unknown_cap_returns_error` | TEST428: Unknown cap returns error | bifaci/relay_switch_test.go:164 |
| test429 | `Test429_relay_switch_find_master_for_cap` | TEST429: Cap routing logic (find_master_for_cap) | bifaci/relay_switch_test.go:209 |
| test430 | `Test430_relay_switch_tie_breaking` | TEST430: Tie-breaking (same cap on multiple masters - first match wins, routing is consistent) | bifaci/relay_switch_test.go:282 |
| test431 | `Test431_relay_switch_continuation_frame_routing` | TEST431: Continuation frame routing (CHUNK, END follow REQ) | bifaci/relay_switch_test.go:353 |
| test432 | `Test432_relay_switch_empty_masters_list_error` | TEST432: Empty masters list creates empty switch, add_master works | bifaci/relay_switch_test.go:430 |
| test433 | `Test433_relay_switch_capability_aggregation_deduplicates` | TEST433: Capability aggregation deduplicates caps | bifaci/relay_switch_test.go:445 |
| test434 | `Test434_relay_switch_limits_negotiation_minimum` | TEST434: Limits negotiation takes minimum | bifaci/relay_switch_test.go:503 |
| test435 | `Test435_relay_switch_urn_matching` | TEST435: URN matching (exact vs accepts()) | bifaci/relay_switch_test.go:552 |
| test436 | `Test436_compute_checksum` | TEST436: Verify FNV-1a checksum function produces consistent results | bifaci/frame_test.go:862 |
| test440 | `Test440_chunk_index_checksum_roundtrip` | TEST440: CHUNK frame with chunk_index and checksum roundtrips through encode/decode | bifaci/io_test.go:950 |
| test441 | `Test441_stream_end_chunk_count_roundtrip` | TEST441: STREAM_END frame with chunk_count roundtrips through encode/decode | bifaci/io_test.go:991 |
| test442 | `Test442_seq_assigner_monotonic_same_rid` | TEST442: SeqAssigner assigns seq 0,1,2,3 for consecutive frames with same RID | bifaci/frame_test.go:879 |
| test443 | `Test443_seq_assigner_independent_rids` | TEST443: SeqAssigner maintains independent counters for different RIDs | bifaci/frame_test.go:908 |
| test444 | `Test444_seq_assigner_skips_non_flow` | TEST444: SeqAssigner skips non-flow frames (Heartbeat, RelayNotify, RelayState, Hello) | bifaci/frame_test.go:934 |
| test445 | `Test445_seq_assigner_remove_by_flow_key` | TEST445: SeqAssigner.remove with FlowKey(rid, None) resets that flow; FlowKey(rid, Some(xid)) is unaffected | bifaci/frame_test.go:962 |
| test446 | `Test446_seq_assigner_mixed_types` | TEST446: SeqAssigner handles mixed frame types (REQ, CHUNK, LOG, END) for same RID | bifaci/frame_test.go:1044 |
| test447 | `Test447_flow_key_with_xid` | TEST447: FlowKey::from_frame extracts (rid, Some(xid)) when routing_id present | bifaci/frame_test.go:1065 |
| test448 | `Test448_flow_key_without_xid` | TEST448: FlowKey::from_frame extracts (rid, None) when routing_id absent | bifaci/frame_test.go:1082 |
| test449 | `Test449_flow_key_equality` | TEST449: FlowKey equality: same rid+xid equal, different xid different key | bifaci/frame_test.go:1096 |
| test450 | `Test450_flow_key_hash_lookup` | TEST450: FlowKey hash: same keys hash equal (HashMap lookup) | bifaci/frame_test.go:1118 |
| test451 | `Test451_reorder_buffer_in_order` | TEST451: ReorderBuffer in-order delivery: seq 0,1,2 delivered immediately | bifaci/frame_test.go:1132 |
| test452 | `Test452_reorder_buffer_out_of_order` | TEST452: ReorderBuffer out-of-order: seq 1 then 0 delivers both in order | bifaci/frame_test.go:1157 |
| test453 | `Test453_reorder_buffer_gap_fill` | TEST453: ReorderBuffer gap fill: seq 0,2,1 delivers 0, buffers 2, then delivers 1+2 | bifaci/frame_test.go:1180 |
| test454 | `Test454_reorder_buffer_stale_seq` | TEST454: ReorderBuffer stale seq is hard error | bifaci/frame_test.go:1207 |
| test455 | `Test455_reorder_buffer_overflow` | TEST455: ReorderBuffer overflow triggers protocol error | bifaci/frame_test.go:1228 |
| test456 | `Test456_reorder_buffer_independent_flows` | TEST456: Multiple concurrent flows reorder independently | bifaci/frame_test.go:1248 |
| test457 | `Test457_reorder_buffer_cleanup` | TEST457: cleanup_flow removes state; new frames start at seq 0 | bifaci/frame_test.go:1276 |
| test458 | `Test458_reorder_buffer_non_flow_bypass` | TEST458: Non-flow frames bypass reorder entirely | bifaci/frame_test.go:1301 |
| test459 | `Test459_reorder_buffer_end_frame` | TEST459: Terminal END frame flows through correctly | bifaci/frame_test.go:1317 |
| test460 | `Test460_reorder_buffer_err_frame` | TEST460: Terminal ERR frame flows through correctly | bifaci/frame_test.go:1335 |
| test473 | `Test473_cap_discard_parses_as_valid_urn` | TEST473: CAP_DISCARD parses as valid CapUrn with in=media: and out=media:void | standard/caps_test.go:59 |
| test474 | `Test474_cap_discard_structure` | TEST474: CAP_DISCARD accepts specific-input/void-output caps | standard/caps_test.go:68 |
| test475 | `Test475_validate_passes_with_identity` | TEST475: CapManifest::validate() passes when CAP_IDENTITY is present | bifaci/manifest_test.go:356 |
| test476 | `Test476_validate_fails_without_identity` | TEST476: CapManifest::validate() fails when CAP_IDENTITY is missing | bifaci/manifest_test.go:367 |
| test497 | `Test497_chunk_corrupted_payload_rejected` | TEST497: Verify CHUNK frame with corrupted payload is rejected by checksum | bifaci/io_test.go:1021 |
| test544 | `Test544_peer_invoker_sends_end_frame` | TEST544: PeerCall::finish sends END frame | bifaci/cartridge_runtime_test.go:2331 |
| test545 | `Test545_demux_peer_response_returns_data` | TEST545: PeerCall::finish returns PeerResponse with data | bifaci/cartridge_runtime_test.go:2364 |
| test546 | `Test546_is_image` | TEST546: is_image returns true only when image marker tag is present | urn/media_urn_test.go:408 |
| test547 | `Test547_is_audio` | TEST547: is_audio returns true only when audio marker tag is present | urn/media_urn_test.go:436 |
| test548 | `Test548_is_video` | TEST548: is_video returns true only when video marker tag is present | urn/media_urn_test.go:464 |
| test549 | `Test549_is_numeric` | TEST549: is_numeric returns true only when numeric marker tag is present | urn/media_urn_test.go:488 |
| test550 | `Test550_is_bool` | TEST550: is_bool returns true only when bool marker tag is present | urn/media_urn_test.go:520 |
| test551 | `Test551_is_file_path` | TEST551: is_file_path returns true for scalar file-path, false for array | urn/media_urn_test.go:549 |
| test552 | `Test552_is_file_path_array` | TEST552: is_file_path_array returns true for list file-path, false for scalar | urn/media_urn_test.go:570 |
| test553 | `Test553_is_any_file_path` | TEST553: is_any_file_path returns true for both scalar and array file-path | urn/media_urn_test.go:587 |
| test555 | `Test555_with_tag_and_without_tag` | TEST555: with_tag adds a tag and without_tag removes it | urn/media_urn_test.go:742 |
| test556 | `Test556_image_media_urn_for_ext` | TEST556: image_media_urn_for_ext creates valid image media URN | urn/media_urn_test.go:766 |
| test557 | `Test557_audio_media_urn_for_ext` | TEST557: audio_media_urn_for_ext creates valid audio media URN | urn/media_urn_test.go:778 |
| test558 | `Test558_predicate_constant_consistency` | TEST558: predicates are consistent with constants — every constant triggers exactly the expected predicates | urn/media_urn_test.go:608 |
| test559 | `Test559_without_tag` | TEST559: without_tag removes tag, ignores in/out, case-insensitive for keys | urn/cap_urn_test.go:1012 |
| test560 | `Test560_with_in_out_spec` | TEST560: with_in_spec and with_out_spec change direction specs | urn/cap_urn_test.go:1041 |
| test561 | `Test561_in_out_media_urn` | TEST561: in_media_urn and out_media_urn parse direction specs into MediaUrn | urn/cap_urn_test.go:1428 |
| test562 | `Test562_canonical_option` | TEST562: canonical_option returns None for None input, canonical string for Some | urn/cap_urn_test.go:1451 |
| test563 | `Test563_find_all_matches` | TEST563: CapMatcher::find_all_matches returns all matching caps sorted by specificity | urn/cap_urn_test.go:1067 |
| test564 | `Test564_are_compatible` | TEST564: CapMatcher::are_compatible detects bidirectional overlap | urn/cap_urn_test.go:1094 |
| test565 | `Test565_tags_to_string` | TEST565: tags_to_string returns only tags portion without prefix | urn/cap_urn_test.go:1124 |
| test566 | `Test566_with_tag_ignores_in_out` | TEST566: with_tag silently ignores in/out keys | urn/cap_urn_test.go:1136 |
| test567 | `Test567_str_variants` | TEST567: conforms_to_str and accepts_str work with string arguments | urn/cap_urn_test.go:1150 |
| test568 | `Test568_dispatch_output_tag_order` | TEST568: is_dispatchable with different tag order in output spec | urn/cap_urn_test.go:1477 |
| test569 | `Test569_unregister_cap_set` | TEST569: unregister_cap_set removes a host and returns true, false if not found | cap_matrix_test.go:873 |
| test570 | `Test570_clear` | TEST570: clear removes all registered sets | cap_matrix_test.go:899 |
| test571 | `Test571_get_all_capabilities` | TEST571: get_all_capabilities returns caps from all hosts | cap_matrix_test.go:920 |
| test572 | `Test572_get_capabilities_for_host` | TEST572: get_capabilities_for_host returns caps for specific host, None for unknown | cap_matrix_test.go:938 |
| test573 | `Test573_iter_hosts_and_caps` | TEST573: iter_hosts_and_caps iterates all hosts with their capabilities | cap_matrix_test.go:959 |
| test574 | `Test574_cap_block_remove_registry` | TEST574: CapBlock::remove_registry removes by name, returns Arc | cap_matrix_test.go:983 |
| test575 | `Test575_cap_block_get_registry` | TEST575: CapBlock::get_registry returns Arc clone by name | cap_matrix_test.go:1009 |
| test576 | `Test576_cap_block_get_registry_names` | TEST576: CapBlock::get_registry_names returns names in insertion order | cap_matrix_test.go:1029 |
| test577 | `Test577_cap_graph_input_output_specs` | TEST577: CapGraph::get_input_specs and get_output_specs return correct sets | cap_matrix_test.go:1048 |
| test578 | `Test578_rule1_duplicate_media_urns` | TEST578: RULE1 - duplicate media_urns rejected | cap/validation_test.go:118 |
| test579 | `Test579_rule2_empty_sources` | TEST579: RULE2 - empty sources rejected | cap/validation_test.go:129 |
| test580 | `Test580_rule3_different_stdin_urns` | TEST580: RULE3 - multiple stdin sources with different URNs rejected | cap/validation_test.go:139 |
| test581 | `Test581_rule3_same_stdin_urns_ok` | TEST581: RULE3 - multiple stdin sources with same URN is OK | cap/validation_test.go:150 |
| test582 | `Test582_rule4_duplicate_source_type` | TEST582: RULE4 - duplicate source type in single arg rejected | cap/validation_test.go:160 |
| test583 | `Test583_rule5_duplicate_position` | TEST583: RULE5 - duplicate position across args rejected | cap/validation_test.go:173 |
| test584 | `Test584_rule6_position_gap` | TEST584: RULE6 - position gap rejected (0, 2 without 1) | cap/validation_test.go:184 |
| test585 | `Test585_rule6_sequential_ok` | TEST585: RULE6 - sequential positions (0, 1, 2) pass | cap/validation_test.go:195 |
| test586 | `Test586_rule7_position_and_cli_flag` | TEST586: RULE7 - arg with both position and cli_flag rejected | cap/validation_test.go:205 |
| test587 | `Test587_rule9_duplicate_cli_flag` | TEST587: RULE9 - duplicate cli_flag across args rejected | cap/validation_test.go:218 |
| test588 | `Test588_rule10_reserved_cli_flags` | TEST588: RULE10 - reserved cli_flags rejected | cap/validation_test.go:229 |
| test589 | `Test589_all_rules_pass` | TEST589: valid cap args with mixed sources pass all rules | cap/validation_test.go:242 |
| test590 | `Test590_cli_flag_only_args` | TEST590: validate_cap_args accepts cap with only cli_flag sources (no positions) | cap/validation_test.go:257 |
| test591 | `Test591_is_more_specific_than` | TEST591: is_more_specific_than returns true when self has more tags for same request | cap/definition_test.go:237 |
| test592 | `Test592_remove_metadata` | TEST592: remove_metadata adds then removes metadata correctly | cap/definition_test.go:262 |
| test593 | `Test593_registered_by_lifecycle` | TEST593: registered_by lifecycle — set, get, clear | cap/definition_test.go:284 |
| test594 | `Test594_metadata_json_lifecycle` | TEST594: metadata_json lifecycle — set, get, clear | cap/definition_test.go:306 |
| test595 | `Test595_with_args_constructor` | TEST595: with_args constructor stores args correctly | cap/definition_test.go:325 |
| test596 | `Test596_with_full_definition_constructor` | TEST596: with_full_definition constructor stores all fields | cap/definition_test.go:344 |
| test597 | `Test597_cap_arg_with_full_definition` | TEST597: CapArg::with_full_definition stores all fields including optional ones | cap/definition_test.go:374 |
| test598 | `Test598_cap_output_lifecycle` | TEST598: CapOutput lifecycle — set_output, set/clear metadata | cap/definition_test.go:399 |
| test599 | `Test599_is_empty` | TEST599: is_empty returns true for empty response, false for non-empty | cap/response_test.go:279 |
| test600 | `Test600_size` | TEST600: size returns exact byte count for all content types | cap/response_test.go:294 |
| test601 | `Test601_get_content_type` | TEST601: get_content_type returns correct MIME type for each variant | cap/response_test.go:309 |
| test602 | `Test602_as_type_binary_error` | TEST602: as_type on binary response returns error (cannot deserialize binary) | cap/response_test.go:321 |
| test603 | `Test603_as_bool_edge_cases` | TEST603: as_bool handles all accepted truthy/falsy variants and rejects garbage | cap/response_test.go:330 |
| test605 | `Test605_all_coercion_paths_build_valid_urns` | TEST605: all_coercion_paths each entry builds a valid parseable CapUrn | standard/caps_test.go:81 |
| test606 | `Test606_coercion_urn_specs` | TEST606: coercion_urn in/out specs match the type's media URN constant | standard/caps_test.go:98 |
| test607 | `Test607_media_urns_for_extension_unknown` | TEST607: media_urns_for_extension returns error for unknown extension | media/spec_test.go:507 |
| test608 | `Test608_media_urns_for_extension_populated` | TEST608: media_urns_for_extension returns URNs after adding a spec with extensions | media/spec_test.go:517 |
| test609 | `Test609_get_extension_mappings` | TEST609: get_extension_mappings returns all registered extension->URN pairs | media/spec_test.go:548 |
| test610 | `Test610_get_cached_spec` | TEST610: get_cached_spec returns None for unknown and Some for known | media/spec_test.go:575 |
| test611 | `Test611_is_embedded_profile_comprehensive` | TEST611: is_embedded_profile recognizes all 9 embedded profiles and rejects non-embedded | media/profile_test.go:18 |
| test612 | `Test612_clear_cache` | TEST612: clear_cache empties all in-memory schemas | media/profile_test.go:34 |
| test613 | `Test613_validate_cached` | TEST613: validate_cached validates against cached standard schemas | media/profile_test.go:42 |
| test614 | `Test614_registry_creation` | TEST614: Verify registry creation succeeds and cache directory exists | media/spec_test.go:595 |
| test615 | `Test615_cache_key_generation` | TEST615: Verify cache key generation is deterministic and distinct for different URNs | media/spec_test.go:602 |
| test616 | `Test616_stored_media_spec_to_def` | TEST616: Verify StoredMediaSpec converts to MediaSpecDef preserving all fields | media/spec_test.go:615 |
| test617 | `Test617_normalize_media_urn` | TEST617: Verify normalize_media_urn produces consistent non-empty results | media/spec_test.go:634 |
| test618 | `Test618_registry_creation` | TEST618: Verify profile schema registry creation succeeds with temp cache | media/profile_test.go:61 |
| test619 | `Test619_embedded_schemas_loaded` | TEST619: Verify all 9 embedded standard schemas are loaded on creation | media/profile_test.go:68 |
| test620 | `Test620_string_validation` | TEST620: Verify string schema validates strings and rejects non-strings | media/profile_test.go:80 |
| test621 | `Test621_integer_validation` | TEST621: Verify integer schema validates integers and rejects floats and strings | media/profile_test.go:87 |
| test622 | `Test622_number_validation` | TEST622: Verify number schema validates integers and floats, rejects strings | media/profile_test.go:95 |
| test623 | `Test623_boolean_validation` | TEST623: Verify boolean schema validates true/false and rejects string "true" | media/profile_test.go:103 |
| test624 | `Test624_object_validation` | TEST624: Verify object schema validates objects and rejects arrays | media/profile_test.go:111 |
| test625 | `Test625_string_array_validation` | TEST625: Verify string array schema validates string arrays and rejects mixed arrays | media/profile_test.go:118 |
| test626 | `Test626_unknown_profile_skips_validation` | TEST626: Verify unknown profile URL skips validation and returns Ok | media/profile_test.go:126 |
| test627 | `Test627_is_embedded_profile` | TEST627: Verify is_embedded_profile recognizes standard and rejects custom URLs | media/profile_test.go:133 |
| test628 | `Test628_media_urn_constants_format` | TEST628: Verify media URN constants all start with "media:" prefix | urn/media_urn_test.go:734 |
| test629 | `Test629_profile_constants_format` | TEST629: Verify profile URL constants all start with capdag.com schema prefix | media/spec_test.go:643 |
| test630 | `Test630_cartridge_repo_creation` | TEST630: Verify CartridgeRepo creation starts with empty cartridge list | bifaci/cartridge_repo_test.go:640 |
| test631 | `Test631_needs_sync_empty_cache` | TEST631: Verify needs_sync returns true with empty cache and non-empty URLs | bifaci/cartridge_repo_test.go:648 |
| test638 | `Test638_no_peer_router_rejects_all` | TEST638: Verify NoPeerRouter rejects all requests with PeerInvokeNotSupported | bifaci/router_test.go:12 |
| test639 | `Test639_wildcard_empty_cap_defaults_to_media_wildcard` | TEST639: cap: (empty) defaults to in=media:;out=media: | urn/cap_urn_test.go:1170 |
| test640 | `Test640_wildcard_002_in_only_defaults_out_to_media` | TEST640: cap:in defaults out to media: | urn/cap_urn_test.go:1496 |
| test641 | `Test641_wildcard_003_out_only_defaults_in_to_media` | TEST641: cap:out defaults in to media: | urn/cap_urn_test.go:1504 |
| test642 | `Test642_wildcard_004_in_out_no_values_become_media` | TEST642: cap:in;out both become media: | urn/cap_urn_test.go:1512 |
| test643 | `Test643_wildcard_005_explicit_asterisk_becomes_media` | TEST643: cap:in=*;out=* becomes media: | urn/cap_urn_test.go:1520 |
| test644 | `Test644_wildcard_006_specific_in_wildcard_out` | TEST644: cap:in=media:;out=* has specific in, wildcard out | urn/cap_urn_test.go:1528 |
| test645 | `Test645_wildcard_007_wildcard_in_specific_out` | TEST645: cap:in=*;out=media:text has wildcard in, specific out | urn/cap_urn_test.go:1536 |
| test646 | `Test646_wildcard_008_invalid_in_spec_fails` | TEST646: cap:in=foo fails (invalid media URN) | urn/cap_urn_test.go:1544 |
| test647 | `Test647_wildcard_009_invalid_out_spec_fails` | TEST647: cap:in=media:;out=bar fails (invalid media URN) | urn/cap_urn_test.go:1550 |
| test648 | `Test648_wildcard_accepts_specific` | TEST648: Wildcard in/out match specific caps | urn/cap_urn_test.go:1180 |
| test649 | `Test649_specificity_scoring` | TEST649: Specificity - wildcard has 0, specific has tag count | urn/cap_urn_test.go:1195 |
| test650 | `Test650_wildcard_012_preserve_other_tags` | TEST650: cap:in;out;op=test preserves other tags | urn/cap_urn_test.go:1556 |
| test651 | `Test651_identity_forms_equivalent` | TEST651: All identity forms produce the same CapUrn | urn/cap_urn_test.go:1207 |
| test652 | `Test652_cap_identity_constant_works` | TEST652: CAP_IDENTITY constant matches identity caps regardless of string form | urn/cap_urn_test.go:1221 |
| test653 | `Test653_identity_routing_isolation` | TEST653: Identity (no tags) does not match specific requests via routing | urn/cap_urn_test.go:1236 |
| test667 | `Test667_verify_chunk_checksum_detects_corruption` | TEST667: verify_chunk_checksum detects corrupted payload | bifaci/frame_test.go:824 |
| test668 | `Test668_ResolveSlotWithPopulatedByteSlotValues` | TEST668: resolve_binding returns byte values when slot is populated with data | planner/argument_binding_test.go:20 |
| test669 | `Test669_ResolveSlotFallsBackToDefault` | TEST669: resolve_binding falls back to cap default value when slot has no data | planner/argument_binding_test.go:48 |
| test670 | `Test670_ResolveRequiredSlotNoValueReturnsErr` | TEST670: resolve_binding returns error when required slot has no value and no default | planner/argument_binding_test.go:69 |
| test671 | `Test671_ResolveOptionalSlotNoValueReturnsNone` | TEST671: resolve_binding returns None when optional slot has no value and no default | planner/argument_binding_test.go:83 |
| test678 | `Test678_find_stream_equivalent_urn` | TEST678: find_stream with exact equivalent URN (same tags, different order) succeeds | bifaci/cartridge_runtime_test.go:2539 |
| test679 | `Test679_find_stream_base_vs_full_fails` | TEST679: find_stream with base URN vs full URN fails — is_equivalent is strict This is the root cause of the cartridge_client.rs bug. Sender sent "media:llm-generation-request" but receiver looked for "media:llm-generation-request;json;record". | bifaci/cartridge_runtime_test.go:2556 |
| test680 | `Test680_require_stream_missing_fails` | TEST680: require_stream with missing URN returns hard StreamError | bifaci/cartridge_runtime_test.go:2567 |
| test681 | `Test681_find_stream_multiple` | TEST681: find_stream with multiple streams returns the correct one | bifaci/cartridge_runtime_test.go:2581 |
| test682 | `Test682_require_stream_returns_data` | TEST682: require_stream_str returns UTF-8 string for text data | bifaci/cartridge_runtime_test.go:2597 |
| test683 | `Test683_find_stream_invalid_urn_returns_nil` | TEST683: find_stream returns None for invalid media URN string (not a parse error — just None) | bifaci/cartridge_runtime_test.go:2611 |
| test823 | `Test823_dispatch_exact_match` | TEST823: is_dispatchable — exact match provider dispatches request | urn/cap_urn_test.go:1256 |
| test824 | `Test824_dispatch_contravariant_input` | TEST824: is_dispatchable — provider with broader input handles specific request (contravariance) | urn/cap_urn_test.go:1265 |
| test825 | `Test825_dispatch_request_unconstrained_input` | TEST825: is_dispatchable — request with unconstrained input dispatches to specific provider media: on the request input axis means "unconstrained" — vacuously true | urn/cap_urn_test.go:1274 |
| test826 | `Test826_dispatch_covariant_output` | TEST826: is_dispatchable — provider output must satisfy request output (covariance) | urn/cap_urn_test.go:1284 |
| test827 | `Test827_dispatch_generic_output_fails` | TEST827: is_dispatchable — provider with generic output cannot satisfy specific request | urn/cap_urn_test.go:1294 |
| test828 | `Test828_dispatch_wildcard_requires_tag_presence` | TEST828: is_dispatchable — wildcard * tag in request, provider missing tag → reject | urn/cap_urn_test.go:1304 |
| test829 | `Test829_dispatch_wildcard_with_tag_present` | TEST829: is_dispatchable — wildcard * tag in request, provider has tag → accept | urn/cap_urn_test.go:1314 |
| test830 | `Test830_dispatch_provider_extra_tags` | TEST830: is_dispatchable — provider extra tags are refinement, always OK | urn/cap_urn_test.go:1324 |
| test831 | `Test831_dispatch_cross_backend_mismatch` | TEST831: is_dispatchable — cross-backend mismatch prevented | urn/cap_urn_test.go:1334 |
| test832 | `Test832_dispatch_asymmetric` | TEST832: is_dispatchable is NOT symmetric | urn/cap_urn_test.go:1344 |
| test833 | `Test833_comparable_symmetric` | TEST833: is_comparable — both directions checked | urn/cap_urn_test.go:1354 |
| test834 | `Test834_comparable_unrelated` | TEST834: is_comparable — unrelated caps are NOT comparable | urn/cap_urn_test.go:1364 |
| test835 | `Test835_equivalent_identical` | TEST835: is_equivalent — identical caps | urn/cap_urn_test.go:1374 |
| test836 | `Test836_equivalent_non_equivalent` | TEST836: is_equivalent — non-equivalent comparable caps | urn/cap_urn_test.go:1384 |
| test837 | `Test837_dispatch_op_mismatch` | TEST837: is_dispatchable — op tag mismatch rejects | urn/cap_urn_test.go:1394 |
| test838 | `Test838_dispatch_request_wildcard_output` | TEST838: is_dispatchable — request with wildcard output accepts any provider output | urn/cap_urn_test.go:1403 |
| test839 | `Test839_peer_response_delivers_logs_before_stream_start` | TEST839: LOG frames arriving BEFORE StreamStart are delivered immediately This tests the critical fix: during a peer call, the peer (e.g., modelcartridge) sends LOG frames for minutes during model download BEFORE sending any data (StreamStart + Chunk). The handler must receive these LOGs in real-time so it can re-emit progress and keep the engine's activity timer alive. Previously, demux_single_stream blocked on awaiting StreamStart before returning PeerResponse, which meant the handler couldn't call recv() until data arrived — causing 120s activity timeouts during long downloads. | bifaci/cartridge_runtime_test.go:2392 |
| test840 | `Test840_peer_response_collect_bytes_discards_logs` | TEST840: PeerResponse::collect_bytes discards LOG frames | bifaci/cartridge_runtime_test.go:2460 |
| test841 | `Test841_peer_response_collect_value_discards_logs` | TEST841: PeerResponse::collect_value discards LOG frames | bifaci/cartridge_runtime_test.go:2488 |
| test842 | `Test842_progress_sender_emits_frames` | TEST842: run_with_keepalive returns closure result (fast operation, no keepalive frames) | bifaci/cartridge_runtime_test.go:2636 |
| test843 | `Test843_progress_sender_from_goroutine` | TEST843: run_with_keepalive returns Ok/Err from closure | bifaci/cartridge_runtime_test.go:2682 |
| test844 | `Test844_progress_sender_multiple_goroutines` | TEST844: run_with_keepalive propagates errors from closure | bifaci/cartridge_runtime_test.go:2715 |
| test845 | `Test845_progress_sender_independent_of_emitter` | TEST845: ProgressSender emits progress and log frames independently of OutputStream | bifaci/cartridge_runtime_test.go:2760 |
| test846 | `Test846_progress_frame_roundtrip` | TEST846: Test progress LOG frame encode/decode roundtrip preserves progress float | bifaci/io_test.go:1055 |
| test847 | `Test847_progress_double_roundtrip` | TEST847: Double roundtrip (modelcartridge → relay → candlecartridge) | bifaci/io_test.go:1108 |
| test848 | `Test848_relay_notify_roundtrip` | TEST848: RelayNotify encode/decode roundtrip preserves manifest and limits | bifaci/io_test.go:886 |
| test849 | `Test849_relay_state_roundtrip` | TEST849: RelayState encode/decode roundtrip preserves resource payload | bifaci/io_test.go:927 |
| test850 | `Test850_all_format_conversion_paths_build_valid_urns` | TEST850: all_format_conversion_paths each entry builds a valid parseable CapUrn | standard/caps_test.go:111 |
| test851 | `Test851_format_conversion_urn_specs` | TEST851: format_conversion_urn in/out specs match the input constants | standard/caps_test.go:125 |
| test852 | `Test852_lub_identical` | TEST852: LUB of identical URNs returns the same URN | urn/media_urn_test.go:648 |
| test853 | `Test853_lub_no_common_tags` | TEST853: LUB of URNs with no common tags returns media: (universal) | urn/media_urn_test.go:656 |
| test854 | `Test854_lub_partial_overlap` | TEST854: LUB keeps common tags, drops differing ones | urn/media_urn_test.go:668 |
| test855 | `Test855_lub_list_vs_scalar` | TEST855: LUB of list and non-list drops list tag | urn/media_urn_test.go:680 |
| test856 | `Test856_lub_empty` | TEST856: LUB of empty input returns universal type | urn/media_urn_test.go:692 |
| test857 | `Test857_lub_single` | TEST857: LUB of single input returns that input | urn/media_urn_test.go:700 |
| test858 | `Test858_lub_three_inputs` | TEST858: LUB with three+ inputs narrows correctly | urn/media_urn_test.go:708 |
| test859 | `Test859_lub_valued_tags` | TEST859: LUB with valued tags (non-marker) that differ | urn/media_urn_test.go:722 |
| test860 | `Test860_seq_assigner_same_rid_different_xids_independent` | TEST860: Same RID with different XIDs get independent seq counters | bifaci/frame_test.go:1008 |
| test890 | `Test890_direction_semantic_matching` | TEST890: Semantic direction matching - generic provider matches specific request | urn/cap_urn_test.go:923 |
| test891 | `Test891_direction_semantic_specificity` | TEST891: Semantic direction specificity - more media URN tags = higher specificity | urn/cap_urn_test.go:980 |
| test892 | `Test892_extensions_serialization` | TEST892: Test extensions serializes/deserializes correctly in MediaSpecDef | media/spec_test.go:422 |
| test893 | `Test893_extensions_with_metadata_and_validation` | TEST893: Test extensions can coexist with metadata and validation | media/spec_test.go:447 |
| test894 | `Test894_multiple_extensions` | TEST894: Test multiple extensions in a media spec | media/spec_test.go:480 |
| test920 | `Test920_cap_documentation_roundtrip` | TEST920: Tests creation of a simple execution plan with a single capability Verifies that single_cap() generates a valid plan with input_slot, cap node, and output node | cap/definition_test.go:514 |
| test921 | `Test921_cap_documentation_omitted_when_nil` | TEST921: Tests creation of a linear chain of capabilities connected in sequence Verifies that linear_chain() correctly links multiple caps with proper edges and topological order | cap/definition_test.go:532 |
| test922 | `Test922_cap_documentation_parses_from_json` | TEST922: Tests creation and validation of an empty execution plan with no nodes Verifies that plans without capabilities are valid and handle zero nodes correctly | cap/definition_test.go:544 |
| test923 | `Test923_cap_documentation_lifecycle` | TEST923: Tests storing and retrieving metadata attached to an execution plan Verifies that arbitrary JSON metadata can be associated with a plan for context preservation | cap/definition_test.go:563 |
| test976 | `Test976_cap_graph_find_best_path` | TEST976: CapGraph::find_best_path returns highest-specificity path over shortest | cap_matrix_test.go:1091 |
| test1105 | `Test1105_TwoStepsSameCapUrnDifferentSlotValues` | TEST1105: Two steps with the same cap_urn get distinct slot values via different node_ids. This is the core disambiguation scenario that step-index keying was designed to solve. | planner/argument_binding_test.go:96 |
| test1106 | `Test1106_SlotFallsThroughToCapSettingsShared` | TEST1106: Slot resolution falls through to cap_settings when no slot_value exists. cap_settings are keyed by cap_urn (shared across steps), so both steps get the same value. | planner/argument_binding_test.go:138 |
| test1107 | `Test1107_SlotValueOverridesCapSettingsPerStep` | TEST1107: step_0 has a slot_value override, step_1 falls through to cap_settings. Proves per-step override works while shared settings remain as fallback. | planner/argument_binding_test.go:174 |
| test1108 | `Test1108_ResolveAllPassesNodeID` | TEST1108: ResolveAll with node_id threads correctly through to each binding. | planner/argument_binding_test.go:216 |
| test1109 | `Test1109_SlotKeyUsesNodeIDNotCapUrn` | TEST1109: Slot key uses node_id, NOT cap_urn — a slot_value keyed by cap_urn must not match. | planner/argument_binding_test.go:267 |
| test1142 | `Test1142_resolved_graph_to_mermaid_renders_shapes_dedupes_edges_and_escapes` | TEST1142: ResolvedGraph.to_mermaid() renders node shapes, deduplicates edges, and escapes labels | orchestrator/orchestrator_test.go:38 |
| test1155 | `Test1155_FromStrandProducesSingleStrandMachine` | TEST1155: Building a machine from one strand produces one strand with one resolved edge. | machine/machine_test.go:152 |
| test1156 | `Test1156_FromStrandsKeepStrandsDisjoint` | TEST1156: Building from multiple strands keeps them disjoint and preserves input strand order. | machine/machine_test.go:169 |
| test1157 | `Test1157_FromStrandsEmptyInputFailsHard` | TEST1157: Building from zero strands fails with NoCapabilitySteps. | machine/machine_test.go:196 |
| test1158 | `Test1158_MachineIsEquivalentIsStrictPositional` | TEST1158: Machine equivalence is strict about strand order and rejects reordered strands. | machine/machine_test.go:210 |
| test1159 | `Test1159_MachineStrandIsEquivalentWalksNodeBijection` | TEST1159: MachineStrand equivalence accepts two separately built but structurally identical strands. | machine/machine_test.go:234 |
| test1160 | `Test1160_InputOutputAnchors` | TEST1160: Creating a MachineRun stores the canonical notation and starts in the pending state. | machine/machine_test.go:253 |
| test1161 | `Test1161_simple_linear_chain_conversion` | TEST1161: Converting a simple linear plan produces resolved edges for the cap-to-cap chain. | orchestrator/orchestrator_test.go:88 |
| test1163 | `Test1163_ParseSingleStrandTwoCapsConnectedViaSharedNode` | TEST1163: Parsing one connected strand yields a single machine strand with both caps connected by the shared node. | machine/machine_test.go:365 |
| test1164 | `Test1164_ParseTwoDisconnectedStrandsYieldsTwoMachineStrands` | TEST1164: Parsing two disconnected strand definitions yields two separate machine strands. | machine/machine_test.go:395 |
| test1165 | `Test1165_ParseUnknownCapInRegistryReturnsAbstractionError` | TEST1165: Parsing fails hard when a referenced cap is missing from the registry cache. | machine/machine_test.go:498 |
| test1166 | `Test1166_ParseDuplicateAliasReturnsError` | TEST1166: Duplicate header aliases are reported as syntax errors. | machine/machine_test.go:468 |
| test1167 | `Test1167_ParseUndefinedAliasReturnsError` | TEST1167: Wiring that references an undefined alias is reported as a syntax error. | machine/machine_test.go:485 |
| test1168 | `Test1168_ParseNodeNameCollidesWithCapAlias` | TEST1168: Parsing rejects node names that collide with declared cap aliases. | machine/machine_test.go:515 |
| test1169 | `Test1169_ForEachSetsIsLoop` | TEST1169: Loop markers in notation set the resolved edge loop flag on the following cap step. | machine/machine_test.go:288 |
| test1170 | `Test1170_CollectIsElided` | TEST1170: Parsing and then serializing machine notation round-trips to the canonical form. | machine/machine_test.go:325 |
| test1171 | `Test1171_ParseEmptyInputReturnsError` | TEST1171: Empty machine notation is rejected as a syntax error. | machine/machine_test.go:441 |
| test1172 | `Test1172_MachineStringRepr` | TEST1172: Serializing a two-step strand emits the expected aliases and node names. | machine/machine_test.go:576 |
| test1173 | `Test1173_ToMachineNotationRoundTrips` | TEST1173: Serializing and reparsing a machine preserves strict machine equivalence. | machine/machine_test.go:536 |
| test1175 | `Test1175_EmptyMachineSerializesToEmpty` | TEST1175: Serializing an empty machine produces an empty string. | machine/machine_test.go:567 |
| test1187 | `Test1187_StrandNonEquivalenceDifferentCap` | TEST1187: Strand resolution fails when a referenced cap is not found in the registry. | machine/machine_test.go:645 |
| test1189 | `Test1189_StrandEquivalenceWithDifferentNodeAllocationOrders` | TEST1189: Strand resolution keeps canonical anchor ordering stable across equivalent inputs. | machine/machine_test.go:596 |
| | | | |
| unnumbered | `TestArgumentsMultiple` | Mirror-specific coverage: Test multiple arguments are correctly serialized in CBOR payload | bifaci/integration_test.go:1592 |
| unnumbered | `TestArgumentsRoundtrip` | Mirror-specific coverage: Test host call with unified CBOR arguments sends correct content_type and payload | bifaci/integration_test.go:1232 |
| unnumbered | `TestAutoChunkingReassembly` | Mirror-specific coverage: Test auto-chunking splits payload larger than max_chunk into CHUNK frames + END frame, and host concatenated() reassembles the full original data | bifaci/integration_test.go:1665 |
| unnumbered | `TestCacheOperations` |  | cap/registry_test.go:76 |
| unnumbered | `TestCapBlockRegistryManagement` |  | cap_matrix_test.go:466 |
| unnumbered | `TestCapCallerBinaryResponse` |  | cap/caller_test.go:198 |
| unnumbered | `TestCapCallerCall` |  | cap/caller_test.go:121 |
| unnumbered | `TestCapCallerCreation` |  | cap/caller_test.go:43 |
| unnumbered | `TestCapCallerResolveOutputSpec` |  | cap/caller_test.go:59 |
| unnumbered | `TestCapCallerWithArguments` |  | cap/caller_test.go:157 |
| unnumbered | `TestCapDescription` |  | cap/definition_test.go:446 |
| unnumbered | `TestCapExists` |  | cap/registry_test.go:119 |
| unnumbered | `TestCapJSONRoundTrip` |  | cap/definition_test.go:581 |
| unnumbered | `TestCapManifestCompatibility` |  | bifaci/manifest_test.go:302 |
| unnumbered | `TestCapManifestValidation` |  | bifaci/manifest_test.go:267 |
| unnumbered | `TestCapManifestWithPageURL` |  | bifaci/manifest_test.go:61 |
| unnumbered | `TestCapRequestHandling` | Additional existing tests below (not part of TEST108-116 sequence) | cap/definition_test.go:430 |
| unnumbered | `TestCapUrn_JSONSerialization` | JSON serialization test (not numbered in Rust) | urn/cap_urn_test.go:1413 |
| unnumbered | `TestCapValidationCoordinator_EndToEnd` |  | cap/schema_validation_test.go:424 |
| unnumbered | `TestCapWithMediaSpecs` |  | cap/definition_test.go:458 |
| unnumbered | `TestCartridgeErrorResponse` | Mirror-specific coverage: Test cartridge ERR frame is received by host as error | bifaci/integration_test.go:707 |
| unnumbered | `TestCartridgeSuddenDisconnect` | Mirror-specific coverage: Test host receives error when cartridge closes connection unexpectedly | bifaci/integration_test.go:1309 |
| unnumbered | `TestChunkingDataIntegrity3x` | Mirror-specific coverage: Test auto-chunking preserves data integrity across chunk boundaries for 3x max_chunk payload | bifaci/integration_test.go:1919 |
| unnumbered | `TestComplexNestedSchemaValidation` |  | cap/schema_validation_test.go:549 |
| unnumbered | `TestConcatenatedVsFinalPayloadDivergence` | Mirror-specific coverage: Test that concatenated() returns full payload while final_payload() returns only last chunk | bifaci/integration_test.go:1895 |
| unnumbered | `TestConstructor` | Mirror-specific coverage: Test simple constructor creates media URN with type tag | urn/media_urn_test.go:224 |
| unnumbered | `TestCustomMediaUrnResolution` |  | cap/schema_validation_test.go:666 |
| unnumbered | `TestEndFrameNoPayload` | Mirror-specific coverage: Test END frame without payload is handled as complete response with empty data | bifaci/integration_test.go:1417 |
| unnumbered | `TestExactMaxChunkSingleEnd` | Mirror-specific coverage: Test payload exactly equal to max_chunk produces single END frame (no CHUNK frames) | bifaci/integration_test.go:1751 |
| unnumbered | `TestExtractEffectivePayloadInvalidCapUrn` | Mirror-specific coverage: Test extract_effective_payload with invalid cap URN returns CapUrn error | bifaci/cartridge_runtime_test.go:422 |
| unnumbered | `TestFileSchemaResolver_ErrorHandling` |  | cap/schema_validation_test.go:538 |
| unnumbered | `TestHeartbeatDuringStreaming` | Mirror-specific coverage: Test cartridge-initiated heartbeat mid-stream is handled transparently by host | bifaci/integration_test.go:1052 |
| unnumbered | `TestHostInitiatedHeartbeatNoPingPong` | Mirror-specific coverage: Test host does not echo back cartridge's heartbeat response (no infinite ping-pong) | bifaci/integration_test.go:1152 |
| unnumbered | `TestInputValidator_WithSchemaValidation` |  | cap/schema_validation_test.go:310 |
| unnumbered | `TestIntegrationBinaryCapHandling` | TestIntegrationBinaryCapHandling verifies binary cap handling | bifaci/integration_test.go:187 |
| unnumbered | `TestIntegrationCallerAndResponseSystem` | TestIntegrationCallerAndResponseSystem verifies the caller and response system | bifaci/integration_test.go:127 |
| unnumbered | `TestIntegrationCapValidation` | TestIntegrationCapValidation verifies cap schema validation | bifaci/integration_test.go:327 |
| unnumbered | `TestIntegrationCapWithMediaSpecs` | TestIntegrationCapWithMediaSpecs verifies caps with custom media specs | bifaci/integration_test.go:277 |
| unnumbered | `TestIntegrationCaseInsensitiveUrns` | TestIntegrationCaseInsensitiveUrns verifies URNs are case-insensitive | bifaci/integration_test.go:86 |
| unnumbered | `TestIntegrationMediaSpecDefConstruction` | TestIntegrationMediaSpecDefConstruction verifies media.MediaSpecDef construction | bifaci/integration_test.go:417 |
| unnumbered | `TestIntegrationMediaUrnResolution` | TestIntegrationMediaUrnResolution verifies media URN resolution | bifaci/integration_test.go:369 |
| unnumbered | `TestIntegrationTextCapHandling` | TestIntegrationTextCapHandling verifies text cap handling | bifaci/integration_test.go:227 |
| unnumbered | `TestIntegrationVersionlessCapCreation` | TestIntegrationVersionlessCapCreation verifies caps can be created without version fields | bifaci/integration_test.go:58 |
| unnumbered | `TestLogFramesDuringRequest` | Mirror-specific coverage: Test LOG frames sent during a request are transparently skipped by host | bifaci/integration_test.go:764 |
| unnumbered | `TestMaxChunkPlusOneSplitsIntoTwo` | Mirror-specific coverage: Test payload of max_chunk + 1 produces exactly one CHUNK frame + one END frame | bifaci/integration_test.go:1816 |
| unnumbered | `TestMediaUrnResolutionWithMediaSpecs` |  | cap/schema_validation_test.go:627 |
| unnumbered | `TestOutputValidator_WithSchemaValidation` |  | cap/schema_validation_test.go:367 |
| unnumbered | `TestParseHeadersWithNoWiringsReturnsNoEdgesError` | TestParseHeadersWithNoWiringsReturnsNoEdgesError verifies the ErrNoEdges case. | machine/machine_test.go:453 |
| unnumbered | `TestParseSimple` | Mirror-specific coverage: Test parsing simple media URN verifies correct structure with no version, subtype, or profile | urn/media_urn_test.go:14 |
| unnumbered | `TestParseWithProfile` | Mirror-specific coverage: Test parsing media URN with profile extracts profile URL correctly | urn/media_urn_test.go:32 |
| unnumbered | `TestParseWithSubtype` | Mirror-specific coverage: Test parsing media URN with marker tags works correctly | urn/media_urn_test.go:22 |
| unnumbered | `TestRegistryGetCap` |  | cap/registry_test.go:49 |
| unnumbered | `TestRegistryValidation` |  | cap/registry_test.go:62 |
| unnumbered | `TestRequestAfterShutdown` | Mirror-specific coverage: Test host request on a closed host returns error | bifaci/integration_test.go:1549 |
| unnumbered | `TestResponseWrapperAsBool` |  | cap/response_test.go:108 |
| unnumbered | `TestResponseWrapperAsFloat` |  | cap/response_test.go:94 |
| unnumbered | `TestResponseWrapperFromText` |  | cap/response_test.go:44 |
| unnumbered | `TestResponseWrapperGetContentType` |  | cap/response_test.go:148 |
| unnumbered | `TestResponseWrapperIsEmpty` |  | cap/response_test.go:138 |
| unnumbered | `TestResponseWrapperMatchesOutputType` |  | cap/response_test.go:159 |
| unnumbered | `TestResponseWrapperValidateAgainstCap` |  | cap/response_test.go:244 |
| unnumbered | `TestSchemaValidator_ArraySchemaValidation` |  | cap/schema_validation_test.go:259 |
| unnumbered | `TestSchemaValidator_ValidateArgumentWithSchema_NilSchema` | Additional Go-specific coverage: nil schema skips direct schema validation | cap/schema_validation_test.go:115 |
| unnumbered | `TestSchemaValidator_ValidateArguments_Integration` |  | cap/schema_validation_test.go:196 |
| unnumbered | `TestSchemaValidator_ValidateOutputWithSchema_Failure` |  | cap/schema_validation_test.go:165 |
| unnumbered | `TestStdinSourceNilHandling` | TestStdinSourceNilHandling tests that nil StdinSource is handled correctly | cap/caller_test.go:340 |
| unnumbered | `TestStreamingSequenceNumbers` | Mirror-specific coverage: Test streaming response sequence numbers are contiguous and start from 0 | bifaci/integration_test.go:1473 |
| unnumbered | `TestWithSubtypeConstructor` | Mirror-specific coverage: Test with_subtype constructor creates media URN with subtype | urn/media_urn_test.go:231 |
---

## Unnumbered Tests

The following tests are cataloged but do not currently participate in numeric test indexing.

- `TestArgumentsMultiple` — bifaci/integration_test.go:1592
- `TestArgumentsRoundtrip` — bifaci/integration_test.go:1232
- `TestAutoChunkingReassembly` — bifaci/integration_test.go:1665
- `TestCacheOperations` — cap/registry_test.go:76
- `TestCapBlockRegistryManagement` — cap_matrix_test.go:466
- `TestCapCallerBinaryResponse` — cap/caller_test.go:198
- `TestCapCallerCall` — cap/caller_test.go:121
- `TestCapCallerCreation` — cap/caller_test.go:43
- `TestCapCallerResolveOutputSpec` — cap/caller_test.go:59
- `TestCapCallerWithArguments` — cap/caller_test.go:157
- `TestCapDescription` — cap/definition_test.go:446
- `TestCapExists` — cap/registry_test.go:119
- `TestCapJSONRoundTrip` — cap/definition_test.go:581
- `TestCapManifestCompatibility` — bifaci/manifest_test.go:302
- `TestCapManifestValidation` — bifaci/manifest_test.go:267
- `TestCapManifestWithPageURL` — bifaci/manifest_test.go:61
- `TestCapRequestHandling` — cap/definition_test.go:430
- `TestCapUrn_JSONSerialization` — urn/cap_urn_test.go:1413
- `TestCapValidationCoordinator_EndToEnd` — cap/schema_validation_test.go:424
- `TestCapWithMediaSpecs` — cap/definition_test.go:458
- `TestCartridgeErrorResponse` — bifaci/integration_test.go:707
- `TestCartridgeSuddenDisconnect` — bifaci/integration_test.go:1309
- `TestChunkingDataIntegrity3x` — bifaci/integration_test.go:1919
- `TestComplexNestedSchemaValidation` — cap/schema_validation_test.go:549
- `TestConcatenatedVsFinalPayloadDivergence` — bifaci/integration_test.go:1895
- `TestConstructor` — urn/media_urn_test.go:224
- `TestCustomMediaUrnResolution` — cap/schema_validation_test.go:666
- `TestEndFrameNoPayload` — bifaci/integration_test.go:1417
- `TestExactMaxChunkSingleEnd` — bifaci/integration_test.go:1751
- `TestExtractEffectivePayloadInvalidCapUrn` — bifaci/cartridge_runtime_test.go:422
- `TestFileSchemaResolver_ErrorHandling` — cap/schema_validation_test.go:538
- `TestHeartbeatDuringStreaming` — bifaci/integration_test.go:1052
- `TestHostInitiatedHeartbeatNoPingPong` — bifaci/integration_test.go:1152
- `TestInputValidator_WithSchemaValidation` — cap/schema_validation_test.go:310
- `TestIntegrationBinaryCapHandling` — bifaci/integration_test.go:187
- `TestIntegrationCallerAndResponseSystem` — bifaci/integration_test.go:127
- `TestIntegrationCapValidation` — bifaci/integration_test.go:327
- `TestIntegrationCapWithMediaSpecs` — bifaci/integration_test.go:277
- `TestIntegrationCaseInsensitiveUrns` — bifaci/integration_test.go:86
- `TestIntegrationMediaSpecDefConstruction` — bifaci/integration_test.go:417
- `TestIntegrationMediaUrnResolution` — bifaci/integration_test.go:369
- `TestIntegrationTextCapHandling` — bifaci/integration_test.go:227
- `TestIntegrationVersionlessCapCreation` — bifaci/integration_test.go:58
- `TestLogFramesDuringRequest` — bifaci/integration_test.go:764
- `TestMaxChunkPlusOneSplitsIntoTwo` — bifaci/integration_test.go:1816
- `TestMediaUrnResolutionWithMediaSpecs` — cap/schema_validation_test.go:627
- `TestOutputValidator_WithSchemaValidation` — cap/schema_validation_test.go:367
- `TestParseHeadersWithNoWiringsReturnsNoEdgesError` — machine/machine_test.go:453
- `TestParseSimple` — urn/media_urn_test.go:14
- `TestParseWithProfile` — urn/media_urn_test.go:32
- `TestParseWithSubtype` — urn/media_urn_test.go:22
- `TestRegistryGetCap` — cap/registry_test.go:49
- `TestRegistryValidation` — cap/registry_test.go:62
- `TestRequestAfterShutdown` — bifaci/integration_test.go:1549
- `TestResponseWrapperAsBool` — cap/response_test.go:108
- `TestResponseWrapperAsFloat` — cap/response_test.go:94
- `TestResponseWrapperFromText` — cap/response_test.go:44
- `TestResponseWrapperGetContentType` — cap/response_test.go:148
- `TestResponseWrapperIsEmpty` — cap/response_test.go:138
- `TestResponseWrapperMatchesOutputType` — cap/response_test.go:159
- `TestResponseWrapperValidateAgainstCap` — cap/response_test.go:244
- `TestSchemaValidator_ArraySchemaValidation` — cap/schema_validation_test.go:259
- `TestSchemaValidator_ValidateArgumentWithSchema_NilSchema` — cap/schema_validation_test.go:115
- `TestSchemaValidator_ValidateArguments_Integration` — cap/schema_validation_test.go:196
- `TestSchemaValidator_ValidateOutputWithSchema_Failure` — cap/schema_validation_test.go:165
- `TestStdinSourceNilHandling` — cap/caller_test.go:340
- `TestStreamingSequenceNumbers` — bifaci/integration_test.go:1473
- `TestWithSubtypeConstructor` — urn/media_urn_test.go:231

---

*Generated from CapDag-Go source tree*
*Total tests: 654*
*Total numbered tests: 586*
*Total unnumbered tests: 68*
*Total numbered tests missing descriptions: 0*
*Total numbering mismatches: 0*
