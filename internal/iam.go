package internal

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

// IAM client struct
type IAM struct {
	Client *iam.IAM
}

// NewIamSess return IAM struct initialized
func NewIamSess(profile, region string) *IAM {
	return &IAM{
		Client: iam.New(getSession(profile, region)),
	}
}

// User iam user struct
type User struct {
	Name          string
	ManagedPolicy string
	InlinePolicy  string
	Group         string
	AccessKey     string
	AccessKeyUsed string
	PWLastUsed    string
	CreateDate    string
}

// Users User struct slice
type Users []User

// ListUsers return Users
// input iam.ListUsersInput
func (c *IAM) ListUsers(input *iam.ListUsersInput) (Users, error) {
	list := Users{}
	output := func(page *iam.ListUsersOutput, lastpage bool) bool {
		for _, o := range page.Users {
			used := "None"
			if o.PasswordLastUsed != nil {
				used = o.PasswordLastUsed.String()
			}

			username := *o.UserName

			mi := &iam.ListAttachedUserPoliciesInput{
				UserName: aws.String(username),
			}

			managed, err := c.ListAttachedUserPolicies(mi)
			if err != nil {
				return false
			}

			ii := &iam.ListUserPoliciesInput{
				UserName: aws.String(username),
			}

			inline, err := c.ListUserPolicies(ii)
			if err != nil {
				return false
			}

			gi := &iam.ListGroupsForUserInput{
				UserName: aws.String(username),
			}

			group, err := c.ListGroupsForUser(gi)
			if err != nil {
				return false
			}

			ai := &iam.ListAccessKeysInput{
				UserName: aws.String(username),
			}

			key, err := c.ListAccessKeys(ai)
			if err != nil {
				return false
			}

			list = append(list, User{
				Name:          username,
				ManagedPolicy: managed,
				InlinePolicy:  inline,
				Group:         group,
				AccessKey:     key,
				PWLastUsed:    used,
				CreateDate:    o.CreateDate.String(),
			})
		}
		return true
	}

	if err := c.Client.ListUsersPages(input, output); err != nil {
		return nil, fmt.Errorf("List users: %v", err)
	}

	return list, nil
}

// ListAttachedUserPolicies return string(managed policy names)
// input iam.ListAttachedUserPoliciesInput
func (c *IAM) ListAttachedUserPolicies(input *iam.ListAttachedUserPoliciesInput) (string, error) {
	output, err := c.Client.ListAttachedUserPolicies(input)
	if err != nil {
		return "", fmt.Errorf("List attached user policies: %v", err)
	}

	var (
		policies []string
		policy   string
	)
	for _, p := range output.AttachedPolicies {
		policies = append(policies, *p.PolicyName)
	}
	policy = strings.Join(policies[:], ",")

	if len(policy) == 0 {
		policy = "None"
	}

	return policy, nil
}

// ListUserPolicies return string(inline policy names)
// input iam.ListAttachedUserPoliciesInput
func (c *IAM) ListUserPolicies(input *iam.ListUserPoliciesInput) (string, error) {
	output, err := c.Client.ListUserPolicies(input)
	if err != nil {
		return "", fmt.Errorf("List user policies: %v", err)
	}

	var (
		policies []string
		policy   string
	)
	for _, p := range output.PolicyNames {
		policies = append(policies, *p)
	}
	policy = strings.Join(policies[:], ",")

	if len(policy) == 0 {
		policy = "None"
	}

	return policy, nil
}

// ListGroupsForUser return string(group names)
// input iam.ListGroupsForUserInput
func (c *IAM) ListGroupsForUser(input *iam.ListGroupsForUserInput) (string, error) {
	output, err := c.Client.ListGroupsForUser(input)
	if err != nil {
		return "", fmt.Errorf("List groups for user: %v", err)
	}

	var (
		groups []string
		group  string
	)
	for _, g := range output.Groups {
		groups = append(groups, *g.GroupName)
	}
	group = strings.Join(groups[:], ",")

	if len(group) == 0 {
		group = "None"
	}

	return group, nil
}

// ListAccessKeys return string(access keys)
// input iam.ListAccessKeysInput
func (c *IAM) ListAccessKeys(input *iam.ListAccessKeysInput) (string, error) {
	output, err := c.Client.ListAccessKeys(input)
	if err != nil {
		return "", fmt.Errorf("List access keys: %v", err)
	}

	var (
		keys []string
		key  string
	)
	for _, k := range output.AccessKeyMetadata {
		keys = append(keys, *k.AccessKeyId)
	}
	key = strings.Join(keys[:], ",")

	if len(key) == 0 {
		key = "None"
	}

	return key, nil
}
