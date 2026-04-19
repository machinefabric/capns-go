package planner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TEST767: Tests ArgumentInfo struct fields — name, media_urn, resolution, description
// Verifies that argument metadata is correctly stored and accessible
func Test767_argument_info_fields(t *testing.T) {
	argInfo := &ArgumentInfo{
		Name:        "width",
		MediaUrn:    "media:integer",
		Description: "Width in pixels",
		Resolution:  ResolutionHasDefault,
		DefaultValue: 200,
		IsRequired:  false,
		Validation:  map[string]any{"min": 50, "max": 2000},
	}

	assert.Equal(t, "width", argInfo.Name)
	assert.Equal(t, ResolutionHasDefault, argInfo.Resolution)
	assert.Equal(t, "has_default", argInfo.Resolution.String())
	assert.Equal(t, 200, argInfo.DefaultValue)
}

// TEST768: Tests PathArgumentRequirements structure for single-step execution paths
// Verifies that argument requirements are correctly organized by step with resolution information
func Test768_path_argument_requirements_structure(t *testing.T) {
	requirements := &PathArgumentRequirements{
		SourceSpec: "media:pdf",
		TargetSpec: "media:png",
		Steps: []*StepArgumentRequirements{
			{
				CapUrn:    "cap:op=generate_thumbnail;in=pdf;out=png",
				StepIndex: 0,
				Title:     "Generate Thumbnail",
				Arguments: []*ArgumentInfo{
					{
						Name:       "file_path",
						MediaUrn:   "media:string",
						Description: "Path to file",
						Resolution: ResolutionFromInputFile,
						IsRequired: true,
					},
				},
				Slots: []*ArgumentInfo{},
			},
		},
		CanExecuteWithoutInput: true,
	}

	assert.True(t, requirements.CanExecuteWithoutInput)
	require.Equal(t, 1, len(requirements.Steps))
	assert.Equal(t, 0, len(requirements.Steps[0].Slots))
	assert.Equal(t, ResolutionFromInputFile, requirements.Steps[0].Arguments[0].Resolution)
}

// TEST769: Tests PathArgumentRequirements tracking of required user-input slots
// Verifies that arguments requiring user input are collected in slots and CanExecuteWithoutInput is false
func Test769_path_with_required_slot(t *testing.T) {
	targetLanguageArg := &ArgumentInfo{
		Name:        "target_language",
		MediaUrn:    "media:string",
		Description: "Target language code",
		Resolution:  ResolutionRequiresUserInput,
		IsRequired:  true,
	}

	requirements := &PathArgumentRequirements{
		SourceSpec: "media:text",
		TargetSpec: "media:translated",
		Steps: []*StepArgumentRequirements{
			{
				CapUrn:    "cap:op=translate;in=text;out=translated",
				StepIndex: 0,
				Title:     "Translate",
				Arguments: []*ArgumentInfo{
					{
						Name:       "file_path",
						MediaUrn:   "media:string",
						Description: "Path to file",
						Resolution: ResolutionFromInputFile,
						IsRequired: true,
					},
					targetLanguageArg,
				},
				Slots: []*ArgumentInfo{targetLanguageArg},
			},
		},
		CanExecuteWithoutInput: false,
	}

	assert.False(t, requirements.CanExecuteWithoutInput)
	require.Equal(t, 1, len(requirements.Steps[0].Slots))
	assert.Equal(t, "target_language", requirements.Steps[0].Slots[0].Name)
}
