package cloudformation

// AWSCloudWatchAlarm_Dimension AWS CloudFormation Resource (AWS::CloudWatch::Alarm.Dimension)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html
type AWSCloudWatchAlarm_Dimension struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html#cfn-cloudwatch-alarm-dimension-name
	Name string `json:"Name,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html#cfn-cloudwatch-alarm-dimension-value
	Value string `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudWatchAlarm_Dimension) AWSCloudFormationType() string {
	return "AWS::CloudWatch::Alarm.Dimension"
}
