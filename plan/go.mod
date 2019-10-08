module github.cerner.com/JS016083/longplan-chc/plan

go 1.13

require (
	github.cerner.com/JS016083/longplan-chc/client v0.0.0
	github.cerner.com/JS016083/longplan-chc/models v0.0.0
)

replace github.cerner.com/JS016083/longplan-chc/models => ../models

replace github.cerner.com/JS016083/longplan-chc/client => ../client
