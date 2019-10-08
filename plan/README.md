# longplan-chc/plan
An simple example of integrating various [HealtheIntent services](https://docs.healtheintent.com/#services) to interact with a consumer's [Health Concerns](https://docs.healtheintent.com/api/v1/health_concern/) and [Longitudinal Plan](https://docs.healtheintent.com/api/v1/longitudinal_plan/).

## How to run?

From the `/plan` directory execute:

```
> go run plan.go
```

A HealtheIntent [tenant](https://docs.healtheintent.com/#uri-structure) mnemonic, [authorization header](https://docs.healtheintent.com/#authentication), millennium person id and [template id](https://docs.healtheintent.com/api/v1/longitudinal_plan/#plan-templates) are required input.

## What does it do?

After providing the required input when prompted, the following is the general outline of actions that will be performed:
- [Find or Create a Consumer](https://docs.healtheintent.com/api/v1/consumer/) entity using the millennium person id provided.
- [Find the Longitudinal Plan Template] specified as input.
- Create the [Health Concerns](https://docs.healtheintent.com/api/v1/health_concern/#create-a-health-concern) and [Goals](https://docs.healtheintent.com/api/v1/longitudinal_plan/#create-a-goal) that are specified from the provided template id for the identified consumer.
- [Relate the Health Concern to the Goals](https://docs.healtheintent.com/api/v1/longitudinal_plan/#relate-a-health-concern-to-a-goal) as defined by the template.