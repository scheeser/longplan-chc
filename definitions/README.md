# longplan-chc/definitions
An simple example of using HealtheIntent [Logitduinal Plan](https://docs.healtheintent.com/api/v1/longitudinal_plan/) and [Health Concern](https://docs.healtheintent.com/api/v1/health_concern/) APIs to create a Health Concern Definition, Goal Definitions and wrap those items up in a Longitudinal Plan Template.

## How to run?

From the `/definitions` directory execute:

```
> go run definitions.go
```

A HealtheIntent [tenant](https://docs.healtheintent.com/#uri-structure) mnemonic and [authorization header](https://docs.healtheintent.com/#authentication) are required input.

## What does it do?

After providing the required input when prompted, the following is the general outline of actions that will be performed:
- [Create a Health Concern definition](https://docs.healtheintent.com/api/v1/health_concern/#create-a-health-concern-definition) for the condition Diabetes Mellitus Type 1
- Parse the CSV with the name `goal_definitions.csv`
- [Create a Goal Definition](https://docs.healtheintent.com/api/v1/longitudinal_plan/#create-a-goal-definition) for every line in the CSV above.
- [Create a Plan Template](https://docs.healtheintent.com/api/v1/longitudinal_plan/#create-a-plan-template) using the Health Concern and Goal Definitions from the previous steps.

Each line from the CSV is expected to be in the followng format:
```
Goal Definition Text,Coding System,Coding Code,Coding Display
```