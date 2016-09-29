// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elb

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

func (c *ELB) WaitUntilAnyInstanceInService(input *DescribeInstanceHealthInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstanceHealth",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAny",
				Argument: "InstanceStates[].State",
				Expected: "InService",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *ELB) WaitUntilInstanceDeregistered(input *DescribeInstanceHealthInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstanceHealth",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "InstanceStates[].State",
				Expected: "OutOfService",
			},
			{
				State:    "success",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidInstance",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *ELB) WaitUntilInstanceInService(input *DescribeInstanceHealthInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstanceHealth",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "InstanceStates[].State",
				Expected: "InService",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
