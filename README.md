# longplan-chc
[Longitudinal Plan API](https://docs.healtheintent.com/api/v1/longitudinal_plan/) interactions for CHC 2019.

## Modules
- models
    - Structs used to represent resources used in interacting with [HealtheIntent API](https://docs.healtheintent.com/) interactions.
- client 
    - Helper functions for executing HTTP calls to HealtheIntent APIs.
- definitions
    - Example of using CSV files as an input source for creating Longitudinal Plan content.
- plan
    - Example locates an individual's consumer record (or creates one if it doesn't exisit), create's instances of the Health Concern and Goals as defined by a provided template for the target consumer.
