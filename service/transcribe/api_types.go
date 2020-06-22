// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Settings for content redaction within a transcription job.
type ContentRedaction struct {
	_ struct{} `type:"structure"`

	// The output transcript file stored in either the default S3 bucket or in a
	// bucket you specify.
	//
	// When you choose redacted Amazon Transcribe outputs only the redacted transcript.
	//
	// When you choose redacted_and_unredacted Amazon Transcribe outputs both the
	// redacted and unredacted transcripts.
	//
	// RedactionOutput is a required field
	RedactionOutput RedactionOutput `type:"string" required:"true" enum:"true"`

	// Request parameter that defines the entities to be redacted. The only accepted
	// value is PII.
	//
	// RedactionType is a required field
	RedactionType RedactionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ContentRedaction) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ContentRedaction) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ContentRedaction"}
	if len(s.RedactionOutput) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RedactionOutput"))
	}
	if len(s.RedactionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("RedactionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Provides information about when a transcription job should be executed.
type JobExecutionSettings struct {
	_ struct{} `type:"structure"`

	// Indicates whether a job should be queued by Amazon Transcribe when the concurrent
	// execution limit is exceeded. When the AllowDeferredExecution field is true,
	// jobs are queued and executed when the number of executing jobs falls below
	// the concurrent execution limit. If the field is false, Amazon Transcribe
	// returns a LimitExceededException exception.
	//
	// If you specify the AllowDeferredExecution field, you must specify the DataAccessRoleArn
	// field.
	AllowDeferredExecution *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of a role that has access to the S3 bucket
	// that contains the input files. Amazon Transcribe assumes this role to read
	// queued media files. If you have specified an output S3 bucket for the transcription
	// results, this role should have access to the output bucket as well.
	//
	// If you specify the AllowDeferredExecution field, you must specify the DataAccessRoleArn
	// field.
	DataAccessRoleArn *string `type:"string"`
}

// String returns the string representation
func (s JobExecutionSettings) String() string {
	return awsutil.Prettify(s)
}

// Describes the input media file in a transcription request.
type Media struct {
	_ struct{} `type:"structure"`

	// The S3 object location of the input media file. The URI must be in the same
	// region as the API endpoint that you are calling. The general form is:
	//
	// s3://<bucket-name>/<keyprefix>/<objectkey>
	//
	// For example:
	//
	// s3://examplebucket/example.mp4
	//
	// s3://examplebucket/mediadocs/example.mp4
	//
	// For more information about S3 object names, see Object Keys (http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide.
	MediaFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Media) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Media) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Media"}
	if s.MediaFileUri != nil && len(*s.MediaFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MediaFileUri", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Identifies the location of a medical transcript.
type MedicalTranscript struct {
	_ struct{} `type:"structure"`

	// The S3 object location of the medical transcript.
	//
	// Use this URI to access the medical transcript. This URI points to the S3
	// bucket you created to store the medical transcript.
	TranscriptFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s MedicalTranscript) String() string {
	return awsutil.Prettify(s)
}

// The data structure that containts the information for a medical transcription
// job.
type MedicalTranscriptionJob struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the job was completed.
	CompletionTime *time.Time `type:"timestamp"`

	// A timestamp that shows when the job was created.
	CreationTime *time.Time `type:"timestamp"`

	// If the TranscriptionJobStatus field is FAILED, this field contains information
	// about why the job failed.
	//
	// The FailureReason field contains one of the following values:
	//
	//    * Unsupported media format- The media format specified in the MediaFormat
	//    field of the request isn't valid. See the description of the MediaFormat
	//    field for a list of valid values.
	//
	//    * The media format provided does not match the detected media format-
	//    The media format of the audio file doesn't match the format specified
	//    in the MediaFormat field in the request. Check the media format of your
	//    media file and make sure the two values match.
	//
	//    * Invalid sample rate for audio file- The sample rate specified in the
	//    MediaSampleRateHertz of the request isn't valid. The sample rate must
	//    be between 8000 and 48000 Hertz.
	//
	//    * The sample rate provided does not match the detected sample rate- The
	//    sample rate in the audio file doesn't match the sample rate specified
	//    in the MediaSampleRateHertz field in the request. Check the sample rate
	//    of your media file and make sure that the two values match.
	//
	//    * Invalid file size: file size too large- The size of your audio file
	//    is larger than what Amazon Transcribe Medical can process. For more information,
	//    see Guidlines and Quotas (https://docs.aws.amazon.com/transcribe/latest/dg/limits-guidelines.html#limits)
	//    in the Amazon Transcribe Medical Guide
	//
	//    * Invalid number of channels: number of channels too large- Your audio
	//    contains more channels than Amazon Transcribe Medical is configured to
	//    process. To request additional channels, see Amazon Transcribe Medical
	//    Endpoints and Quotas (https://docs.aws.amazon.com/general/latest/gr/transcribe-medical.html)
	//    in the Amazon Web Services General Reference
	FailureReason *string `type:"string"`

	// The language code for the language spoken in the source audio file. US English
	// (en-US) is the only supported language for medical transcriptions. Any other
	// value you enter for language code results in a BadRequestException error.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Describes the input media file in a transcription request.
	Media *Media `type:"structure"`

	// The format of the input media file.
	MediaFormat MediaFormat `type:"string" enum:"true"`

	// The sample rate, in Hertz, of the source audio containing medical information.
	//
	// If you don't specify the sample rate, Amazon Transcribe Medical determines
	// it for you. If you choose to specify the sample rate, it must match the rate
	// detected by Amazon Transcribe Medical. In most cases, you should leave the
	// MediaSampleHertz blank and let Amazon Transcribe Medical determine the sample
	// rate.
	MediaSampleRateHertz *int64 `min:"8000" type:"integer"`

	// The name for a given medical transcription job.
	MedicalTranscriptionJobName *string `min:"1" type:"string"`

	// Object that contains object.
	Settings *MedicalTranscriptionSetting `type:"structure"`

	// The medical specialty of any clinicians providing a dictation or having a
	// conversation. PRIMARYCARE is the only available setting for this object.
	// This specialty enables you to generate transcriptions for the following medical
	// fields:
	//
	//    * Family Medicine
	Specialty Specialty `type:"string" enum:"true"`

	// A timestamp that shows when the job started processing.
	StartTime *time.Time `type:"timestamp"`

	// An object that contains the MedicalTranscript. The MedicalTranscript contains
	// the TranscriptFileUri.
	Transcript *MedicalTranscript `type:"structure"`

	// The completion status of a medical transcription job.
	TranscriptionJobStatus TranscriptionJobStatus `type:"string" enum:"true"`

	// The type of speech in the transcription job. CONVERSATION is generally used
	// for patient-physician dialogues. DICTATION is the setting for physicians
	// speaking their notes after seeing a patient. For more information, see how-it-works-med
	Type Type `type:"string" enum:"true"`
}

// String returns the string representation
func (s MedicalTranscriptionJob) String() string {
	return awsutil.Prettify(s)
}

// Provides summary information about a transcription job.
type MedicalTranscriptionJobSummary struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the job was completed.
	CompletionTime *time.Time `type:"timestamp"`

	// A timestamp that shows when the medical transcription job was created.
	CreationTime *time.Time `type:"timestamp"`

	// If the TranscriptionJobStatus field is FAILED, a description of the error.
	FailureReason *string `type:"string"`

	// The language of the transcript in the source audio file.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The name of a medical transcription job.
	MedicalTranscriptionJobName *string `min:"1" type:"string"`

	// Indicates the location of the transcription job's output.
	//
	// The CUSTOMER_BUCKET is the S3 location provided in the OutputBucketName field
	// when the
	OutputLocationType OutputLocationType `type:"string" enum:"true"`

	// The medical specialty of the transcription job. Primary care is the only
	// valid value.
	Specialty Specialty `type:"string" enum:"true"`

	// A timestamp that shows when the job began processing.
	StartTime *time.Time `type:"timestamp"`

	// The status of the medical transcription job.
	TranscriptionJobStatus TranscriptionJobStatus `type:"string" enum:"true"`

	// The speech of the clinician in the input audio.
	Type Type `type:"string" enum:"true"`
}

// String returns the string representation
func (s MedicalTranscriptionJobSummary) String() string {
	return awsutil.Prettify(s)
}

// Optional settings for the StartMedicalTranscriptionJob operation.
type MedicalTranscriptionSetting struct {
	_ struct{} `type:"structure"`

	// Instructs Amazon Transcribe Medical to process each audio channel separately
	// and then merge the transcription output of each channel into a single transcription.
	//
	// Amazon Transcribe Medical also produces a transcription of each item detected
	// on an audio channel, including the start time and end time of the item and
	// alternative transcriptions of item. The alternative transcriptions also come
	// with confidence scores provided by Amazon Transcribe Medical.
	//
	// You can't set both ShowSpeakerLabels and ChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException
	ChannelIdentification *bool `type:"boolean"`

	// The maximum number of alternatives that you tell the service to return. If
	// you specify the MaxAlternatives field, you must set the ShowAlternatives
	// field to true.
	MaxAlternatives *int64 `min:"2" type:"integer"`

	// The maximum number of speakers to identify in the input audio. If there are
	// more speakers in the audio than this number, multiple speakers are identified
	// as a single speaker. If you specify the MaxSpeakerLabels field, you must
	// set the ShowSpeakerLabels field to true.
	MaxSpeakerLabels *int64 `min:"2" type:"integer"`

	// Determines whether alternative transcripts are generated along with the transcript
	// that has the highest confidence. If you set ShowAlternatives field to true,
	// you must also set the maximum number of alternatives to return in the MaxAlternatives
	// field.
	ShowAlternatives *bool `type:"boolean"`

	// Determines whether the transcription job uses speaker recognition to identify
	// different speakers in the input audio. Speaker recongition labels individual
	// speakers in the audio file. If you set the ShowSpeakerLabels field to true,
	// you must also set the maximum number of speaker labels in the MaxSpeakerLabels
	// field.
	//
	// You can't set both ShowSpeakerLabels and ChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException.
	ShowSpeakerLabels *bool `type:"boolean"`
}

// String returns the string representation
func (s MedicalTranscriptionSetting) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MedicalTranscriptionSetting) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MedicalTranscriptionSetting"}
	if s.MaxAlternatives != nil && *s.MaxAlternatives < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxAlternatives", 2))
	}
	if s.MaxSpeakerLabels != nil && *s.MaxSpeakerLabels < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxSpeakerLabels", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Provides optional settings for the StartTranscriptionJob operation.
type Settings struct {
	_ struct{} `type:"structure"`

	// Instructs Amazon Transcribe to process each audio channel separately and
	// then merge the transcription output of each channel into a single transcription.
	//
	// Amazon Transcribe also produces a transcription of each item detected on
	// an audio channel, including the start time and end time of the item and alternative
	// transcriptions of the item including the confidence that Amazon Transcribe
	// has in the transcription.
	//
	// You can't set both ShowSpeakerLabels and ChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException.
	ChannelIdentification *bool `type:"boolean"`

	// The number of alternative transcriptions that the service should return.
	// If you specify the MaxAlternatives field, you must set the ShowAlternatives
	// field to true.
	MaxAlternatives *int64 `min:"2" type:"integer"`

	// The maximum number of speakers to identify in the input audio. If there are
	// more speakers in the audio than this number, multiple speakers are identified
	// as a single speaker. If you specify the MaxSpeakerLabels field, you must
	// set the ShowSpeakerLabels field to true.
	MaxSpeakerLabels *int64 `min:"2" type:"integer"`

	// Determines whether the transcription contains alternative transcriptions.
	// If you set the ShowAlternatives field to true, you must also set the maximum
	// number of alternatives to return in the MaxAlternatives field.
	ShowAlternatives *bool `type:"boolean"`

	// Determines whether the transcription job uses speaker recognition to identify
	// different speakers in the input audio. Speaker recognition labels individual
	// speakers in the audio file. If you set the ShowSpeakerLabels field to true,
	// you must also set the maximum number of speaker labels MaxSpeakerLabels field.
	//
	// You can't set both ShowSpeakerLabels and ChannelIdentification in the same
	// request. If you set both, your request returns a BadRequestException.
	ShowSpeakerLabels *bool `type:"boolean"`

	// Set to mask to remove filtered text from the transcript and replace it with
	// three asterisks ("***") as placeholder text. Set to remove to remove filtered
	// text from the transcript without using placeholder text.
	VocabularyFilterMethod VocabularyFilterMethod `type:"string" enum:"true"`

	// The name of the vocabulary filter to use when transcribing the audio. The
	// filter that you specify must have the same language code as the transcription
	// job.
	VocabularyFilterName *string `min:"1" type:"string"`

	// The name of a vocabulary to use when processing the transcription job.
	VocabularyName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Settings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Settings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Settings"}
	if s.MaxAlternatives != nil && *s.MaxAlternatives < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxAlternatives", 2))
	}
	if s.MaxSpeakerLabels != nil && *s.MaxSpeakerLabels < 2 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxSpeakerLabels", 2))
	}
	if s.VocabularyFilterName != nil && len(*s.VocabularyFilterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFilterName", 1))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Identifies the location of a transcription.
type Transcript struct {
	_ struct{} `type:"structure"`

	// The S3 object location of the redacted transcript.
	//
	// Use this URI to access the redacated transcript. If you specified an S3 bucket
	// in the OutputBucketName field when you created the job, this is the URI of
	// that bucket. If you chose to store the transcript in Amazon Transcribe, this
	// is a shareable URL that provides secure access to that location.
	RedactedTranscriptFileUri *string `min:"1" type:"string"`

	// The S3 object location of the the transcript.
	//
	// Use this URI to access the transcript. If you specified an S3 bucket in the
	// OutputBucketName field when you created the job, this is the URI of that
	// bucket. If you chose to store the transcript in Amazon Transcribe, this is
	// a shareable URL that provides secure access to that location.
	TranscriptFileUri *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Transcript) String() string {
	return awsutil.Prettify(s)
}

// Describes an asynchronous transcription job that was created with the StartTranscriptionJob
// operation.
type TranscriptionJob struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the job was completed.
	CompletionTime *time.Time `type:"timestamp"`

	// An object that describes content redaction settings for the transcription
	// job.
	ContentRedaction *ContentRedaction `type:"structure"`

	// A timestamp that shows when the job was created.
	CreationTime *time.Time `type:"timestamp"`

	// If the TranscriptionJobStatus field is FAILED, this field contains information
	// about why the job failed.
	//
	// The FailureReason field can contain one of the following values:
	//
	//    * Unsupported media format - The media format specified in the MediaFormat
	//    field of the request isn't valid. See the description of the MediaFormat
	//    field for a list of valid values.
	//
	//    * The media format provided does not match the detected media format -
	//    The media format of the audio file doesn't match the format specified
	//    in the MediaFormat field in the request. Check the media format of your
	//    media file and make sure that the two values match.
	//
	//    * Invalid sample rate for audio file - The sample rate specified in the
	//    MediaSampleRateHertz of the request isn't valid. The sample rate must
	//    be between 8000 and 48000 Hertz.
	//
	//    * The sample rate provided does not match the detected sample rate - The
	//    sample rate in the audio file doesn't match the sample rate specified
	//    in the MediaSampleRateHertz field in the request. Check the sample rate
	//    of your media file and make sure that the two values match.
	//
	//    * Invalid file size: file size too large - The size of your audio file
	//    is larger than Amazon Transcribe can process. For more information, see
	//    Limits (https://docs.aws.amazon.com/transcribe/latest/dg/limits-guidelines.html#limits)
	//    in the Amazon Transcribe Developer Guide.
	//
	//    * Invalid number of channels: number of channels too large - Your audio
	//    contains more channels than Amazon Transcribe is configured to process.
	//    To request additional channels, see Amazon Transcribe Limits (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits-amazon-transcribe)
	//    in the Amazon Web Services General Reference.
	FailureReason *string `type:"string"`

	// Provides information about how a transcription job is executed.
	JobExecutionSettings *JobExecutionSettings `type:"structure"`

	// The language code for the input speech.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// An object that describes the input media for the transcription job.
	Media *Media `type:"structure"`

	// The format of the input media file.
	MediaFormat MediaFormat `type:"string" enum:"true"`

	// The sample rate, in Hertz, of the audio track in the input media file.
	MediaSampleRateHertz *int64 `min:"8000" type:"integer"`

	// Optional settings for the transcription job. Use these settings to turn on
	// speaker recognition, to set the maximum number of speakers that should be
	// identified and to specify a custom vocabulary to use when processing the
	// transcription job.
	Settings *Settings `type:"structure"`

	// A timestamp that shows with the job was started processing.
	StartTime *time.Time `type:"timestamp"`

	// An object that describes the output of the transcription job.
	Transcript *Transcript `type:"structure"`

	// The name of the transcription job.
	TranscriptionJobName *string `min:"1" type:"string"`

	// The status of the transcription job.
	TranscriptionJobStatus TranscriptionJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s TranscriptionJob) String() string {
	return awsutil.Prettify(s)
}

// Provides a summary of information about a transcription job.
type TranscriptionJobSummary struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the job was completed.
	CompletionTime *time.Time `type:"timestamp"`

	// The content redaction settings of the transcription job.
	ContentRedaction *ContentRedaction `type:"structure"`

	// A timestamp that shows when the job was created.
	CreationTime *time.Time `type:"timestamp"`

	// If the TranscriptionJobStatus field is FAILED, a description of the error.
	FailureReason *string `type:"string"`

	// The language code for the input speech.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// Indicates the location of the output of the transcription job.
	//
	// If the value is CUSTOMER_BUCKET then the location is the S3 bucket specified
	// in the outputBucketName field when the transcription job was started with
	// the StartTranscriptionJob operation.
	//
	// If the value is SERVICE_BUCKET then the output is stored by Amazon Transcribe
	// and can be retrieved using the URI in the GetTranscriptionJob response's
	// TranscriptFileUri field.
	OutputLocationType OutputLocationType `type:"string" enum:"true"`

	// A timestamp that shows when the job started processing.
	StartTime *time.Time `type:"timestamp"`

	// The name of the transcription job.
	TranscriptionJobName *string `min:"1" type:"string"`

	// The status of the transcription job. When the status is COMPLETED, use the
	// GetTranscriptionJob operation to get the results of the transcription.
	TranscriptionJobStatus TranscriptionJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s TranscriptionJobSummary) String() string {
	return awsutil.Prettify(s)
}

// Provides information about a vocabulary filter.
type VocabularyFilterInfo struct {
	_ struct{} `type:"structure"`

	// The language code of the words in the vocabulary filter.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary was last updated.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the vocabulary filter. The name must be unique in the account
	// that holds the filter.
	VocabularyFilterName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s VocabularyFilterInfo) String() string {
	return awsutil.Prettify(s)
}

// Provides information about a custom vocabulary.
type VocabularyInfo struct {
	_ struct{} `type:"structure"`

	// The language code of the vocabulary entries.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary was last modified.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the vocabulary.
	VocabularyName *string `min:"1" type:"string"`

	// The processing state of the vocabulary. If the state is READY you can use
	// the vocabulary in a StartTranscriptionJob request.
	VocabularyState VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s VocabularyInfo) String() string {
	return awsutil.Prettify(s)
}
